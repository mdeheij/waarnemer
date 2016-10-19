package api

import (
	"time"

	gin "github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	log "github.com/mdeheij/logwrap"
	configuration "github.com/mdeheij/monitoring/configuration"
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
	// If configuration values are empty, Gin will listen and serve on 0.0.0.0:8080.

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	servicesGroup := r.Group("/api/service/", AuthRequired())
	{
		servicesGroup.GET("/start", servicesStart)
		servicesGroup.GET("/stop", servicesStop)
		servicesGroup.GET("/updatelist", servicesUpdateList)
		servicesGroup.GET("/update/:identifier", servicesUpdate)
		servicesGroup.GET("/reschedule/:identifier", servicesRescheduleCheck)
		servicesGroup.GET("/list", servicesGetServicesAsJSON)
		servicesGroup.GET("/list/:identifier", servicesGetServicesWithIdentifier)
		servicesGroup.OPTIONS("/list", servicesGetServicesAsJSON) //TODO: is this used, and if so, refactor
	}

	publicGroup := r.Group("/public")
	{
		publicGroup.GET("/status/:group", servicesGetPublicServices)
	}
	log.Notice("Starting webserver")

	log.Info("API listening on http://" + configuration.C.API.Address)
	r.Run(configuration.C.API.Address)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
