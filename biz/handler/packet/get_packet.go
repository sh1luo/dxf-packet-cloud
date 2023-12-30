package packet

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"packet_cloud/biz/handler/common"
	"packet_cloud/biz/model"
	"packet_cloud/biz/model/hertz/packet"
)

// GetPacket .
// @router /v1/packet/get [GET]
func GetPacket(ctx context.Context, c *app.RequestContext) {
	var err error
	var req packet.GetPacketReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(packet.GetPacketResp)

	writeLock.RLock()
	defer writeLock.RUnlock()

	packets, err := model.ReadPackets()
	if err != nil {
		log.Printf("[GetPacket] username=%s, time=%s, error=%s\n", req.Username, req.Time, err)
		common.MakeResponse(501, "获取云数据包失败", c)
		return
	}

	for i := 0; i < len(packets); i++ {
		packets[i].UserPackets = make([]*packet.UserPacket, 0)
	}

	resp.CloudPackets = packets
	resp.Code = 200
	resp.Msg = "获取云数据包成功"

	log.Printf("[GetPacket] username=%s, time=%s\n", req.Username, req.Time)

	c.JSON(consts.StatusOK, resp)
}
