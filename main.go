package main

import (
	"context"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/roderm/audio-panel/api"
	"github.com/roderm/audio-panel/device"
	jws "github.com/roderm/json-rpc/websocket"
	"golang.org/x/net/websocket"
)

func main() {

	ds := device.NewDeviceStore(context.Background())
	go ds.AddDevice(device.DeviceConfig{
		DeviceType:    "pioneer",
		DeviceAddress: "192.168.178.28",
	})
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
