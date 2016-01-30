package handlers

import (
	"fmt"
	"github.com/bartholdbos/golegram"
)

type Instance struct {
	bot *golegram.Bot
}

var instance Instance

func Telegram(targets []int32, message string) {
	var err error
	fmt.Println("Sending to following targets: ", targets)

	if instance.bot == nil {
		//TODO: get config from file
		instance.bot, err = golegram.NewBot("94110015:AAE8TIIoQxyu4KdWRnGZ2_yvI9C6-1w1eF0")
	}
	if err == nil {
		for _, target := range targets {
			result, err := instance.bot.SendMessage(target, message)
			if err != nil {
				fmt.Println(err)
				fmt.Println(result)
			} else {
				fmt.Println("Geen error bij zenden van Telegram")
			}
		}
	}
}
