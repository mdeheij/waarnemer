package services

import (
	"encoding/json"
	"fmt"
	"github.com/mdeheij/monitoring/configuration"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Services map containing all the services
var Services = NewCMap()

// DaemonActive indicates whether or not there are checking services active.
var DaemonActive = false

// DebugMode whether or not to show verbose output.
var DebugMode = false

// SilenceAll whether or not to TODO: implement this.
var SilenceAll = false

var garbage string

// Service describes all information that is required for a check.
type Service struct {
	Identifier       string       `json:"identifier"`
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
	Lock             bool
	RTime            int64
	Output           string
}

// Enable enables a service.
func (service Service) Enable() {
	service.Enabled = true
	Services.Set(service.Identifier, service)
}

// Disable disables a service.
func (service Service) Disable() {
	service.Enabled = false
	Services.Set(service.Identifier, service)
}

// Reschedule sets the last check date so early that it has to be checked again.
func (service Service) Reschedule() {
	// TODO: verify whether or not spaces matter in a UNIX date.
	service.LastCheck, _ = time.Parse(time.UnixDate, "Sat Mar  7 11:06:39.1234 PST 1990")
	Services.Set(service.Identifier, service)
}

// EnableDebug sets debugmode to true
func EnableDebug() {
	DebugMode = true
}

// DebugMessage prints text when DebugMode is true.
func DebugMessage(text interface{}) {
	if DebugMode {
		fmt.Println(text)
	}
}

// UpdateList compares the current map with fresh JSON getServices() and updates values.
func UpdateList() {
	//Do not start a new check while updating
	jsonServices := getServices() //[]Service
	jsonServicesMap := NewCMap()  //Concurrent string-Service map

	for _, newService := range jsonServices {
		oldService, _ := Services.Get(newService.Identifier)

		newService.Lock = oldService.Lock
		newService.LastCheck = oldService.LastCheck
		newService.Health = oldService.Health
		newService.ThresholdCounter = oldService.ThresholdCounter
		newService.Output = oldService.Output
		newService.RTime = oldService.RTime

		jsonServicesMap.Set(newService.Identifier, newService)
		DebugMessage("Reloaded " + oldService.Identifier + " as " + newService.Identifier)
	}

	Services = jsonServicesMap
}

// Update updates a service from fresh getServices().
func (service Service) Update() string {
	// Do not start a new check while updating.
	service.Lock = true

	for _, newService := range getServices() {
		if newService.Identifier == service.Identifier {
			// Copy in-memory attributes of service to new service.
			newService.LastCheck = service.LastCheck
			newService.Health = service.Health
			newService.ThresholdCounter = service.ThresholdCounter
			newService.Output = service.Output
			newService.RTime = service.RTime

			// Push new service to `Services` map.
			fmt.Println("Setting service :112 -> ", service.Identifier, newService)
			Services.Set(service.Identifier, newService)

			return "(!!) Reloaded " + service.Identifier + " from " + newService.Identifier
		}
	}

	return "ERROR: SERVICE NOT FOUND"
}

// StatusColor generates a command line colour based on health.
func StatusColor(text string, health int) string {

	switch health {
	case 0:
		return "\x1b[32;1m" + text + "\x1b[0m"
	case 1:
		return "\x1b[33;1m" + text + "\x1b[0m"
	case 2:
		return "\x1b[31;1m" + text + "\x1b[0m"
	default:
		return "\x1b[33;1m ERR " + text + "\x1b[0m"
	}
}

// Init intializes the service module.
func Init() {
	reloadServices()
	go checkDispatcher()
}

// Start starts daemon checking.
func Start() {
	DebugMessage("Starting...")
	DaemonActive = true

	if DebugMode {
		a := NewAction(Service{Host: configuration.Config.Hostname, Identifier: "monitoring.daemon", Threshold: 3, Health: 1, Output: "Monitoring started!", Action: ActionConfig{Name: "telegram", Telegramtarget: []int32{configuration.Config.TelegramNotificationTarget}}})

		a.Run()
	}
}

//Stop stops the daemon from checking.
func Stop() {
	DebugMessage("Stopping...")
	DaemonActive = false

	a := NewAction(Service{})
	a.Run()
}

func (service Service) getJSON() string {
	b, err := json.Marshal(service)

	if err != nil {
		panic(err)
	}

	return string(b)
}

//TestConfiguration checks if configuration can be loaded and shows amount of services
func TestConfiguration() {
	fmt.Println("Length:", len(getServices()), "services")
}

func (service Service) spawnChild() int {
	args := service.Command
	args = strings.Replace(args, "$HOST$", service.Host, -1)
	args = strings.Replace(args, "$TIMEOUT$", strconv.Itoa(service.Timeout), -1)

	status, output, rtime := CheckService(args)
	service.Output = output

	if status > 0 { //It's going down
		oldHealth := service.Health
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

	service.Lock = false
	service.RTime = rtime
	service.LastCheck = time.Now()
	Services.Set(service.Identifier, service)

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
				//editableService := service
				//fmt.Print("ID:", key, "  ")
				//	DebugMessage(diff)

				if service.Enabled {
					if diff > service.Interval && service.Lock == false {
						//lock current check
						service.Lock = true

						//update status in map
						//Services[key] = service
						Services.Set(key, service)

						//spawn check for service
						go service.spawnChild()

					} else {
						//fmt.Print(service.Identifier + " will check in: ")
						//DebugMessage(nextCheck)
					}
				}
			}
		} else {
			//DebugMessage("Not checking because DaemonActive is not true")
		}

		for i := 0; i < 1; i++ {
			//	DebugMessage("● ")

			time.Sleep(time.Second)
			//DebugMessage("Next run in " + strconv.Itoa(3-i) + "..")
		}
	}
}

func getServices() []Service {
	raw, err := ioutil.ReadFile(configuration.Config.BaseFolder + "services.json")

	if err != nil {
		DebugMessage("Cannot read file!")
		panic(err)
	} else {
		DebugMessage(string(raw))
	}

	var s []Service
	errUnmarshal := json.Unmarshal(raw, &s)

	if errUnmarshal != nil {
		DebugMessage("Cannot parse JSON file!")
		panic(err)
	}

	return s
}

func reloadServices() {
	DebugMessage("Reading JSON")
	var count int

	DebugMessage("Telegram Bot Token: (" + configuration.Config.TelegramBotToken + ")")

	for _, service := range getServices() {
		service.Health = -1 //you know nothing, monitoring
		Services.Set(service.Identifier, service)
		DebugMessage("Loaded " + service.Identifier)

		count++
	}
}
