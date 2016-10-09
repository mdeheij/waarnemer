package services

import (
	//"flag"
	//	"fmt"
	//"os"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/mdeheij/monitoring/configuration"
	"testing"
	//"time"
)

func NoTestReload(t *testing.T) {
	identifier := "github.web"

	sampleService := Service{Identifier: identifier, Host: "hardcoded", Command: "hardcoded", Timeout: 5, Interval: 15}

	sampleService.Update()

	service, _ := Services.Get(identifier)

	if service.Host == "hardcoded" {
		t.Fail()
	}
}

func NoTestServiceCheck(t *testing.T) {
	onlineCheck := Service{Identifier: "github.web", Host: "localhost", Command: "ping -H $HOST$ -t $TIMEOUT$", Timeout: 5, Interval: 15}
	statusOnlineCheck := onlineCheck.spawnChild()
	if statusOnlineCheck > 0 {
		t.Fail()
	}

	offlineCheck := Service{Identifier: "this.should.be.broken", Host: "", Command: "ping -H $HOST$ -t $TIMEOUT$", Timeout: 0, Interval: 15}
	statusOfflineCheck := offlineCheck.spawnChild()
	if statusOfflineCheck != 2 {
		t.Fail()
	}
}
