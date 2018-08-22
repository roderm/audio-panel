package api

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

type Request struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      interface{} `json:"id"`
}

type Response struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   Error       `json:"error"`
	Id      interface{} `json:"id"`
}
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WebsocketHandler struct {
	clients []*websocket.Conn
	routes  map[string]func(interface{}) interface{}
	subs    map[string]func(interface{}) chan interface{}
}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{
		routes: make(map[string]func(interface{}) interface{}),
		subs:   make(map[string]func(interface{}) chan interface{}),
	}
}

func (wh *WebsocketHandler) Add(action string, cb func(interface{}) interface{}) {
	wh.routes[action] = cb
}
func (wh *WebsocketHandler) AddSubscription(action string, cb func(interface{}) chan interface{}) {
	wh.subs[action] = cb
}
func (wh *WebsocketHandler) execRequest(m *Request, conn *websocket.Conn) {
	if wh.routes[m.Method] != nil {
		data := wh.routes[m.Method](m.Params)
		r := Response{
			Id:     m.Id,
			Result: data,
		}
		wh.WriteToClient(r, conn)
	} else if wh.subs[m.Method] != nil {
		subChan := wh.subs[m.Method](m)
		for {
			data := <-subChan
			r := Response{
				Id:     m.Id,
				Result: data,
			}
			wh.WriteToClient(r, conn)
		}
	} else {
		log.Println("No action found")
	}
}
func (wh *WebsocketHandler) WriteToClient(r Response, conn *websocket.Conn) {
	json_str, err := json.Marshal(r)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = conn.Write(json_str)
	if err != nil {
		log.Println(err.Error())
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
			log.Println(err.Error())
			continue
		}
		message := &Request{}
		err = json.Unmarshal(msg[:n], message)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go wh.execRequest(message, conn)
	}
}
