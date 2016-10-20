package action

import (
	"fmt"
	"strconv"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/services/handlers"
	"github.com/mdeheij/monitoring/services/model"
)

func Handle(service model.Service) {
	if configuration.C.NoActionHandling == false {
		if service.Acknowledged != true {

			switch service.Action.Name {
			case "telegram":
				handlers.Telegram(service.Action.Telegramtarget, buildMessage(service))
			case "none":
				log.Notice("Skipping action for ", service.Identifier)
			default:
				handlers.Telegram(service.Action.Telegramtarget, buildMessage(service))
			}

		}
	} else {
		fmt.Println("Silenced.")
	}
}

func buildMessage(service model.Service) (msg string) {
	thresholdCounting := strconv.Itoa(service.ThresholdCounter) + "/" + strconv.Itoa(service.Threshold)
	actionTypeString := ""

	switch service.Health {
	case 2:
		actionTypeString = "🔴"
	case 0:
		actionTypeString = "✅"
	case 1:
		actionTypeString = "⚠️"
	}

	return fmt.Sprintf("%s %s (%s) %s\n %s", actionTypeString, service.Identifier, service.Host, thresholdCounting, service.Output)
}
