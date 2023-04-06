package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter(context *gin.Engine) {
	// 定义一个路由组
	helloRouter := context.Group("/hello")
	// 往该路由组里添加路由
	helloRouter.GET("/user", test)
}

func test(context *gin.Context) {

	data :=
		context.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "没问题",
			"data":    "假装有数据",
		})
}
