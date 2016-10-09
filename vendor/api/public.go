package api

import (
	"services"

	"github.com/gin-gonic/gin"
)

//PublicService is a stripped down struct which can be used for public status pages
type PublicService struct {
	Identifier  string
	Description string
	Online      bool
}

func getPublicServices(group string) []PublicService {

	// var allowedIdentifierList = configuration.GetPublicGroup(group) //TODO: Regex or array?!

	var publicServices []PublicService

	// for item := range services.Services.IterBuffered() {
	// 	originalService := item.Val
	// 	log.Warning("Processing Public..", allowedIdentifierList)
	//
	// 	if utils.StringInSlice(originalService.Identifier, allowedIdentifierList) {
	// 		publicService := PublicService{Identifier: originalService.Identifier, Description: originalService.Description}
	//
	// 		if originalService.Health < 2 {
	// 			publicService.Online = true
	// 		}
	//
	// 		publicServices = append(publicServices, publicService)
	// 	}
	// }
	return publicServices
}

//servicesGetServicesEmbed returns the services as a nice embed for Grafana
func servicesGetPublicServices(c *gin.Context) {
	group := c.Param("group")

	//test, _ := services.Services.Get("company.dns.ns1")

	c.JSON(200, gin.H{
		"daemon":   services.DaemonActive,
		"services": getPublicServices(group),
	})
}
