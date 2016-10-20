package checker

import (
	"strconv"
	"strings"
	"time"

	"github.com/mdeheij/monitoring/services/action"
	"github.com/mdeheij/monitoring/services/model"
	"github.com/mdeheij/monitoring/services/model/health"
)

func SpawnCheck(service *model.Service) int {
	args := service.Command
	args = strings.Replace(args, "$HOST$", service.Host, -1)
	args = strings.Replace(args, "$TIMEOUT$", strconv.Itoa(service.Timeout), -1)

	status, output, rtime := CheckService(service.Timeout, args)
	service.Output = output

	if status > 0 { //It's going down
		oldHealth := service.Health
		service.ThresholdCounter++

		if oldHealth == -1 { //cold check, now its down
			service.Health = health.WARNING //set warning state
		}

		if oldHealth == 0 {
			service.Health = health.WARNING //(re)set warning state
		}

		if oldHealth == health.WARNING && service.ThresholdCounter >= service.Threshold {
			service.Health = health.CRITICAL //Service is down
			action.Handle(*service)          //Ready for action
		}
	} else {
		oldHealth := service.Health
		service.Health = health.OK
		service.ThresholdCounter = 0
		if oldHealth == health.CRITICAL {
			action.Handle(*service)
		}
	}

	service.Release()
	service.RTime = rtime
	service.LastCheck = time.Now()
	model.Services.Set(service.Identifier, *service)

	return status
}
