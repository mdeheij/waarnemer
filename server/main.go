package server

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/configuration"
	"strconv"
)

//DebugMode sets verbose output
var DebugMode bool

//Setup initializes routers, login and services.
func Setup(debug bool, autostart bool) {
	DebugMode = debug

	if !DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.LoadHTMLGlob(configuration.Config.ResourceFolder + "templates/*")
	r.Static("/assets", configuration.Config.ResourceFolder+"assets")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/admin/")
	})

	loginInit(r)
	servicesInit(r, debug, autostart)

	// If configuration values are empty, Gin will listen and serve on 0.0.0.0:8080.
	r.Run(fmt.Sprintf("%s:%s", configuration.Config.ServerAddress, strconv.Itoa(configuration.Config.ServerPort)))
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

func guiPage(c *gin.Context) {
	c.HTML(200, "services.tmpl", gin.H{
		"title":   "Monitoring",
		"angular": "{{",
	})
}
