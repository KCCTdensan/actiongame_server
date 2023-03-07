// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/actiongame.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{0}
}

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloRes) Reset() {
	*x = HelloRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRes) ProtoMessage() {}

func (x *HelloRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRes.ProtoReflect.Descriptor instead.
func (*HelloRes) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserReq) Reset() {
	*x = GetUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReq) ProtoMessage() {}

func (x *GetUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReq.ProtoReflect.Descriptor instead.
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserRes) Reset() {
	*x = GetUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRes) ProtoMessage() {}

func (x *GetUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRes.ProtoReflect.Descriptor instead.
func (*GetUserRes) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserRes) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JoinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JoinReq) Reset() {
	*x = JoinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReq) ProtoMessage() {}

func (x *JoinReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReq.ProtoReflect.Descriptor instead.
func (*JoinReq) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{5}
}

func (x *JoinReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserConn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enum type (con/discon)
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserConn) Reset() {
	*x = UserConn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserConn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConn) ProtoMessage() {}

func (x *UserConn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConn.ProtoReflect.Descriptor instead.
func (*UserConn) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{6}
}

func (x *UserConn) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Pos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Pos) Reset() {
	*x = Pos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pos) ProtoMessage() {}

func (x *Pos) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pos.ProtoReflect.Descriptor instead.
func (*Pos) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{7}
}

func (x *Pos) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Pos) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type UserMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pos *Pos   `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *UserMove) Reset() {
	*x = UserMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actiongame_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMove) ProtoMessage() {}

func (x *UserMove) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actiongame_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMove.ProtoReflect.Descriptor instead.
func (*UserMove) Descriptor() ([]byte, []int) {
	return file_proto_actiongame_proto_rawDescGZIP(), []int{8}
}

func (x *UserMove) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserMove) GetPos() *Pos {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_proto_actiongame_proto protoreflect.FileDescriptor

var file_proto_actiongame_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x21, 0x0a, 0x03, 0x50, 0x6f, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x79, 0x22, 0x38, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x32, 0xb6,
	0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x30, 0x01, 0x12, 0x27,
	0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x76, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_actiongame_proto_rawDescOnce sync.Once
	file_proto_actiongame_proto_rawDescData = file_proto_actiongame_proto_rawDesc
)

func file_proto_actiongame_proto_rawDescGZIP() []byte {
	file_proto_actiongame_proto_rawDescOnce.Do(func() {
		file_proto_actiongame_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_actiongame_proto_rawDescData)
	})
	return file_proto_actiongame_proto_rawDescData
}

var file_proto_actiongame_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_actiongame_proto_goTypes = []interface{}{
	(*Empty)(nil),      // 0: proto.Empty
	(*HelloReq)(nil),   // 1: proto.HelloReq
	(*HelloRes)(nil),   // 2: proto.HelloRes
	(*GetUserReq)(nil), // 3: proto.GetUserReq
	(*GetUserRes)(nil), // 4: proto.GetUserRes
	(*JoinReq)(nil),    // 5: proto.JoinReq
	(*UserConn)(nil),   // 6: proto.UserConn
	(*Pos)(nil),        // 7: proto.Pos
	(*UserMove)(nil),   // 8: proto.UserMove
}
var file_proto_actiongame_proto_depIdxs = []int32{
	7, // 0: proto.UserMove.pos:type_name -> proto.Pos
	1, // 1: proto.Game.Hello:input_type -> proto.HelloReq
	3, // 2: proto.Game.GetUser:input_type -> proto.GetUserReq
	5, // 3: proto.Game.Join:input_type -> proto.JoinReq
	7, // 4: proto.Game.Move:input_type -> proto.Pos
	2, // 5: proto.Game.Hello:output_type -> proto.HelloRes
	4, // 6: proto.Game.GetUser:output_type -> proto.GetUserRes
	6, // 7: proto.Game.Join:output_type -> proto.UserConn
	8, // 8: proto.Game.Move:output_type -> proto.UserMove
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_actiongame_proto_init() }
func file_proto_actiongame_proto_init() {
	if File_proto_actiongame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_actiongame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserConn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_actiongame_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMove); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_actiongame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_actiongame_proto_goTypes,
		DependencyIndexes: file_proto_actiongame_proto_depIdxs,
		MessageInfos:      file_proto_actiongame_proto_msgTypes,
	}.Build()
	File_proto_actiongame_proto = out.File
	file_proto_actiongame_proto_rawDesc = nil
	file_proto_actiongame_proto_goTypes = nil
	file_proto_actiongame_proto_depIdxs = nil
}
