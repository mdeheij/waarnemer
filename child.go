package main

import "os"
import "os/exec"
import "syscall"
import "fmt"
import "bytes"
import "time"
import "strconv"
import "strings"

func main2() {
	cmd := exec.Command("go", "version")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))

}

func DebugExec(executable string, args []string) (status int) {
	now := time.Now()
	time := strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute())

	status, output := execute(executable, args)

	var tekentje string
	if status > 0 {
		tekentje = StatusColor("●", false)
	} else {
		tekentje = StatusColor("●", true)
	}

	fmt.Println(tekentje + " [" + time + "] (" + strconv.Itoa(status) + ") - " + output)
	return status
}

func execute(cmdName string, cmdArgs []string) (status int, output string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	//	cmd := exec.Command("cat", "8.8.1.6", "-c 1") //complete bullshit for exit code simulation
	cmdOutput := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = errOutput

	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			//fmt.Printf("OutputHENK: %s\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		// Success
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		//fmt.Printf("OutputPIET: %s\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}

	//fmt.Println(string(cmdOutput.Bytes()))
	//errorMsg := string(errOutput.Bytes())
	//fmt.Println(waitStatus.ExitStatus())
	//fmt.Println("Error: " + errorMsg)

	outputString := string(cmdOutput.Bytes())
	shortOutputStrings := strings.Split(outputString, "\n")

	return waitStatus.ExitStatus(), shortOutputStrings[0]
}
