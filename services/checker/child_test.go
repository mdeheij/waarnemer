package checker

import "testing"

func TestSomeCommands(t *testing.T) {
	CommandExec(t, "rpe -u root -H server.local -e '/etc/monitoring/diskspace data'")
	CommandExec(t, "rpe -u root -H server.local -e '/etc/monitoring/diskspace /'")
	CommandExec(t, "curl -H http://server.local -t 5")
}

func CommandExec(t *testing.T, cmd string) {
	CheckService(5, cmd)
}
