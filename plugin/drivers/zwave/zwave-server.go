package main

import (
	"fmt"
	"context"
	"log"

	"github.com/stampzilla/gozwave"
	"github.com/stampzilla/gozwave/events"
	pb "github.com/roderm/audio-panel/proto"
)

type ZwaveServerConfig struct {
	Port string
}
type ZwaveServer struct {
	controller *gozwave.Controller
	ctx        context.Context
}

func NewZwaveServer(ctx contex.Context, config ZwaveServerConfig, lgr log.Logger, identifier string) {
	srv := &ZwaveServer{
		controller: z,
		ctx: ctx
	}
	srv.controller, err := gozwave.Connect(config.Port, "")
	if err != nil {
		return srv, err
	}
	srv.listen()
	return srv, err
}

func (s *ZwaveServer) listen() {
	go func() {
		for {
			select {
			case event := <-s.controller.GetNextEvent():
				switch e := event.(type) {
				case events.NodeDiscoverd:
					znode := s.controller.Nodes.Get(e.Address)
					znode.RLock()
					log.Printf("Node: %#v\n", znode)
					znode.RUnlock()

				case events.NodeUpdated:
					znode := s.controller.Nodes.Get(e.Address)
					znode.RLock()
					log.Printf("Node: %#v\n", znode)
					znode.RUnlock()
				}
			}
		}
	}()
}
func (s *ZwaveServer) getDevice(ident string) (*gozwave.Node, error) {
	id, err := strconv.Atoi(ident)
	if err != nil {
		return nil, err
	}
	node := s.controller.Nodes.Get(id)
	if node == nil {
		return nil, fmt.Errorf("Node not found.")
	}
	return s.controller.Nodes.Get(id), nil
}
func (s *ZwaveServer) PropertyUpdate(u *pb.PropertyUpdate) error{
	d, err := s.getDevice(u.ItemIdentifier)
}
func (s *ZwaveServer) OnUpdate(cb func(*pb.PropertyUpdate)){

}
func (s *ZwaveServer) GetDevice() *pb.Device{

}