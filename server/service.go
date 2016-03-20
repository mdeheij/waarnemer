package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/gin-csrf"
	"github.com/mdeheij/monitoring/services"
	"strconv"
	"strings"
)

func servicesInit(r *gin.Engine, debug bool, autostart bool) {
	if debug {
		services.EnableDebug()
	}
	if autostart {
		services.Start()
	}

	services.Init()

	frontendGroup := r.Group("/admin", AuthRequired())
	{
		frontendGroup.GET("/", guiPage)
		frontendGroup.GET("/token", func(c *gin.Context) {
			c.String(200, csrf.GetToken(c))
		})
	}
	monitoringGroup := r.Group("/api/service/", AuthRequired())
	{
		// monitoringGroup.GET("/", servicesPage)
		monitoringGroup.POST("/start", servicesStart)
		monitoringGroup.POST("/stop", servicesStop)
		monitoringGroup.POST("/updatelist", servicesUpdateList)
		monitoringGroup.POST("/update/:identifier", servicesUpdate)
		monitoringGroup.POST("/reschedule/:identifier", servicesRescheduleCheck)
		monitoringGroup.GET("/list", servicesGetServicesAsJSON)
		monitoringGroup.GET("/list/:identifier", servicesGetServicesWithIdentifier)
	}
	embedGroup := r.Group("/embed")
	{
		embedGroup.GET("/", servicesGetServicesEmbed)
	}
}

func servicesStart(c *gin.Context) {
	services.Start()

	c.JSON(200, gin.H{
		"task":   "start",
		"status": services.DaemonActive,
	})
}

func servicesUpdateList(c *gin.Context) {
	services.UpdateList()

	c.JSON(200, gin.H{
		"task": "updateList",
	})
}

func servicesRescheduleCheck(c *gin.Context) {
	identifier := c.Param("identifier")
	var result string

	if identifier != "" {
		service, _ := services.Services.Get(identifier)
		service.Reschedule()
		result = "Command sent"
	} else {
		result = "No parameter to update!"
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}
func servicesUpdate(c *gin.Context) {
	identifier := c.Param("identifier")
	var result string
	var lastCheckOld string
	var lastCheckNew string

	if identifier != "" {
		service, _ := services.Services.Get(identifier)
		lastCheckOld = service.LastCheck.String()
		//lastCheckOld = services.Services[identifier].LastCheck.String()
		lastCheckNew = service.LastCheck.String()
		result = service.Update()
	} else {
		result = "No parameter to update!"
	}

	c.JSON(200, gin.H{
		"result":       result,
		"lastCheckOld": lastCheckOld,
		"lastCheckNew": lastCheckNew,
	})
}

func servicesStop(c *gin.Context) {
	services.Stop()
	c.JSON(200, gin.H{
		"task":   "stop",
		"status": services.DaemonActive,
	})

}

func servicesGetServicesAsJSON(c *gin.Context) {
	result, _ := json.Marshal(services.Services)
	c.String(200, string(result))
}

//servicesGetServicesWithIdentifier returns the services which match the given identifier as a json
func servicesGetServicesWithIdentifier(c *gin.Context) {
	identifier := c.Param("identifier")

	if identifier != "" {
		var s []services.Service

		for item := range services.Services.Iter() {
			service := item.Val

			service.Host = strings.Replace(service.Host, "http://", "", -1)
			service.Host = strings.Replace(service.Host, "https://", "", -1)
			tempHost := strings.Split(service.Host, "/")
			service.Host = tempHost[0]

			if strings.HasPrefix(service.Identifier, identifier) {
				s = append(s, service)
			}
		}

		result, _ := json.Marshal(s)

		if len(s) > 0 {
			c.String(200, string(result))
			return
		}
	}

	c.JSON(200, gin.H{})
}

func getProblematicServices() []services.Service {
	var s []services.Service

	for item := range services.Services.IterBuffered() {
		service := item.Val

		if service.Health > 0 {
			s = append(s, service)
		}
	}

	return s
}

//servicesGetServicesEmbed returns the services as a nice embed for Grafana
func servicesGetServicesEmbed(c *gin.Context) {
	//hostname := c.Param("hostname")

	c.HTML(200, "embed.html", gin.H{
		"title":    "Monitoring",
		"services": getProblematicServices(),
		"subtitle": "Running: " + strconv.FormatBool(services.DaemonActive),
		"angular":  "{{",
	})
	//TODO: think how I'm going to build this
	//	c.JSON(200, gin.H{})
}
