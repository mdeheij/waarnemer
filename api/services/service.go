package services

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/services"
)

//Reschedule a service to be checked as soon as possible
func Reschedule(c *gin.Context) {
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

//Show returns information of specified service
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

//Update reloads specified service from configuration
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

		if service.Health > 0 {
			s = append(s, service)
		}
	}
	return s
}
