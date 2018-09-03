package device

import (
	pb "github.com/roderm/audio-panel/proto"
)

type CommandSet struct {
	Driver       string                     `json:"driver"`
	Zones        []interface{}              `json:"zones"`
	InputSources []pb.AVR_Zone_Source       `json:"input_sources"`
	ListenMods   []pb.AVR_Zone_ListeningMod `json:"listening_mods"`
}

type NewDeviConfic struct {
	Driver string      `json:"driver"`
	Config interface{} `json:"driver"`
}
type DeviceConfig struct {
	Setup         string `json:"command_set"`
	DeviceAddress string `json:"device_address"`
	DevicePort    string `json:"device_port"`
}

type IDevice interface {
	SetPower(string, bool) error
	Mute(string, bool) error
	SetVolume(string, int32) error
	SetSource(string, string) error
	SetListeningMod(string, string) error
	OnUpdate(func(*pb.AVR))
	GetAvr() *pb.AVR
}
