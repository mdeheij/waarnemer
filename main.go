package main

import (
	"flag"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/server"
)

var debug bool
var autostart bool
var config string

func init() {

	flag.BoolVar(&debug, "debug", false, "Enable debugging")
	flag.BoolVar(&autostart, "autostart", true, "Start checking automatically")
	flag.StringVar(&config, "config", "/etc/monitoring/config.json", "Path to config file")
	flag.Parse()
}

func main() {
	configuration.Init(config)
	server.Setup(debug, autostart)
}
