package server

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	//"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/statistics"
	"gopkg.in/gorp.v1"
	"strconv"
)

var dbmap *gorp.DbMap
var DebugMode bool

//var updates = make(map[string]statistics.Update)

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", configuration.Config.DatabaseConfig.Username+":"+configuration.Config.DatabaseConfig.Password+"@tcp("+configuration.Config.DatabaseConfig.Host+":"+configuration.Config.DatabaseConfig.Port+")/"+configuration.Config.DatabaseConfig.Database+"?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err, "sql.Open failed")
	}
	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{configuration.Config.DatabaseConfig.Engine, configuration.Config.DatabaseConfig.Encoding}}

	dbmap.AddTableWithName(statistics.Update{}, "serverupdate").SetKeys(false, "id")
	return dbmap
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	return string(data)
}

// func getJsonVar(js *simplejson.Json, str string) string {
// 	result, err := js.Get(str).String()
// 	check(err)
// 	return result
// }

func Setup(debug bool) {

	if debug == true {
		DebugMode = true
	} else {
		DebugMode = false
		gin.SetMode(gin.ReleaseMode)
	}

	dbmap = initDb()

	r := gin.Default()
	r.LoadHTMLGlob(configuration.Config.ResourceFolder + "templates/*")
	r.Static("/assets", configuration.Config.ResourceFolder+"assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", nil)
	})

	loginInit(r)
	servicesInit(r, debug)
	statsInit(r)

	bindTarget := configuration.Config.ServerAddress + ":" + strconv.Itoa(configuration.Config.ServerPort)
	//fmt.Println("http://" + bindTarget)

	r.Run(bindTarget) // listen and serve on 0.0.0.0:8080
}
