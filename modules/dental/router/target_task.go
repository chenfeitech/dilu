package router

import (
	"dilu/modules/dental/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerTargetTaskRouter)
}

// 默认需登录认证的路由
func registerTargetTaskRouter(v1 *gin.RouterGroup) {
	r := v1.Group("target-task").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiTargetTask.Get)
		r.POST("/create", apis.ApiTargetTask.Create)
		r.POST("/update", apis.ApiTargetTask.Update)
		r.POST("/page", apis.ApiTargetTask.QueryPage)
		r.POST("/del", apis.ApiTargetTask.Del)
	}
}