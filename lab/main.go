package main

import (
	"fmt"
	"github.com/mdeheij/monitoring"
)

func main() {
	fmt.Println("Starting Lab")
	monitoring.Init()

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
