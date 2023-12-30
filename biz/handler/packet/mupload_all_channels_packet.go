package packet

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"packet_cloud/biz/model"
	"packet_cloud/biz/model/hertz/packet"
)

// MUploadAllChannelsPacket .
// @router /v1/packet/mupload [POST]
func MUploadAllChannelsPacket(ctx context.Context, c *app.RequestContext) {
	var err error
	var req packet.MUploadAllChannelsPacketReq
	err = c.BindAndValidate(&req)
	if err != nil || req.McloudPacket == nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(packet.MUploadAllChannelsPacketResp)

	writeLock.Lock()
	defer writeLock.Unlock()

	packets, err := model.ReadPackets()
	if err != nil {
		log.Println("[UploadPacket] read packets error", err)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	m := map[int]string{
		0: "跨1",
		1: "跨2",
		2: "跨3A",
		3: "跨3B",
		4: "跨4",
		5: "跨5",
		6: "跨6",
		7: "跨7",
		8: "跨8",
	}

	for idx, channel := range req.McloudPacket.Channel {
		inserted := &packet.CloudPacket{
			Id:          0,
			Region:      m[idx],
			Name:        req.McloudPacket.Name,
			Channel:     channel,
			Uploader:    req.McloudPacket.Uploader,
			Time:        req.McloudPacket.Time,
			UserPackets: req.McloudPacket.UserPackets,
		}

		if len(packets) > 0 {
			inserted.Id = packets[len(packets)-1].Id + 1
		} else {
			inserted.Id = 1
		}
		packets = append(packets, inserted)
	}

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
