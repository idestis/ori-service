// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activityevents.proto

package activityevents

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Deployment struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_39d0251f5d748552, []int{0}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EventRequest struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            *Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DeploymentId         int32      `protobuf:"varint,4,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39d0251f5d748552, []int{1}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EventRequest) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *EventRequest) GetDeploymentId() int32 {
	if m != nil {
		return m.DeploymentId
	}
	return 0
}

type Timestamp struct {
	// Will hold time when client send data (for calculations)
	Send int64 `protobuf:"varint,1,opt,name=send,proto3" json:"send,omitempty"`
	// Will hold time when server received data (for calculations)
	Receive              int64    `protobuf:"varint,2,opt,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_39d0251f5d748552, []int{2}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSend() int64 {
	if m != nil {
		return m.Send
	}
	return 0
}

func (m *Timestamp) GetReceive() int64 {
	if m != nil {
		return m.Receive
	}
	return 0
}

type EventResponse struct {
	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	//
	ExecutionTime        int64    `protobuf:"varint,3,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39d0251f5d748552, []int{3}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EventResponse) GetExecutionTime() int64 {
	if m != nil {
		return m.ExecutionTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Deployment)(nil), "activityevents.Deployment")
	proto.RegisterType((*EventRequest)(nil), "activityevents.EventRequest")
	proto.RegisterType((*Timestamp)(nil), "activityevents.Timestamp")
	proto.RegisterType((*EventResponse)(nil), "activityevents.EventResponse")
}

func init() { proto.RegisterFile("activityevents.proto", fileDescriptor_39d0251f5d748552) }

var fileDescriptor_39d0251f5d748552 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4b, 0x4f, 0xc2, 0x40,
	0x10, 0xc7, 0x59, 0x8a, 0x0f, 0x06, 0xe8, 0x61, 0xe2, 0xa1, 0x12, 0x4c, 0xc8, 0x1a, 0x13, 0x4e,
	0xc4, 0xe0, 0xc1, 0x78, 0x14, 0xf5, 0xa0, 0xe1, 0x60, 0x36, 0xde, 0xb1, 0xb6, 0x13, 0xb3, 0x89,
	0x7d, 0xc8, 0x2c, 0x44, 0xbe, 0x88, 0x1f, 0xc2, 0x4f, 0x69, 0xba, 0x7d, 0x69, 0x7d, 0xdc, 0xfa,
	0x9f, 0xfe, 0x67, 0xe7, 0xf7, 0x9f, 0x5d, 0x38, 0xf0, 0x03, 0xa3, 0x37, 0xda, 0x6c, 0x69, 0x43,
	0xb1, 0xe1, 0x69, 0xba, 0x4a, 0x4c, 0x82, 0xee, 0xf7, 0xaa, 0x1c, 0x01, 0x5c, 0x53, 0xfa, 0x92,
	0x6c, 0x23, 0x8a, 0x0d, 0xba, 0xd0, 0xd6, 0xa1, 0x27, 0xc6, 0x62, 0xb2, 0xa3, 0xda, 0x3a, 0x94,
	0xef, 0x02, 0xfa, 0x37, 0x99, 0x51, 0xd1, 0xeb, 0x9a, 0xf8, 0x87, 0x01, 0x3d, 0xd8, 0x8b, 0x88,
	0xd9, 0x7f, 0x26, 0xaf, 0x3d, 0x16, 0x93, 0xae, 0x2a, 0x25, 0x9e, 0x43, 0xd7, 0xe8, 0x88, 0xd8,
	0xf8, 0x51, 0xea, 0x39, 0x63, 0x31, 0xe9, 0xcd, 0x0e, 0xa7, 0x0d, 0xa4, 0x87, 0xd2, 0xa0, 0x6a,
	0x2f, 0x1e, 0xc3, 0x20, 0xac, 0x88, 0x96, 0x3a, 0xf4, 0x3a, 0x76, 0x5a, 0xbf, 0x2e, 0xde, 0x86,
	0xf2, 0x02, 0xba, 0x55, 0x33, 0x22, 0x74, 0x98, 0xe2, 0x1c, 0xcb, 0x51, 0xf6, 0x3b, 0x03, 0x5b,
	0x51, 0x40, 0x7a, 0x93, 0x83, 0x39, 0xaa, 0x94, 0xf2, 0x11, 0x06, 0x45, 0x24, 0x4e, 0x93, 0x98,
	0xe9, 0xb7, 0x4c, 0xbc, 0x0e, 0x02, 0x62, 0xb6, 0xad, 0xfb, 0xaa, 0x94, 0x78, 0x02, 0x2e, 0xbd,
	0x51, 0xb0, 0x36, 0x3a, 0x89, 0x97, 0x19, 0xb1, 0x0d, 0xe6, 0xa8, 0x41, 0x55, 0xcd, 0xa0, 0x66,
	0x1f, 0x02, 0xdc, 0xcb, 0x22, 0xa9, 0x1d, 0xc5, 0xb8, 0x80, 0xde, 0xd5, 0x8a, 0x7c, 0x43, 0x56,
	0xe3, 0xa8, 0xb9, 0x89, 0xaf, 0x4b, 0x1e, 0x1e, 0xfd, 0xf1, 0x37, 0xe7, 0x95, 0x2d, 0xbc, 0x03,
	0x58, 0x68, 0x36, 0xc5, 0xd9, 0xc3, 0xa6, 0xbd, 0xbe, 0xd0, 0xe1, 0xbf, 0x83, 0x64, 0xeb, 0x54,
	0xcc, 0x71, 0xde, 0x60, 0xbd, 0x17, 0x4f, 0xbb, 0xf6, 0xad, 0x9c, 0x7d, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0x49, 0x8b, 0x17, 0x43, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ActivityEventsClient is the client API for ActivityEvents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityEventsClient interface {
	// A simple RPC.
	//
	// Create a new event of a given deployment.
	CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	// A server-to-client streaming RPC.
	// Obtains the Events available within the given Deployment.
	// Results are streamed rather than returned at once
	ListEvents(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (ActivityEvents_ListEventsClient, error)
}

type activityEventsClient struct {
	cc *grpc.ClientConn
}

func NewActivityEventsClient(cc *grpc.ClientConn) ActivityEventsClient {
	return &activityEventsClient{cc}
}

func (c *activityEventsClient) CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/activityevents.ActivityEvents/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityEventsClient) ListEvents(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (ActivityEvents_ListEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ActivityEvents_serviceDesc.Streams[0], "/activityevents.ActivityEvents/ListEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &activityEventsListEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActivityEvents_ListEventsClient interface {
	Recv() (*EventRequest, error)
	grpc.ClientStream
}

type activityEventsListEventsClient struct {
	grpc.ClientStream
}

func (x *activityEventsListEventsClient) Recv() (*EventRequest, error) {
	m := new(EventRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActivityEventsServer is the server API for ActivityEvents service.
type ActivityEventsServer interface {
	// A simple RPC.
	//
	// Create a new event of a given deployment.
	CreateEvent(context.Context, *EventRequest) (*EventResponse, error)
	// A server-to-client streaming RPC.
	// Obtains the Events available within the given Deployment.
	// Results are streamed rather than returned at once
	ListEvents(*Deployment, ActivityEvents_ListEventsServer) error
}

// UnimplementedActivityEventsServer can be embedded to have forward compatible implementations.
type UnimplementedActivityEventsServer struct {
}

func (*UnimplementedActivityEventsServer) CreateEvent(ctx context.Context, req *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedActivityEventsServer) ListEvents(req *Deployment, srv ActivityEvents_ListEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}

func RegisterActivityEventsServer(s *grpc.Server, srv ActivityEventsServer) {
	s.RegisterService(&_ActivityEvents_serviceDesc, srv)
}

func _ActivityEvents_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityEventsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activityevents.ActivityEvents/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityEventsServer).CreateEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityEvents_ListEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Deployment)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActivityEventsServer).ListEvents(m, &activityEventsListEventsServer{stream})
}

type ActivityEvents_ListEventsServer interface {
	Send(*EventRequest) error
	grpc.ServerStream
}

type activityEventsListEventsServer struct {
	grpc.ServerStream
}

func (x *activityEventsListEventsServer) Send(m *EventRequest) error {
	return x.ServerStream.SendMsg(m)
}

var _ActivityEvents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "activityevents.ActivityEvents",
	HandlerType: (*ActivityEventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _ActivityEvents_CreateEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEvents",
			Handler:       _ActivityEvents_ListEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "activityevents.proto",
}
