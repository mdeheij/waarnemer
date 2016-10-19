package system

import (
	"errors"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/services/loader"
	"github.com/mdeheij/monitoring/services/model"
)

//TestConfiguration checks if configuration can be loaded and shows amount of services
func TestConfiguration() error {
	services := loader.FindServices()
	length := len(services)
	log.Info("Length:", length, "services")
	if length > 0 {
		return nil
	} else {
		return errors.New("Amount of services is invalid.")
	}
}

//Reload compares current map with memory and updates added, changed or removed services
func Reload() {
	//Do not start a new check while updating
	jsonServices := loader.FindServices() //[]Service
	jsonServicesMap := model.NewCMap()    //Concurrent string-Service map

	for _, newService := range jsonServices {
		//oldService := Services[newService.Identifier]
		oldService, _ := model.Services.Get(newService.Identifier)

		if oldService.Identifier != newService.Identifier {
			newService.Health = 2 //TODO: const enum
		} else {
			newService.Health = oldService.Health
		}

		newService.Claimed = oldService.Claimed
		newService.LastCheck = oldService.LastCheck
		newService.ThresholdCounter = oldService.ThresholdCounter
		newService.Output = oldService.Output
		newService.RTime = oldService.RTime

		jsonServicesMap.Set(newService.Identifier, newService)
		log.Warning("Reloaded " + oldService.Identifier + " as " + newService.Identifier)
	}

	model.Services = jsonServicesMap
}
