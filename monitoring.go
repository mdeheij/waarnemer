package monitoring

import (
	"database/sql"
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

type Service struct {
	Id               int    `db:"serviceID"`
	Ok               bool   `db:"ok"`
	Host             string `db:"host"`
	Label            string `db:"label"`
	Executable       string `db:"executable"`
	Arguments        string `db:"arguments"`
	Timeout          int    `db:"timeout"`
	Interval         int    `db:"intervaltime"`
	Threshold        int    `db:"warningthreshold"`
	ThresholdCounter int
	LastCheck        time.Time
	Lock             bool
	Comment          string
	Test             sql.NullBool
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

func childSpawn(service Service) {

	executable := "checks/" + service.Executable
	args := strings.Split(service.Arguments, ",")
	DebugExec(executable, args)
	if int(service.LastCheck.Unix())+service.Interval >= int(time.Now().Unix()) {
		fmt.Println("Tijd om te checken!")
	}
	service.LastCheck = time.Now()
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
			fmt.Print("\nID:", key, "  ")
			//	fmt.Println(diff)

			if diff > service.Interval {
				fmt.Println("---Tijd om te checken!")
				service.Comment = "YES HET WERKT!!111"
				service.LastCheck = time.Now()
			}

			Services[key] = service //update the struct in map m with new data
		}
		for i := 0; i < 5; i++ {
			fmt.Print("● ")
			time.Sleep(1 * time.Second)
			//fmt.Println("Next run in " + strconv.Itoa(3-i) + "..")
		}
		fmt.Println("")

		Services[278].print()

	}
}

func Init() {
	fmt.Println("Started.")
	reloadServices()

	//Services[278].print()

	go slicePoll()

}

func reloadServices() {
	fmt.Println("Connecting db.")
	db = sqlx.MustConnect("mysql", "serverstat@tcp(localhost:3306)/serverstat")

	rows, err := db.Queryx("SELECT serviceID, ok, host, label, timeout, executable, arguments, intervaltime, warningthreshold FROM service WHERE enabled = true LIMIT 3")
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
