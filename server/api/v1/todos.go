package v1

import "github.com/gin-gonic/gin"

type ToDo struct {
	router *gin.Engine
}

func ToDos(router *gin.Engine) {
	toDo := new(router)
	group := toDo.router.Group("/api/v1")

	group.GET("/todos", toDo.GetAll)
	group.GET("/todos/:id", toDo.GetDetails)
	group.POST("/todos", toDo.Create)
	group.PUT("/todos/:id", toDo.Update)
	group.DELETE("/todos/:id", toDo.Delete)
}

func (toDo *ToDo) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{"message": "GETed All"})
}

func (toDo *ToDo) GetDetails(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "GETed /todos/" + id})
}

func (toDo *ToDo) Create(c *gin.Context) {
	c.JSON(200, gin.H{"message": "POSTed to /todos"})
}

func (toDo *ToDo) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "PUTed to /todos/" + id})
}

func (toDo *ToDo) Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "DELETEd to /todos/" + id})
}

func new(router *gin.Engine) ToDo {
	return ToDo{
		router: router,
	}
}
