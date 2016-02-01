package configuration

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"io/ioutil"
	"os"
)

var Config Configuration

type Configuration struct {
	Hostname                   string
	BaseFolder                 string
	ServerAddress              string
	ServerPort                 int
	AllowedUsers               []string
	SecureCookieName           string
	SecureCookie               string
	ChecksFolder               string
	ConfigFolder               string
	ConfigFile                 string
	TelegramBotToken           string `json:"TelegramBotToken"`
	TelegramNotificationTarget int32  `json:"TelegramNotificationTarget"`
	DatabaseConfig             MySQLConfig
	CookieConfig               sessions.Options
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

func Init() {
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(raw, &Config)
	if err != nil {
		panic(err.Error())
	}

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	} else {
		Config.Hostname = name
	}

}
