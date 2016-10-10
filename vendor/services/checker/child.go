package checker

import (
	"bytes"
	"configuration"
	"log"
	"message"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/mgutz/str"
)

//CheckService accepts a command and returns a status code, output (stdout) and the reaction time
func CheckService(command string) (status int, output string, rtime int64) {
	// Initialization.
	now := time.Now()

	//Convert executable to Go's os/exec.Command
	commandSlice := str.ToArgv(configuration.C.ChecksFolder + "/" + command)
	status, output = Execute(commandSlice[0], commandSlice[1:len(commandSlice)]...)

	//TODO: implement error logging
	elapsedTime := time.Since(now)
	elapsedTimeHuman := elapsedTime.Nanoseconds() / 1000000

	symbol := message.StatusColor("●", status)
	log.Debug(symbol + " (" + strconv.Itoa(status) + ") - " + command + " -" + output)

	return status, output, elapsedTimeHuman
}

//Execute a command by specifying executable and arguments, returns statuscode and output summary
func Execute(cmdName string, cmdArgs ...string) (status int, output string) {
	// TODO: Check if cmdName file exists
	cmd := exec.Command(cmdName, cmdArgs...)

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
