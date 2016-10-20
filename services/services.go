package services

import (
	"errors"

	log "github.com/mdeheij/logwrap"

	"github.com/mdeheij/monitoring/services/loader"
	"github.com/mdeheij/monitoring/services/model"
)

var daemonActive = false

func Load() {
	for _, service := range loader.FindServices() {
		service.Health = -1 //you know nothing, monitoring
		model.Services.Set(service.Identifier, service)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//UpdateService updates only the given service from actual configuration
func UpdateService(old model.Service) error {
	//Do not start a new check while updating by claiming it.
	old.Claim()
	for _, new := range loader.FindServices() {
		if new.Identifier == old.Identifier {
			new.CopyAttributes(&old)
			//push new service to Services map
			model.Services.Set(old.Identifier, new)

			log.Warning("Reloaded", new.Identifier, "from", old.Identifier)
			return nil
		}
	}

	return errors.New("Service not found.")
}

func GetService(identifier string) (model.Service, bool) {
	return model.Services.Get(identifier) //TODO: add error handling
}

func GetAllServices() model.ConcurrentMap {
	return model.Services
}
