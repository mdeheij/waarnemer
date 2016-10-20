package daemon

import (
	"time"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/services"
	"github.com/mdeheij/monitoring/services/checker"
	"github.com/mdeheij/monitoring/services/model"
)

var spawned bool

//IsActive returns true when daemon is actually checking services
var IsActive bool

//SpawnDispatcher starts a checkDispatcher
func Spawn() {
	if !spawned {
		log.Warning("Spawning the daemon!")
		services.Load()
		go checkDispatcher()
	} else {
		log.Error("Daemon is already active!")
	}
}

//Start daemon's checking
func Start() {
	log.Notice("Starting..")
	IsActive = true
}

//Stop daemon from checking
func Stop() {
	log.Warning("Stopping..")
	IsActive = false
}

func checkDispatcher() {
	spawned = true
	for {
		if IsActive == true {
			for item := range model.Services.IterBuffered() {
				service := item.Val
				key := item.Key

				diff := int(time.Now().Unix()) - int(service.LastCheck.Unix())

				if service.Enabled {
					if diff > service.Interval && service.Claimed == false {
						//lock current check
						service.Claim()

						//update status in map
						model.Services.Set(key, service)

						//spawn check for service
						go checker.SpawnCheck(&service)
					}
				}
			}
		}

		for i := 0; i < 1; i++ {
			time.Sleep(time.Second)
		}

	}
}
