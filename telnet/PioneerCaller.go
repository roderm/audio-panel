package telnet

import (
	"context"
	"fmt"
	"github.com/ziutek/telnet"
	"strconv"
	"strings"
	"time"
)

type PioneerCaller struct {
	ctx          context.Context
	sendCommands chan string
	RecCommands  chan string
	conn         *telnet.Conn
	connected    bool
}

func NewPioneerCaller(ctx context.Context, host string, port int) (*PioneerCaller, error) {
	var err error
	ret := &PioneerCaller{
		ctx:          ctx,
		sendCommands: make(chan string),
		RecCommands:  make(chan string)}
	addr := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))
	fmt.Printf("Address to call: %s \n", addr)
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
					c.RecCommands <- strings.TrimSpace(msg)
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

func (c *PioneerCaller) Send(command string) {
	c.sendCommands <- command
}

func (c *PioneerCaller) unload() {
	fmt.Println("Dead of telnet")
	c.conn.Close()
}
