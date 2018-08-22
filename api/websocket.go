package api

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
)

type Message struct {
	RequestID string      `json:"requestId"`
	Method    string      `json:"method"`
	Action    string      `json:"action"`
	Data      interface{} `json:"data"`
}

type Response struct {
	RequestID string      `json:"requestId"`
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
}

type WebsocketHandler struct {
	clients []*websocket.Conn
	routes  map[string]func(*Message) interface{}
	subs    map[string]func(*Message) chan interface{}
}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{
		routes: make(map[string]func(*Message) interface{}),
		subs:   make(map[string]func(*Message) chan interface{}),
	}
}

func (wh *WebsocketHandler) Add(action string, cb func(*Message) interface{}) {
	wh.routes[action] = cb
}
func (wh *WebsocketHandler) AddSubscription(action string, cb func(*Message) chan interface{}) {
	wh.subs[action] = cb
}
func (wh *WebsocketHandler) execRequest(m *Message, conn *websocket.Conn) {
	fmt.Println(m)
	if wh.routes[m.Action] != nil {
		data := wh.routes[m.Action](m)
		r := Response{
			RequestID: m.RequestID,
			ErrorCode: 0,
			Data:      data,
		}
		wh.WriteToClient(r, conn)
	} else if wh.subs[m.Action] != nil {
		subChan := wh.subs[m.Action](m)
		for {
			data := <-subChan
			r := Response{
				RequestID: m.RequestID,
				ErrorCode: 0,
				Data:      data,
			}
			wh.WriteToClient(r, conn)
		}
	} else {
		fmt.Errorf("No action found")
	}
}
func (wh *WebsocketHandler) WriteToClient(r Response, conn *websocket.Conn) {
	json_str, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	_, err = conn.Write(json_str)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func (wh *WebsocketHandler) Handle(conn *websocket.Conn) {
	for {
		var msg = make([]byte, 1024)
		n, err := conn.Read(msg)
		if err == io.EOF {
			conn.Close()
			return
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		message := &Message{}
		err = json.Unmarshal(msg[:n], message)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go wh.execRequest(message, conn)
	}
}

/*
func SubscribeDevice(dev device.IDevice) {
	avrs = append(avrs, dev)
	dev.OnUpdate(func() {
		avr := dev.GetAvr()
		avrJs, err := json.Marshal(avr)
		if err != nil {
			return
		}
		for _, cl := range clients {
			_, err = cl.Write(avrJs)
			if err != nil {
				fmt.Errorf(err.Error())
			}

		}
	})
}
*/
