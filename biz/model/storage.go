package model

import (
	"github.com/bytedance/sonic"
	"os"
	"packet_cloud/biz/model/hertz/packet"
)

// ReadPackets 从文件中读取云包
func ReadPackets() ([]*packet.CloudPacket, error) {
	p := make([]*packet.CloudPacket, 0)

	//t := time.Now()
	bs, err := os.ReadFile("./packets")
	//log.Println("[ReadPackets] read file cost:", time.Since(t).Seconds())
	if err != nil {
		return p, err
	}
	//t = time.Now()
	err = sonic.Unmarshal(bs, &p)
	//log.Println("[ReadPackets] sonic.Unmarshal all packets to memory cost:", time.Since(t).Seconds())
	if err != nil {
		return p, err
	}

	return p, nil
}

func SavePackets(p []*packet.CloudPacket) error {
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
