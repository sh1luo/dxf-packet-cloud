package packet

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
	"net/http"
	"packet_cloud/service/readwriter"
)

// OnlineEdit .
// @router /edit [GET]
func OnlineEdit(ctx context.Context, c *app.RequestContext) {

	rw := readwriter.NewReadWriter(readwriter.LFS)
	if rw == nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	packets, err := rw.ReadPacket()
	if err != nil {
		log.Println("[OnlineEdit] read file error", err)
		return
	}

	for _, packet := range packets {
		for _, userPacket := range packet.UserPackets {
			userPacket.Content = "内容暂时不展示"
		}
	}

	c.HTML(http.StatusOK, "packet/online_edit.html", utils.H{"packets": packets})
}
