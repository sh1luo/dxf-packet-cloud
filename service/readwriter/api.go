package readwriter

import "packet_cloud/biz/model/hertz/packet"

type StorageMedia int

const (
	LFS StorageMedia = iota
	MySQL
)

type ReadWriter interface {
	ReadPacket() ([]*packet.CloudPacket, error)
	SavePacket([]*packet.CloudPacket) error
}

func NewReadWriter(media StorageMedia) ReadWriter {
	switch media {
	case LFS:
		return &LocalFileSystem{}
	case MySQL:

	default:

	}
	return nil
}
