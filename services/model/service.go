package model

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/mdeheij/monitoring/services/checker"
)

var garbage string

//Services map containing all the services
var Services = NewCMap()

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

//reloadServiceCopy: copies in-memory attributes of service to new service
func (new Service) CopyMemoryAttributes(original *Service) { //TODO: rename this

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

func (service *Service) SpawnCheck() int {

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
