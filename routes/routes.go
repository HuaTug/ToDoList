package routes

import (
	"ToDoList/api"
	"ToDoList/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 路由配置

func NewRouter() *gin.Engine {
	r := gin.Default() //生成了一个WSGI应用程序实例
	store := cookie.NewStore([]byte("something-very-secret"))
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
		}
	}
	return r
}

/*r.Use(middleware.Cors())
v1 := r.Group("api/v1")
{
	v1.GET("ping", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	// 用户操作
	v1.POST("user/register", api.UserRegister)
	v1.POST("user/login", api.UserLogin)
	authed := v1.Group("/") //需要登陆保护
	authed.Use(middleware.JWT())
	{
		//任务操作
		authed.GET("tasks", api.ListTasks)
		authed.POST("task", api.CreateTask)
		authed.GET("task/:id", api.ShowTask)
		authed.DELETE("task/:id", api.DeleteTask)
		authed.PUT("task/:id", api.UpdateTask)
		authed.POST("search", api.SearchTasks)
	}
}
return r*/
