package router

import (
	"api-gatware/handler/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("shopList", api.ShopList)
	return r
}
