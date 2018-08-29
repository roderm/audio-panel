package device

import (
	pb "github.com/roderm/audio-panel/proto"
)

type IDevice interface {
	Reachable() bool
	SetPower(bool)
	Mute(int32, bool)
	SetVolume(int32, int32)
	SetSource(int32, string)
	SetListeningMod(int32, string)
	OnUpdate(func(*pb.AVR))
	GetAvr() *pb.AVR
}
