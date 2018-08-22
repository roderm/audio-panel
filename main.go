package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/roderm/audio-panel/api"
	"golang.org/x/net/websocket"
)

func main() {

	wsHandler := api.NewWebsocketHandler()
	wsHandler.Add("sayhello", func(m *api.Message) interface{} {
		type myHelloMsg struct {
			Text string `json:"text"`
		}
		return myHelloMsg{Text: "Hello you"}
	})
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
