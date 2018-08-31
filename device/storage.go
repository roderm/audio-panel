package device

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/roderm/audio-panel/proto"
	"io/ioutil"
	"log"
	"os"
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
		log.Panicf("Couldn't create device: %s", err.Error())
		return 0, err
	}
	lock.Lock()
	deviceId = +1
	avr := device.GetAvr()
	avr.Id = int64(deviceId)
	lock.Unlock()
	updateFunc := func(dev *pb.AVR) {
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

func createDevice(ctx context.Context, config DeviceConfig) (IDevice, error) {
	if _, err := os.Stat(config.Setup); err != nil {
		return nil, fmt.Errorf("path to configfile(%s) not found.", config.Setup)
	}

	fileContent, err := ioutil.ReadFile(config.Setup)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read configfile(%s).", config.Setup)
	}
	var setup CommandSet
	err = json.Unmarshal(fileContent, &setup)
	if err != nil {
		return nil, fmt.Errorf("Bad Syntax in configfile(%s), wasn't able to parse json.", config.Setup)
	}

	switch setup.Driver {
	case "pioneer":
		return NewPioneerDriver(ctx, config, setup)
	default:
		return nil, errors.New("No device configured")
	}
}
