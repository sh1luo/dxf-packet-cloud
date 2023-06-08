package model

import (
	"github.com/bytedance/sonic"
	"os"
	"packet_cloud/biz/model/hertz/packet"
)

// ReadPackets 从文件中读取云包
func ReadPackets() ([]*packet.Packet, error) {
	p := make([]*packet.Packet, 0)

	bs, err := os.ReadFile("./packets")
	if err != nil {
		return p, err
	}
	err = sonic.Unmarshal(bs, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func SaveFile(p []*packet.Packet) error {
	bs, err := sonic.Marshal(p)
	if err != nil {
		return err
	}
	err = os.WriteFile("./packets", bs, 0644)
	if err != nil {
		return err
	}

	return nil
}
