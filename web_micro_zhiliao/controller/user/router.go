package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(router *gin.RouterGroup) {
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"mag":  "success",
		})
	}, nil)
}
