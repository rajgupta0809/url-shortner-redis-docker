package main

import (
	"fmt"
	"log"

	"github.com/raj-gupta/url-shortner-redis-docker/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/:url", routes.ResolveUrl)
	router.POST("/api/:url", routes.ShortenUrl)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("err", err)
	}

	app := gin.Default()

	setupRoutes(app)

	log.Fatal(app.Run("3000"))
}
