package main

import (
	"fmt"
	"github.com/mdeheij/monitoring"
	"os"
	/*	"log"

		"path/filepath"*/)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 0 {
		if argsWithoutProg[0] == "demo" {
			fmt.Println("demo mode")
			monitoring.Init()
			monitoring.Start()

			var input string
			fmt.Scanln(&input)
			fmt.Println("stopping")

			monitoring.Stop()

			fmt.Scanln(&input)
			fmt.Println("starting")

			monitoring.Start()

			fmt.Scanln(&input)
			fmt.Println("stopping")

			monitoring.Stop()

		}
	} else {

		fmt.Println("Starting standalone")
		monitoring.Init()
		monitoring.Start()

		var input string
		fmt.Scanln(&input)
		fmt.Println("done")
	}

}
