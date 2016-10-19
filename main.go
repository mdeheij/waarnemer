package main

import (
	"github.com/mdeheij/monitoring/api"
	"github.com/mdeheij/monitoring/system"
)

func main() {
	system.Boot()
	api.Setup()
}
