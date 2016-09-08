package routes

import (
	"github.com/borentaylor05/streamline/server/api/v1"
	"github.com/gin-gonic/gin"
)

func AttachRoutes(router *gin.Engine) {
	v1.ToDos(router)
	v1.Retailers(router)
}
