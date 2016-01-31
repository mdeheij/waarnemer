package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//	"reflect"
	"git.gate.sh/mdeheij/monitoring/configuration"
	"strconv"
	"strings"
	"time"
)

var garbage string

var Services = make(map[string]Service)

var DaemonActive = false
var SilenceAll = false

type ServicesConfiguration struct {
	BaseFolder       string
	TelegramBotToken string `json:"TelegramBotToken"`
}

var ServicesConfig ServicesConfiguration

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

func (service Service) Enable() {
	//enable
	service.Enabled = true
	Services[service.Identifier] = service
	//commit naar db
}

func (service Service) Disable() {
	//enable
	service.Enabled = false
	Services[service.Identifier] = service
	//commit naar db
}

func (service Service) Reschedule() {
	service.LastCheck, _ = time.Parse(time.UnixDate, "Sat Mar  7 11:06:39.1234 PST 1990")
	Services[service.Identifier] = service
}

func UpdateList() {
	//Do not start a new check while updating
	jsonServices := getServices() //dit is geen map
	jsonServicesMap := make(map[string]Service)

	for _, newService := range jsonServices {
		oldService := Services[newService.Identifier]

		newService.Lock = oldService.Lock
		newService.LastCheck = oldService.LastCheck
		newService.Health = oldService.Health
		newService.ThresholdCounter = oldService.ThresholdCounter
		newService.Output = oldService.Output
		newService.RTime = oldService.RTime

		jsonServicesMap[newService.Identifier] = newService
		fmt.Println("Reloaded " + oldService.Identifier + " as " + newService.Identifier)
	}

	Services = jsonServicesMap

}

func (service Service) Update() string {
	//Do not start a new check while updating
	service.Lock = true
	for _, newService := range getServices() {
		if newService.Identifier == service.Identifier {
			//Copy in-memory attributes of service to new service
			newService.LastCheck = service.LastCheck
			newService.Health = service.Health
			newService.ThresholdCounter = service.ThresholdCounter
			newService.Output = service.Output
			newService.RTime = service.RTime

			//push new service to Services map
			Services[service.Identifier] = newService

			return "(!!) Reloaded " + service.Identifier + " from " + newService.Identifier
		}
	}
	return "ERROR: SERVICE NOT FOUND"

}

func (service Service) getJSON() string {
	b, err := json.Marshal(service)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

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

func (service Service) spawnChild() {

	args := service.Command
	args = strings.Replace(args, "$HOST$", service.Host, -1)
	args = strings.Replace(args, "$TIMEOUT$", strconv.Itoa(service.Timeout), -1)

	//fmt.Println("::::SpawnChild::Checking for " + service.Identifier + " - " + args)
	status, output, rtime := CheckService(args)
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

	service.Lock = false
	service.RTime = rtime
	service.LastCheck = time.Now()
	Services[service.Identifier] = service
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func checkDispatcher() {

	for {
		if DaemonActive == true {
			for key, service := range Services {
				diff := int(time.Now().Unix()) - int(service.LastCheck.Unix())
				//editableService := service
				//fmt.Print("ID:", key, "  ")
				//	fmt.Println(diff)

				if service.Enabled {
					if diff > service.Interval && service.Lock == false {
						//lock current check
						service.Lock = true

						//update status in map
						Services[key] = service

						//spawn check for service
						go service.spawnChild()

					} else {
						//fmt.Print(service.Identifier + " will check in: ")
						//fmt.Println(nextCheck)
					}
				}
			}
		} else {
			//fmt.Println("Not checking because DaemonActive is not true")
		}

		for i := 0; i < 1; i++ {
			//	fmt.Println("● ")
			time.Sleep(1 * time.Second)
			//fmt.Println("Next run in " + strconv.Itoa(3-i) + "..")
		}

	}
}

func Init() {

	reloadServices()
	go checkDispatcher()
}
func Start() {
	fmt.Println("Starting..")
	DaemonActive = true
	a := NewAction(Service{Host: configuration.Config.Hostname, Identifier: "monitoring.daemon", Threshold: 3, Health: 1, Output: "Monitoring started!", Action: ActionConfig{Name: "telegram", Telegramtarget: []int32{configuration.Config.TelegramNotificationTarget}}})
	a.Run()
}
func Stop() {
	fmt.Println("Stopping..")
	DaemonActive = false
	a := NewAction(Service{})
	a.Run()
}

func reloadServices() {
	fmt.Println("Reading JSON")
	var count int
	fmt.Println("━━━━━━━━━━[Loading services]━━━━━━━━━━━━━━━┉┉┉┉┉┉┈┈┈ ")
	fmt.Println("TG BOT TOKEN: (" + configuration.Config.TelegramBotToken + ")")
	for _, service := range getServices() {
		service.Health = -1 //you know nothing, monitoring
		Services[service.Identifier] = service
		fmt.Println("Loaded " + service.Identifier)
		count++
	}
	fmt.Println("-----")
	fmt.Print("Total: ")
	fmt.Println(count)

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┉┉┉┉┉┉┈┈┈ ")

}

func getServices() []Service {
	raw, err := ioutil.ReadFile(configuration.Config.BaseFolder + "services.json")
	if err != nil {
		panic(err)
	}

	var s []Service
	json.Unmarshal(raw, &s)
	return s
}
