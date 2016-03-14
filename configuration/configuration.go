package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"io/ioutil"
	"os"
)

//Config instance of Configuration struct
var Config Configuration

//Configuration struct
type Configuration struct {
	Hostname                   string
	BaseFolder                 string
	ResourceFolder             string
	ServerAddress              string
	ServerPort                 int
	Users                      []User
	SecureCookieName           string
	SecureCookie               string
	TelegramBotToken           string `json:"TelegramBotToken"`
	TelegramNotificationTarget string `json:"TelegramNotificationTarget"`
	CookieConfig               sessions.Options
	// Public                     []PublicGroup
	//TODO: coming soon in a new release
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
	}
	tempContent, err = ioutil.ReadFile("config.json")
	if err == nil {
		configContent = tempContent
	}
	tempContent, err = ioutil.ReadFile("/etc/monitoring/config.json")
	if err == nil {
		configContent = tempContent
	}

	errUnmarshal := json.Unmarshal(configContent, &Config)
	if errUnmarshal != nil {
		fmt.Println("Cannot load configuration! Make sure the configuration file matches your version of monitoring.")
		panic(err.Error())
	}

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	} else {
		Config.Hostname = name
	}

}
