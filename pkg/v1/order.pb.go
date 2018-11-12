// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order_Vendor int32

const (
	Order_BLIBLI    Order_Vendor = 0
	Order_BUKALAPAK Order_Vendor = 1
	Order_MURAHID   Order_Vendor = 2
)

var Order_Vendor_name = map[int32]string{
	0: "BLIBLI",
	1: "BUKALAPAK",
	2: "MURAHID",
}

var Order_Vendor_value = map[string]int32{
	"BLIBLI":    0,
	"BUKALAPAK": 1,
	"MURAHID":   2,
}

func (x Order_Vendor) String() string {
	return proto.EnumName(Order_Vendor_name, int32(x))
}

func (Order_Vendor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1, 0}
}

type Person struct {
	Nama                 string   `protobuf:"bytes,1,opt,name=nama,proto3" json:"nama,omitempty"`
	Alamat               string   `protobuf:"bytes,2,opt,name=alamat,proto3" json:"alamat,omitempty"`
	Telepon              string   `protobuf:"bytes,3,opt,name=telepon,proto3" json:"telepon,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetNama() string {
	if m != nil {
		return m.Nama
	}
	return ""
}

func (m *Person) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *Person) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Order struct {
	Vendor               Order_Vendor `protobuf:"varint,1,opt,name=vendor,proto3,enum=v1.Order_Vendor" json:"vendor,omitempty"`
	Pembeli              *Person      `protobuf:"bytes,2,opt,name=pembeli,proto3" json:"pembeli,omitempty"`
	Barang               *Product     `protobuf:"bytes,3,opt,name=barang,proto3" json:"barang,omitempty"`
	Status               string       `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string       `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetVendor() Order_Vendor {
	if m != nil {
		return m.Vendor
	}
	return Order_BLIBLI
}

func (m *Order) GetPembeli() *Person {
	if m != nil {
		return m.Pembeli
	}
	return nil
}

func (m *Order) GetBarang() *Product {
	if m != nil {
		return m.Barang
	}
	return nil
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Update struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Update) Reset()         { *m = Update{} }
func (m *Update) String() string { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()    {}
func (*Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update.Unmarshal(m, b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update.Marshal(b, m, deterministic)
}
func (m *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(m, src)
}
func (m *Update) XXX_Size() int {
	return xxx_messageInfo_Update.Size(m)
}
func (m *Update) XXX_DiscardUnknown() {
	xxx_messageInfo_Update.DiscardUnknown(m)
}

var xxx_messageInfo_Update proto.InternalMessageInfo

func (m *Update) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Update) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type OrderId struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderId) Reset()         { *m = OrderId{} }
func (m *OrderId) String() string { return proto.CompactTextString(m) }
func (*OrderId) ProtoMessage()    {}
func (*OrderId) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *OrderId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderId.Unmarshal(m, b)
}
func (m *OrderId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderId.Marshal(b, m, deterministic)
}
func (m *OrderId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderId.Merge(m, src)
}
func (m *OrderId) XXX_Size() int {
	return xxx_messageInfo_OrderId.Size(m)
}
func (m *OrderId) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderId.DiscardUnknown(m)
}

var xxx_messageInfo_OrderId proto.InternalMessageInfo

func (m *OrderId) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type UserId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Orders struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Orders) Reset()         { *m = Orders{} }
func (m *Orders) String() string { return proto.CompactTextString(m) }
func (*Orders) ProtoMessage()    {}
func (*Orders) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{5}
}

func (m *Orders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Orders.Unmarshal(m, b)
}
func (m *Orders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Orders.Marshal(b, m, deterministic)
}
func (m *Orders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orders.Merge(m, src)
}
func (m *Orders) XXX_Size() int {
	return xxx_messageInfo_Orders.Size(m)
}
func (m *Orders) XXX_DiscardUnknown() {
	xxx_messageInfo_Orders.DiscardUnknown(m)
}

var xxx_messageInfo_Orders proto.InternalMessageInfo

func (m *Orders) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterEnum("v1.Order_Vendor", Order_Vendor_name, Order_Vendor_value)
	proto.RegisterType((*Person)(nil), "v1.Person")
	proto.RegisterType((*Order)(nil), "v1.Order")
	proto.RegisterType((*Update)(nil), "v1.Update")
	proto.RegisterType((*OrderId)(nil), "v1.OrderId")
	proto.RegisterType((*UserId)(nil), "v1.UserId")
	proto.RegisterType((*Orders)(nil), "v1.Orders")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xe1, 0x6a, 0xdb, 0x30,
	0x10, 0x8e, 0xdd, 0x46, 0xc1, 0xe7, 0xb6, 0x64, 0xc7, 0x18, 0x26, 0x30, 0xe8, 0xdc, 0x0d, 0xc2,
	0x06, 0xd9, 0xea, 0x3d, 0x81, 0xb3, 0xc2, 0x66, 0x9a, 0xb1, 0xe2, 0x92, 0xfd, 0x1d, 0x4a, 0x74,
	0x14, 0x83, 0x63, 0x19, 0x49, 0xcd, 0xd3, 0xed, 0x75, 0xf6, 0x1e, 0xc3, 0x27, 0x25, 0x2b, 0xdd,
	0xbf, 0xbb, 0xef, 0xfb, 0xee, 0xee, 0xbb, 0x93, 0x20, 0xd5, 0x46, 0x91, 0x59, 0xf4, 0x46, 0x3b,
	0x8d, 0xf1, 0xfe, 0x7a, 0x76, 0xbe, 0x95, 0x4e, 0xb6, 0xfa, 0xc1, 0x43, 0xb9, 0x02, 0x71, 0x47,
	0xc6, 0xea, 0x0e, 0x11, 0x4e, 0x3b, 0xb9, 0x93, 0x59, 0x74, 0x19, 0xcd, 0x93, 0x9a, 0x63, 0x7c,
	0x05, 0x42, 0xb6, 0x72, 0x27, 0x5d, 0x16, 0x33, 0x1a, 0x32, 0xcc, 0x60, 0xe2, 0xa8, 0xa5, 0x5e,
	0x77, 0xd9, 0x09, 0x13, 0x87, 0x14, 0x5f, 0xc2, 0x98, 0x76, 0xb2, 0x69, 0xb3, 0x53, 0xc6, 0x7d,
	0x92, 0xff, 0x89, 0x60, 0xfc, 0x63, 0x30, 0x82, 0x73, 0x10, 0x7b, 0xea, 0x94, 0x36, 0x3c, 0xe7,
	0xa2, 0x98, 0x2e, 0xf6, 0xd7, 0x0b, 0xa6, 0x16, 0x3f, 0x19, 0xaf, 0x03, 0x8f, 0x6f, 0x61, 0xd2,
	0xd3, 0x6e, 0x43, 0x6d, 0xc3, 0xc3, 0xd3, 0x02, 0x06, 0xa9, 0x37, 0x5b, 0x1f, 0x28, 0xbc, 0x02,
	0xb1, 0x91, 0x46, 0x76, 0x0f, 0x6c, 0x24, 0x2d, 0x52, 0x16, 0x19, 0xad, 0x1e, 0xb7, 0xae, 0x0e,
	0xd4, 0xb0, 0x86, 0x75, 0xd2, 0x3d, 0xda, 0xe0, 0x2a, 0x64, 0xf8, 0x1a, 0x60, 0x6b, 0x48, 0x3a,
	0x52, 0xbf, 0xa4, 0xcb, 0xc6, 0xcc, 0x25, 0x01, 0x29, 0x5d, 0xfe, 0x09, 0x84, 0xf7, 0x84, 0x00,
	0x62, 0xb9, 0xaa, 0x96, 0xab, 0x6a, 0x3a, 0xc2, 0x73, 0x48, 0x96, 0xeb, 0xdb, 0x72, 0x55, 0xde,
	0x95, 0xb7, 0xd3, 0x08, 0x53, 0x98, 0x7c, 0x5f, 0xd7, 0xe5, 0xb7, 0xea, 0x66, 0x1a, 0x0f, 0x15,
	0xeb, 0x5e, 0x49, 0x47, 0x78, 0x01, 0x71, 0xa3, 0xc2, 0x2d, 0xe3, 0x46, 0x3d, 0xb1, 0x10, 0x3f,
	0xb5, 0x90, 0x5f, 0xc1, 0x84, 0xb7, 0xaf, 0xd4, 0x70, 0x54, 0xed, 0xc3, 0x50, 0x77, 0x48, 0xf3,
	0x0c, 0xc4, 0xda, 0xb2, 0xe6, 0x59, 0xdb, 0xfc, 0x03, 0x08, 0x2e, 0xb7, 0xf8, 0x06, 0x04, 0xcb,
	0x6d, 0x16, 0x5d, 0x9e, 0xcc, 0xd3, 0x22, 0x39, 0x1e, 0xb6, 0x0e, 0x44, 0xf1, 0x3b, 0x82, 0x33,
	0x46, 0xee, 0xc9, 0xec, 0x9b, 0x2d, 0xe1, 0x1c, 0xd2, 0x2f, 0xbc, 0xad, 0x7f, 0x9b, 0x7f, 0x25,
	0xb3, 0xb3, 0x21, 0xac, 0xc9, 0xf6, 0xba, 0xb3, 0x94, 0x8f, 0xf0, 0x1d, 0x24, 0x5f, 0xc9, 0x85,
	0x51, 0xfc, 0x10, 0xde, 0xd0, 0x0c, 0x8e, 0x35, 0x36, 0x1f, 0xe1, 0x47, 0x78, 0xe1, 0xf7, 0xbf,
	0xe7, 0xed, 0x7c, 0x5b, 0x2f, 0x67, 0xf8, 0xbf, 0xbe, 0xef, 0x21, 0xbd, 0xa1, 0x96, 0x0e, 0x0e,
	0xd2, 0x63, 0xb7, 0x4a, 0x3d, 0xd7, 0x6e, 0x04, 0xff, 0xd8, 0xcf, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x31, 0x5b, 0x58, 0xd3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Response, error)
	GetOrders(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Orders, error)
	UpdateStatusOrder(ctx context.Context, in *Update, opts ...grpc.CallOption) (*Response, error)
	DeleteOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*Response, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/v1.OrderService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateStatusOrder(ctx context.Context, in *Update, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.OrderService/UpdateStatusOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.OrderService/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	CreateOrder(context.Context, *Order) (*Response, error)
	GetOrders(context.Context, *UserId) (*Orders, error)
	UpdateStatusOrder(context.Context, *Update) (*Response, error)
	DeleteOrder(context.Context, *OrderId) (*Response, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OrderService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrders(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateStatusOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Update)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateStatusOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OrderService/UpdateStatusOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateStatusOrder(ctx, req.(*Update))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OrderService/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrderService_GetOrders_Handler,
		},
		{
			MethodName: "UpdateStatusOrder",
			Handler:    _OrderService_UpdateStatusOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
