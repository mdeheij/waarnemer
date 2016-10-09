package checker

import
//"flag"
//	"fmt"
//"github.com/davecgh/go-spew/spew"
//"github.com/mdeheij/monitoring/configuration"
"testing"

//"time"

func TestSomeCommands(t *testing.T) {
	CommandExec(t, "rpe -u mdeheij -H wolkopslag.nl -e '/home/mdeheij/diskspace data'")
	CommandExec(t, "rpe -u mdeheij -H wolkopslag.nl -e '/home/mdeheij/diskspace /'")
	CommandExec(t, "curl -H http://google.nl -t 5")
}

func CommandExec(t *testing.T, cmd string) {
	CheckService(cmd)
}
