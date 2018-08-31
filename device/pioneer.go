package device

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	pb "github.com/roderm/audio-panel/proto"
	"github.com/roderm/audio-panel/telnet"
	"log"
	"regexp"
	"strconv"
	"time"
)

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
	console    *telnet.PioneerCaller
	avr        *pb.AVR
	updateSubs []func(*pb.AVR)
}

func NewPioneerDriver(ctx context.Context, config DeviceConfig, cmdConfig CommandSet) (IDevice, error) {
	p := &PioneerDriver{
		ctx: ctx,
		avr: &pb.AVR{
			Zones: make(map[string]*pb.AVR_Zone),
		},
	}

	port, err := strconv.Atoi(config.DevicePort)
	if err != nil {
		return nil, err
	}
	p.console, err = telnet.NewPioneerCaller(ctx, config.DeviceAddress, port)
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
		p.avr.Zones[zs.Name] = &pb.AVR_Zone{
			Name:   zs.Name,
			Power:  false,
			IsMain: zs.IsMain,
			Mute:   false,
		}
		for cmd, set := range zs.Commands {
			if resp, ok := set.Command["response"]; ok {
				if _, ok := set.Command["datatype"]; ok {
					p.console.Subscribe(resp, func(cmd string) func(interface{}) {
						return func(val interface{}) {
							switch cmd {
							case "power":
								p.avr.Zones[zs.Name].Power = (toInt(val) > 0)
							case "volume":
								p.avr.Zones[zs.Name].Volume = int32((float32(100) / float32(zs.MaxVol)) * float32(toInt(val)))
							case "mute":
								p.avr.Zones[zs.Name].Mute = (toInt(val) > 0)
							case "input":
								p.avr.Zones[zs.Name].CurrentSource = val.(string)
							case "listening_mod":
								p.avr.Zones[zs.Name].CurrentListeningMod = val.(string)
							default:
								return
							}
							p.notifyUpdate()
						}
					}(cmd))
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
	go func() {
		for {
			select {
			case <-p.ctx.Done():
				return
			default:
				for _, z := range p.Zones {
					for _, c := range z.Commands {
						if resp, ok := c.Command["get"]; ok {
							p.console.Send(resp)
						}
					}
				}
				time.Sleep(time.Minute * 5)
			}
		}
	}()
}
func mapToConfigDatatype(dt string, val interface{}) (interface{}, error) {
	switch dt {
	case "string":
		return val.(string), nil
	case "int":
		return strconv.Atoi(val.(string))
	}
	return val, nil
}

func (p *PioneerDriver) simpleSet(zone string, key string, value interface{}) error {
	z, err := p.getZoneSetup(zone)
	if err != nil {
		return err
	}
	if cmd, ok := z.Commands[key]; ok {
		if set, ok := cmd.Command["set"]; ok {
			p.sendCommand(set, value)
			return nil
		} else {
			return fmt.Errorf("not enough driver info available to set %s", key)
		}
	}
	return fmt.Errorf("Zone doesn't support %s", key)
}
func (p *PioneerDriver) sendCommand(command string, value interface{}) {
	r := regexp.MustCompile(`(?P<PLACEHOLDER>#{1,})`)
	PLACEHOLDER := r.FindString(command)
	format := fmt.Sprintf(`%%0%dv`, len(PLACEHOLDER))
	cmd := r.ReplaceAllString(command, fmt.Sprintf(format, value))
	p.console.Send(cmd)
}

func (p *PioneerDriver) notifyUpdate() {
	for _, f := range p.updateSubs {
		go f(p.avr)
	}
}

func (p *PioneerDriver) getZoneSetup(zone string) (*pioneerZoneSetup, error) {
	for _, z := range p.Zones {
		if z.Name == zone {
			return &z, nil
		}
	}
	return nil, fmt.Errorf("No zone found for %s", zone)
}

func (p *PioneerDriver) SetPower(zone string, on bool) error {
	onInt := 0
	if on {
		onInt = 1
	}
	return p.simpleSet(zone, "power", onInt)
}

func (p *PioneerDriver) Mute(zone string, mute bool) error {
	muteInt := 1
	if mute {
		muteInt = 1
	}
	return p.simpleSet(zone, "mute", muteInt)
}

func (p *PioneerDriver) SetSource(zone string, source string) error {
	return p.simpleSet(zone, "input", source)
}
func (p *PioneerDriver) SetListeningMod(zone string, mode string) error {
	return p.simpleSet(zone, "input", mode)
}
func (p *PioneerDriver) SetVolume(zone string, volume int32) error {
	z, err := p.getZoneSetup(zone)
	if err != nil {
		panic(err)
	}
	vol := int(float32(z.MaxVol*int(volume)) / 100.0)
	if cmd, ok := z.Commands["volume"]; ok {
		_, hasSet := cmd.Command["set"]
		_, hasUp := cmd.Command["up"]
		_, hasDown := cmd.Command["down"]
		_, hasResponse := cmd.Command["response"]
		switch true {
		case hasSet:
			p.sendCommand(cmd.Command["set"], vol)
		case hasUp && hasDown && hasResponse:
			p.console.Once(cmd.Command["response"], func(val interface{}) {
				current := val.(int)
				switch true {
				case vol > current:
					for i := 0; i < vol-current; i++ {
						p.console.Send(cmd.Command["up"])
					}
				case vol < current:
					for i := 0; i < vol-current; i++ {
						p.console.Send(cmd.Command["down"])
					}
				}
			})
		default:
			return fmt.Errorf("not enough driver info available to set volume")
		}
		return nil
	} else {
		return fmt.Errorf("zone doesn't support volume")
	}
}

func (p *PioneerDriver) OnUpdate(f func(*pb.AVR)) {
	p.updateSubs = append(p.updateSubs, f)
}
func (p *PioneerDriver) GetAvr() *pb.AVR {
	return p.avr
}

func toInt(iface interface{}) int {
	i, _ := strconv.Atoi(iface.(string))
	return i
}
