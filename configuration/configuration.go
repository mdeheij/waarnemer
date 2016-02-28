package configuration

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"io/ioutil"
	"os"
)

var Config Configuration

type Configuration struct {
	Hostname         string
	BaseFolder       string
	ResourceFolder   string
	ServerAddress    string
	ServerPort       int
	Users            []User
	SecureCookieName string
	SecureCookie     string
	// ChecksFolder               string
	// ConfigFolder               string
	// ConfigFile                 string
	TelegramBotToken           string `json:"TelegramBotToken"`
	TelegramNotificationTarget int32  `json:"TelegramNotificationTarget"`
	DatabaseConfig             MySQLConfig
	CookieConfig               sessions.Options
}

type User struct {
	Username string
	Hash     string
}

type MySQLConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	Engine   string
	Encoding string
}

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

	// if configfile != "/etc/monitoring/config.json" {
	// 	configContent, err = ioutil.ReadFile(configfile)
	// } else {
	// 	configContent, err = ioutil.ReadFile("config.json")
	// }
	//check if the chosen method of reading config file worked
	if err != nil {
		panic(err.Error())
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
