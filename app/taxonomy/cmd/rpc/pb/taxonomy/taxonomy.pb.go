// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: rpc/pb/taxonomy.proto

package taxonomy

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

type Taxonomy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Term []string `protobuf:"bytes,2,rep,name=term,proto3" json:"term,omitempty"`
}

func (x *Taxonomy) Reset() {
	*x = Taxonomy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_taxonomy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Taxonomy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Taxonomy) ProtoMessage() {}

func (x *Taxonomy) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_taxonomy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Taxonomy.ProtoReflect.Descriptor instead.
func (*Taxonomy) Descriptor() ([]byte, []int) {
	return file_rpc_pb_taxonomy_proto_rawDescGZIP(), []int{0}
}

func (x *Taxonomy) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Taxonomy) GetTerm() []string {
	if x != nil {
		return x.Term
	}
	return nil
}

type ReqTaxonomyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReqTaxonomyId) Reset() {
	*x = ReqTaxonomyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_taxonomy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqTaxonomyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqTaxonomyId) ProtoMessage() {}

func (x *ReqTaxonomyId) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_taxonomy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqTaxonomyId.ProtoReflect.Descriptor instead.
func (*ReqTaxonomyId) Descriptor() ([]byte, []int) {
	return file_rpc_pb_taxonomy_proto_rawDescGZIP(), []int{1}
}

func (x *ReqTaxonomyId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReqGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqGetAll) Reset() {
	*x = ReqGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_taxonomy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetAll) ProtoMessage() {}

func (x *ReqGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_taxonomy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetAll.ProtoReflect.Descriptor instead.
func (*ReqGetAll) Descriptor() ([]byte, []int) {
	return file_rpc_pb_taxonomy_proto_rawDescGZIP(), []int{2}
}

type RespGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Taxonomies []*Taxonomy `protobuf:"bytes,1,rep,name=taxonomies,proto3" json:"taxonomies,omitempty"`
}

func (x *RespGetAll) Reset() {
	*x = RespGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_taxonomy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespGetAll) ProtoMessage() {}

func (x *RespGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_taxonomy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespGetAll.ProtoReflect.Descriptor instead.
func (*RespGetAll) Descriptor() ([]byte, []int) {
	return file_rpc_pb_taxonomy_proto_rawDescGZIP(), []int{3}
}

func (x *RespGetAll) GetTaxonomies() []*Taxonomy {
	if x != nil {
		return x.Taxonomies
	}
	return nil
}

var File_rpc_pb_taxonomy_proto protoreflect.FileDescriptor

var file_rpc_pb_taxonomy_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x22, 0x2e, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x22, 0x1f, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x22,
	0x40, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x32, 0x0a,
	0x0a, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x74, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x0a, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x65,
	0x73, 0x32, 0x75, 0x0a, 0x0a, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x65, 0x73, 0x12,
	0x32, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2e, 0x72, 0x65, 0x71, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x64, 0x1a,
	0x12, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e,
	0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x74, 0x61,
	0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_pb_taxonomy_proto_rawDescOnce sync.Once
	file_rpc_pb_taxonomy_proto_rawDescData = file_rpc_pb_taxonomy_proto_rawDesc
)

func file_rpc_pb_taxonomy_proto_rawDescGZIP() []byte {
	file_rpc_pb_taxonomy_proto_rawDescOnce.Do(func() {
		file_rpc_pb_taxonomy_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_pb_taxonomy_proto_rawDescData)
	})
	return file_rpc_pb_taxonomy_proto_rawDescData
}

var file_rpc_pb_taxonomy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_pb_taxonomy_proto_goTypes = []interface{}{
	(*Taxonomy)(nil),      // 0: taxonomy.taxonomy
	(*ReqTaxonomyId)(nil), // 1: taxonomy.reqTaxonomyId
	(*ReqGetAll)(nil),     // 2: taxonomy.reqGetAll
	(*RespGetAll)(nil),    // 3: taxonomy.respGetAll
}
var file_rpc_pb_taxonomy_proto_depIdxs = []int32{
	0, // 0: taxonomy.respGetAll.taxonomies:type_name -> taxonomy.taxonomy
	1, // 1: taxonomy.taxonomies.get:input_type -> taxonomy.reqTaxonomyId
	2, // 2: taxonomy.taxonomies.getAll:input_type -> taxonomy.reqGetAll
	0, // 3: taxonomy.taxonomies.get:output_type -> taxonomy.taxonomy
	3, // 4: taxonomy.taxonomies.getAll:output_type -> taxonomy.respGetAll
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_pb_taxonomy_proto_init() }
func file_rpc_pb_taxonomy_proto_init() {
	if File_rpc_pb_taxonomy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_pb_taxonomy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Taxonomy); i {
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
		file_rpc_pb_taxonomy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqTaxonomyId); i {
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
		file_rpc_pb_taxonomy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetAll); i {
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
		file_rpc_pb_taxonomy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespGetAll); i {
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
			RawDescriptor: file_rpc_pb_taxonomy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_pb_taxonomy_proto_goTypes,
		DependencyIndexes: file_rpc_pb_taxonomy_proto_depIdxs,
		MessageInfos:      file_rpc_pb_taxonomy_proto_msgTypes,
	}.Build()
	File_rpc_pb_taxonomy_proto = out.File
	file_rpc_pb_taxonomy_proto_rawDesc = nil
	file_rpc_pb_taxonomy_proto_goTypes = nil
	file_rpc_pb_taxonomy_proto_depIdxs = nil
}
