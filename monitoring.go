package monitoring

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"time"
)

var garbage string

var db *sqlx.DB
var Services = make(map[int]Service)

var ChecksFolder string = "/home/mdeheij/projects/src/github.com/mdeheij/monitoring/checks/"

type Service struct {
	Id               int    `db:"serviceID"`
	Ok               bool   `db:"ok"`
	Enabled          bool   `db:"enabled"`
	Host             string `db:"host"`
	Label            string `db:"label"`
	Executable       string `db:"executable"`
	Arguments        string `db:"arguments"`
	Acknowledged     string `db:"ack"`
	Timeout          int    `db:"timeout"`
	Interval         int    `db:"intervaltime"`
	Threshold        int    `db:"warningthreshold"`
	ThresholdCounter int
	LastCheck        time.Time
	Lock             bool
	Comment          string
	DisplayHealth    string
	DisplayTools     string
	DisplayTimeDiff  int
}

func (service Service) enable() {
	//enable
	service.Enabled = true
	Services[service.Id] = service
	//commit naar db
}

func (service Service) disable() {
	//enable
	service.Enabled = false
	Services[service.Id] = service
	//commit naar db
}

func (service Service) print() {
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┉┉┉┉┉┉┈┈┈ ")
	fmt.Print("┃ ID:     \t")
	fmt.Println(service.Id)
	fmt.Print("┃ OK:     \t")
	fmt.Println(StatusColor("●", service.Ok))
	fmt.Println("┃ Host:   \t" + service.Host)
	fmt.Println("┃ Label:  \t" + service.Label)
	fmt.Println("┃ Command: \t" + service.Executable)
	fmt.Println("┃ Last check: \t" + strconv.Itoa(int(service.LastCheck.Unix())))
	fmt.Printf("┃ Timeout: \t%v\n┃ Interval:\t%v\n┃ Threshold:\t%v/%v", service.Timeout, service.Interval, service.ThresholdCounter, service.Threshold)
	fmt.Println("")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┉┉┉┉┉┉┈┈┈ ")
}

func StatusColor(text string, positive bool) string {
	if positive {
		return "\x1b[32;1m" + text + "\x1b[0m"
	} else {
		return "\x1b[31;1m" + text + "\x1b[0m"
	}
}

func (service Service) spawnChild() {

	executable := ChecksFolder + service.Executable
	args := strings.Split(service.Arguments, ",")
	status := DebugExec(executable, args)
	if status > 0 {
		service.Ok = false
		//go service.handleDown --> func (service Service) handleDown
	} else {
		service.Ok = true
		//dit moet wel ff in SQL komen wellicht? maar niet fatal.
		// hmm misschien ook niet.
		// even over slapen.
	}
	/*
		timeNextCheck := int(service.LastCheck.Unix()) + service.Interval
		if timeNextCheck >= int(time.Now().Unix()) {
			fmt.Println("Checking")
		} else {
			//fmt.Println(timeNextCheck - int(time.Now().Unix()))
		}*/
	service.LastCheck = time.Now()
	Services[service.Id] = service
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func slicePoll() {
	for {

		for key, service := range Services {
			diff := int(time.Now().Unix()) - int(service.LastCheck.Unix())
			//editableService := service
			//fmt.Print("ID:", key, "  ")
			//	fmt.Println(diff)
			if service.Enabled {
				if diff > service.Interval {
					service.print()
					fmt.Println("--- Checking")
					//lock current check
					service.Lock = true
					service.Comment = "Locked!"
					Services[key] = service
					go service.spawnChild()

					//commit to map

				} else {
					//nextCheck := service.Interval - diff
					//	fmt.Print(service.Label + " will check in: ")
					//	fmt.Println(nextCheck)
				}
			}
			//update the struct in map m with new data
		}
		for i := 0; i < 1; i++ {
			//fmt.Println("● ")
			time.Sleep(1 * time.Second)
			//fmt.Println("Next run in " + strconv.Itoa(3-i) + "..")
		}
		fmt.Println("")

	}
}

func Init() {
	fmt.Println("Started.")
	reloadServices()
	go slicePoll()
}

func reloadServices() {
	fmt.Println("Connecting db.")
	db = sqlx.MustConnect("mysql", "serverstat:fawkejrutwt573458239gWRRFHWG@tcp(shared.mike.solutions:3306)/serverstat")

	rows, err := db.Queryx("SELECT serviceID, ok, enabled, host, label, timeout, executable, arguments, intervaltime, warningthreshold FROM service")
	checkError(err)
	enabledServicesCounter := 0

	for rows.Next() {
		var service Service
		err = rows.StructScan(&service)
		checkError(err)

		Services[service.Id] = service
		enabledServicesCounter++
	}
	fmt.Print(enabledServicesCounter)
	fmt.Println(" services loaded.")
}

func nutteloos() {
	garbage = strconv.Itoa(6)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
	/*	go f("goroutine")
		f("direct")*/

	/*	for rows.Next() {
		var service Service
		err = rows.StructScan(&service)
		checkError(err)
		go childSpawn(service)
	}*/
}

/*func geenmain() {
	garbage = strconv.Itoa(6)
	// Must.... functions will panic on fail
	db = sqlx.MustConnect("mysql", "serverstat@tcp(localhost:3306)/serverstat")

	rows, err := db.Queryx("SELECT serviceID, ok, host, label, timeout, executable, arguments, lastcheck, intervaltime, warning_threshold, warning_threshold_counter FROM service ")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("id: %s, %s", action.Name, action.Command.String)

	for rows.Next() {
		var service Service
		err = rows.StructScan(&service)
		checkError(err)
		service.print()
	}

}*/

/*func main()
	db, err := sqlx.Connect("mysql", "user=serverstat dbname=serverstat sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	//args := []string{""}
	args := []string{"-c 2", "-i", "0.2", "-w 2", "bartbart333.tk"}
	debugExec("ping", args)
	//args = []string{"aasasfasf;;;nu.nl"}
	//debugExec("cat", args)
}*/

/*type Service struct {
	serviceID                 string  `db:"serviceID"`
	parentID                  string  `db:"parentID"`
	ok                        bool    `db:"ok"`
	host                      string  `db:"host"`
	label                     string  `db:"label"`
	command                   string  `db:"command"`sd
	timeout                   int     `db:"timeout"`
	intervaltime              int     `db:"intervaltime"`
	warning_threshold         int     `db:"warning_threshold"`
	warning_threshold_counter int     `db:"warning_threshold_counter"`
	enabled                   bool    `db:"enabled"`
	errorMsg                  string  `db:"error"`
	ack                       bool    `db:"ack"`
	output                    string  `db:"output"`
	rtime                     float64 `db:"rtime"`
	lastcheck                 string  `db:"lastcheck"`
	action                    int     `db:"action"`
}

func livePoll() {
	for {

		rows, err := db.Queryx("SELECT serviceID, ok, host, label, timeout, executable, arguments, lastcheck, intervaltime, warning_threshold, warning_threshold_counter FROM service WHERE enabled = true AND serviceID  = 278 ")
		checkError(err)

		for rows.Next() {
			var service Service
			err = rows.StructScan(&service)
			checkError(err)
			go childSpawn(service)
		}

		for i := 0; i < 4; i++ {
			time.Sleep(1 * time.Second)
			//fmt.Println("Next run in " + strconv.Itoa(3-i) + "..")
		}
		fmt.Println("")

	}
}
*/
