package checker

import (
	"bytes"
	"os/exec"
	"strconv"
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
	log.Debug(symbol + " (" + strconv.Itoa(status) + ") - " + command + " -" + output)

	return status, output, elapsedTimeHuman
}

//Execute a command by specifying executable and arguments, returns statuscode and output summary
func Execute(timeout int, cmdName string, cmdArgs ...string) (status int, output string) {
	// TODO: Check if cmdName file exists
	cmd := exec.Command(cmdName, cmdArgs...)

	cmdOutput := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	// fail := false
	killed := false
	cmd.Stdout = cmdOutput
	cmd.Stderr = errOutput
	log.Notice(cmdName, "Voor de run")
	// var waitStatus syscall.WaitStatus

	cmd.Start()

	timer := time.AfterFunc(3*time.Second, func() {
		log.Warning("MURDER MURDER MURDER, KILL KILL KILL", cmdName)
		if cmd.Process != nil {
			err := cmd.Process.Kill()
			if err != nil {
				log.Error(cmdName, "CANNOT BE KILLED")
			}
			log.Notice("pid", cmd.Process.Pid)
			killed = true
			return
		} else {
			log.Warning("Tried to kill", cmdName, "but it didn't exists (anymore).")
		}
	})

	log.Notice(cmdName, "Na de timer")

	//
	// if err := cmd.Run(); err != nil {
	// 	log.Notice(cmdName, "Positie 1")
	// 	if err != nil {
	// 		log.Notice(cmdName, "Positie 1:faal")
	// 		fail = true
	// 	}
	// 	log.Notice(cmdName, "Positie 2")
	//
	// 	if exitError, ok := err.(*exec.ExitError); ok {
	// 		log.Notice(cmdName, "Positie 2b")
	// 		waitStatus = exitError.Sys().(syscall.WaitStatus)
	// 	}
	// } else {
	// 	log.Notice(cmdName, "Positie 2c::success")
	// 	// Success
	// 	waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	// 	log.Notice(cmdName, "Positie 3")
	// }
	log.Notice(cmdName, "Na de run")

	timer.Stop()

	if killed {
		return 5, "Timeout for check"
	}

	return 404, "End of check"
	// outputString := string(cmdOutput.Bytes())
	// shortOutputStrings := strings.Split(outputString, "\n")
	// statusCode := waitStatus.ExitStatus()
	//
	// if waitStatus.ExitStatus() == 0 && fail {
	// 	statusCode = 420
	// }
	//
	// return statusCode, shortOutputStrings[0]
}
