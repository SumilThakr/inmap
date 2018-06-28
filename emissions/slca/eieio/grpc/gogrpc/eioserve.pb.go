// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eioserve.proto

package eieiorpc

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

type Selectors struct {
	Codes                []string  `protobuf:"bytes,1,rep,name=Codes,proto3" json:"Codes,omitempty"`
	Names                []string  `protobuf:"bytes,2,rep,name=Names,proto3" json:"Names,omitempty"`
	Values               []float32 `protobuf:"fixed32,3,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Selectors) Reset()         { *m = Selectors{} }
func (m *Selectors) String() string { return proto.CompactTextString(m) }
func (*Selectors) ProtoMessage()    {}
func (*Selectors) Descriptor() ([]byte, []int) {
	return fileDescriptor_eioserve_63bc52d84ac2c67f, []int{0}
}
func (m *Selectors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selectors.Unmarshal(m, b)
}
func (m *Selectors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selectors.Marshal(b, m, deterministic)
}
func (dst *Selectors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selectors.Merge(dst, src)
}
func (m *Selectors) XXX_Size() int {
	return xxx_messageInfo_Selectors.Size(m)
}
func (m *Selectors) XXX_DiscardUnknown() {
	xxx_messageInfo_Selectors.DiscardUnknown(m)
}

var xxx_messageInfo_Selectors proto.InternalMessageInfo

func (m *Selectors) GetCodes() []string {
	if m != nil {
		return m.Codes
	}
	return nil
}

func (m *Selectors) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Selectors) GetValues() []float32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Selection struct {
	DemandGroup          string   `protobuf:"bytes,1,opt,name=DemandGroup,proto3" json:"DemandGroup,omitempty"`
	DemandSector         string   `protobuf:"bytes,2,opt,name=DemandSector,proto3" json:"DemandSector,omitempty"`
	ProductionGroup      string   `protobuf:"bytes,3,opt,name=ProductionGroup,proto3" json:"ProductionGroup,omitempty"`
	ProductionSector     string   `protobuf:"bytes,4,opt,name=ProductionSector,proto3" json:"ProductionSector,omitempty"`
	ImpactType           string   `protobuf:"bytes,5,opt,name=ImpactType,proto3" json:"ImpactType,omitempty"`
	DemandType           string   `protobuf:"bytes,6,opt,name=DemandType,proto3" json:"DemandType,omitempty"`
	Year                 int32    `protobuf:"varint,7,opt,name=Year,proto3" json:"Year,omitempty"`
	Population           string   `protobuf:"bytes,8,opt,name=Population,proto3" json:"Population,omitempty"`
	Pollutant            int32    `protobuf:"varint,9,opt,name=Pollutant,proto3" json:"Pollutant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selection) Reset()         { *m = Selection{} }
func (m *Selection) String() string { return proto.CompactTextString(m) }
func (*Selection) ProtoMessage()    {}
func (*Selection) Descriptor() ([]byte, []int) {
	return fileDescriptor_eioserve_63bc52d84ac2c67f, []int{1}
}
func (m *Selection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selection.Unmarshal(m, b)
}
func (m *Selection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selection.Marshal(b, m, deterministic)
}
func (dst *Selection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selection.Merge(dst, src)
}
func (m *Selection) XXX_Size() int {
	return xxx_messageInfo_Selection.Size(m)
}
func (m *Selection) XXX_DiscardUnknown() {
	xxx_messageInfo_Selection.DiscardUnknown(m)
}

var xxx_messageInfo_Selection proto.InternalMessageInfo

func (m *Selection) GetDemandGroup() string {
	if m != nil {
		return m.DemandGroup
	}
	return ""
}

func (m *Selection) GetDemandSector() string {
	if m != nil {
		return m.DemandSector
	}
	return ""
}

func (m *Selection) GetProductionGroup() string {
	if m != nil {
		return m.ProductionGroup
	}
	return ""
}

func (m *Selection) GetProductionSector() string {
	if m != nil {
		return m.ProductionSector
	}
	return ""
}

func (m *Selection) GetImpactType() string {
	if m != nil {
		return m.ImpactType
	}
	return ""
}

func (m *Selection) GetDemandType() string {
	if m != nil {
		return m.DemandType
	}
	return ""
}

func (m *Selection) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Selection) GetPopulation() string {
	if m != nil {
		return m.Population
	}
	return ""
}

func (m *Selection) GetPollutant() int32 {
	if m != nil {
		return m.Pollutant
	}
	return 0
}

type Year struct {
	Years                []int32  `protobuf:"varint,1,rep,packed,name=Years,proto3" json:"Years,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Year) Reset()         { *m = Year{} }
func (m *Year) String() string { return proto.CompactTextString(m) }
func (*Year) ProtoMessage()    {}
func (*Year) Descriptor() ([]byte, []int) {
	return fileDescriptor_eioserve_63bc52d84ac2c67f, []int{2}
}
func (m *Year) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Year.Unmarshal(m, b)
}
func (m *Year) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Year.Marshal(b, m, deterministic)
}
func (dst *Year) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Year.Merge(dst, src)
}
func (m *Year) XXX_Size() int {
	return xxx_messageInfo_Year.Size(m)
}
func (m *Year) XXX_DiscardUnknown() {
	xxx_messageInfo_Year.DiscardUnknown(m)
}

var xxx_messageInfo_Year proto.InternalMessageInfo

func (m *Year) GetYears() []int32 {
	if m != nil {
		return m.Years
	}
	return nil
}

type Point struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_eioserve_63bc52d84ac2c67f, []int{3}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Rectangle struct {
	LL                   *Point   `protobuf:"bytes,1,opt,name=LL,proto3" json:"LL,omitempty"`
	LR                   *Point   `protobuf:"bytes,2,opt,name=LR,proto3" json:"LR,omitempty"`
	UR                   *Point   `protobuf:"bytes,3,opt,name=UR,proto3" json:"UR,omitempty"`
	UL                   *Point   `protobuf:"bytes,4,opt,name=UL,proto3" json:"UL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_eioserve_63bc52d84ac2c67f, []int{4}
}
func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (dst *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(dst, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetLL() *Point {
	if m != nil {
		return m.LL
	}
	return nil
}

func (m *Rectangle) GetLR() *Point {
	if m != nil {
		return m.LR
	}
	return nil
}

func (m *Rectangle) GetUR() *Point {
	if m != nil {
		return m.UR
	}
	return nil
}

func (m *Rectangle) GetUL() *Point {
	if m != nil {
		return m.UL
	}
	return nil
}

type ColorInfo struct {
	RGB                  [][]byte `protobuf:"bytes,1,rep,name=RGB,proto3" json:"RGB,omitempty"`
	Legend               string   `protobuf:"bytes,2,opt,name=Legend,proto3" json:"Legend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColorInfo) Reset()         { *m = ColorInfo{} }
func (m *ColorInfo) String() string { return proto.CompactTextString(m) }
func (*ColorInfo) ProtoMessage()    {}
func (*ColorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eioserve_63bc52d84ac2c67f, []int{5}
}
func (m *ColorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColorInfo.Unmarshal(m, b)
}
func (m *ColorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColorInfo.Marshal(b, m, deterministic)
}
func (dst *ColorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColorInfo.Merge(dst, src)
}
func (m *ColorInfo) XXX_Size() int {
	return xxx_messageInfo_ColorInfo.Size(m)
}
func (m *ColorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ColorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ColorInfo proto.InternalMessageInfo

func (m *ColorInfo) GetRGB() [][]byte {
	if m != nil {
		return m.RGB
	}
	return nil
}

func (m *ColorInfo) GetLegend() string {
	if m != nil {
		return m.Legend
	}
	return ""
}

func init() {
	proto.RegisterType((*Selectors)(nil), "eieiorpc.Selectors")
	proto.RegisterType((*Selection)(nil), "eieiorpc.Selection")
	proto.RegisterType((*Year)(nil), "eieiorpc.Year")
	proto.RegisterType((*Point)(nil), "eieiorpc.Point")
	proto.RegisterType((*Rectangle)(nil), "eieiorpc.Rectangle")
	proto.RegisterType((*ColorInfo)(nil), "eieiorpc.ColorInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EIEIOrpcClient is the client API for EIEIOrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EIEIOrpcClient interface {
	DemandGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	DemandSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	ProdGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	ProdSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	Years(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Year, error)
	DefaultSelection(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selection, error)
	Populations(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	MapInfo(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*ColorInfo, error)
	Geometry(ctx context.Context, in *Selection, opts ...grpc.CallOption) (EIEIOrpc_GeometryClient, error)
}

type eIEIOrpcClient struct {
	cc *grpc.ClientConn
}

func NewEIEIOrpcClient(cc *grpc.ClientConn) EIEIOrpcClient {
	return &eIEIOrpcClient{cc}
}

func (c *eIEIOrpcClient) DemandGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/DemandGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) DemandSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/DemandSectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) ProdGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/ProdGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) ProdSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/ProdSectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) Years(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Year, error) {
	out := new(Year)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/Years", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) DefaultSelection(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selection, error) {
	out := new(Selection)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/DefaultSelection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) Populations(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/Populations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) MapInfo(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*ColorInfo, error) {
	out := new(ColorInfo)
	err := c.cc.Invoke(ctx, "/eieiorpc.EIEIOrpc/MapInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIEIOrpcClient) Geometry(ctx context.Context, in *Selection, opts ...grpc.CallOption) (EIEIOrpc_GeometryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EIEIOrpc_serviceDesc.Streams[0], "/eieiorpc.EIEIOrpc/Geometry", opts...)
	if err != nil {
		return nil, err
	}
	x := &eIEIOrpcGeometryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EIEIOrpc_GeometryClient interface {
	Recv() (*Rectangle, error)
	grpc.ClientStream
}

type eIEIOrpcGeometryClient struct {
	grpc.ClientStream
}

func (x *eIEIOrpcGeometryClient) Recv() (*Rectangle, error) {
	m := new(Rectangle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EIEIOrpcServer is the server API for EIEIOrpc service.
type EIEIOrpcServer interface {
	DemandGroups(context.Context, *Selection) (*Selectors, error)
	DemandSectors(context.Context, *Selection) (*Selectors, error)
	ProdGroups(context.Context, *Selection) (*Selectors, error)
	ProdSectors(context.Context, *Selection) (*Selectors, error)
	Years(context.Context, *Selection) (*Year, error)
	DefaultSelection(context.Context, *Selection) (*Selection, error)
	Populations(context.Context, *Selection) (*Selectors, error)
	MapInfo(context.Context, *Selection) (*ColorInfo, error)
	Geometry(*Selection, EIEIOrpc_GeometryServer) error
}

func RegisterEIEIOrpcServer(s *grpc.Server, srv EIEIOrpcServer) {
	s.RegisterService(&_EIEIOrpc_serviceDesc, srv)
}

func _EIEIOrpc_DemandGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).DemandGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/DemandGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).DemandGroups(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_DemandSectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).DemandSectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/DemandSectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).DemandSectors(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_ProdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).ProdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/ProdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).ProdGroups(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_ProdSectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).ProdSectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/ProdSectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).ProdSectors(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_Years_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).Years(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/Years",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).Years(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_DefaultSelection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).DefaultSelection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/DefaultSelection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).DefaultSelection(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_Populations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).Populations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/Populations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).Populations(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_MapInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIEIOrpcServer).MapInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eieiorpc.EIEIOrpc/MapInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIEIOrpcServer).MapInfo(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIEIOrpc_Geometry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Selection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EIEIOrpcServer).Geometry(m, &eIEIOrpcGeometryServer{stream})
}

type EIEIOrpc_GeometryServer interface {
	Send(*Rectangle) error
	grpc.ServerStream
}

type eIEIOrpcGeometryServer struct {
	grpc.ServerStream
}

func (x *eIEIOrpcGeometryServer) Send(m *Rectangle) error {
	return x.ServerStream.SendMsg(m)
}

var _EIEIOrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eieiorpc.EIEIOrpc",
	HandlerType: (*EIEIOrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DemandGroups",
			Handler:    _EIEIOrpc_DemandGroups_Handler,
		},
		{
			MethodName: "DemandSectors",
			Handler:    _EIEIOrpc_DemandSectors_Handler,
		},
		{
			MethodName: "ProdGroups",
			Handler:    _EIEIOrpc_ProdGroups_Handler,
		},
		{
			MethodName: "ProdSectors",
			Handler:    _EIEIOrpc_ProdSectors_Handler,
		},
		{
			MethodName: "Years",
			Handler:    _EIEIOrpc_Years_Handler,
		},
		{
			MethodName: "DefaultSelection",
			Handler:    _EIEIOrpc_DefaultSelection_Handler,
		},
		{
			MethodName: "Populations",
			Handler:    _EIEIOrpc_Populations_Handler,
		},
		{
			MethodName: "MapInfo",
			Handler:    _EIEIOrpc_MapInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Geometry",
			Handler:       _EIEIOrpc_Geometry_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eioserve.proto",
}

func init() { proto.RegisterFile("eioserve.proto", fileDescriptor_eioserve_63bc52d84ac2c67f) }

var fileDescriptor_eioserve_63bc52d84ac2c67f = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x64, 0xcb, 0xb1, 0xc6, 0x6e, 0x62, 0xb6, 0xa5, 0x88, 0x12, 0x5a, 0xb1, 0xbd, 0x88,
	0x1e, 0x4c, 0x49, 0x09, 0xf4, 0x03, 0x7a, 0x68, 0x12, 0x8c, 0x41, 0x6d, 0xcc, 0xa6, 0x29, 0xf1,
	0x71, 0x2b, 0x4f, 0x82, 0x40, 0xd6, 0x8a, 0x95, 0x54, 0xc8, 0xaf, 0xe8, 0xa5, 0x7f, 0xa6, 0xff,
	0xae, 0xec, 0xac, 0x6c, 0xb9, 0x8d, 0x21, 0x4d, 0x4e, 0xda, 0x79, 0xfb, 0xde, 0xcc, 0x68, 0xe6,
	0x49, 0xb0, 0x87, 0xa9, 0x2a, 0x51, 0xff, 0xc0, 0x71, 0xa1, 0x55, 0xa5, 0x58, 0x1f, 0x53, 0x4c,
	0x95, 0x2e, 0x12, 0x7e, 0x06, 0xfe, 0x39, 0x66, 0x98, 0x54, 0x4a, 0x97, 0xec, 0x09, 0x78, 0xc7,
	0x6a, 0x81, 0x65, 0xe0, 0x84, 0x9d, 0xc8, 0x17, 0x36, 0x30, 0xe8, 0x17, 0xb9, 0xc4, 0x32, 0x70,
	0x2d, 0x4a, 0x01, 0x7b, 0x0a, 0xbd, 0x6f, 0x32, 0xab, 0xb1, 0x0c, 0x3a, 0x61, 0x27, 0x72, 0x45,
	0x13, 0xf1, 0xdf, 0xee, 0x2a, 0x63, 0xaa, 0x72, 0x16, 0xc2, 0xe0, 0x04, 0x97, 0x32, 0x5f, 0x4c,
	0xb4, 0xaa, 0x8b, 0xc0, 0x09, 0x9d, 0xc8, 0x17, 0x9b, 0x10, 0xe3, 0x30, 0xb4, 0xe1, 0x39, 0x35,
	0x11, 0xb8, 0x44, 0xf9, 0x0b, 0x63, 0x11, 0xec, 0xcf, 0xb4, 0x5a, 0xd4, 0x94, 0xd3, 0x66, 0xea,
	0x10, 0xed, 0x5f, 0x98, 0xbd, 0x82, 0x51, 0x0b, 0x35, 0x19, 0xbb, 0x44, 0xbd, 0x85, 0xb3, 0xe7,
	0x00, 0xd3, 0x65, 0x21, 0x93, 0xea, 0xeb, 0x4d, 0x81, 0x81, 0x47, 0xac, 0x0d, 0xc4, 0xdc, 0xdb,
	0x2e, 0xe8, 0xbe, 0x67, 0xef, 0x5b, 0x84, 0x31, 0xe8, 0xce, 0x51, 0xea, 0x60, 0x37, 0x74, 0x22,
	0x4f, 0xd0, 0xd9, 0x68, 0x66, 0xaa, 0xa8, 0x33, 0x69, 0xea, 0x04, 0x7d, 0xab, 0x69, 0x11, 0x76,
	0x00, 0xfe, 0x4c, 0x65, 0x59, 0x5d, 0xc9, 0xbc, 0x0a, 0x7c, 0x12, 0xb6, 0x00, 0x3f, 0xb0, 0x19,
	0xcd, 0xc4, 0xcd, 0xd3, 0xee, 0xc1, 0x13, 0x36, 0xe0, 0x2f, 0xc1, 0x9b, 0xa9, 0x34, 0xaf, 0xd8,
	0x10, 0x9c, 0x4b, 0x1a, 0xa5, 0x2b, 0x9c, 0x4b, 0x13, 0xcd, 0x69, 0x6a, 0xae, 0x70, 0xe6, 0xfc,
	0xa7, 0x03, 0xbe, 0xc0, 0xa4, 0x92, 0xf9, 0x75, 0x86, 0xec, 0x05, 0xb8, 0x71, 0x4c, 0xd4, 0xc1,
	0xe1, 0xfe, 0x78, 0xb5, 0xf4, 0x31, 0xa5, 0x11, 0x6e, 0x1c, 0x13, 0x41, 0x90, 0x7a, 0x2b, 0x41,
	0x18, 0xc2, 0x85, 0xa0, 0x69, 0x6f, 0x23, 0x5c, 0x58, 0x42, 0x4c, 0x33, 0xde, 0x4a, 0x88, 0xf9,
	0x11, 0xf8, 0xc7, 0x2a, 0x53, 0x7a, 0x9a, 0x5f, 0x29, 0x36, 0x82, 0x8e, 0x98, 0x7c, 0xa2, 0xf7,
	0x1a, 0x0a, 0x73, 0x34, 0x3e, 0x8a, 0xf1, 0x1a, 0xf3, 0x45, 0xb3, 0xf9, 0x26, 0x3a, 0xfc, 0xd5,
	0x85, 0xfe, 0xe9, 0xf4, 0x74, 0x7a, 0xa6, 0x8b, 0x84, 0xbd, 0x5f, 0x99, 0x84, 0xb6, 0x5c, 0xb2,
	0xc7, 0x6d, 0xa1, 0xb5, 0xd7, 0x9e, 0xdd, 0x02, 0x95, 0x2e, 0xf9, 0x0e, 0xfb, 0x00, 0x8f, 0x36,
	0xcd, 0x74, 0x3f, 0xf1, 0x5b, 0x00, 0xe3, 0x9b, 0x07, 0x94, 0x7d, 0x07, 0x03, 0xa3, 0x7c, 0x48,
	0xd1, 0x71, 0xb3, 0xfe, 0xed, 0xa2, 0xbd, 0x16, 0x34, 0x2c, 0xbe, 0xc3, 0x3e, 0xc2, 0xe8, 0x04,
	0xaf, 0x64, 0x9d, 0x55, 0xed, 0x87, 0xf7, 0x7f, 0xf5, 0x52, 0x95, 0x37, 0xad, 0xae, 0x2d, 0x7a,
	0xbf, 0x56, 0x8f, 0x60, 0xf7, 0xb3, 0x2c, 0x68, 0xb5, 0x77, 0xc9, 0xd6, 0x26, 0xa0, 0xb1, 0xf6,
	0x27, 0xa8, 0x96, 0x58, 0xe9, 0x9b, 0x3b, 0x75, 0x6b, 0x37, 0xf3, 0x9d, 0xd7, 0xce, 0xf7, 0x1e,
	0xfd, 0xc0, 0xde, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xfb, 0xe4, 0x09, 0xd2, 0x04, 0x00,
	0x00,
}