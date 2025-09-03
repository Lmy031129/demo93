package api

import (
	"api-gatware/basic/globar"
	__ "api-gatware/basic/proto"
	"api-gatware/handler/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopList(c *gin.Context) {
	var req request.ShopReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	list, err := globar.Client.ShopList(c, &__.ShopListReq{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查询成功",
		"data":    list,
	})
}
