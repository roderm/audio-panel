package iface

import (
	pb "github.com/roderm/audio-panel-protobuf/go/msg/device"
)

type IDevice interface {
	PropertyUpdate(*pb.PropertyUpdate) error
	OnUpdate(func(*pb.PropertyUpdate))
	GetDevice() *pb.Device
}
