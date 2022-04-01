package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router = HandleRoutes()

func HandleRoutes() *gin.Engine {

	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
		return
	})

	return r
}
