package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/system/status"
)

//Homepage gives a simple splash page
func Homepage(c *gin.Context) {
	c.HTML(200, "routes", gin.H{
		"routes": status.Get().Api.Routes,
	})
}
