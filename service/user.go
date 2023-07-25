package service

import (
	"ToDoList/model"
	"ToDoList/pkg/utils"
	"ToDoList/serializer"
	"github.com/jinzhu/gorm"
)

// UserService 用户注册服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" `
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" `
}

func (service *UserService) Register() *serializer.Response {
	//code := e.SUCCESS
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	// 表单验证
	if count == 1 {
		//code = e.ErrorExistUser
		return &serializer.Response{
			/*			Status: code,
						Msg:    e.GetMsg(code),*/
			Status: 400,
			Msg:    "已经有人了，无需再注册",
		}
	}
	user.UserName = service.UserName
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		//util.LogrusObj.Info(err)
		//code = e.ErrorFailEncryption
		return &serializer.Response{
			/*			Status: code,
						Msg:    e.GetMsg(code),*/
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		//util.LogrusObj.Info(err)
		//code = e.ErrorDatabase
		return &serializer.Response{
			/*			Status: code,
						Msg:    e.GetMsg(code),*/
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return &serializer.Response{
		/*		Status: code,
				Msg:    e.GetMsg(code),*/
		Status: 200,
		Msg:    "用户注册成功",
	}
}

// Login 用户登陆函数
func (service *UserService) Login() serializer.Response {
	var user model.User
	//先去找一下这个user，看看数据库中有没有这个人
	//code := e.SUCCESS
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		// 如果查询不到，返回相应的错误
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			//util.LogrusObj.Info(err)
			//code = e.ErrorNotExistUser
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先登录",
			}
		}
		//util.LogrusObj.Info(err)
		//如果不是用户不存在，是其他不可抗拒的因素导致的错误
		//code = e.ErrorDatabase
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	//对其密码进行验证
	if !user.CheckPassword(service.Password) {
		//code = e.ErrorNotCompare
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//发一个token，为了其他功能需要身份验证所给前端存储的
	//创建一个备忘录，这个功能就要token，不然都不知道是谁创建的备忘录
	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		//util.LogrusObj.Info(err)
		//code = e.ErrorAuthToken
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
