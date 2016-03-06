package services

import (
	"fmt"
	"github.com/mdeheij/monitoring/services/handlers"
	"strconv"
	"time"
)

type ActionConfig struct {
	Name           string
	Telegramtarget []int32
	Rpecommand     string
}

type ActionHandler struct {
	name    string
	service Service
}

func NewAction(service Service) *ActionHandler {
	var a ActionHandler = ActionHandler{name: "telegram"}
	a.service = service
	return &a
}

func (a ActionHandler) buildMessage() (msg string) {
	timestamp := a.service.LastCheck.Format(time.Stamp)

	thresholdCounting := strconv.Itoa(a.service.ThresholdCounter) + "/" + strconv.Itoa(a.service.Threshold)

	actionTypeString := ""
	switch a.service.Health {
	case 2:
		actionTypeString = "🔴 PROBLEM"
	case 0:
		actionTypeString = "🔵 RECOVERY"
	case 1:
		actionTypeString = "⚠️ WARNING"
	}

	return "*" + actionTypeString + " - " + timestamp + "*\n" + a.service.Identifier + " (" + a.service.Host + ")" + "\nThreshold: " + thresholdCounting + "\nOutput: _" + a.service.Output + "_"
}

func (a ActionHandler) Run() {
	if SilenceAll == false {
		if a.service.Acknowledged != true {

			switch a.service.Action.Name {
			case "telegram":
				handlers.Telegram(a.service.Action.Telegramtarget, a.buildMessage())
			case "rpe":
				handlers.RemotePluginExecutor(a.service.Host)
			case "none":
				DebugMessage(a.service.Identifier + ", doing nothing.")
			default:
				handlers.Telegram(a.service.Action.Telegramtarget, a.buildMessage())
			}

		}
	} else {
		fmt.Println("Silenced.")
	}
}
