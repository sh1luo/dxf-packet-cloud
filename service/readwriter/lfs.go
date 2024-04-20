package readwriter

import (
	"github.com/bytedance/sonic"
	"os"
	"packet_cloud/biz/model/hertz/packet"
	"sync"
)

const (
	fileRelativePath = "./packets"
)

var (
	syncLock sync.RWMutex
)

type LocalFileSystem struct {
}

func (s *LocalFileSystem) ReadPacket() ([]*packet.CloudPacket, error) {
	packets := make([]*packet.CloudPacket, 0)

	syncLock.RLock()
	defer syncLock.RUnlock()

	bytes, err := os.ReadFile(fileRelativePath)
	if err != nil {
		return nil, err
	}
	err = sonic.Unmarshal(bytes, &packets)
	if err != nil {
		return nil, err
	}
	return packets, nil
}

func (s *LocalFileSystem) SavePacket(packets []*packet.CloudPacket) error {
	bytes, err := sonic.Marshal(packets)
	if err != nil {
		return err
	}

	syncLock.Lock()
	defer syncLock.Unlock()

	err = os.WriteFile(fileRelativePath, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
