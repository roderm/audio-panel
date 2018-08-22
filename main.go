package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	jws "github.com/roderm/json-rpc/websocket"
	"golang.org/x/net/websocket"
	"time"
)

func main() {

	wsHandler := jws.NewHandler()
	wsHandler.Add("sayhello", func(params interface{}) interface{} {
		type myHelloMsg struct {
			Text string `json:"text"`
		}
		return myHelloMsg{Text: "Hello you"}
	})

	wsHandler.AddSubscription("foo", func(params interface{}) chan interface{} {
		type bar struct {
			Text string
		}
		mychan := make(chan interface{})
		go func() {
			for {
				time.Sleep(time.Second * 10)
				mychan <- bar{Text: "I'm still here"}
			}
		}()
		return mychan
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
