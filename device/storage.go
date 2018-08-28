package device

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/roderm/audio-panel/proto"
	"sync"
)

var deviceId int = 0
var lock sync.Mutex

type DeviceStore struct {
	ctx        context.Context
	devs       []IDevice
	updateSubs []func(*pb.AVR)
}

func NewDeviceStore(ctx context.Context) *DeviceStore {
	return &DeviceStore{ctx: ctx}
}

func (d *DeviceStore) notifySubscritions(dev *pb.AVR, id int64) {
	dev.Id = int64(id)
	for _, s := range d.updateSubs {
		s(dev)
	}
}
func (d *DeviceStore) AddDevice(config DeviceConfig) (int, error) {
	device, err := createDevice(d.ctx, config)
	if err != nil {
		return 0, err
	}
	lock.Lock()
	deviceId = +1
	avr := device.GetAvr()
	avr.Id = int64(deviceId)
	lock.Unlock()
	updateFunc := func(dev *pb.AVR) {
		fmt.Println(avr.GetZones()[0].GetVolume())
		d.notifySubscritions(dev, avr.Id)
	}
	device.OnUpdate(updateFunc)
	d.devs = append(d.devs, device)
	// trigger new device
	updateFunc(avr)
	return int(avr.Id), nil
}

func (d *DeviceStore) GetDevices() []IDevice {
	return d.devs
}

func (d *DeviceStore) SubscribeUpdate(f func(*pb.AVR)) {
	d.updateSubs = append(d.updateSubs, f)
}

type DeviceConfig struct {
	DeviceType    string `json:"device_type"`
	DeviceAddress string `json:"device_address"`
	DevicePort    string `json:"device_port"`
}

func createDevice(ctx context.Context, config DeviceConfig) (IDevice, error) {
	switch config.DeviceType {
	case "pioneer":
		return NewPioneerDevice(ctx, config.DeviceAddress), nil
	default:
		return nil, errors.New("No device configured")
	}
}
