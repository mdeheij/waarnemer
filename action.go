package monitoring

import (
	//"fmt"
	"github.com/bartholdbos/golegram"
	"strconv"
	"time"
)

type Action struct {
	actionType    int
	errorMsg      string
	service       Service
	actionHandler string
}

var s Service

func NewAction(service Service) *Action {
	var a Action = Action{actionHandler: "telegram"}
	a.service = service
	return &a
}

func (a Action) messageBuilder() (msg string) {
	timestamp := a.service.LastCheck.Format(time.RFC822)

	thresholdCounting := strconv.Itoa(a.service.ThresholdCounter) + "/" + strconv.Itoa(a.service.Threshold)

	actionTypeString := ""
	switch a.service.Health {
	case 2:
		actionTypeString = "🔴 PROBLEM"
	case 0:
		actionTypeString = "🔵 RECOVERY"
	case 1:
		actionTypeString = "⚠️ WARNING" //warning, no idea when you actually want to show this
	}

	return actionTypeString + " - " + timestamp + "\n" + a.service.Label + " (" + a.service.Host + ")" + "\nThreshold: " + thresholdCounting + "\nOutput: " + a.errorMsg
}

func (a Action) SendTelegram() {
	//shit, need to generate a new one because i just pushed my token to a public repo
	bot, _ := Golegram.NewBot("94110015:AAE8TIIoQxyu4KdWRnGZ2_yvI9C6-1w1eF0")
	_, _ = bot.SendMessage(4009810, a.messageBuilder())

}

func (a Action) Run() {

	a.SendTelegram()
}
