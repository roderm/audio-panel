package grpc

import (
	"context"

	pb "github.com/roderm/audio-panel-protobuf/go/box_controll"
	"github.com/roderm/audio-panel-protobuf/go/msg/device"
	"github.com/roderm/audio-panel-protobuf/go/msg/filters"
	"google.golang.org/grpc"
)

func NewGrpcInstance() *grpc.Server {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBoxControllServer(grpcServer, &grpcServerApi{})
	return grpcServer
}

type grpcServerApi struct {
}

func (s *grpcServerApi) GetBoxes(f *filters.BoxFilter, stream pb.BoxControll_GetBoxesServer) error {
	return nil
}

func (s *grpcServerApi) GetDevices(f *filters.DeviceFilter, stream pb.BoxControll_GetDevicesServer) error {
	return nil
}

func (s *grpcServerApi) OnUpdate(f *filters.UpdateFilter, stream pb.BoxControll_OnUpdateServer) error {
	return nil
}

func (s *grpcServerApi) StateUpdate(context.Context, *device.UpdateReq) (*device.UpdateRes, error) {
	return nil, nil
}
