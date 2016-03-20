package handlers

import (
	"fmt"

	"github.com/bartholdbos/golegram"
	"github.com/mdeheij/monitoring/configuration"
)

type instanceHolder struct {
	bot *golegram.Bot
}

var instance instanceHolder

//Telegram sends a Telegram message to one or more users by their unique ID
func Telegram(targets []string, message string) {
	var err error

	fmt.Println("[Telegram] Sending to following target(s): ", targets)

	if instance.bot == nil {
		//TODO: get config from file
		instance.bot, err = golegram.NewBot(configuration.Config.TelegramBotToken)
		//instance.bot, err = golegram.NewBot("94110015:AAE8TIIoQxyu4KdWRnGZ2_yvI9C6-1w1eF0")
	}
	if err == nil {
		for _, target := range targets {
			result, err := instance.bot.SendMessage(target, message)
			if err != nil {
				fmt.Println("[Telegram] Message error")
				fmt.Println(err)
				fmt.Println(result)
			}
		}
	}
}
