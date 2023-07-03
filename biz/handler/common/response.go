package common

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func MakeResponse(code int32, message string, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"code": code,
		"msg":  message,
	})
}
