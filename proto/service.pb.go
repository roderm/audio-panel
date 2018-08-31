// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package proto // import "github.com/roderm/audio-panel/proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5162c1a8919e83af, []int{0}
}
func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (dst *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(dst, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type AVR struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id" json:"id"`
	Power                bool                 `protobuf:"varint,2,opt,name=power" json:"power"`
	Zones                map[string]*AVR_Zone `protobuf:"bytes,3,rep,name=zones" json:"zones" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	InputSources         map[string]string    `protobuf:"bytes,4,rep,name=inputSources" json:"inputSources" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AVR) Reset()         { *m = AVR{} }
func (m *AVR) String() string { return proto.CompactTextString(m) }
func (*AVR) ProtoMessage()    {}
func (*AVR) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5162c1a8919e83af, []int{1}
}
func (m *AVR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AVR.Unmarshal(m, b)
}
func (m *AVR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AVR.Marshal(b, m, deterministic)
}
func (dst *AVR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AVR.Merge(dst, src)
}
func (m *AVR) XXX_Size() int {
	return xxx_messageInfo_AVR.Size(m)
}
func (m *AVR) XXX_DiscardUnknown() {
	xxx_messageInfo_AVR.DiscardUnknown(m)
}

var xxx_messageInfo_AVR proto.InternalMessageInfo

func (m *AVR) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AVR) GetPower() bool {
	if m != nil {
		return m.Power
	}
	return false
}

func (m *AVR) GetZones() map[string]*AVR_Zone {
	if m != nil {
		return m.Zones
	}
	return nil
}

func (m *AVR) GetInputSources() map[string]string {
	if m != nil {
		return m.InputSources
	}
	return nil
}

type AVR_Zone struct {
	Power                bool                     `protobuf:"varint,1,opt,name=power" json:"power"`
	Name                 string                   `protobuf:"bytes,2,opt,name=name" json:"name"`
	Volume               int32                    `protobuf:"varint,3,opt,name=volume" json:"volume"`
	IsMain               bool                     `protobuf:"varint,4,opt,name=isMain" json:"isMain"`
	CurrentSource        string                   `protobuf:"bytes,5,opt,name=currentSource" json:"currentSource"`
	SyncVol              bool                     `protobuf:"varint,6,opt,name=syncVol" json:"syncVol"`
	SyncSrc              bool                     `protobuf:"varint,7,opt,name=syncSrc" json:"syncSrc"`
	Mute                 bool                     `protobuf:"varint,8,opt,name=mute" json:"mute"`
	CurrentListeningMod  string                   `protobuf:"bytes,9,opt,name=currentListeningMod" json:"currentListeningMod"`
	ListeningMods        []*AVR_Zone_ListeningMod `protobuf:"bytes,10,rep,name=listeningMods" json:"listeningMods"`
	Sources              []*AVR_Zone_Source       `protobuf:"bytes,11,rep,name=sources" json:"sources"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AVR_Zone) Reset()         { *m = AVR_Zone{} }
func (m *AVR_Zone) String() string { return proto.CompactTextString(m) }
func (*AVR_Zone) ProtoMessage()    {}
func (*AVR_Zone) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5162c1a8919e83af, []int{1, 0}
}
func (m *AVR_Zone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AVR_Zone.Unmarshal(m, b)
}
func (m *AVR_Zone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AVR_Zone.Marshal(b, m, deterministic)
}
func (dst *AVR_Zone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AVR_Zone.Merge(dst, src)
}
func (m *AVR_Zone) XXX_Size() int {
	return xxx_messageInfo_AVR_Zone.Size(m)
}
func (m *AVR_Zone) XXX_DiscardUnknown() {
	xxx_messageInfo_AVR_Zone.DiscardUnknown(m)
}

var xxx_messageInfo_AVR_Zone proto.InternalMessageInfo

func (m *AVR_Zone) GetPower() bool {
	if m != nil {
		return m.Power
	}
	return false
}

func (m *AVR_Zone) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AVR_Zone) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *AVR_Zone) GetIsMain() bool {
	if m != nil {
		return m.IsMain
	}
	return false
}

func (m *AVR_Zone) GetCurrentSource() string {
	if m != nil {
		return m.CurrentSource
	}
	return ""
}

func (m *AVR_Zone) GetSyncVol() bool {
	if m != nil {
		return m.SyncVol
	}
	return false
}

func (m *AVR_Zone) GetSyncSrc() bool {
	if m != nil {
		return m.SyncSrc
	}
	return false
}

func (m *AVR_Zone) GetMute() bool {
	if m != nil {
		return m.Mute
	}
	return false
}

func (m *AVR_Zone) GetCurrentListeningMod() string {
	if m != nil {
		return m.CurrentListeningMod
	}
	return ""
}

func (m *AVR_Zone) GetListeningMods() []*AVR_Zone_ListeningMod {
	if m != nil {
		return m.ListeningMods
	}
	return nil
}

func (m *AVR_Zone) GetSources() []*AVR_Zone_Source {
	if m != nil {
		return m.Sources
	}
	return nil
}

type AVR_Zone_ListeningMod struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code" json:"Code"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AVR_Zone_ListeningMod) Reset()         { *m = AVR_Zone_ListeningMod{} }
func (m *AVR_Zone_ListeningMod) String() string { return proto.CompactTextString(m) }
func (*AVR_Zone_ListeningMod) ProtoMessage()    {}
func (*AVR_Zone_ListeningMod) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5162c1a8919e83af, []int{1, 0, 0}
}
func (m *AVR_Zone_ListeningMod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AVR_Zone_ListeningMod.Unmarshal(m, b)
}
func (m *AVR_Zone_ListeningMod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AVR_Zone_ListeningMod.Marshal(b, m, deterministic)
}
func (dst *AVR_Zone_ListeningMod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AVR_Zone_ListeningMod.Merge(dst, src)
}
func (m *AVR_Zone_ListeningMod) XXX_Size() int {
	return xxx_messageInfo_AVR_Zone_ListeningMod.Size(m)
}
func (m *AVR_Zone_ListeningMod) XXX_DiscardUnknown() {
	xxx_messageInfo_AVR_Zone_ListeningMod.DiscardUnknown(m)
}

var xxx_messageInfo_AVR_Zone_ListeningMod proto.InternalMessageInfo

func (m *AVR_Zone_ListeningMod) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AVR_Zone_ListeningMod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AVR_Zone_Source struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code" json:"Code"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AVR_Zone_Source) Reset()         { *m = AVR_Zone_Source{} }
func (m *AVR_Zone_Source) String() string { return proto.CompactTextString(m) }
func (*AVR_Zone_Source) ProtoMessage()    {}
func (*AVR_Zone_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5162c1a8919e83af, []int{1, 0, 1}
}
func (m *AVR_Zone_Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AVR_Zone_Source.Unmarshal(m, b)
}
func (m *AVR_Zone_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AVR_Zone_Source.Marshal(b, m, deterministic)
}
func (dst *AVR_Zone_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AVR_Zone_Source.Merge(dst, src)
}
func (m *AVR_Zone_Source) XXX_Size() int {
	return xxx_messageInfo_AVR_Zone_Source.Size(m)
}
func (m *AVR_Zone_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_AVR_Zone_Source.DiscardUnknown(m)
}

var xxx_messageInfo_AVR_Zone_Source proto.InternalMessageInfo

func (m *AVR_Zone_Source) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AVR_Zone_Source) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AVR_PowerToggle struct {
	IP                   string   `protobuf:"bytes,1,opt,name=IP" json:"IP"`
	On                   bool     `protobuf:"varint,2,opt,name=On" json:"On"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AVR_PowerToggle) Reset()         { *m = AVR_PowerToggle{} }
func (m *AVR_PowerToggle) String() string { return proto.CompactTextString(m) }
func (*AVR_PowerToggle) ProtoMessage()    {}
func (*AVR_PowerToggle) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5162c1a8919e83af, []int{2}
}
func (m *AVR_PowerToggle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AVR_PowerToggle.Unmarshal(m, b)
}
func (m *AVR_PowerToggle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AVR_PowerToggle.Marshal(b, m, deterministic)
}
func (dst *AVR_PowerToggle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AVR_PowerToggle.Merge(dst, src)
}
func (m *AVR_PowerToggle) XXX_Size() int {
	return xxx_messageInfo_AVR_PowerToggle.Size(m)
}
func (m *AVR_PowerToggle) XXX_DiscardUnknown() {
	xxx_messageInfo_AVR_PowerToggle.DiscardUnknown(m)
}

var xxx_messageInfo_AVR_PowerToggle proto.InternalMessageInfo

func (m *AVR_PowerToggle) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *AVR_PowerToggle) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

type ZoneUpdate struct {
	AVR_IP               string   `protobuf:"bytes,1,opt,name=AVR_IP,json=AVRIP" json:"AVR_IP"`
	Zone_Name            string   `protobuf:"bytes,2,opt,name=Zone_Name,json=ZoneName" json:"Zone_Name"`
	State                bool     `protobuf:"varint,3,opt,name=State" json:"State"`
	Volume               int32    `protobuf:"varint,4,opt,name=Volume" json:"Volume"`
	Source               int32    `protobuf:"varint,5,opt,name=Source" json:"Source"`
	SyncVol              bool     `protobuf:"varint,6,opt,name=SyncVol" json:"SyncVol"`
	SyncSrc              bool     `protobuf:"varint,7,opt,name=SyncSrc" json:"SyncSrc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZoneUpdate) Reset()         { *m = ZoneUpdate{} }
func (m *ZoneUpdate) String() string { return proto.CompactTextString(m) }
func (*ZoneUpdate) ProtoMessage()    {}
func (*ZoneUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5162c1a8919e83af, []int{3}
}
func (m *ZoneUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZoneUpdate.Unmarshal(m, b)
}
func (m *ZoneUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZoneUpdate.Marshal(b, m, deterministic)
}
func (dst *ZoneUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZoneUpdate.Merge(dst, src)
}
func (m *ZoneUpdate) XXX_Size() int {
	return xxx_messageInfo_ZoneUpdate.Size(m)
}
func (m *ZoneUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ZoneUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ZoneUpdate proto.InternalMessageInfo

func (m *ZoneUpdate) GetAVR_IP() string {
	if m != nil {
		return m.AVR_IP
	}
	return ""
}

func (m *ZoneUpdate) GetZone_Name() string {
	if m != nil {
		return m.Zone_Name
	}
	return ""
}

func (m *ZoneUpdate) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

func (m *ZoneUpdate) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *ZoneUpdate) GetSource() int32 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *ZoneUpdate) GetSyncVol() bool {
	if m != nil {
		return m.SyncVol
	}
	return false
}

func (m *ZoneUpdate) GetSyncSrc() bool {
	if m != nil {
		return m.SyncSrc
	}
	return false
}

func init() {
	proto.RegisterType((*Null)(nil), "null")
	proto.RegisterType((*AVR)(nil), "AVR")
	proto.RegisterMapType((map[string]string)(nil), "AVR.InputSourcesEntry")
	proto.RegisterMapType((map[string]*AVR_Zone)(nil), "AVR.ZonesEntry")
	proto.RegisterType((*AVR_Zone)(nil), "AVR.Zone")
	proto.RegisterType((*AVR_Zone_ListeningMod)(nil), "AVR.Zone.ListeningMod")
	proto.RegisterType((*AVR_Zone_Source)(nil), "AVR.Zone.Source")
	proto.RegisterType((*AVR_PowerToggle)(nil), "AVR_PowerToggle")
	proto.RegisterType((*ZoneUpdate)(nil), "ZoneUpdate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AVRService service

type AVRServiceClient interface {
	GetAVRs(ctx context.Context, in *Null, opts ...grpc.CallOption) (AVRService_GetAVRsClient, error)
	Power(ctx context.Context, in *AVR_PowerToggle, opts ...grpc.CallOption) (*AVR, error)
	UpdateZone(ctx context.Context, in *ZoneUpdate, opts ...grpc.CallOption) (*ZoneUpdate, error)
	Subscribe(ctx context.Context, in *AVR, opts ...grpc.CallOption) (AVRService_SubscribeClient, error)
}

type aVRServiceClient struct {
	cc *grpc.ClientConn
}

func NewAVRServiceClient(cc *grpc.ClientConn) AVRServiceClient {
	return &aVRServiceClient{cc}
}

func (c *aVRServiceClient) GetAVRs(ctx context.Context, in *Null, opts ...grpc.CallOption) (AVRService_GetAVRsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AVRService_serviceDesc.Streams[0], c.cc, "/AVRService/GetAVRs", opts...)
	if err != nil {
		return nil, err
	}
	x := &aVRServiceGetAVRsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AVRService_GetAVRsClient interface {
	Recv() (*AVR, error)
	grpc.ClientStream
}

type aVRServiceGetAVRsClient struct {
	grpc.ClientStream
}

func (x *aVRServiceGetAVRsClient) Recv() (*AVR, error) {
	m := new(AVR)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aVRServiceClient) Power(ctx context.Context, in *AVR_PowerToggle, opts ...grpc.CallOption) (*AVR, error) {
	out := new(AVR)
	err := grpc.Invoke(ctx, "/AVRService/Power", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aVRServiceClient) UpdateZone(ctx context.Context, in *ZoneUpdate, opts ...grpc.CallOption) (*ZoneUpdate, error) {
	out := new(ZoneUpdate)
	err := grpc.Invoke(ctx, "/AVRService/UpdateZone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aVRServiceClient) Subscribe(ctx context.Context, in *AVR, opts ...grpc.CallOption) (AVRService_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AVRService_serviceDesc.Streams[1], c.cc, "/AVRService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &aVRServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AVRService_SubscribeClient interface {
	Recv() (*AVR, error)
	grpc.ClientStream
}

type aVRServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *aVRServiceSubscribeClient) Recv() (*AVR, error) {
	m := new(AVR)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AVRService service

type AVRServiceServer interface {
	GetAVRs(*Null, AVRService_GetAVRsServer) error
	Power(context.Context, *AVR_PowerToggle) (*AVR, error)
	UpdateZone(context.Context, *ZoneUpdate) (*ZoneUpdate, error)
	Subscribe(*AVR, AVRService_SubscribeServer) error
}

func RegisterAVRServiceServer(s *grpc.Server, srv AVRServiceServer) {
	s.RegisterService(&_AVRService_serviceDesc, srv)
}

func _AVRService_GetAVRs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Null)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AVRServiceServer).GetAVRs(m, &aVRServiceGetAVRsServer{stream})
}

type AVRService_GetAVRsServer interface {
	Send(*AVR) error
	grpc.ServerStream
}

type aVRServiceGetAVRsServer struct {
	grpc.ServerStream
}

func (x *aVRServiceGetAVRsServer) Send(m *AVR) error {
	return x.ServerStream.SendMsg(m)
}

func _AVRService_Power_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AVR_PowerToggle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AVRServiceServer).Power(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AVRService/Power",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AVRServiceServer).Power(ctx, req.(*AVR_PowerToggle))
	}
	return interceptor(ctx, in, info, handler)
}

func _AVRService_UpdateZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AVRServiceServer).UpdateZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AVRService/UpdateZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AVRServiceServer).UpdateZone(ctx, req.(*ZoneUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _AVRService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AVR)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AVRServiceServer).Subscribe(m, &aVRServiceSubscribeServer{stream})
}

type AVRService_SubscribeServer interface {
	Send(*AVR) error
	grpc.ServerStream
}

type aVRServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *aVRServiceSubscribeServer) Send(m *AVR) error {
	return x.ServerStream.SendMsg(m)
}

var _AVRService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AVRService",
	HandlerType: (*AVRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Power",
			Handler:    _AVRService_Power_Handler,
		},
		{
			MethodName: "UpdateZone",
			Handler:    _AVRService_UpdateZone_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAVRs",
			Handler:       _AVRService_GetAVRs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _AVRService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_service_5162c1a8919e83af) }

var fileDescriptor_service_5162c1a8919e83af = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0x56, 0x9a, 0x3f, 0x6b, 0x4e, 0xb7, 0xdf, 0xf6, 0xf3, 0xc6, 0x64, 0xca, 0x05, 0x55, 0x61,
	0xa8, 0x42, 0x22, 0x1d, 0x43, 0x42, 0x68, 0x42, 0x42, 0x61, 0x42, 0xa8, 0x12, 0x63, 0x95, 0x03,
	0xb9, 0xd8, 0xcd, 0x94, 0x26, 0x56, 0x89, 0x48, 0xed, 0xca, 0x49, 0x86, 0xc6, 0x23, 0xf0, 0x48,
	0x3c, 0x07, 0x6f, 0xc1, 0x4b, 0x20, 0xdb, 0xe9, 0xe2, 0x88, 0x5d, 0x70, 0x55, 0x7f, 0xe7, 0x7c,
	0xdf, 0xf1, 0xc9, 0xa7, 0xcf, 0x85, 0xfd, 0xb5, 0xe0, 0x15, 0x9f, 0x96, 0x54, 0x5c, 0xe7, 0x29,
	0x0d, 0x14, 0x1a, 0x7b, 0xe0, 0xb0, 0xba, 0x28, 0xc6, 0xbf, 0x5d, 0xb0, 0xc3, 0x98, 0xa0, 0xff,
	0xa0, 0x97, 0x67, 0xd8, 0x1a, 0x59, 0x13, 0x9b, 0xf4, 0xf2, 0x0c, 0x1d, 0x80, 0xbb, 0xe6, 0xdf,
	0xa8, 0xc0, 0xbd, 0x91, 0x35, 0xe9, 0x13, 0x0d, 0xd0, 0x11, 0xb8, 0xdf, 0x39, 0xa3, 0x25, 0xb6,
	0x47, 0xf6, 0x64, 0x70, 0xb2, 0x1b, 0x84, 0x31, 0x09, 0x2e, 0x65, 0xe5, 0x1d, 0xab, 0xc4, 0x0d,
	0xd1, 0x5d, 0x74, 0x0a, 0xdb, 0x39, 0x5b, 0xd7, 0x55, 0xc4, 0x6b, 0x91, 0xd2, 0x12, 0x3b, 0x8a,
	0x7d, 0xa8, 0xd8, 0x33, 0xa3, 0xa1, 0x45, 0x1d, 0xee, 0xf0, 0x97, 0x0d, 0x8e, 0x9c, 0xd8, 0x6e,
	0x60, 0x99, 0x1b, 0x20, 0x70, 0x58, 0xb2, 0xa2, 0x6a, 0x2d, 0x9f, 0xa8, 0x33, 0x3a, 0x04, 0xef,
	0x9a, 0x17, 0xf5, 0x8a, 0x62, 0x7b, 0x64, 0x4d, 0x5c, 0xd2, 0x20, 0x59, 0xcf, 0xcb, 0xf3, 0x24,
	0x67, 0xd8, 0x51, 0x23, 0x1a, 0x84, 0x1e, 0xc3, 0x4e, 0x5a, 0x0b, 0x41, 0x59, 0x73, 0x29, 0x76,
	0xd5, 0xb0, 0x6e, 0x11, 0x61, 0xd8, 0x2a, 0x6f, 0x58, 0x1a, 0xf3, 0x02, 0x7b, 0x4a, 0xbe, 0x81,
	0x9b, 0x4e, 0x24, 0x52, 0xbc, 0xd5, 0x76, 0x22, 0x91, 0xca, 0xed, 0x56, 0x75, 0x45, 0x71, 0x5f,
	0x95, 0xd5, 0x19, 0x1d, 0xc3, 0x7e, 0x33, 0xf8, 0x43, 0x5e, 0x56, 0x94, 0xe5, 0x6c, 0x79, 0xce,
	0x33, 0xec, 0xab, 0x3b, 0xef, 0x6a, 0xa1, 0xd7, 0xb0, 0x53, 0x18, 0xb8, 0xc4, 0x60, 0xf8, 0x27,
	0xbd, 0x09, 0x4c, 0x3a, 0xe9, 0x92, 0xd1, 0x53, 0xd8, 0x2a, 0x1b, 0xdf, 0x07, 0x4a, 0xb7, 0xd7,
	0xea, 0xf4, 0xa7, 0x91, 0x0d, 0x61, 0xf8, 0x12, 0xb6, 0x3b, 0x37, 0x23, 0x70, 0xce, 0x78, 0x46,
	0x95, 0xe5, 0x3e, 0x51, 0x67, 0x59, 0xfb, 0x68, 0x38, 0x2e, 0xcf, 0xc3, 0x63, 0xf0, 0x1a, 0x97,
	0xfe, 0x55, 0x71, 0x06, 0xd0, 0xe6, 0x04, 0xed, 0x81, 0xfd, 0x95, 0xde, 0x34, 0x22, 0x79, 0x44,
	0x0f, 0xc1, 0xbd, 0x4e, 0x8a, 0x5a, 0x8b, 0x06, 0x27, 0xfe, 0xed, 0xce, 0x44, 0xd7, 0x4f, 0x7b,
	0xaf, 0xac, 0xe1, 0x1b, 0xf8, 0xff, 0xaf, 0xf8, 0xdc, 0x31, 0xeb, 0xc0, 0x9c, 0xe5, 0x1b, 0x03,
	0xc6, 0xcf, 0x61, 0x37, 0x8c, 0xc9, 0xd5, 0x5c, 0x46, 0xe9, 0x13, 0x5f, 0x2e, 0x0b, 0x2a, 0x83,
	0x3f, 0x9b, 0x37, 0xea, 0xde, 0x6c, 0x2e, 0xf1, 0x05, 0x6b, 0x52, 0xdf, 0xbb, 0x60, 0xe3, 0x9f,
	0x96, 0xde, 0xfc, 0xf3, 0x3a, 0x4b, 0x2a, 0x8a, 0xee, 0x81, 0x27, 0x27, 0xdc, 0x4a, 0xdc, 0x30,
	0x26, 0xb3, 0x39, 0x7a, 0x00, 0xbe, 0x24, 0x5d, 0x19, 0xdf, 0xdd, 0x97, 0x05, 0x89, 0xe5, 0x3e,
	0x51, 0x95, 0x54, 0x3a, 0x9e, 0x7d, 0xa2, 0x81, 0x4c, 0x67, 0xac, 0x53, 0xeb, 0xe8, 0xd4, 0xc6,
	0xb7, 0xa9, 0x35, 0x62, 0xe9, 0x12, 0xaf, 0xcd, 0x63, 0xd4, 0xcd, 0x63, 0xd4, 0xe6, 0x31, 0xea,
	0xe6, 0xb1, 0x81, 0x27, 0x3f, 0x2c, 0x80, 0x30, 0x26, 0x91, 0x7e, 0xfa, 0x92, 0xf8, 0x9e, 0x56,
	0x61, 0x4c, 0x4a, 0xe4, 0x06, 0xf2, 0xf9, 0x0f, 0x1d, 0xe9, 0xf3, 0xb1, 0x25, 0xed, 0x57, 0xa6,
	0x20, 0x15, 0x16, 0xd3, 0x20, 0x4d, 0x41, 0x4f, 0x00, 0xb4, 0x03, 0xea, 0x6d, 0x0e, 0x82, 0xd6,
	0x92, 0xa1, 0x09, 0xd0, 0x7d, 0xf0, 0xa3, 0x7a, 0x51, 0xa6, 0x22, 0x5f, 0x50, 0xa4, 0xa4, 0x9b,
	0x3b, 0xde, 0x1e, 0x5d, 0x3e, 0x5a, 0xe6, 0xd5, 0x97, 0x7a, 0x11, 0xa4, 0x7c, 0x35, 0x15, 0x3c,
	0xa3, 0x62, 0x35, 0x4d, 0xea, 0x2c, 0xe7, 0xcf, 0xd6, 0x09, 0xa3, 0xc5, 0x54, 0xfd, 0x33, 0x2d,
	0x3c, 0xf5, 0xf3, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x02, 0x60, 0xcf, 0xb7, 0x04,
	0x00, 0x00,
}
