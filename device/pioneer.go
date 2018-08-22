package device

import (
	"fmt"
	pb "github.com/roderm/audio-panel/proto"
	"github.com/roderm/audio-panel/telnet"
	"regexp"
	"strconv"
)

const maxVol = 160 // 185
const maxHZ = 80

type PioneerDevice struct {
	device      pb.AVR
	nc          *telnet.PioneerCaller
	updateFuncs []func()
}

func NewPioneerDevice(ip string) IDevice {
	p := PioneerDevice{}
	p.nc, _ = telnet.NewPioneerCaller(ip, 8102)
	p.device = pb.AVR{
		IP:    ip,
		Zones: make(map[int32]*pb.AVR_Zone),
	}
	p.device.Zones[0] = &pb.AVR_Zone{
		IsMain: true,
		Name:   "main",
		Volume: 80,
	}
	p.device.Zones[1] = &pb.AVR_Zone{
		IsMain: true,
		Name:   "hdzone",
		Volume: 80,
	}
	p.initCommands()
	return &p
}

func (d *PioneerDevice) initCommands() {
	d.startListener()
	d.nc.Send("?V")
	d.nc.Send("?PWR")
	d.nc.Send("?HZV")
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

func (d *PioneerDevice) SetSource(zone int32, src int32) {

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

func (d *PioneerDevice) OnUpdate(fn func()) {
	d.updateFuncs = append(d.updateFuncs, fn)
}
func (d *PioneerDevice) GetAvr() *pb.AVR {
	return &d.device
}

func (d *PioneerDevice) startListener() {
	d.nc.StartListen()
	go func() {
		for {
			cmd := <-d.nc.RecCommands
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
			default:
				fmt.Printf("Unknown command received %s \n", COMMAND)
				break
			}
			go func() {
				for _, fn := range d.updateFuncs {
					fn()
				}
			}()
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
