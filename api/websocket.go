package api

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

func WebsocketHandler(conn *websocket.Conn) {
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("We read", n, "bytes, and they were", string(buf))
	}
}
