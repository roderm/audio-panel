package device

import (
	"context"
	"fmt"
	pb "github.com/roderm/audio-panel/proto"
	"github.com/roderm/audio-panel/telnet"
	"regexp"
	"strconv"
	"time"
)

const maxVol = 160 // 185
const maxHZ = 80

var listeningMods = map[string]string{
	"0041": "Extended Stereo",
}

type PioneerDevice struct {
	ctx         context.Context
	device      pb.AVR
	nc          *telnet.PioneerCaller
	updateFuncs []func(*pb.AVR)
}

func NewPioneerDevice(ctx context.Context, ip string) IDevice {
	p := PioneerDevice{
		ctx: ctx,
	}
	p.nc, _ = telnet.NewPioneerCaller(ctx, ip, 8102)
	p.device = pb.AVR{
		IP:    ip,
		Zones: make(map[int32]*pb.AVR_Zone),
	}
	p.device.Zones[0] = &pb.AVR_Zone{
		IsMain: true,
		Name:   "main",
		Volume: 80,
		Muted:  false,
	}
	p.device.Zones[1] = &pb.AVR_Zone{
		IsMain: true,
		Name:   "hdzone",
		Volume: 80,
		Muted:  false,
	}
	p.initCommands()
	return &p
}

func (d *PioneerDevice) initCommands() {
	d.startListener()
	go func() {
		for {
			select {
			case <-d.ctx.Done():
				return
			default:
				d.nc.Send("?V")
				d.nc.Send("?PWR")
				d.nc.Send("?HZV")
				d.nc.Send("?M")
				d.nc.Send("?HZM")
				d.nc.Send("?S")
				time.Sleep(time.Second * 10)
			}
		}
	}()
}

func (d *PioneerDevice) Reachable() bool {
	return true
}
func (d *PioneerDevice) SetPower(on bool) {
	if on {
		d.nc.Send("PWR1")
	} else {
		d.nc.Send("PWR0")
	}
}

func (d *PioneerDevice) Mute(zone int32, on bool) {
	var mt string
	switch zone {
	case 0:
		mt = "M"
	case 1:
		mt = "HZM"
	}
	if on {
		d.nc.Send(fmt.Sprintf("%sO", mt))
	} else {
		d.nc.Send(fmt.Sprintf("%sF", mt))
	}
}

func (d *PioneerDevice) SetSource(zone int32, src string) {

}

func (d *PioneerDevice) SetListeningMod(zone int32, src string) {

}

func (d *PioneerDevice) SetVolume(zone int32, volume int32) {
	switch zone {
	case 0:
		vol := int(float32(maxVol*volume)/100) - 1
		d.nc.Send(fmt.Sprintf("%03dVL", vol))
	case 1:
		vol := int(float32(maxHZ*volume) / 100)
		d.nc.Send(fmt.Sprintf("%02dHZV", vol))
	}
}

func (d *PioneerDevice) OnUpdate(fn func(*pb.AVR)) {
	d.updateFuncs = append(d.updateFuncs, fn)
}
func (d *PioneerDevice) GetAvr() *pb.AVR {
	return &d.device
}

func (d *PioneerDevice) startListener() {
	d.nc.StartListen()
	go func() {
		for {
			select {
			case <-d.ctx.Done():
				return
			case cmd := <-d.nc.RecCommands:
				go func(command string) {
					var expression = regexp.MustCompile(`(?P<COMMAND>^[A-Z]+[\d]??[A-Z]+)(?P<VALUE>[\d]{2,})`)
					COMMAND := expression.ReplaceAllString(cmd, "${COMMAND}")
					VALUE := expression.ReplaceAllString(cmd, "${VALUE}")
					switch COMMAND {
					case "PWR":
						d.device.Power = toBool(VALUE)
					case "VOL": // Main-Zone Volume
						d.device.Zones[0].Volume = int32(float32((toInt(VALUE)+1)*100) / maxVol)
					case "XV": // Zone 2 Volume
						d.device.Zones[1].Volume = int32(float32((toInt(VALUE)+1)*100) / maxHZ)
					case "YV": // Zone 3 Volume
						// d.device.Zones[2].Volume = int32(float32((toInt(VALUE)+1)*100) / maxHZ)
					case "MUT":
						d.device.Zones[0].Muted = toBool(VALUE)
					case "HZMUT":
						d.device.Zones[0].Muted = toBool(VALUE)
					case "SR":
						d.device.Zones[0].CurrentListeningMod = VALUE
					case "FN":
						d.device.Zones[0].CurrentSource = VALUE
					default:
						return
					}
					for _, fn := range d.updateFuncs {
						fn(&d.device)
					}
				}(cmd)
			}
		}
	}()
}

func toBool(val string) bool {
	return (toInt(val) > 0)
}
func toInt(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}
