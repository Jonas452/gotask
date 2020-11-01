package adapter

import (
	"encoding/json"
	"gotask/internal/app/adapter/repository"
	"gotask/internal/app/application/usecase"
	"gotask/internal/app/domain"
	"io/ioutil"

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
	c.JSON(200, task)
}

func (ctrl Controller) getAllTasks(c *gin.Context) {
	tasks := usecase.GetAllTasks(taskRepository)
	c.JSON(200, tasks)
}

func (ctrl Controller) createTask(c *gin.Context) {

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err)
		return
	}

	var jTask domain.Task
	err = json.Unmarshal(jsonData, &jTask)
	if err != nil {
		c.JSON(400, err)
		return
	}

	task, err := usecase.CreateTask(taskRepository, jTask)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, task)
}

func (ctrl Controller) updateTask(c *gin.Context) {
	id := c.Param("id")
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(400, err)
		return
	}

	var jTask domain.Task
	err = json.Unmarshal(jsonData, &jTask)
	if err != nil {
		c.JSON(400, err)
		return
	}

	task, err := usecase.UpdateTask(taskRepository, id, jTask)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, task)
}

func (ctrl Controller) deleteTask(c *gin.Context) {
	id := c.Param("id")

	err := usecase.DeleteTask(taskRepository, id)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, true)
}
