package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"one_ser.com/app/gateway/http"
	"one_ser.com/app/gateway/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors(), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		v1.POST("/user/register", http.UserRegister)
		// v1.POST("/user/login", http.UserLogin)

		v1.POST("/user/login", http.Userlogin)
		v1.GET("/user/info", http.GetUserInfo)

	}
	{
		//背包操作
		v1.POST("/backpack/add", http.BackPackAdd)
		v1.GET("/backpack/getAll", http.BackPackGetAll)
	}
	//聊天操作
	{
		v1.GET("/chat/sendClient", http.SendClientMsg)

	}
	//群组操作
	{
		v1.POST("/community/creat", http.CreatCommunity)
	}
	return r
}
