// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: graphql.proto

package graphqlpb

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type ParameterType int32

const (
	ParameterType_Variable ParameterType = 0
	ParameterType_String   ParameterType = 1
	ParameterType_Integer  ParameterType = 2
	ParameterType_Float    ParameterType = 3
	ParameterType_Block    ParameterType = 4
	ParameterType_Boolean  ParameterType = 5
	ParameterType_Null     ParameterType = 6
	ParameterType_Enum     ParameterType = 7
	ParameterType_List     ParameterType = 8
	ParameterType_Object   ParameterType = 9
)

// Enum value maps for ParameterType.
var (
	ParameterType_name = map[int32]string{
		0: "Variable",
		1: "String",
		2: "Integer",
		3: "Float",
		4: "Block",
		5: "Boolean",
		6: "Null",
		7: "Enum",
		8: "List",
		9: "Object",
	}
	ParameterType_value = map[string]int32{
		"Variable": 0,
		"String":   1,
		"Integer":  2,
		"Float":    3,
		"Block":    4,
		"Boolean":  5,
		"Null":     6,
		"Enum":     7,
		"List":     8,
		"Object":   9,
	}
)

func (x ParameterType) Enum() *ParameterType {
	p := new(ParameterType)
	*p = x
	return p
}

func (x ParameterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParameterType) Descriptor() protoreflect.EnumDescriptor {
	return file_graphql_proto_enumTypes[0].Descriptor()
}

func (ParameterType) Type() protoreflect.EnumType {
	return &file_graphql_proto_enumTypes[0]
}

func (x ParameterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ParameterType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ParameterType(num)
	return nil
}

// Deprecated: Use ParameterType.Descriptor instead.
func (ParameterType) EnumDescriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_DEFAULT  Type = 0
	Type_MUTATION Type = 1
	Type_QUERY    Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "DEFAULT",
		1: "MUTATION",
		2: "QUERY",
	}
	Type_value = map[string]int32{
		"DEFAULT":  0,
		"MUTATION": 1,
		"QUERY":    2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_graphql_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_graphql_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Type(num)
	return nil
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{1}
}

type Upstream int32

const (
	Upstream_UPSTREAM_UNSPECIFIED Upstream = 0
	Upstream_UPSTREAM_SERVER      Upstream = 1
	Upstream_UPSTREAM_CLIENT      Upstream = 2
)

// Enum value maps for Upstream.
var (
	Upstream_name = map[int32]string{
		0: "UPSTREAM_UNSPECIFIED",
		1: "UPSTREAM_SERVER",
		2: "UPSTREAM_CLIENT",
	}
	Upstream_value = map[string]int32{
		"UPSTREAM_UNSPECIFIED": 0,
		"UPSTREAM_SERVER":      1,
		"UPSTREAM_CLIENT":      2,
	}
)

func (x Upstream) Enum() *Upstream {
	p := new(Upstream)
	*p = x
	return p
}

func (x Upstream) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Upstream) Descriptor() protoreflect.EnumDescriptor {
	return file_graphql_proto_enumTypes[2].Descriptor()
}

func (Upstream) Type() protoreflect.EnumType {
	return &file_graphql_proto_enumTypes[2]
}

func (x Upstream) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Upstream) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Upstream(num)
	return nil
}

// Deprecated: Use Upstream.Descriptor instead.
func (Upstream) EnumDescriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{2}
}

type Oneof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ignore *bool   `protobuf:"varint,4,opt,name=ignore" json:"ignore,omitempty"`
	Name   *string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
}

func (x *Oneof) Reset() {
	*x = Oneof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oneof) ProtoMessage() {}

func (x *Oneof) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oneof.ProtoReflect.Descriptor instead.
func (*Oneof) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{0}
}

func (x *Oneof) GetIgnore() bool {
	if x != nil && x.Ignore != nil {
		return *x.Ignore
	}
	return false
}

func (x *Oneof) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Required   *bool        `protobuf:"varint,1,opt,name=required" json:"required,omitempty"`
	Params     *string      `protobuf:"bytes,2,opt,name=params" json:"params,omitempty"`
	Dirs       *string      `protobuf:"bytes,3,opt,name=dirs" json:"dirs,omitempty"`
	Ignore     *bool        `protobuf:"varint,4,opt,name=ignore" json:"ignore,omitempty"`
	Name       *string      `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Parameters []*Parameter `protobuf:"bytes,6,rep,name=parameters" json:"parameters,omitempty"`
	Directives []*Directive `protobuf:"bytes,7,rep,name=directives" json:"directives,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetRequired() bool {
	if x != nil && x.Required != nil {
		return *x.Required
	}
	return false
}

func (x *Field) GetParams() string {
	if x != nil && x.Params != nil {
		return *x.Params
	}
	return ""
}

func (x *Field) GetDirs() string {
	if x != nil && x.Dirs != nil {
		return *x.Dirs
	}
	return ""
}

func (x *Field) GetIgnore() bool {
	if x != nil && x.Ignore != nil {
		return *x.Ignore
	}
	return false
}

func (x *Field) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Field) GetParameters() []*Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *Field) GetDirectives() []*Directive {
	if x != nil {
		return x.Directives
	}
	return nil
}

type Rpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       *Type        `protobuf:"varint,1,opt,name=type,enum=danielvladco.protobuf.graphql.Type" json:"type,omitempty"`
	Ignore     *bool        `protobuf:"varint,2,opt,name=ignore" json:"ignore,omitempty"`
	Name       *string      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Directives []*Directive `protobuf:"bytes,4,rep,name=directives" json:"directives,omitempty"`
}

func (x *Rpc) Reset() {
	*x = Rpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rpc) ProtoMessage() {}

func (x *Rpc) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rpc.ProtoReflect.Descriptor instead.
func (*Rpc) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{2}
}

func (x *Rpc) GetType() Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Type_DEFAULT
}

func (x *Rpc) GetIgnore() bool {
	if x != nil && x.Ignore != nil {
		return *x.Ignore
	}
	return false
}

func (x *Rpc) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Rpc) GetDirectives() []*Directive {
	if x != nil {
		return x.Directives
	}
	return nil
}

type Svc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       *Type        `protobuf:"varint,1,opt,name=type,enum=danielvladco.protobuf.graphql.Type" json:"type,omitempty"`
	Ignore     *bool        `protobuf:"varint,2,opt,name=ignore" json:"ignore,omitempty"`
	Name       *string      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Upstream   *Upstream    `protobuf:"varint,4,opt,name=upstream,enum=danielvladco.protobuf.graphql.Upstream" json:"upstream,omitempty"`
	Directives []*Directive `protobuf:"bytes,5,rep,name=directives" json:"directives,omitempty"`
}

func (x *Svc) Reset() {
	*x = Svc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Svc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Svc) ProtoMessage() {}

func (x *Svc) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Svc.ProtoReflect.Descriptor instead.
func (*Svc) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{3}
}

func (x *Svc) GetType() Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Type_DEFAULT
}

func (x *Svc) GetIgnore() bool {
	if x != nil && x.Ignore != nil {
		return *x.Ignore
	}
	return false
}

func (x *Svc) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Svc) GetUpstream() Upstream {
	if x != nil && x.Upstream != nil {
		return *x.Upstream
	}
	return Upstream_UPSTREAM_UNSPECIFIED
}

func (x *Svc) GetDirectives() []*Directive {
	if x != nil {
		return x.Directives
	}
	return nil
}

type Directive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Parameter []*Parameter `protobuf:"bytes,2,rep,name=parameter" json:"parameter,omitempty"`
}

func (x *Directive) Reset() {
	*x = Directive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directive) ProtoMessage() {}

func (x *Directive) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directive.ProtoReflect.Descriptor instead.
func (*Directive) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{4}
}

func (x *Directive) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Directive) GetParameter() []*Parameter {
	if x != nil {
		return x.Parameter
	}
	return nil
}

type Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string        `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Type  *ParameterType `protobuf:"varint,2,req,name=type,enum=danielvladco.protobuf.graphql.ParameterType" json:"type,omitempty"`
	Value *string        `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (x *Parameter) Reset() {
	*x = Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameter) ProtoMessage() {}

func (x *Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameter.ProtoReflect.Descriptor instead.
func (*Parameter) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{5}
}

func (x *Parameter) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Parameter) GetType() ParameterType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ParameterType_Variable
}

func (x *Parameter) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var file_graphql_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*Rpc)(nil),
		Field:         65030,
		Name:          "danielvladco.protobuf.graphql.rpc",
		Tag:           "bytes,65030,opt,name=rpc",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*Svc)(nil),
		Field:         65030,
		Name:          "danielvladco.protobuf.graphql.svc",
		Tag:           "bytes,65030,opt,name=svc",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*Field)(nil),
		Field:         65030,
		Name:          "danielvladco.protobuf.graphql.field",
		Tag:           "bytes,65030,opt,name=field",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.OneofOptions)(nil),
		ExtensionType: (*Oneof)(nil),
		Field:         65030,
		Name:          "danielvladco.protobuf.graphql.oneof",
		Tag:           "bytes,65030,opt,name=oneof",
		Filename:      "graphql.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// optional danielvladco.protobuf.graphql.Rpc rpc = 65030;
	E_Rpc = &file_graphql_proto_extTypes[0]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// optional danielvladco.protobuf.graphql.Svc svc = 65030;
	E_Svc = &file_graphql_proto_extTypes[1]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional danielvladco.protobuf.graphql.Field field = 65030;
	E_Field = &file_graphql_proto_extTypes[2]
)

// Extension fields to descriptor.OneofOptions.
var (
	// optional danielvladco.protobuf.graphql.Oneof oneof = 65030;
	E_Oneof = &file_graphql_proto_extTypes[3]
)

var File_graphql_proto protoreflect.FileDescriptor

var file_graphql_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x33, 0x0a, 0x05, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x69, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c,
	0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x48, 0x0a,
	0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x0a, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12,
	0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e,
	0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65,
	0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x22, 0xf9,
	0x01, 0x0a, 0x03, 0x53, 0x76, 0x63, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61,
	0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x55, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x48, 0x0a, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61,
	0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x0a,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x09, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x22, 0x77, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x83, 0x01, 0x0a,
	0x0d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c,
	0x0a, 0x08, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x42,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c,
	0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x10, 0x09, 0x2a, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x55, 0x54, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x02,
	0x2a, 0x4e, 0x0a, 0x08, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x14,
	0x55, 0x50, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x50, 0x53, 0x54, 0x52, 0x45,
	0x41, 0x4d, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x55,
	0x50, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x02,
	0x3a, 0x56, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x86, 0xfc, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e,
	0x52, 0x70, 0x63, 0x52, 0x03, 0x72, 0x70, 0x63, 0x3a, 0x57, 0x0a, 0x03, 0x73, 0x76, 0x63, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x86, 0xfc, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65,
	0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x53, 0x76, 0x63, 0x52, 0x03, 0x73, 0x76,
	0x63, 0x3a, 0x5b, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x86, 0xfc, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x5b,
	0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x86, 0xfc, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x4f,
	0x6e, 0x65, 0x6f, 0x66, 0x52, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c,
	0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x71, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x70,
	0x62,
}

var (
	file_graphql_proto_rawDescOnce sync.Once
	file_graphql_proto_rawDescData = file_graphql_proto_rawDesc
)

func file_graphql_proto_rawDescGZIP() []byte {
	file_graphql_proto_rawDescOnce.Do(func() {
		file_graphql_proto_rawDescData = protoimpl.X.CompressGZIP(file_graphql_proto_rawDescData)
	})
	return file_graphql_proto_rawDescData
}

var file_graphql_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_graphql_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_graphql_proto_goTypes = []interface{}{
	(ParameterType)(0),                // 0: danielvladco.protobuf.graphql.ParameterType
	(Type)(0),                         // 1: danielvladco.protobuf.graphql.Type
	(Upstream)(0),                     // 2: danielvladco.protobuf.graphql.Upstream
	(*Oneof)(nil),                     // 3: danielvladco.protobuf.graphql.Oneof
	(*Field)(nil),                     // 4: danielvladco.protobuf.graphql.Field
	(*Rpc)(nil),                       // 5: danielvladco.protobuf.graphql.Rpc
	(*Svc)(nil),                       // 6: danielvladco.protobuf.graphql.Svc
	(*Directive)(nil),                 // 7: danielvladco.protobuf.graphql.Directive
	(*Parameter)(nil),                 // 8: danielvladco.protobuf.graphql.Parameter
	(*descriptor.MethodOptions)(nil),  // 9: google.protobuf.MethodOptions
	(*descriptor.ServiceOptions)(nil), // 10: google.protobuf.ServiceOptions
	(*descriptor.FieldOptions)(nil),   // 11: google.protobuf.FieldOptions
	(*descriptor.OneofOptions)(nil),   // 12: google.protobuf.OneofOptions
}
var file_graphql_proto_depIdxs = []int32{
	8,  // 0: danielvladco.protobuf.graphql.Field.parameters:type_name -> danielvladco.protobuf.graphql.Parameter
	7,  // 1: danielvladco.protobuf.graphql.Field.directives:type_name -> danielvladco.protobuf.graphql.Directive
	1,  // 2: danielvladco.protobuf.graphql.Rpc.type:type_name -> danielvladco.protobuf.graphql.Type
	7,  // 3: danielvladco.protobuf.graphql.Rpc.directives:type_name -> danielvladco.protobuf.graphql.Directive
	1,  // 4: danielvladco.protobuf.graphql.Svc.type:type_name -> danielvladco.protobuf.graphql.Type
	2,  // 5: danielvladco.protobuf.graphql.Svc.upstream:type_name -> danielvladco.protobuf.graphql.Upstream
	7,  // 6: danielvladco.protobuf.graphql.Svc.directives:type_name -> danielvladco.protobuf.graphql.Directive
	8,  // 7: danielvladco.protobuf.graphql.Directive.parameter:type_name -> danielvladco.protobuf.graphql.Parameter
	0,  // 8: danielvladco.protobuf.graphql.Parameter.type:type_name -> danielvladco.protobuf.graphql.ParameterType
	9,  // 9: danielvladco.protobuf.graphql.rpc:extendee -> google.protobuf.MethodOptions
	10, // 10: danielvladco.protobuf.graphql.svc:extendee -> google.protobuf.ServiceOptions
	11, // 11: danielvladco.protobuf.graphql.field:extendee -> google.protobuf.FieldOptions
	12, // 12: danielvladco.protobuf.graphql.oneof:extendee -> google.protobuf.OneofOptions
	5,  // 13: danielvladco.protobuf.graphql.rpc:type_name -> danielvladco.protobuf.graphql.Rpc
	6,  // 14: danielvladco.protobuf.graphql.svc:type_name -> danielvladco.protobuf.graphql.Svc
	4,  // 15: danielvladco.protobuf.graphql.field:type_name -> danielvladco.protobuf.graphql.Field
	3,  // 16: danielvladco.protobuf.graphql.oneof:type_name -> danielvladco.protobuf.graphql.Oneof
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	13, // [13:17] is the sub-list for extension type_name
	9,  // [9:13] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_graphql_proto_init() }
func file_graphql_proto_init() {
	if File_graphql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graphql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oneof); i {
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
		file_graphql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_graphql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rpc); i {
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
		file_graphql_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Svc); i {
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
		file_graphql_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directive); i {
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
		file_graphql_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameter); i {
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
			RawDescriptor: file_graphql_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_graphql_proto_goTypes,
		DependencyIndexes: file_graphql_proto_depIdxs,
		EnumInfos:         file_graphql_proto_enumTypes,
		MessageInfos:      file_graphql_proto_msgTypes,
		ExtensionInfos:    file_graphql_proto_extTypes,
	}.Build()
	File_graphql_proto = out.File
	file_graphql_proto_rawDesc = nil
	file_graphql_proto_goTypes = nil
	file_graphql_proto_depIdxs = nil
}
