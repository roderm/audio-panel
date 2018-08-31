package telnet

import (
	"context"
	"fmt"
	"github.com/ziutek/telnet"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type subscription struct {
	id int
	cb func(interface{})
}
type PioneerCaller struct {
	ctx          context.Context
	sendCommands chan string
	conn         *telnet.Conn
	connected    bool
	subs         map[string][]subscription
	subsIds      int
	idLock       sync.Mutex
}

func NewPioneerCaller(ctx context.Context, host string, port int) (*PioneerCaller, error) {
	var err error
	ret := &PioneerCaller{
		ctx:          ctx,
		subs:         make(map[string][]subscription),
		sendCommands: make(chan string),
	}
	addr := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))
	ret.conn, err = telnet.Dial("tcp", addr)
	return ret, err
}
func (c *PioneerCaller) StartListen() {
	go func() {
		buf := make([]byte, 512)
		for {
			select {
			case <-c.ctx.Done():
				c.unload()
				return
			default:
				n, _ := c.conn.Read(buf) // Use raw read to find issue #15.
				msg := string(buf[:n])
				go func() {
					var expression = regexp.MustCompile(`(?P<COMMAND>^[A-Z]+[\d]??[A-Z]+)(?P<VALUE>[\d]{1,})`)
					COMMAND := strings.TrimSpace(expression.ReplaceAllString(string(msg), "${COMMAND}"))
					VALUE := strings.TrimSpace(expression.ReplaceAllString(string(msg), "${VALUE}"))
					placeholder := ""
					for i := 0; i < len(VALUE); i++ {
						placeholder += "#"
					}
					resp := fmt.Sprintf("%s%s", COMMAND, placeholder)
					for _, s := range c.subs[resp] {
						s.cb(VALUE)
					}
				}()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-c.ctx.Done():
				c.unload()
				return
			case cmd := <-c.sendCommands:
				command := []byte(fmt.Sprintf("%s\n\r", cmd))
				n, err := c.conn.Write(command)
				if err != nil {
					fmt.Println(err)
					break
				}
				if expected, actual := int64(len(command)), n; int(expected) != actual {
					err := fmt.Errorf("Transmission problem: tried sending %d bytes, but actually only sent %d bytes.", expected, actual)
					fmt.Printf(err.Error())
					return
				}
				time.Sleep(time.Second)
			}

		}
	}()
}

func (c *PioneerCaller) Once(response string, cb func(interface{})) {
	s := c.getNewSub()
	s.cb = func(v interface{}) {
		for i, sub := range c.subs[response] {
			if sub.id == s.id {
				c.subs[response] = append(c.subs[response][:i], c.subs[response][i+1])
			}
		}
		cb(v)
	}
	c.subs[response] = append(c.subs[response], s)
}

func (c *PioneerCaller) Subscribe(response string, cb func(interface{})) func() {
	s := c.getNewSub()
	s.cb = cb
	f := func() {
		for i, sub := range c.subs[response] {
			if sub.id == s.id {
				c.subs[response] = append(c.subs[response][:i], c.subs[response][i+1])
			}
		}
	}
	c.subs[response] = append(c.subs[response], s)
	return f
}
func (c *PioneerCaller) Send(command string) {
	c.sendCommands <- command
}

func (c *PioneerCaller) unload() {
	fmt.Println("Dead of telnet")
	c.conn.Close()
}

func (c *PioneerCaller) getNewSub() subscription {
	c.idLock.Lock()
	defer c.idLock.Unlock()
	c.subsIds += 1
	return subscription{
		id: c.subsIds,
	}

}
