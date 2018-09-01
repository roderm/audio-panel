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
	handler.Add("set_mute", setMute)
	handler.Add("set_power", setPower)
	handler.AddSubscription("subscribe_update", subDevices)
}

func getDevices(param interface{}) (interface{}, jws.Error) {
	type Result struct {
		Devices []*pb.AVR
	}
	ret := Result{}
	for _, d := range storage.GetDevices() {
		ret.Devices = append(ret.Devices, d.GetAvr())
	}
	return ret, jws.Error{Code: 0}
}

func setPower(params interface{}) (interface{}, jws.Error) {
	type Power struct {
		Device int64
		Zone   string
		Power  bool
	}
	type Result struct {
		Ok bool
	}

	var requests []Power
	p, ok := params.([]interface{})
	if !ok {
		log.Printf("Param for set Mute is no Array \n")
	}
	for _, r := range p {
		var myval Power
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
				err := d.SetPower(v.Zone, v.Power)
				if err != nil {
					return Result{Ok: false}, jws.Error{Code: 500, Message: err.Error()}
				}
				return Result{Ok: true}, jws.Error{Code: 0}
			}
		}
	}
	return Result{Ok: false}, jws.Error{Code: 0}
}

func setMute(params interface{}) (interface{}, jws.Error) {
	type Mute struct {
		Device int64
		Zone   string
		Mute   bool
	}
	type Result struct {
		Ok bool
	}

	var requests []Mute
	p, ok := params.([]interface{})
	if !ok {
		log.Printf("Param for set Mute is no Array \n")
	}
	for _, r := range p {
		var myval Mute
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
				err := d.Mute(v.Zone, v.Mute)
				if err != nil {
					return Result{Ok: false}, jws.Error{Code: 500, Message: err.Error()}
				}
				return Result{Ok: true}, jws.Error{Code: 0}
			}
		}
	}
	return Result{Ok: false}, jws.Error{Code: 0}
}

func setVolume(params interface{}) (interface{}, jws.Error) {
	type Volume struct {
		Device int64
		Zone   string
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
				err := d.SetVolume(v.Zone, v.Volume)
				if err != nil {
					return Result{Ok: false}, jws.Error{Code: 500, Message: err.Error()}
				}
				return Result{Ok: true}, jws.Error{Code: 0}
			}
		}
	}
	return Result{Ok: false}, jws.Error{Code: 500, Message: "Unknown errors"}
}

func subDevices(interface{}) chan interface{} {
	mychan := make(chan interface{})
	storage.SubscribeUpdate(func(update *pb.AVR) {
		mychan <- update
	})
	return mychan
}
