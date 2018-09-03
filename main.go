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
	"path/filepath"
)

type Config struct {
	TplPath string                `json:"viewspath"`
	Devices []device.DeviceConfig `json:"devices"`
}

func main() {
	conf := readConfig()
	ds := device.NewDeviceStore(context.Background())
	for _, dev := range conf.Devices {
		go func() {
			_, err := ds.AddDevice(dev)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wsHandler := jws.NewHandler(context.Background())
	api.AddDeviceApi(wsHandler, ds)
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	// Serve frontend static files
	tpls, _ := os.Stat(conf.TplPath)
	if tpls.IsDir() {
		router.Use(static.Serve("/", static.LocalFile(conf.TplPath, true)))
	} else {
		fmt.Println("No filesystem is served")
	}
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

func getPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath

}
