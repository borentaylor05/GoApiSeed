package routes

import "github.com/gin-gonic/gin"

func CreateRouter() (router *gin.Engine) {
	router = gin.Default()
	AttachRoutes(router)
	return
}
