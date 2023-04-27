package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"packet_cloud/biz/model"
)

// OnlineEdit .
// @router /edit [GET]
func OnlineEdit(ctx context.Context, c *app.RequestContext) {
	packets, err := model.ReadFile()
	if err != nil {
		log.Println("[OnlineEdit] read file error", err)
		return
	}
	c.HTML(http.StatusOK, "packet/online_edit.html", packets)
}
