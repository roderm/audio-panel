package main

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	pl "github.com/roderm/audio-panel/plugin/iface"
	pb "github.com/roderm/audio-panel/proto"
	"log"
	"regexp"
	"strconv"
	"time"
)

type PioneerConfig struct {
	Setup         string `json:"Setup"`
	DeviceAddress string `json:"DeviceAddress"`
	DevicePort    string `json:"DevicePort"`
}
type CommandSet struct {
	Driver       string                     `json:"driver"`
	Zones        []interface{}              `json:"zones"`
	InputSources []pb.AVR_Zone_Source       `json:"input_sources"`
	ListenMods   []pb.AVR_Zone_ListeningMod `json:"listening_mods"`
}
type pioneerCommandSetup struct {
	Command map[string]string
}
type pioneerZoneSetup struct {
	Name     string                         `json:"name"`
	IsMain   bool                           `json:"main"`
	MaxVol   int                            `json:"maxVol"`
	Commands map[string]pioneerCommandSetup `json:"commands"`
}
type PioneerDriver struct {
	ctx        context.Context
	Zones      []pioneerZoneSetup
	console    *PioneerCaller
	avr        *pb.Device
	updateSubs []func(*pb.PropertyUpdate)
}

func NewPioneerDriver(ctx context.Context, config PioneerConfig, cmdConfig CommandSet, id string) (pl.IDevice, error) {
	p := &PioneerDriver{
		ctx: ctx,
		avr: &pb.Device{
			Identifier: id,
			//Zones: make(map[string]*pb.AVR_Zone),
		},
	}

	port, err := strconv.Atoi(config.DevicePort)
	if err != nil {
		return nil, err
	}
	p.console, err = NewPioneerCaller(ctx, config.DeviceAddress, port)
	if err != nil {
		return nil, err
	}
	for _, z := range cmdConfig.Zones {
		var zs pioneerZoneSetup
		err := mapstructure.Decode(z, &zs)
		zoneMap := z.(map[string]interface{})
		err = mapstructure.Decode(zoneMap["commands"], &zs.Commands)
		for cmd, set := range zoneMap["commands"].(map[string]interface{}) {
			var c pioneerCommandSetup
			err := mapstructure.Decode(set, &c.Command)
			if err != nil {
				log.Println(err)
			}
			zs.Commands[cmd] = c
		}
		if err != nil {
			return nil, err
		}
		p.Zones = append(p.Zones, zs)
		// register listeners
		pbZone := &pb.Item{
			Name:       zs.Name,
			Type:       "sound_amplifier",
			Identifier: zs.Name,
			Properties: []*pb.Property{},
		}
		p.avr.Items = append(p.avr.Items, pbZone)

		for cmd, set := range zs.Commands {
			prop := &pb.Property{
				Name: cmd,
			}
			pbZone.Properties = append(pbZone.Properties, prop)
			if resp, ok := set.Command["response"]; ok {
				if _, ok := set.Command["datatype"]; ok {
					p.console.Subscribe(resp, func(zone *pb.Item, prop *pb.Property, cmd string) func(interface{}) {
						return func(val interface{}) {
							if prop == nil {
								return
							}
							nprop, err := p.newProperty(zone.Name, cmd, val)
							if err != nil {
								fmt.Println(err)
								return
							}
							prop.Value = nprop.Value
							p.notifyUpdate(&pb.PropertyUpdate{
								ItemIdentifier: zone.Identifier,
								Property:       nprop,
							})
						}
					}(pbZone, prop, cmd))
				} else {
					log.Printf("No datatype addr for %s: %+v", cmd, set)
				}
			} else {
				log.Printf("Commandset for %s not available: %+v", cmd, set)
			}
		}
	}
	p.console.StartListen()
	p.startTriggers()
	return p, err
}

func (p *PioneerDriver) startTriggers() {
	for _, z := range p.Zones {
		for _, c := range z.Commands {
			cycle := 1800
			if cycleStr, ok := c.Command["pullcycle"]; ok {
				cycle, _ = strconv.Atoi(cycleStr)
			}
			if resp, ok := c.Command["get"]; ok {
				go func() {
					for {
						select {
						case <-p.ctx.Done():
							return
						default:
							p.console.Send(resp)
							time.Sleep(time.Second * time.Duration(cycle))
						}
					}
				}()
			}
		}
	}
}

func (p *PioneerDriver) newProperty(zone string, cmd string, value interface{}) (*pb.Property, error) {
	zs, err := p.getZoneSetup(zone)
	if err != nil {
		return nil, err
	}
	if command, ok := zs.Commands[cmd]; ok {
		if datatype, ok := command.Command["datatype"]; ok {
			switch datatype {
			case "bool":
				return &pb.Property{Name: cmd, Value: &pb.Property_Boolean{Boolean: (toInt(value) == 0)}}, nil
			case "number":
				return &pb.Property{Name: cmd, Value: &pb.Property_Number{Number: int64(toInt(value))}}, nil
			case "string":
				return &pb.Property{Name: cmd, Value: &pb.Property_Text{Text: value.(string)}}, nil
			case "percentage":
				min, min_ok := command.Command["min"]
				max, max_ok := command.Command["max"]
				if min_ok && max_ok {
					min, min_err := strconv.Atoi(min)
					max, max_err := strconv.Atoi(max)
					if min_err == nil && max_err == nil {
						return &pb.Property{Name: cmd, Value: &pb.Property_Decimal{Decimal: float32(((max - min) / 100) * toInt(value))}}, nil
					}
				}
				return nil, fmt.Errorf("min or max not properly set for %s", cmd)
			default:
				return nil, fmt.Errorf("Unknown datatype %s", datatype)
			}
		} else {
			return nil, fmt.Errorf("Datatype not set for %s", cmd)
		}
	} else {
		return nil, fmt.Errorf("Unknown property %s", cmd)
	}
}

func (p *PioneerDriver) set(zone *pioneerZoneSetup, key string, value interface{}) error {
	if cmd, ok := zone.Commands[key]; ok {
		if set, ok := cmd.Command["set"]; ok {
			p.sendCommand(set, value)
			return nil
		} else {
			return fmt.Errorf("not enough driver info available to set %s", key)
		}
	}
	return fmt.Errorf("Zone doesn't support %s", key)
}
func (p *PioneerDriver) simpleSet(zone string, key string, value interface{}) error {
	z, err := p.getZoneSetup(zone)
	if err != nil {
		return err
	}
	return p.set(z, key, value)
}
func (p *PioneerDriver) sendCommand(command string, value interface{}) {
	r := regexp.MustCompile(`(?P<PLACEHOLDER>#{1,})`)
	PLACEHOLDER := r.FindString(command)
	format := fmt.Sprintf(`%%0%dv`, len(PLACEHOLDER))
	cmd := r.ReplaceAllString(command, fmt.Sprintf(format, value))
	p.console.Send(cmd)
}

func (p *PioneerDriver) notifyUpdate(u *pb.PropertyUpdate) {
	for _, f := range p.updateSubs {
		go f(u)
	}
}

func (p *PioneerDriver) getPbZone(zone string) (*pb.Item, error) {
	for _, z := range p.avr.Items {
		if z.Identifier == zone {
			return z, nil
		}
	}
	return nil, fmt.Errorf("Zone not found")
}
func (p *PioneerDriver) getZoneSetup(zone string) (*pioneerZoneSetup, error) {
	for _, z := range p.Zones {
		if z.Name == zone {
			return &z, nil
		}
	}
	return nil, fmt.Errorf("No zone found for %s", zone)
}

func (p *PioneerDriver) setPercentage(z *pioneerZoneSetup, c *pioneerCommandSetup, val float32) error {
	str_min, ok_min := c.Command["min"]
	str_max, ok_max := c.Command["max"]
	if !ok_min || !ok_max {
		return fmt.Errorf("min or max is not properly set.")
	}
	min, min_err := strconv.Atoi(str_min)
	max, max_err := strconv.Atoi(str_max)
	if min_err != nil || max_err != nil {
		return fmt.Errorf("min or max must be integers.")
	}
	target := int((float32(max-min) * val) / 100.0)
	if set, ok := c.Command["set"]; ok {
		p.sendCommand(set, target)
		return nil
	}

	_, hasUp := c.Command["up"]
	_, hasDown := c.Command["down"]
	_, hasResponse := c.Command["response"]
	if hasUp && hasDown && hasResponse {
		p.console.Once(c.Command["response"], func(val interface{}) {
			current := val.(int)
			switch true {
			case target > current:
				for i := 0; i < target-current; i++ {
					p.console.Send(c.Command["up"])
				}
			case target < current:
				for i := 0; i < target-current; i++ {
					p.console.Send(c.Command["down"])
				}
			}
		})
		return nil
	}
	return fmt.Errorf("Not enough arguments to set percentage")
}
func (p *PioneerDriver) PropertyUpdate(u *pb.PropertyUpdate) error {
	prop := u.GetProperty()
	z, err := p.getZoneSetup(u.GetItemIdentifier())
	if err != nil {
		return err
	}
	if zoneProperty, ok := z.Commands[prop.GetName()]; ok {
		if dt, ok := zoneProperty.Command["datatype"]; ok {
			switch dt {
			case "bool":
				onVal := "F"
				if prop.GetBoolean() {
					onVal = "O"
				}
				return p.set(z, prop.GetName(), onVal)
			case "percentage":
				return p.setPercentage(z, &zoneProperty, prop.GetDecimal())
			case "number":
				return p.set(z, prop.GetName(), prop.GetNumber())
			case "string":
				return p.set(z, prop.GetName(), prop.GetText())
			default:
				return fmt.Errorf("Unknown datatype \"%s\"", dt)
			}
		}
	}
	return fmt.Errorf("Unknown property \"%s\"", prop.GetName())
}
func (p *PioneerDriver) OnUpdate(f func(*pb.PropertyUpdate)) {
	p.updateSubs = append(p.updateSubs, f)
}
func (p *PioneerDriver) GetDevice() *pb.Device {
	return p.avr
}

func toInt(iface interface{}) int {
	i, _ := strconv.Atoi(iface.(string))
	return i
}
