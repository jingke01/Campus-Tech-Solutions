package controller

import (
	"github.com/gin-gonic/gin"
	"temp/global"
	"temp/model"
)

func Register(c *gin.Context) {

	var req model.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {

	}
	global.DB.where("")
}
