package api

import (
	"ToDoList/pkg/utils"
	"ToDoList/service"
	"github.com/gin-gonic/gin"
	_ "github.com/sirupsen/logrus"
)

// CreateTask @Tags TASK

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

/*
// ListTasks @Tags TASK

func ListTasks(c *gin.Context) {
	listService := service.ListTasksService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// ShowTask @Tags TASK

func ShowTask(c *gin.Context) {
	showTaskService := service.ShowTaskService{}
	res := showTaskService.Show(c.Param("id"))
	c.JSON(200, res)
}

// DeleteTask @Tags TASK

func DeleteTask(c *gin.Context) {
	deleteTaskService := service.DeleteTaskService{}
	res := deleteTaskService.Delete(c.Param("id"))
	c.JSON(200, res)
}

// UpdateTask @Tags TASK

func UpdateTask(c *gin.Context) {
	updateTaskService := service.UpdateTaskService{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)

	}
}

// SearchTasks @Tags TASK

func SearchTasks(c *gin.Context) {
	searchTaskService := service.SearchTaskService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
*/
