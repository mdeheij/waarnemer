package service

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/services"
	"github.com/mdeheij/monitoring/services/model/health"
)

func RescheduleCheck(c *gin.Context) {
	identifier := c.Param("identifier")
	var result string

	if identifier != "" {
		service, _ := services.GetService(identifier)
		service.Reschedule()
		result = "Command sent"
	} else {
		result = "No identifier specified."
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func Show(c *gin.Context) {
	identifier := c.Param("identifier")
	var service interface{}

	if identifier != "" {
		service, _ = services.GetService(identifier)
	} else {
		c.String(404, "Service not found.")
	}

	c.JSON(200, service)
}

func Update(c *gin.Context) {
	identifier := c.Param("identifier")
	var result string
	if identifier != "" {
		service, _ := services.GetService(identifier)

		err := services.UpdateService(service)
		if err != nil {
			result = "Could not update service."
			log.Error(result, identifier)
		}

	} else {
		result = "No parameter to update!"
	}
	c.JSON(200, gin.H{
		"result": result,
	})
}

func List(c *gin.Context) {
	result, _ := json.Marshal(services.GetAllServices())
	c.String(200, string(result))
}

func GetCritical() []interface{} { //TODO: find different name for this func
	var s []interface{}

	for item := range services.GetAllServices().IterBuffered() {
		service := item.Val

		if service.Health > health.OK {
			s = append(s, service)
		}
	}
	return s
}
