package server

import (
	"encoding/base64"
	"fmt"
	//"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/configuration"
	"strconv"
)

var DebugMode bool

//var updates = make(map[string]statistics.Update)

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

	r := gin.Default()
	r.LoadHTMLGlob(configuration.Config.ResourceFolder + "templates/*")
	r.Static("/assets", configuration.Config.ResourceFolder+"assets")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/services/")
	})

	loginInit(r)
	servicesInit(r, debug)

	bindTarget := configuration.Config.ServerAddress + ":" + strconv.Itoa(configuration.Config.ServerPort)
	//fmt.Println("http://" + bindTarget)

	r.Run(bindTarget) // listen and serve on 0.0.0.0:8080
}
