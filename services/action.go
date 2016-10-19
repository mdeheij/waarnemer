package services

import (
	"fmt"
	"strconv"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/services/handlers"
)

//ActionConfig is defined in a service to be handled when considered down.
type ActionConfig struct {
	Name           string
	Telegramtarget []string
	Rpecommand     string
}

type ActionHandler struct {
	name    string
	service *Service
}

func NewAction(service *Service) *ActionHandler {
	var a = ActionHandler{name: "telegram"}
	a.service = service
	return &a
}

func (a ActionHandler) buildMessage() (msg string) {
	//	timestamp := a.service.LastCheck.Format(time.Stamp)

	thresholdCounting := strconv.Itoa(a.service.ThresholdCounter) + "/" + strconv.Itoa(a.service.Threshold)

	actionTypeString := ""
	switch a.service.Health {
	case 2:
		actionTypeString = "🔴"
	case 0:
		actionTypeString = "✅"
	case 1:
		actionTypeString = "⚠️"
	}

	return fmt.Sprintf("%s %s (%s) %s\n %s", actionTypeString, a.service.Identifier, a.service.Host, thresholdCounting, a.service.Output)
}

//Run on ActionHandler will execute specified action for service when considered down.
func (a ActionHandler) Run() {
	if configuration.C.NoActionHandling == false {
		if a.service.Acknowledged != true {

			switch a.service.Action.Name {
			case "telegram":
				handlers.Telegram(a.service.Action.Telegramtarget, a.buildMessage())
			case "none":
				log.Notice("Skipping action for ", a.service.Identifier)
			default:
				handlers.Telegram(a.service.Action.Telegramtarget, a.buildMessage())
			}

		}
	} else {
		fmt.Println("Silenced.")
	}
}
