package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	router *gin.Engine
}

func ToDos(router *gin.Engine) {
	toDo := new(router)
	router.GET("/todos", toDo.rootHandler)
	router.GET("/todos/:id", toDo.handler)
}

func (toDo *ToDo) rootHandler(c *gin.Context) {
	c.String(http.StatusOK, "Got to root!")
}

func (toDo *ToDo) handler(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "Got to todo: "+id)
}

func new(router *gin.Engine) ToDo {
	return ToDo{
		router: router,
	}
}
