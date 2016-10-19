package services

import (
	"fmt"
	"testing"

	"github.com/mdeheij/monitoring/configuration"
)

func TestAction(t *testing.T) {
	//fmt.Println(ServicesConfig)
	//fmt.Println("[Testing]")

	t.SkipNow() //TODO: fix this test.
	configuration.Init("config.yaml")

	tgSlice := []string{configuration.C.Actions.Telegram.Target}
	ac := ActionConfig{Name: "telegram", Telegramtarget: tgSlice}
	s := &Service{Host: "go test", Identifier: "TestAction", Threshold: 3, Health: 1, Output: "OK", Action: ac}

	a := NewAction(s)
	a.Run()

	s = &Service{Host: "dev", Identifier: "localhost", Threshold: 3, Health: 1, Output: "OK", Action: ActionConfig{Name: "rpe"}}

	a = NewAction(s)
	a.Run()

	fmt.Println("A test message should be sent. Please verify.")
}
