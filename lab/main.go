package main

import (
	"fmt"
	"github.com/otium/queue"
	"os/exec"
)

var temp error
var newgetal int

func main() {
	q := queue.NewQueue(func(val interface{}) {
		cmd := exec.Command("ping", "8.8.8.8")
		err := cmd.Start()
		temp = err
		fmt.Println(cmd.Output())

	}, 20)
	for i := 0; i < 200; i++ {
		q.Push(i)
	}
	q.Wait()
}
