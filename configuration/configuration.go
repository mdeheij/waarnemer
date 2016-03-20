package configuration

import (
	"encoding/json"
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
	TelegramNotificationTarget int32  `json:"TelegramNotificationTarget"`
	CookieConfig               sessions.Options
}

//User struct used for login
type User struct {
	Username string
	Hash     string
}

//Init initializes the configuration.
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
		panic(err.Error())
	}

	name, err := os.Hostname()

	if err != nil {
		panic(err)
	} else {
		Config.Hostname = name
	}
}
