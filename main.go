package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/roderm/audio-panel/api"
	"github.com/roderm/audio-panel/device"
	jws "github.com/roderm/json-rpc/websocket"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"os"
)

type Config struct {
	Devices []device.DeviceConfig `json:"devices"`
}

func main() {
	conf := readConfig()
	ds := device.NewDeviceStore(context.Background())
	for _, dev := range conf.Devices {
		go ds.AddDevice(dev)
	}
	wsHandler := jws.NewHandler(context.Background())
	api.AddDeviceApi(wsHandler, ds)
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	// Setup route group for the API
	router.GET("/api", func(c *gin.Context) {
		handler := websocket.Handler(wsHandler.Handle)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// Start and run the server
	router.Run(":3000")
}

func readConfig() Config {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		panic("no config file given")
	}
	configFile := argsWithoutProg[0]
	if _, err := os.Stat(configFile); err != nil {
		panic(fmt.Sprintf("path to configfile(%s) not found.", configFile))
	}
	fileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("Couldn't read configfile(%s).", configFile))
	}
	var conf Config
	err = json.Unmarshal(fileContent, &conf)
	if err != nil {
		panic(fmt.Sprintf("Bad Syntax in configfile(%s), wasn't able to parse json.", configFile))
	}
	return conf
}
