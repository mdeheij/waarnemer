package main

import (
	"api"
	"configuration"
	"flag"
	"log"
	"os"
	"services"
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

func main() {
	configuration.Init(config)
	configuration.Debug = debug

	if testConfig {
		err := services.TestConfiguration()
		if err != nil {
			log.Panic(err)
		} else {
			os.Exit(0)
		}
	}

	if noActionHandling {
		configuration.C.NoActionHandling = true
	}

	services.Init()

	if autoStart {
		services.Start()
	}

	api.Setup()
}
