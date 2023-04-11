package busRouter

import (
	"github.com/gin-gonic/gin"
	"web_micro_zhiliao/controller/product"
	"web_micro_zhiliao/controller/seckill"
	"web_micro_zhiliao/controller/user"
)

func InitRouter(router *gin.Engine) {
	userGroup := router.Group("/user")
	productGroup := router.Group("/product")
	seckillGroup := router.Group("/seckill")

	user.Router(userGroup)
	product.Router(productGroup)
	seckill.Router(seckillGroup)

}
