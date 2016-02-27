package services

import (
	//"flag"
	"fmt"
	//"os"
	"github.com/mdeheij/monitoring/configuration"
	"testing"
	//"time"
)

func TestServiceInit(t *testing.T) {
	DebugMode = false
	fmt.Println("ServicesConfig")
	configuration.Init("../config.json")
	reloadServices()
}

func TestServiceCheck(t *testing.T) {
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
