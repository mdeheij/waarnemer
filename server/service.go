package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/services"
	"strconv"
	"strings"
)

func servicesInit(r *gin.Engine, debug bool) {
	if debug == true {
		services.EnableDebug()
	}
	services.Start()
	services.Init()

	monitoringGroup := r.Group("/services", AuthRequired())
	{
		monitoringGroup.GET("/", servicesPage)
		monitoringGroup.GET("/start", servicesStart)
		monitoringGroup.GET("/stop", servicesStop)
		monitoringGroup.GET("/updatelist", servicesUpdateList)
		monitoringGroup.GET("/update/:identifier", servicesUpdate)
		monitoringGroup.GET("/reschedule/:identifier", servicesRescheduleCheck)
		monitoringGroup.GET("/list.json", servicesGetServicesAsJSON)
		monitoringGroup.GET("/list/:identifier", servicesGetServicesWithIdentifier)
	}
}

func servicesPage(c *gin.Context) {

	//services.Service

	c.HTML(200, "services.tmpl", gin.H{
		"title":    "Monitoring",
		"subtitle": "Running: " + strconv.FormatBool(services.DaemonActive),
		"angular":  "{{",
	})

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
		result = service.Update()
		lastCheckNew = service.LastCheck.String()
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
