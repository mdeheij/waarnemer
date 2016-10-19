package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/services"
)

//PublicService is a stripped down struct which can be used for public status pages
type PublicService struct {
	Identifier  string
	Description string
	Online      bool
}

//servicesGetServicesEmbed returns the services as a nice embed for Grafana
func servicesGetPublicServices(c *gin.Context) {
	// group := c.Param("group")

	//test, _ := services.Services.Get("company.dns.ns1")

	c.JSON(200, gin.H{
		"daemon": services.DaemonActive,
		// "services": getPublicServices(group),
	})
}
