package api

import "github.com/gin-gonic/gin"

type CRUD interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type User interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}
