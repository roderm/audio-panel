package api

import (
	"fmt"
	"github.com/roderm/audio-panel/device"
	pb "github.com/roderm/audio-panel/proto"
	jws "github.com/roderm/json-rpc/websocket"
)

var storage *device.DeviceStore

func AddApi(handler *jws.Handler, strg *device.DeviceStore) {
	storage = strg
	handler.Add("get_devices", getDevices)
	handler.Add("set", setProperty)
	handler.AddSubscription("subscribe_update", subDevices)
}

func getDevices(param interface{}) (interface{}, jws.Error) {
	type Result struct {
		Devices []*pb.Device
	}
	ret := Result{}
	fmt.Println(storage.GetReceivers())
	for _, d := range storage.GetReceivers() {
		ret.Devices = append(ret.Devices, d.GetDevice())
	}
	return ret, jws.Error{Code: 0}
}

func setProperty(params interface{}) (interface{}, jws.Error) {
	var faults []error
	p, ok := params.([]interface{})
	if !ok {
		return nil, jws.Error{Code: 10, Message: "Parameter is not an array"}
	}
	if len(p) == 0 {
		return nil, jws.Error{Code: 11, Message: "Parameter array is empty"}
	}
	for _, r := range p {
		update, err := getUpdateRequest(r)
		if err == nil {
			if dev, ok := storage.GetReceivers()[update.DeviceIdentifier]; ok {
				err := dev.PropertyUpdate(update)
				if err != nil {
					faults = append(faults, err)
				}
			}
		} else {
			faults = append(faults, fmt.Errorf("Error from setProperty: %s", err.Error()))
		}
	}
	if len(faults) > 0 {
		msg := ""
		for _, e := range faults {
			msg = fmt.Sprintf("%s%s\n", msg, e)
		}
		return nil, jws.Error{Code: 30, Message: msg}
	}
	return nil, jws.Error{Code: 0}
}

func getUpdateRequest(v interface{}) (*pb.PropertyUpdate, error) {
	update := new(pb.PropertyUpdate)
	// err := mapstructure.Decode(r, &update)
	r_map, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Cannot convert request to map")
	}
	update.DeviceIdentifier, ok = r_map["DeviceIdentifier"].(string)
	if !ok {
		return nil, fmt.Errorf("Cannot get DeviceIdentifier")
	}
	update.ItemIdentifier, ok = r_map["ItemIdentifier"].(string)
	if !ok {
		return nil, fmt.Errorf("Cannot get DeviceIdentifier")
	}
	p_map, ok := r_map["Property"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Cannot convert property to map")
	}
	update.Property = &pb.Property{}
	update.Property.Name, ok = p_map["Name"].(string)
	if !ok {
		return nil, fmt.Errorf("Cannot get Propertyname")
	}
	map_val, ok := p_map["Value"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Cannot convert property to Value")
	}
	if v, ok := map_val["Boolean"]; ok {
		update.Property.Value = &pb.Property_Boolean{Boolean: v.(bool)}
		return update, nil
	}
	if v, ok := map_val["Text"]; ok {
		update.Property.Value = &pb.Property_Text{Text: v.(string)}
		return update, nil
	}
	if v, ok := map_val["Number"]; ok {
		update.Property.Value = &pb.Property_Number{Number: v.(int64)}
		return update, nil
	}
	if v, ok := map_val["Decimal"]; ok {
		update.Property.Value = &pb.Property_Decimal{Decimal: v.(float32)}
		return update, nil
	}
	return update, fmt.Errorf("Uknown type of val")
}
func subDevices(interface{}) chan interface{} {
	mychan := make(chan interface{})
	storage.SubscribeUpdate(func(update *pb.PropertyUpdate) {
		mychan <- update
	})
	return mychan
}
