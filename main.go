package main

import (
	"flag"
	//"fmt"
	"os"

	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/server"
	"github.com/mdeheij/monitoring/services"
)

var debug bool
var testconfig bool
var config string

func init() {

	flag.BoolVar(&debug, "debug", false, "Enable debugging")
	flag.BoolVar(&testconfig, "test", false, "Test configuration instead of running")
	flag.StringVar(&config, "config", "/etc/monitoring/config.json", "Path to config file")
	flag.Parse()
}

func main() {

	if testconfig {
		configuration.Init(config)
		err := services.TestConfiguration()
		if err != nil {
			panic(err)
		} else {
			os.Exit(0)
		}
	}

	configuration.Init(config)
	server.Setup(debug, true)
}
