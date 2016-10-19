package system

import (
	"flag"
	"os"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/system/daemon"
)

var debug bool
var testConfig bool
var apiEnabled bool
var autoStart bool
var noActionHandling bool
var config string

func init() {
	log.Notice("Initializing..")
	flag.BoolVar(&debug, "debug", false, "Enable debugging")
	flag.BoolVar(&testConfig, "test", false, "Test configuration instead of running")
	flag.BoolVar(&apiEnabled, "api", true, "Enable Web API")
	flag.BoolVar(&autoStart, "autostart", true, "Autostart service checking")
	flag.BoolVar(&noActionHandling, "no-action-handling", false, "Do not do anything when services are down")
	flag.StringVar(&config, "config", "/etc/monitoring/config.yaml", "Path to config file")
	flag.Parse()
}

//Boot handles flags and initializes configuration
func Boot() {
	configuration.Init(config)
	configuration.Debug = debug
	configuration.AutoStart = autoStart

	if testConfig {
		err := TestConfiguration()
		if err != nil {
			log.Panic(err)
		} else {
			os.Exit(0)
		}
	}

	if noActionHandling {
		configuration.C.NoActionHandling = true
	}

	if configuration.AutoStart {
		daemon.Start()
	}

	daemon.Spawn()
}
