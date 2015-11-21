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
var daemonActive = false

type Service struct {
	Id               int    `db:"serviceID"`
	Health           int    `db:"health"`
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
	fmt.Println(StatusColor("●", service.Health))
	fmt.Println("┃ Host:   \t" + service.Host)
	fmt.Println("┃ Label:  \t" + service.Label)
	fmt.Println("┃ Command: \t" + service.Executable)
	fmt.Println("┃ Last check: \t" + strconv.Itoa(int(service.LastCheck.Unix())))
	fmt.Printf("┃ Timeout: \t%v\n┃ Interval:\t%v\n┃ Threshold:\t%v/%v", service.Timeout, service.Interval, service.ThresholdCounter, service.Threshold)
	fmt.Println("")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┉┉┉┉┉┉┈┈┈ ")
}

func StatusColor(text string, health int) string {

	switch health {
	case 0:
		return "\x1b[32;1m" + text + "\x1b[0m"
	case 1:
		return "\x1b[33;1m" + text + "\x1b[0m"
	case 2:
		return "\x1b[31;1m" + text + "\x1b[0m"
	default:
		return "\x1b[33;1m NEEDEFAULT" + text + "\x1b[0m"
	}
}

func (service Service) spawnChild() {

	executable := ChecksFolder + service.Executable
	args := strings.Split(service.Arguments, ",")
	status, output := CheckService(executable, args)
	if status > 0 { //It's going down

		oldHealth := service.Health
		service.Health = 2
		service.Comment = "Output: " + output
		service.ThresholdCounter++

		if oldHealth == 0 {
			service.Health = 1 //(re)set warning state
		}

		if oldHealth == 1 && service.ThresholdCounter >= service.Threshold {
			service.Health = 2      //It's officially down!
			a := NewAction(service) //Ready for action
			a.errorMsg = output
			a.Run()
		}

		//go service.handleDown --> func (service Service) handleDown
	} else {
		oldHealth := service.Health
		service.Health = 0
		service.ThresholdCounter = 0
		if oldHealth == 2 {
			a := NewAction(service) //Ready for recovery notify
			a.errorMsg = output
			a.Run()
		}
		// dit moet wel ff in SQL komen wellicht?
		// hmm misschien ook niet.
		// even over slapen.
		// 1 vote voor nee
	}

	service.Lock = false
	service.LastCheck = time.Now()
	service.print()
	Services[service.Id] = service
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func slicePoll() {

	for {
		if daemonActive == true {
			for key, service := range Services {
				diff := int(time.Now().Unix()) - int(service.LastCheck.Unix())
				//editableService := service
				//fmt.Print("ID:", key, "  ")
				//	fmt.Println(diff)
				if service.Enabled {
					if diff > service.Interval {
						//service.print()
						//lock current check
						service.Lock = true
						service.Comment = "Locked!"
						Services[key] = service
						go service.spawnChild()

						//commit to map

					} else {
						/*	nextCheck := service.Interval - diff
							fmt.Print(service.Label + " will check in: ")
							fmt.Println(nextCheck)*/
					}
				}
				//update the struct in map m with new data
			}
			for i := 0; i < 1; i++ {
				//	fmt.Println("● ")
				time.Sleep(1 * time.Second)
				//fmt.Println("Next run in " + strconv.Itoa(3-i) + "..")
			}
		}
	}
}

func Init() {
	fmt.Println("Started.")
	reloadServices()
	go slicePoll()
}
func Start() {
	fmt.Println("Starting..")
	daemonActive = true
}
func Stop() {
	fmt.Println("Stopping..")
	daemonActive = false
}

func reloadServices() {
	fmt.Println("Connecting db.")
	db = sqlx.MustConnect("mysql", "serverstat@tcp(localhost:3306)/serverstat")

	rows, err := db.Queryx("SELECT serviceID, enabled, host, label, timeout, executable, arguments, intervaltime, warningthreshold FROM service")
	checkError(err)
	enabledServicesCounter := 0

	for rows.Next() {
		var service Service
		err = rows.StructScan(&service)
		checkError(err)
		service.Health = -1 //you know nothing, monitoring
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
}
