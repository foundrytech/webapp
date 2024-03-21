package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jicodes/webapp/handlers"
	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/middlewares"
)

func init () {
	initializers.LoadVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CheckRequestMethod())
	r.Use(middlewares.CheckPayload())

	r.GET("/healthz", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("X-Content-Type-Options", "nosniff")

		if err := initializers.DB.Exec("SELECT 1").Error; err == nil {
			c.Status(http.StatusOK)
			log.Println("Health check passed: 200 OK, DB is up and running")
		} else {
			c.Status(http.StatusServiceUnavailable)
			log.Fatal("Health check failed: 503 Service Unavailable, DB is down or not reachable")
		}
	})
	
	r.POST("/v1/user", handlers.CreateUser)
	r.GET("/v1/user/self", middlewares.BasicAuth(), handlers.GetUser)
	r.PUT("/v1/user/self", middlewares.BasicAuth(), handlers.UpdateUser)

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}