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
	ctx             context.Context
	ReceiverPlugins map[string]func(context.Context, interface{}) (pl.IReceiver, error)
	LightPlugins    map[string]func(context.Context, interface{}) (pl.ILight, error)
	receivers       map[int64]pl.IReceiver
	lights          map[int64]pl.ILight
	updateSubs      []func(*pb.AVR)
}

func NewDeviceStore(ctx context.Context) *DeviceStore {
	return &DeviceStore{
		ctx:             ctx,
		ReceiverPlugins: make(map[string]func(context.Context, interface{}) (pl.IReceiver, error)),
		LightPlugins:    make(map[string]func(context.Context, interface{}) (pl.ILight, error)),
		receivers:       make(map[int64]pl.IReceiver),
		lights:          make(map[int64]pl.ILight),
	}
}

func (d *DeviceStore) notifySubscritions(dev *pb.AVR, id int64) {
	dev.Id = int64(id)
	for _, s := range d.updateSubs {
		s(dev)
	}
}
func (d *DeviceStore) AddDevice(config DeviceConfig) (int64, error) {
	var err error
	getId := func() int64 {
		lock.Lock()
		deviceId = +1
		defer lock.Unlock()
		return int64(deviceId)
	}
	// Check if plugin is loaded
	switch config.PluginType {
	case "receiver":
		if _, ok := d.ReceiverPlugins[config.DriverPath]; !ok {
			err = d.addReceiverPlugin(config.DriverPath)
			if err != nil {
				return 0, err
			}
		}
		fmt.Println(config.Config)
		device, err := d.ReceiverPlugins[config.DriverPath](d.ctx, config.Config)
		if err != nil {
			return 0, err
		}
		did := getId()
		device.OnUpdate(func(id int64) func(dev *pb.AVR) {
			return func(dev *pb.AVR) {
				d.notifySubscritions(dev, id)
			}
		}(did))
		d.receivers[did] = device
		return did, err
	case "light":
		err = errors.New("No ligth implemented yet")
	default:
		err = errors.New("No known type")
	}
	return 0, err
}

func (d *DeviceStore) GetReceivers() map[int64]pl.IReceiver {
	return d.receivers
}

func (d *DeviceStore) SubscribeUpdate(f func(*pb.AVR)) {
	d.updateSubs = append(d.updateSubs, f)
}

func (d *DeviceStore) addReceiverPlugin(path string) error {
	plug, err := plugin.Open(path)
	if err != nil {
		return err
	}
	sym, err := plug.Lookup("NewDriver")
	if err != nil {
		return err
	}

	newFunc, ok := sym.(func(context.Context, interface{}) (pl.IReceiver, error))
	if !ok {
		return errors.New("Failed to load Plugin")
	}
	d.ReceiverPlugins[path] = newFunc
	return nil
}

/*
func createReceiver(ctx context.Context, config DeviceConfig) (pl.IReceiver, error) {
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
*/
