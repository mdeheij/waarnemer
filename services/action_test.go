package services

import (
	"flag"
	"fmt"
	"testing"
)

var target int

func init() {
	flag.IntVar(&target, "target", 0, "The telegram target")
	flag.Parse()
}

func NoTestAction(t *testing.T) {
	//fmt.Println(ServicesConfig)
	//fmt.Println("[Testing]")

	if target == 0 {
		fmt.Println("Telgeram Target ID is required. Use --target CHAT_ID")
		t.Fail()
	}
	//tgtg := []int32{4009810, -62946040}
	cijfer := int32(target)
	tgtg := []int32{cijfer}
	ac := ActionConfig{Name: "telegram", Telegramtarget: tgtg}
	s := Service{Host: "localhost", Identifier: "test", Threshold: 3, Health: 1, Output: "OK", Action: ac}

	a := NewAction(s)
	a.Run()

	s = Service{Host: "dev", Identifier: "localhost", Threshold: 3, Health: 1, Output: "OK", Action: ActionConfig{Name: "rpe"}}

	a = NewAction(s)
	a.Run()
}
