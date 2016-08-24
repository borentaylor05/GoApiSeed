package routes

import "github.com/gin-gonic/gin"

func AttachRoutes(router *gin.Engine) {
	ToDos(router)
}
