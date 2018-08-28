package api

import (
	"github.com/mitchellh/mapstructure"
	"github.com/roderm/audio-panel/device"
	pb "github.com/roderm/audio-panel/proto"
	jws "github.com/roderm/json-rpc/websocket"
	"log"
)

var storage *device.DeviceStore

func AddDeviceApi(handler *jws.Handler, strg *device.DeviceStore) {
	storage = strg
	handler.Add("get_devices", getDevices)
	handler.Add("set_volume", setVolume)

	handler.AddSubscription("subscribe_update", subDevices)
}

func getDevices(param interface{}) interface{} {
	type Result struct {
		Devices []*pb.AVR
	}
	ret := Result{}
	for _, d := range storage.GetDevices() {
		ret.Devices = append(ret.Devices, d.GetAvr())
	}
	return ret
}

func setVolume(params interface{}) interface{} {
	type Volume struct {
		Device int64
		Zone   int32
		Volume int32
	}
	type Result struct {
		Ok bool
	}

	var requests []Volume
	p, ok := params.([]interface{})
	if !ok {
		log.Printf("Param for set Volume is no Array \n")
	}
	for _, r := range p {
		var myval Volume
		err := mapstructure.Decode(r, &myval)
		if err == nil {
			requests = append(requests, myval)
		} else {
			log.Println(err)
		}
	}

	for _, v := range requests {
		devs := storage.GetDevices()
		for _, d := range devs {
			if d.GetAvr().Id == v.Device {
				d.SetVolume(v.Zone, v.Volume)
				return Result{Ok: true}
			}
		}
	}
	return Result{Ok: false}
}

func subDevices(interface{}) chan interface{} {
	mychan := make(chan interface{})
	storage.SubscribeUpdate(func(update *pb.AVR) {
		mychan <- update
	})
	return mychan
}
