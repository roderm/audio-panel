package device

import (
	pb "github.com/roderm/audio-panel/proto"
)

type IDevice interface {
	SetPower(bool)
	Mute(int32, bool)
	SetVolume(int32, int32)
	SetSource(int32, int32)
	OnUpdate(func())
	GetAvr() *pb.AVR
}
