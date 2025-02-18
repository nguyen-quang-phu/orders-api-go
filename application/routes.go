package application

import "github.com/gin-gonic/gin"

func loadRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello World"})
	})

	return router
}
