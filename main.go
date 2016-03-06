package main

import (
	"flag"
	"fmt"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/server"
	"github.com/mdeheij/monitoring/services"
	"os"
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
		services.TestConfiguration()
		fmt.Println("TODO: implement this")
		os.Exit(2)
	}

	configuration.Init(config)
	server.Setup(debug, true)
}
