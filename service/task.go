package service

import (
	"ToDoList/model"
	"ToDoList/serializer"
	"time"
)

// 展示任务详情的服务

type ShowTaskService struct {
} //应为这个服务是GET请求，所以是空的

// 删除任务的服务

type DeleteTaskService struct {
}

// 更新任务的服务

type UpdateTaskService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

// 创建任务的服务

type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

// 搜索任务的服务

type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

type ListTasksService struct {
	PageNum  int `json:"page_num" from:"page_num"`
	PageSize int `json:"page_size" from:"page_size"`
}

//新增一条备忘录

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	code := 200
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0, //表示为未完成
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		//Data:   serializer.BuildTasks(task),
		Msg: "创建成功",
	}
}

//列表返回用户所有备忘录

func (service *ListTasksService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if service.PageSize == 0 { //分页的判定操作
		service.PageSize = 15
	}
	//涉及到多表查询
	model.DB.Model(model.Task{}).Preload("User").Where("uid = ?", id).Count(&total).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).
		Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

//展示一条备忘录

func (service *ShowTaskService) Show(id string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, id).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
			//Error:  err.Error(),
		}
	}
	//task.AddView() // 增加点击数
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		//Msg:    e.GetMsg(code),
	}
}

//删除备忘录

func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

//更新备忘录操作

func (service *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,

		Data: serializer.BuildTask(task),
		Msg:  "更新完成",
	}
}

// 查询备忘录操作

func (service *SearchTaskService) Search(uId uint) serializer.Response {
	var tasks []model.Task
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uId).
		Where("title LIKE ? OR content LIKE ?",
				"%"+service.Info+"%", "%"+service.Info+"%").
		Find(&tasks) //模糊查询的方法
	return serializer.Response{
		Status: 200,
		Msg:    " ",
		Data:   serializer.BuildTasks(tasks),
	}
}
