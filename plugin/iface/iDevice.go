package iface

import (
	pb "github.com/roderm/audio-panel/proto"
)

type IReceiver interface {
	SetPower(string, bool) error
	Mute(string, bool) error
	SetVolume(string, int32) error
	SetSource(string, string) error
	SetListeningMod(string, string) error
	OnUpdate(func(*pb.AVR))
	GetAvr() *pb.AVR
}

type ILight interface {
	SetPower(string, bool) error
	SetBrightness(string, int32) error
}
