package model

import (
	"encoding/json"
	"io/ioutil"
	"packet_cloud/biz/model/hertz/packet"
	"sync"
)

var (
	Packets []*packet.Packet
	Mu      sync.RWMutex
)

func ReadFile() ([]*packet.Packet, error) {
	// 判断当前周是否存在文件，如果不存在则创建

	p := make([]*packet.Packet, 0)

	bs, err := ioutil.ReadFile("./packets")
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(bs, &p)
	if err != nil {
		return p, err
	}

	Mu.Lock()
	Packets = p
	Mu.Unlock()

	return Packets, nil
}

func SaveFile(p []*packet.Packet) error {
	bs, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./packets", bs, 0644)
	if err != nil {
		return err
	}

	return nil
}
