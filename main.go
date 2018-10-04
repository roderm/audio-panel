package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/roderm/audio-panel/device"
	"github.com/roderm/audio-panel/grpc"
)

type Config struct {
	TplPath string                `json:"viewspath"`
	Devices []device.DeviceConfig `json:"devices"`
}

func main() {
	conf := readConfig()
	ds := device.NewDeviceStore(context.Background())
	for _, dev := range conf.Devices {
		go func(dev device.DeviceConfig) {
			id, err := ds.AddDevice(dev)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("New device added with id %s \n", id)
			}
		}(dev)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	grpcServer := grpc.NewGrpcInstance()
	grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	/*
		wsHandler := jws.NewHandler(context.Background())
		api.AddApi(wsHandler, ds)
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
	*/
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
