package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

var garbage string

type Service struct {
	Id               int    `db:"serviceID"`
	Ok               bool   `db:"ok"`
	Host             string `db:"host"`
	Label            string `db:"label"`
	Executable       string `db:"executable"`
	Timeout          int    `db:"timeout"`
	Interval         int    `db:"intervaltime"`
	Threshold        int    `db:"warning_threshold"`
	ThresholdCounter int    `db:"warning_threshold_counter"`
	LastCheck        string `db:"lastcheck"`
	Test             sql.NullBool
	StatusIcon       string
}

func StatusColor(text string, positive bool) string {
	if positive {
		return "\x1b[32;1m" + text + "\x1b[0m"
	} else {
		return "\x1b[31;1m" + text + "\x1b[0m"
	}
}

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func childSpawn(executable string, arguments string) {
	args := strings.Split(arguments, ",")

	DebugExec(executable, args)
}

func (service Service) print() {
	fmt.Print("ID:     \t")
	fmt.Println(service.Id)

	fmt.Print("OK:     \t")
	fmt.Println(StatusColor("●", service.Ok))
	fmt.Println("Host:   \t" + service.Host)
	fmt.Println("Label:  \t" + service.Label)
	fmt.Println("Command: \t" + service.Executable)
	fmt.Println("Last check: \t" + service.LastCheck)
	fmt.Printf("Timeout: \t%v\nInterval:\t%v\nThreshold:\t%v/%v", service.Timeout, service.Interval, service.ThresholdCounter, service.Threshold)
	fmt.Println("********")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("********")
	fmt.Println("")
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Started.")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nrc.nl,-c 1,-i 0.2")
	go childSpawn("timeout", "5,ping,nu.nl -c 1")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("curl", "undone.nl,-I")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("ping", "nu.nl,-c 1,-i 0.2")
	go childSpawn("cat", "/etc/hosts")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
	/*	go f("goroutine")
		f("direct")*/
}

func geenmain() {
	garbage = strconv.Itoa(6)
	// Must.... functions will panic on fail
	db := sqlx.MustConnect("mysql", "serverstat@tcp(localhost:3306)/serverstat")
	//var service Service
	// We'll get most recent item and map it into our struct
	//err := db.Get(&service, "SELECT host, label FROM service ORDER BY serviceID DESC LIMIT 1")

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

}

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
*/
