package telnet

import (
	"fmt"
	"github.com/ziutek/telnet"
	"strconv"
	"strings"
	"time"
)

type PioneerCaller struct {
	done         chan bool
	sendCommands chan string
	RecCommands  chan string
	conn         *telnet.Conn
}

func NewPioneerCaller(host string, port int) (*PioneerCaller, error) {
	var err error
	ret := &PioneerCaller{
		done:         make(chan bool),
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
			case <-c.done:
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
			case <-c.done:
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
					fmt.Errorf(err.Error())
					return
				}
				time.Sleep(time.Millisecond * 100)
			}

		}
	}()
}

func (c *PioneerCaller) Send(command string) {
	c.sendCommands <- command
}

func (c *PioneerCaller) Unload() {
	c.done <- true
	c.conn.Close()
}
