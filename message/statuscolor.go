package message

import (
	"fmt"
	"strconv"

	"github.com/mdeheij/monitoring/services/model/health"
)

//StatusColor generates a command line colour based on health
func StatusColor(text string, health int) string {

	switch health {
	case 0:
		return "\x1b[32;1m" + text + "\x1b[0m"
	case 1:
		return "\x1b[33;1m" + text + "\x1b[0m"
	case 2:
		return "\x1b[31;1m" + text + "\x1b[0m"
	case 3:
		return "\x1b[35;1m" + text + "\x1b[0m"
	default:
		return "\x1b[31;1m" + text + "\x1b[0m"
	}
}

func BuildNotificationMessage(identifier string, currentHealth int, host string, thresholdCounter int, threshold int, output string) (msg string) {
	thresholdCounting := strconv.Itoa(thresholdCounter) + "/" + strconv.Itoa(threshold)
	actionTypeString := ""

	switch currentHealth {
	case health.CRITICAL:
		actionTypeString = "❌"
	case health.OK:
		actionTypeString = "✅"
	case health.WARNING:
		actionTypeString = "⚠️"
	}

	return fmt.Sprintf("%s %s (%s) %s\n %s", actionTypeString, identifier, host, thresholdCounting, output)
}
