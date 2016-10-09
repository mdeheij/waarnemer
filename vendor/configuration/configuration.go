package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/contrib/sessions"
)

//Config instance of Configuration struct
var Config Configuration

//IsLoaded returns true when the configuration is loaded
var IsLoaded bool

//Debug mode activated
var Debug bool

//Configuration struct
type Configuration struct {
	Hostname                   string
	ChecksFolder               string
	ServicesFolder             string
	ServerAddress              string
	ServerPort                 int
	Users                      []User
	SecureCookieName           string
	SecureCookie               string
	TelegramBotToken           string `json:"TelegramBotToken"`
	TelegramNotificationTarget string `json:"TelegramNotificationTarget"`
	CookieConfig               sessions.Options
	Public                     []PublicGroup
	NoActionHandling           bool
}

//User struct used for login
type User struct {
	Username string
	Hash     string
}

//PublicGroup struct used to define public available groups
type PublicGroup struct {
	Name     string
	Services []string
}

//Init ializes the configuration
func Init(configfile string) {
	var configContent []byte
	var tempContent []byte
	var err error

	tempContent, err = ioutil.ReadFile(configfile)
	if err == nil {
		configContent = tempContent
	} else {
		tempContent, err = ioutil.ReadFile("config.json")
		if err == nil {
			configContent = tempContent
		} else {
			panic("Not found in folder! Panic!")
		}
	}

	errUnmarshal := json.Unmarshal(configContent, &Config)
	if errUnmarshal != nil {
		fmt.Println("Cannot load configuration! Make sure the configuration file matches your version of monitoring.")
		panic(errUnmarshal.Error())
	}

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	} else {
		Config.Hostname = name
	}

	IsLoaded = true
	log.Notice("Services: ", Config.ServicesFolder)
	log.Notice("Checks: ", Config.ChecksFolder)
}
