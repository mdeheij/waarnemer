package action

import (
	"fmt"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/message"
	"github.com/mdeheij/monitoring/services/handlers"
	"github.com/mdeheij/monitoring/services/model"
)

//Handle dispatches the correct action handling
func Handle(service model.Service) {
	if configuration.C.NoActionHandling == false {
		if service.Acknowledged != true {

			switch service.Action.Name {
			case "telegram":
				handleTelegram(service)
			case "none":
				log.Notice("Skipping action for ", service.Identifier)
			default:
				handleTelegram(service)
			}

		}
	} else {
		fmt.Println("Silenced.")
	}
}

func handleTelegram(service model.Service) {
	message := message.BuildNotificationMessage(service.Identifier, service.Health, service.Host, service.ThresholdCounter, service.Threshold, service.Output)
	handlers.Telegram(service.Action.Telegramtarget, message)
}
