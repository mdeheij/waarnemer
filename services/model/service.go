package model

import (
	"time"

	"github.com/mdeheij/monitoring/services/model/health"
)

var garbage string

//Services map containing all the services
var Services = NewCMap()

//ActionConfig is defined in a service to be handled when considered down.
type ActionConfig struct {
	Name           string
	Telegramtarget []string
	Rpecommand     string
}

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
	service.Health = health.UNKNOWN //health is not important because it is disabled now
	Services.Set(service.Identifier, *service)
}

//Claim a service while checking
func (service *Service) Claim() {
	service.Claimed = true
}

//Release a service after checking
func (service *Service) Release() {
	service.Claimed = false
}

//Reschedule Set last check date so early that it has to be rechecked ASAP
func (service *Service) Reschedule() {
	service.LastCheck, _ = time.Parse(time.UnixDate, "Sat Mar  7 11:06:39.1234 PST 1990")
	Services.Set(service.Identifier, *service)
}

//CopyMemoryAttributes copies in-memory attributes of a service to a new service
func (new Service) CopyMemoryAttributes(original *Service) { //TODO: rename this
	new.Claim()

	new.LastCheck = original.LastCheck
	new.Health = original.Health
	new.ThresholdCounter = original.ThresholdCounter
	new.Output = original.Output
	new.RTime = original.RTime

	new.Release()
}
