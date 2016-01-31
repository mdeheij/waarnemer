package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/services"
	"strconv"
	"strings"
)

func servicesInit(r *gin.Engine) {
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
		services.Services[identifier].Reschedule()
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
		lastCheckOld = services.Services[identifier].LastCheck.String()
		result = services.Services[identifier].Update()
		lastCheckNew = services.Services[identifier].LastCheck.String()
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
func servicesGetServicesWithIdentifier(c *gin.Context) {

	identifier := c.Param("identifier")

	if identifier != "" {
		var s []services.Service

		for _, item := range services.Services {

			item.Host = strings.Replace(item.Host, "http://", "", -1)
			item.Host = strings.Replace(item.Host, "https://", "", -1)
			tempHost := strings.Split(item.Host, "/")
			item.Host = tempHost[0]

			if strings.HasPrefix(item.Identifier, identifier) {
				s = append(s, item)
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
