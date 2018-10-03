package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/roderm/audio-panel/proto"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const apiPath = "api/v1"

type StateValue struct {
	Value bool `json:"value"`
}

type BrightnessValue struct {
	Value int `json:"value"`
	Max   int `json:"max"`
	Min   int `json:"min"`
}

type NanoLeafState struct {
	On         StateValue      `json:"on"`
	Brightness BrightnessValue `json:"brightness"`
}
type NanoLeafResponse struct {
	Name  string        `json:"name"`
	State NanoLeafState `json:"state"`
}
type nanoLeafConfig struct {
	DeviceAddress string
	DevicePort    string
	AuthToken     string
	Pullcycle     int
}
type NanoLeafDriver struct {
	cl     *http.Client
	ctx    context.Context
	conf   nanoLeafConfig
	light  *pb.Device
	update []func(*pb.PropertyUpdate)
}

func NewNanoleafDriver(ctx context.Context, config nanoLeafConfig, id string) *NanoLeafDriver {
	dr := &NanoLeafDriver{
		light: &pb.Device{
			Identifier: id,
			Items:      []*pb.Item{&pb.Item{Name: "main"}},
		},
		ctx: ctx,
		cl: &http.Client{
			Timeout: time.Second * 10,
		},
	}
	return dr
}

func (d *NanoLeafDriver) startPulls() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			r, err := d.cl.Get(fmt.Sprintf("%s/%s/%s", d.getAddr(), apiPath, d.conf.AuthToken))
			if err != nil {
				fmt.Println(err)
				continue
			}
			func(response *http.Response) {
				defer response.Body.Close()
				var resp NanoLeafResponse
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					fmt.Println(err)
					return
				}
				err = json.Unmarshal(contents, &resp)
				if err != nil {
					fmt.Println(err)
					return
				}
			}(r)
			time.Sleep(time.Second * time.Duration(d.conf.Pullcycle))
		}
	}
}
func (d *NanoLeafDriver) checkUpdate(n NanoLeafResponse) {
	var item *pb.Item
	item, err := d.getItem(n.Name)
	if err != nil {
		item = &pb.Item{
			Name: n.Name,
		}
		d.light.Items = append(d.light.Items, item)
	}
	updateProp := func(item *pb.Item, cmd string, value interface{}) {
		noUpdate := true
		prop, err := d.getProp(item, cmd)
		if err != nil {
			prop = &pb.Property{
				Name: cmd,
			}
			item.Properties = append(item.Properties, prop)
		}
		switch reflect.TypeOf(value).Kind() {
		case reflect.Bool:
			noUpdate = noUpdate && prop.Value.(*pb.Property_Boolean).Boolean == value.(bool)
			prop.Value = &pb.Property_Boolean{Boolean: value.(bool)}
		case reflect.Int:
			prop.Value = &pb.Property_Number{Number: value.(int64)}
		case reflect.String:
			prop.Value = &pb.Property_Text{Text: value.(string)}
		}
		if !noUpdate {
			for _, fn := range d.update {
				go fn(&pb.PropertyUpdate{
					DeviceIdentifier: d.light.Identifier,
					ItemIdentifier:   item.Identifier,
					Property:         prop,
				})
			}
		}
	}
	updateProp(item, "on", n.State.On.Value)
	updateProp(item, "brightness", n.State.Brightness.Value)

}

func (d *NanoLeafDriver) getItem(item string) (*pb.Item, error) {
	for _, i := range d.light.Items {
		if item == i.Name {
			return i, nil
		}
	}
	return nil, fmt.Errorf("Item not found")
}
func (d *NanoLeafDriver) getProp(item *pb.Item, prop string) (*pb.Property, error) {
	for _, p := range item.Properties {
		if p.Name == prop {
			return p, nil
		}
	}
	return nil, fmt.Errorf("Property not found")
}
func (d *NanoLeafDriver) getAddr() string {
	return fmt.Sprintf("http://%s:%s", d.conf.DeviceAddress, d.conf.DevicePort)
}
func (d *NanoLeafDriver) SetPower(on bool) error {
	type reqBody struct {
		On StateValue `json:"on"`
	}
	body := reqBody{
		On: StateValue{
			Value: on,
		},
	}
	return d.put("state", body)
}

func (d *NanoLeafDriver) setBrightness(value int32) error {
	type BrightValue struct {
		Value    int32 `json:"value"`
		Duration int32 `json:"duration"`
	}
	type reqBody struct {
		Brightness BrightValue `json:"brightness"`
	}
	body := reqBody{
		Brightness: BrightValue{
			Value:    value,
			Duration: 10,
		},
	}
	return d.put("state", body)
}

func (d *NanoLeafDriver) put(path string, body interface{}) error {
	jsBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	r, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/%s/%s/%s", d.getAddr(), apiPath, d.conf.AuthToken, path),
		strings.NewReader(string(jsBody)),
	)
	if err != nil {
		return err
	}
	r.ContentLength = int64(len(string(jsBody)))
	response, err := d.cl.Do(r)
	if err != nil {
		return err
	} else {
		switch response.StatusCode {
		case http.StatusNoContent:
			return nil
		case http.StatusUnauthorized:
			return errors.New("Auth token is invalid")
		case http.StatusBadRequest:
			return errors.New("Bad request => check body")
		case http.StatusNotFound:
			return errors.New("Resource not found")
		default:
			return errors.New("Uknown error")
		}

	}
}

func (d *NanoLeafDriver) PropertyUpdate(update *pb.PropertyUpdate) error {
	return nil
}
func (d *NanoLeafDriver) OnUpdate(cb func(*pb.PropertyUpdate)) {
	d.update = append(d.update, cb)
}
func (d *NanoLeafDriver) GetDevice() *pb.Device {
	return d.light
}
