package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/services/checker"
	yaml "gopkg.in/yaml.v2"
)

var garbage string

//Services map containing all the services
var Services = NewCMap()

//DaemonActive are we currently checking services?
var DaemonActive = false

//Service struct containing a checks parameters
type Service struct {
	Identifier       string       `json:"identifier"`
	Description      string       `json:"description"`
	Enabled          bool         `json:"enabled"`
	Acknowledged     bool         `json:"acknowledged"`
	Host             string       `json:"host"`
	Command          string       `json:"command"`
	Timeout          int          `json:"timeout"`
	Interval         int          `json:"interval"`
	Action           ActionConfig `json:"action"`
	Threshold        int          `json:"threshold"`
	ThresholdCounter int
	Health           int
	LastCheck        time.Time
	Claimed          bool
	RTime            int64
	Output           string
	Exposed          bool
}

//Enable a service
func (service *Service) Enable() {
	service.Enabled = true
	Services.Set(service.Identifier, *service)
}

//Disable a service
func (service *Service) Disable() {
	service.Enabled = false
	service.Health = -1 //health is not important because it is disabled now
	Services.Set(service.Identifier, *service)
}

//Lock a service
func (service *Service) Claim() {
	service.Claimed = true
}

//Unlock a service
func (service *Service) Release() {
	service.Claimed = false
}

//Reschedule Set last check date so early that it has to be rechecked ASAP
func (service *Service) Reschedule() {
	service.LastCheck, _ = time.Parse(time.UnixDate, "Sat Mar  7 11:06:39.1234 PST 1990")
	Services.Set(service.Identifier, *service)
}

func GetPublicServices(group string) {
	//TODO: build this
}

//UpdateList compare current map with fresh JSON getServices() and update values
func UpdateList() {
	//Do not start a new check while updating
	jsonServices := getServices() //[]Service
	jsonServicesMap := NewCMap()  //Concurrent string-Service map

	for _, newService := range jsonServices {
		//oldService := Services[newService.Identifier]
		oldService, _ := Services.Get(newService.Identifier)

		if oldService.Identifier != newService.Identifier {
			newService.Health = 2
		} else {
			newService.Health = oldService.Health
		}

		newService.Claimed = oldService.Claimed
		newService.LastCheck = oldService.LastCheck
		newService.ThresholdCounter = oldService.ThresholdCounter
		newService.Output = oldService.Output
		newService.RTime = oldService.RTime

		jsonServicesMap.Set(newService.Identifier, newService)
		log.Warning("Reloaded " + oldService.Identifier + " as " + newService.Identifier)
	}

	Services = jsonServicesMap

}

//Update a service from fresh getServices()
func (service Service) Update() error {
	//Do not start a new check while updating
	service.Claim()
	for _, newService := range getServices() {
		if newService.Identifier == service.Identifier {

			newService.copyMemoryAttributes(&service)
			//push new service to Services map
			log.Info("Setting service -> ", service.Identifier, newService)
			Services.Set(service.Identifier, newService)

			log.Notice("(!!) Reloaded " + service.Identifier + " from " + newService.Identifier)
			return nil
		}
	}

	return errors.New("Service not found.")
}

//reloadServiceCopy: copies in-memory attributes of service to new service
func (new Service) copyMemoryAttributes(original *Service) {

	new.Claim()
	//Copy in-memory attributes of service to new service
	new.LastCheck = original.LastCheck
	new.Health = original.Health
	new.ThresholdCounter = original.ThresholdCounter
	new.Output = original.Output
	new.RTime = original.RTime

	new.Release()

}

func (service Service) getJSON() string {
	b, err := json.Marshal(service)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (service *Service) spawnChild() int {

	args := service.Command
	args = strings.Replace(args, "$HOST$", service.Host, -1)
	args = strings.Replace(args, "$TIMEOUT$", strconv.Itoa(service.Timeout), -1)

	//log.Warning("::::SpawnChild::Checking for " + service.Identifier + " - " + args)
	status, output, rtime := checker.CheckService(args)
	service.Output = output

	if status > 0 { //It's going down
		oldHealth := service.Health
		//service.Health = 2
		//service.Comment = "Output: " + output
		service.ThresholdCounter++

		if oldHealth == -1 { //cold check, now its down
			service.Health = 1 //set warning state
		}

		if oldHealth == 0 {
			service.Health = 1 //(re)set warning state
		}

		if oldHealth == 1 && service.ThresholdCounter >= service.Threshold {
			service.Health = 2       //It's officially down!
			NewAction(service).Run() //Ready for action
		}
	} else {
		oldHealth := service.Health
		service.Health = 0
		service.ThresholdCounter = 0
		if oldHealth == 2 {
			a := NewAction(service) //Ready for recovery notify
			a.Run()
		}
	}

	service.Release()
	service.RTime = rtime
	service.LastCheck = time.Now()
	Services.Set(service.Identifier, *service)

	return status
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func checkDispatcher() {

	for {
		if DaemonActive == true {
			for item := range Services.IterBuffered() {
				service := item.Val
				key := item.Key

				diff := int(time.Now().Unix()) - int(service.LastCheck.Unix())

				if service.Enabled {
					if diff > service.Interval && service.Claimed == false {
						//lock current check
						service.Claim()

						//update status in map
						//Services[key] = service
						Services.Set(key, service)

						//spawn check for service
						go service.spawnChild()

					}
				}
			}
		}

		for i := 0; i < 1; i++ {
			time.Sleep(time.Second)
		}

	}
}

//Init service module
func Init() {
	reloadServices()
	go checkDispatcher()
}

//Start daemon's checking
func Start() {
	log.Notice("Starting..")
	DaemonActive = true
}

//Stop daemon from checking
func Stop() {
	log.Warning("Stopping..")
	DaemonActive = false
}

func reloadServices() {
	log.Notice("Reading JSON")

	for _, service := range getServices() {
		service.Health = -1 //you know nothing, monitoring
		Services.Set(service.Identifier, service)
	}
}

//TestConfiguration checks if configuration can be loaded and shows amount of services
func TestConfiguration() error {
	services := getServices()
	length := len(services)
	log.Info("Length:", length, "services")
	if length > 0 {
		return nil
	} else {
		return errors.New("Amount of services is invalid.")
	}
}

func servicesBuilder(path string) []Service {
	var fileServices []Service
	err := yaml.Unmarshal(readFile(path), &fileServices)
	if err != nil {
		panic(err)
	}
	return fileServices
}

func readFile(path string) []byte {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic("Cannot read file: ", path, err.Error())
	}
	return raw
}

func readServiceFiles(searchDir string) []Service {
	findingList := []string{}
	var allServices []Service
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		findingList = append(findingList, path)
		if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
			log.Notice("Loaded ", path)
			allServices = append(allServices, servicesBuilder(path)...)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return allServices
}

func getServices() []Service {
	services := readServiceFiles(configuration.C.Paths.Services)
	if len(services) < 1 {
		log.Panic("No services found! This makes me useless! Panic!")
	}
	return services
}
