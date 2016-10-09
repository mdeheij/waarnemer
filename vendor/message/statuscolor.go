package message

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
