package api

import (
	"time"

	gin "github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/api/pages"
	"github.com/mdeheij/monitoring/api/services"
	"github.com/mdeheij/monitoring/api/system"
	configuration "github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/system/status"
)

//AuthRequired is authentication middleware for user authenticaton.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, value := range configuration.C.UserTokens {
			if ("Bearer " + value) == c.Request.Header.Get("Authorization") {
				c.Header("Content-Type", "application/json; charset=utf8") //TODO: This should be a seperate middleware
				c.Next()
				return
			}
		}

		log.Warning(c.ClientIP(), "No authorization header match.")
		c.String(401, "Unauthorized.")
		c.Abort()
	}
}

//Setup initializes routers, login and services.
func Setup() {

	if configuration.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	r.SetHTMLTemplate(pages.Template)
	r.GET("/", pages.Routes)

	serviceGroup := r.Group("/api/", AuthRequired())
	{
		serviceGroup.GET("/services/:identifier/update", services.Update)
		serviceGroup.GET("/services/:identifier/reschedule", services.Reschedule)
		serviceGroup.GET("/services/:identifier", services.Show)
		serviceGroup.GET("/services", services.List)
	}

	systemAPI := r.Group("/api/system/", AuthRequired())
	{
		systemAPI.GET("/start", system.Start)
		systemAPI.GET("/stop", system.Stop)
		systemAPI.GET("/reload", system.Reload)
	}

	log.Notice("Starting webserver..")
	log.Info("API listening on http://" + configuration.C.API.Address)

	status.Api.Routes = r.Routes()

	r.Run(configuration.C.API.Address)
}
