package device

import (
	"context"
	"errors"
	"fmt"
	pl "github.com/roderm/audio-panel/plugin/iface"
	pb "github.com/roderm/audio-panel/proto"
	"plugin"
	"sync"
)

var deviceId int = 0
var lock sync.Mutex

type DeviceConfig struct {
	DriverPath string      `json:"driver"`
	PluginType string      `json:"type"`
	Config     interface{} `json:"config"`
}
type DeviceStore struct {
	ctx           context.Context
	DevicePlugins map[string]func(context.Context, interface{}, string) (pl.IDevice, error)
	devices       map[string]pl.IDevice
	updateSubs    []func(*pb.PropertyUpdate)
}

func NewDeviceStore(ctx context.Context) *DeviceStore {
	return &DeviceStore{
		ctx:           ctx,
		DevicePlugins: make(map[string]func(context.Context, interface{}, string) (pl.IDevice, error)),
		devices:       make(map[string]pl.IDevice),
	}
}

func (d *DeviceStore) notifySubscritions(dev *pb.PropertyUpdate) {
	for _, s := range d.updateSubs {
		s(dev)
	}
}
func (d *DeviceStore) AddDevice(config DeviceConfig) (string, error) {
	var err error
	lock.Lock()
	defer lock.Unlock()
	getId := func() int64 {
		deviceId += 1
		return int64(deviceId)
	}
	// Check if plugin is loaded

	if _, ok := d.DevicePlugins[config.DriverPath]; !ok {
		err = d.addPlugin(config.DriverPath)
		if err != nil {
			return "", err
		}
	}
	cid := getId()
	did := fmt.Sprintf("Device_%d", cid)
	device, err := d.DevicePlugins[config.DriverPath](d.ctx, config.Config, did)
	if err != nil {
		return did, err
	}
	device.OnUpdate(func(id string) func(*pb.PropertyUpdate) {
		return func(update *pb.PropertyUpdate) {
			update.DeviceIdentifier = id
			d.notifySubscritions(update)
		}
	}(did))
	d.devices[did] = device
	return did, err
}

func (d *DeviceStore) GetReceivers() map[string]pl.IDevice {
	return d.devices
}

func (d *DeviceStore) SubscribeUpdate(f func(*pb.PropertyUpdate)) {
	d.updateSubs = append(d.updateSubs, f)
}

func (d *DeviceStore) addPlugin(path string) error {
	plug, err := plugin.Open(path)
	if err != nil {
		return err
	}
	sym, err := plug.Lookup("NewDriver")
	if err != nil {
		return err
	}

	newFunc, ok := sym.(func(context.Context, interface{}, string) (pl.IDevice, error))
	if !ok {
		return errors.New("Failed to load Plugin")
	}
	d.DevicePlugins[path] = newFunc
	return nil
}
