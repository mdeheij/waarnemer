package checker

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/message"
	"github.com/mgutz/str"
)

//CheckService accepts a command and returns a status code, output (stdout) and the reaction time
func CheckService(timeout int, command string) (status int, output string, rtime int64) {
	// Initialization.
	now := time.Now()

	//Convert executable to Go's os/exec.Command
	commandSlice := str.ToArgv(configuration.C.Paths.Checks + "/" + command)
	status, output = Execute(timeout, commandSlice[0], commandSlice[1:len(commandSlice)]...)

	//TODO: implement error logging
	elapsedTime := time.Since(now)
	elapsedTimeHuman := elapsedTime.Nanoseconds() / 1000000

	symbol := message.StatusColor("●", status)
	log.Debug(symbol + " (" + strconv.Itoa(status) + ") - " + command + " - " + output)

	return status, output, elapsedTimeHuman
}

//Execute a command by specifying executable and arguments, returns statuscode and output summary
func Execute(timeout int, cmdName string, cmdArgs ...string) (status int, output string) {
	// TODO: Check if cmdName file exists
	cmd := exec.Command(cmdName, cmdArgs...)

	cmdOutput := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}

	fail := false
	killed := false

	cmd.Stdout = cmdOutput
	cmd.Stderr = errOutput

	var waitStatus syscall.WaitStatus

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Start()

	timer := time.AfterFunc(time.Duration(timeout+1)*time.Second, func() {
		if cmd.Process != nil {
			pgid, err := syscall.Getpgid(cmd.Process.Pid)
			if err == nil {
				syscall.Kill(-pgid, 15) // note the minus sign
				log.Notice("Killed PID", pgid, "because of timeout of ", timeout, "seconds while running", cmdName)
				killed = true
			}
		} else {
			log.Error("Tried to kill but it didn't exist (anymore).", cmdName, cmdArgs)
		}
	})

	if err := cmd.Wait(); err != nil {
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

	timer.Stop()

	if killed {
		return 124, "CRITICAL - Script timeout while running check."
	}

	outputString := string(cmdOutput.Bytes())
	shortOutputStrings := strings.Split(outputString, "\n")
	statusCode := waitStatus.ExitStatus()

	if waitStatus.ExitStatus() == 0 && fail {
		statusCode = 420
	}

	return statusCode, shortOutputStrings[0]
}
