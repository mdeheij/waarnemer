package status

import "github.com/mdeheij/monitoring/system/daemon"

type ApiStatus struct {
	Routes interface{}
}

var Api ApiStatus

type Status struct {
	DaemonActive bool
	Api          ApiStatus
}

//GetStatus returns the current status of system
func Get() (s Status) {
	return Status{
		DaemonActive: daemon.IsActive,
		Api:          Api,
	}
}
