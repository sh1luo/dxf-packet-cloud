package packet

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"packet_cloud/biz/model"
	"packet_cloud/biz/model/hertz/packet"
)

// UploadPacket .
// @router /v1/packet/upload [POST]
func UploadPacket(ctx context.Context, c *app.RequestContext) {
	var err error
	var req packet.UploadPacketReq
	err = c.BindAndValidate(&req)
	if err != nil || req.CloudPacket == nil || req.CloudPacket.Name == "" || req.CloudPacket.UserPackets == nil ||
		req.CloudPacket.Region == "" || req.CloudPacket.Channel == "" || req.CloudPacket.Uploader == "" || req.CloudPacket.Time == "" {
		c.String(consts.StatusBadRequest, "invalid params")
		return
	}

	resp := new(packet.UploadPacketResp)

	inserted := &packet.CloudPacket{
		Id:          0,
		Region:      req.CloudPacket.Region,
		Name:        req.CloudPacket.Name,
		Channel:     req.CloudPacket.Channel,
		Uploader:    req.CloudPacket.Uploader,
		Time:        req.CloudPacket.Time,
		UserPackets: req.CloudPacket.UserPackets,
	}

	writeLock.Lock()
	defer writeLock.Unlock()

	packets, err := model.ReadPackets()
	if err != nil {
		log.Println("[UploadPacket] read packets error", err)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	if len(packets) > 0 {
		inserted.Id = packets[len(packets)-1].Id + 1
	} else {
		inserted.Id = 1
	}
	packets = append(packets, inserted)

	err = model.SavePackets(packets)

	if err != nil {
		resp.Code = 501
		resp.Msg = "上传失败"
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.Code = 200
	resp.Msg = "上传成功"

	c.JSON(consts.StatusOK, resp)
}
