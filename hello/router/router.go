package router

import "github.com/gin-gonic/gin"

// SetRouterGroup 加载各模块的路由
func SetRouterGroup(router *gin.Engine) {
	helloRouter := router.Group("/hello")
	helloRouter.GET("/get", 监听地址)
}
func 监听地址(context *gin.Context) {

}
