package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	pl "github.com/roderm/audio-panel/plugin/iface"
	"net/http"
	"strings"
	"time"
)

func main() {
	ip := flag.String("ip", "", "Lights IP address")
	flag.Parse()
	if len(*ip) > 0 {
		getAuth(*ip)
	}
}
func getAuth(ip string) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Post(fmt.Sprintf("http://%s/api/v1/new", ip), "application/json", strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func NewDriver(ctx context.Context, config interface{}, id string) (pl.IDevice, error) {
	var conf nanoLeafConfig
	err := mapstructure.Decode(config, &conf)
	if err != nil {
		return nil, err
	}
	return NewNanoleafDriver(ctx, conf, id), nil
}

// curl -d "{}" -H "Content-Type: application/json" -X POST http://192.168.178.38/api/v1/new
