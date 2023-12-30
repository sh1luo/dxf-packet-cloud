package packet

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"packet_cloud/biz/model"
	"packet_cloud/biz/model/hertz/packet"
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

	//log.Println("Delete, id=", req.GetId())

	resp := new(packet.DeletePacketResp)

	writeLock.Lock()
	defer writeLock.Unlock()

	deletedID := make([]int32, 0)

	packets, err := model.ReadPackets()
	if err != nil {
		log.Println("[DeletePacket] read file error:", err)
		resp.Code = 501
		resp.Msg = "删除失败，读数据包失败"
		c.JSON(consts.StatusOK, resp)
		return
	}

	notDeletedPackets := make([]*packet.CloudPacket, 0)
	for idx, p := range packets {
		if p.Id >= req.GetFrom() && p.Id <= req.GetTo() {
			deletedID = append(deletedID, p.Id)
		} else {
			notDeletedPackets = append(notDeletedPackets, packets[idx])
		}
	}

	if len(deletedID) > 0 {
		resp.Code = 200
		resp.Msg = fmt.Sprintf("删除成功: %v", deletedID)
		c.JSON(consts.StatusOK, resp)

		err = model.SavePackets(notDeletedPackets)
		if err != nil {
			resp.Code = 501
			resp.Msg = "删除后保存失败"
			c.JSON(consts.StatusOK, resp)
		}

		return
	}

	resp.Code = -1
	resp.Msg = fmt.Sprintf("未找到删除数据，from = %d, to = %d", req.GetFrom(), req.GetTo())

	c.JSON(consts.StatusOK, resp)
}
