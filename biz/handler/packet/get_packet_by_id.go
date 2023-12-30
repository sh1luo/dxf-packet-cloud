package packet

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"packet_cloud/biz/handler/common"
	"packet_cloud/biz/model"
	"packet_cloud/biz/model/hertz/packet"
	"packet_cloud/util"
)

// GetPacketByID .
// @router /v1/packet/get/:id [GET]
func GetPacketByID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req packet.GetPacketByIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(packet.GetPacketByIDResp)

	writeLock.RLock()
	defer writeLock.RUnlock()

	packets, err := model.ReadPackets()
	if err != nil {
		log.Printf("[GetPacketByID] username=%s, time=%s, error=%s\n", req.Username, req.Time, err)
		common.MakeResponse(501, "获取云数据包失败", c)
		return
	}

	for _, p := range packets {
		if p.Id == req.GetId() {
			bs, err := sonic.Marshal(p)
			if err != nil {
				log.Printf("[GetPacketByID] username=%s, time=%s, id=%d, error=%s\n", req.Username, req.Time, req.GetId(), err)
				common.MakeResponse(501, "获取云数据包失败", c)
				return
			}
			encrypted, err := util.AESCBCEncrypt(bs)
			if err != nil {
				log.Printf("[GetPacketByID] username=%s, time=%s, id=%d, error=%s\n", req.Username, req.Time, req.GetId(), err)
				common.MakeResponse(501, "获取云数据包失败", c)
				return
			}

			resp.UserPackets = encrypted
			resp.Code = 200
			resp.Msg = "获取云数据包成功"
			log.Printf("[GetPacketByID] username=%s, time=%s, id=%d\n", req.Username, req.Time, req.GetId())
			c.JSON(consts.StatusOK, resp)
			return
		}
	}

	log.Printf("[GetPacketByID] packet not found, username=%s, time=%s, id=%d\n", req.Username, req.Time, req.GetId())

	c.JSON(consts.StatusOK, resp)
}
