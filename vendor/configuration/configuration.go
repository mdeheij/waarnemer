package configuration

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

//Config instance of Configuration struct
var C Configuration

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
	TelegramBotToken           string
	TelegramNotificationTarget string
	NoActionHandling           bool
}

//User struct used for login
type User struct {
	Username string
	Token    string
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
		tempContent, err = ioutil.ReadFile("config.yaml")
		if err == nil {
			configContent = tempContent
		} else {
			panic("Not found in folder! Panic!")
		}
	}

	m := Configuration{}
	err = yaml.Unmarshal(configContent, &m)

	errUnmarshal := yaml.Unmarshal(configContent, &C)

	if errUnmarshal != nil {
		log.Panic("Cannot load configuration! Make sure the configuration file matches your version of monitoring.")
		panic(errUnmarshal.Error())
	}

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	} else {
		C.Hostname = name
	}

	IsLoaded = true
	log.Notice("Services: ", C.ServicesFolder)
	log.Notice("Checks: ", C.ChecksFolder)
}
