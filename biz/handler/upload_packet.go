// Code generated by hertztool.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"packet_cloud/service/readwriter"

	"github.com/cloudwego/hertz/pkg/app"
	packet "packet_cloud/biz/model/hertz/packet"
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

	var insertedID int32 = 0
	packets, err := readwriter.ReadPacket(readwriter.LFS)
	if err != nil {
		log.Println("[UploadPacket] read packets error", err)
		c.JSON(consts.StatusInternalServerError, err)
		return
	}

	if len(packets) > 0 {
		insertedID = packets[len(packets)-1].Id + 1
	} else {
		insertedID = 1
	}

	inserted := &packet.CloudPacket{
		Id:          insertedID,
		Region:      req.CloudPacket.Region,
		Name:        req.CloudPacket.Name,
		Channel:     req.CloudPacket.Channel,
		Uploader:    req.CloudPacket.Uploader,
		Time:        req.CloudPacket.Time,
		UserPackets: req.CloudPacket.UserPackets,
	}
	packets = append(packets, inserted)

	err = readwriter.SavePacket(packets, readwriter.LFS)
	if err != nil {
		log.Println("[UploadPacket] save packets error", err)
		c.JSON(consts.StatusInternalServerError, err)
		return
	}

	c.JSON(consts.StatusOK, &packet.UploadPacketResp{
		Code: 0,
		Msg:  "上传成功",
	})
}
