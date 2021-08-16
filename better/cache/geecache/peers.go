package geecache

import (
	pb "github.com/oushuifa/golang/better/cache/geecachepb"
)

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(input *pb.Request, output *pb.Response) error
}
