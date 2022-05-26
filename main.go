package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sideneckx/template-go-server/api_route/user_route"
)

func main() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Token")
	router.Use(cors.New(corsConfig))

	//middleware
	// router.Use(api_middleware.AuthorizeMiddleWare())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "template server"})
	})

	user_route.NewRoute(*router)
	server := "0.0.0.0:" + "8080"
	router.Run(server)
}
