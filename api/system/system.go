package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/system"
	"github.com/mdeheij/monitoring/system/daemon"
	"github.com/mdeheij/monitoring/system/status"
)

func Start(c *gin.Context) {
	daemon.Start()
	c.JSON(200, gin.H{
		"task":   "start",
		"status": status.Get(),
	})
}

func Stop(c *gin.Context) {
	daemon.Stop()
	c.JSON(200, gin.H{
		"task":   "stop",
		"status": status.Get(),
	})
}

func Reload(c *gin.Context) {
	system.Reload()
	c.JSON(200, gin.H{
		"task": "reload",
	})
}
