package adapter

import (
	"gotask/internal/app/adapter/repository"
	"gotask/internal/app/application/usecase"
	"gotask/internal/app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is a controller
type Controller struct{}

var (
	taskRepository = repository.Task{}
)

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	r.GET("/tasks/:id", ctrl.getTask)
	r.GET("/tasks", ctrl.getAllTasks)
	r.POST("/tasks", ctrl.createTask)
	r.PUT("/tasks/:id", ctrl.updateTask)
	r.DELETE("/tasks/:id", ctrl.deleteTask)

	return r
}

func (ctrl Controller) getTask(c *gin.Context) {
	id := c.Param("id")
	task := usecase.GetTask(taskRepository, id)
	c.JSON(http.StatusOK, task)
}

func (ctrl Controller) getAllTasks(c *gin.Context) {
	tasks := usecase.GetAllTasks(taskRepository)
	c.JSON(http.StatusOK, tasks)
}

func (ctrl Controller) createTask(c *gin.Context) {
	var jTask domain.Task
	if err := c.BindJSON(&jTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := usecase.CreateTask(taskRepository, jTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ctrl Controller) updateTask(c *gin.Context) {
	id := c.Param("id")
	var jTask domain.Task
	if err := c.BindJSON(&jTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jTask.ID = id
	task, err := usecase.UpdateTask(taskRepository, jTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ctrl Controller) deleteTask(c *gin.Context) {
	id := c.Param("id")

	err := usecase.DeleteTask(taskRepository, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, true)
}
