package services

//import "os"
import (
	"bytes"
	"github.com/mdeheij/monitoring/configuration"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func CheckService(command string) (status int, output string, rtime int64) {
	now := time.Now()
	timeStampHuman := now.Format(time.Stamp)
	//status, outputMsg := execute(executable, args)
	path := "checks/"
	testArgs := make([]string, 1)
	testArgs[0] = configuration.Config.BaseFolder + path + command

	wrapperLocation := configuration.Config.BaseFolder + path + "wrapper"
	DebugMessage("WrapperLocation: " + wrapperLocation)
	status, outputMsg := execute(wrapperLocation, testArgs)

	elapsedTime := time.Since(now)
	elapsedTimeHuman := elapsedTime.Nanoseconds() / 1000000

	var symbol string
	if status > 0 {
		symbol = StatusColor("●", 2)
	} else {
		symbol = StatusColor("●", 0)
	}

	DebugMessage(symbol + " [" + timeStampHuman + "] {" + elapsedTime.String() + "} (" + strconv.Itoa(status) + ") - " + command + " -" + outputMsg)
	return status, outputMsg, elapsedTimeHuman
}

func execute(cmdName string, cmdArgs []string) (status int, output string) {
	// TODO: Check if cmdName file exists, returning some high error if not.
	cmd := exec.Command(cmdName, cmdArgs...)
	//	cmd := exec.Command("cat", "8.8.1.6", "-c 1") //complete bullshit for exit code simulation
	cmdOutput := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	fail := false
	cmd.Stdout = cmdOutput
	cmd.Stderr = errOutput

	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if err != nil {
			fail = true
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
		}
	} else {
		// Success
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	}

	outputString := string(cmdOutput.Bytes())
	shortOutputStrings := strings.Split(outputString, "\n")
	statusCode := waitStatus.ExitStatus()

	if waitStatus.ExitStatus() == 0 && fail {
		statusCode = 420
	}

	return statusCode, shortOutputStrings[0]
}
