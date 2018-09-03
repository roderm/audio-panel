package iface

import (
	pb "github.com/roderm/audio-panel/proto"
)

type IDevice interface {
	PropertyUpdate(*pb.PropertyUpdate) error
	OnUpdate(func(*pb.PropertyUpdate))
	GetDevice() *pb.Device
}
