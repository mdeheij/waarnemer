package api

import (
	"configuration"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

//AuthRequired is authentication middleware for user authenticaton.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		if err == nil {
			c.Header("Content-Type", "application/json; charset=utf8") //TODO: This should be a seperate middleware
			c.Next()
		} else {
			c.String(401, "Unauthorized.")
			c.Abort()
		}
	}
}

func main() {
	Setup()
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
		// monitoringGroup.GET("/", servicesPage)
		servicesGroup.POST("/start", servicesStart)
		servicesGroup.POST("/stop", servicesStop)
		servicesGroup.POST("/updatelist", servicesUpdateList)
		servicesGroup.POST("/update/:identifier", servicesUpdate)
		servicesGroup.POST("/reschedule/:identifier", servicesRescheduleCheck)
		servicesGroup.OPTIONS("/list", servicesGetServicesAsJSON)
		servicesGroup.GET("/list", servicesGetServicesAsJSON)
		servicesGroup.GET("/list/:identifier", servicesGetServicesWithIdentifier)
	}

	publicGroup := r.Group("/public")
	{
		publicGroup.GET("/status/:group", servicesGetPublicServices)
	}
	log.Notice("Starting webserver")

	log.Info("API listening on http://" + configuration.C.Api.Address)
	r.Run(configuration.C.Api.Address)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
