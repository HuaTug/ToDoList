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

// ListTasks @Tags TASK

func ListTasks(c *gin.Context) {
	var listService service.ListTasksService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// ShowTask @Tags TASK

func ShowTask(c *gin.Context) {
	var ShowTask service.ShowTaskService
	if err := c.ShouldBind(&ShowTask); err == nil {
		res := ShowTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}

}

// DeleteTask @Tags TASK

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}

}

// UpdateTask @Tags TASK

func UpdateTask(c *gin.Context) {
	var UpDateTaskService service.UpdateTaskService
	if err := c.ShouldBind(&UpDateTaskService); err == nil {
		res := UpDateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)

	}
}

// SearchTasks @Tags TASK

func SearchTasks(c *gin.Context) {
	var searchTaskService service.SearchTaskService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
