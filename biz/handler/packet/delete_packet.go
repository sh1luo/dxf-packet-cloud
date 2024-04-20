package packet

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"net/http"
	"packet_cloud/biz/model/hertz/packet"
	"packet_cloud/service/readwriter"
	"slices"
)

// DeletePacket .
// @router /v1/packet/delete/:id [DELETE]
func DeletePacket(ctx context.Context, c *app.RequestContext) {
	var err error
	var req packet.DeletePacketReq
	err = c.BindAndValidate(&req)
	if err != nil {
		log.Println(err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var (
		resp = new(packet.DeletePacketResp)
	)

	rw := readwriter.NewReadWriter(readwriter.LFS)
	if rw == nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	packets, err := rw.ReadPacket()
	if err != nil {
		log.Println("[DeletePacket] read file error:", err)
		c.JSON(consts.StatusInternalServerError, err)
		return
	}

	remaining := slices.DeleteFunc(packets, func(packet *packet.CloudPacket) bool {
		if packet.Id >= req.GetFrom() && packet.Id <= req.GetTo() {
			return true
		}
		return false
	})

	err = rw.SavePacket(remaining)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, err)
		return
	}

	resp.Code = 0
	resp.Msg = "删除成功"

	c.JSON(consts.StatusOK, resp)
}
