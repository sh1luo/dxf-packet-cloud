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

func ReadFile() error {
	// 判断当前周是否存在文件，如果不存在则创建
	if Packets == nil {
		bs, err := ioutil.ReadFile("./packets")
		if err != nil {
			return err
		}
		err = json.Unmarshal(bs, &Packets)
		if err != nil {
			return err
		}
	}

	return nil
}

func SaveFile() error {
	bs, err := json.Marshal(Packets)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./packets", bs, 0644)
	if err != nil {
		return err
	}

	return nil
}
