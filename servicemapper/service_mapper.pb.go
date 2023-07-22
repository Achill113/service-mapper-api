// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: service_mapper/service_mapper.proto

package servicemapper

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

type CreateServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TenantId string  `protobuf:"bytes,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Version  *string `protobuf:"bytes,3,opt,name=version,proto3,oneof" json:"version,omitempty"`
}

func (x *CreateServiceRequest) Reset() {
	*x = CreateServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mapper_service_mapper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceRequest) ProtoMessage() {}

func (x *CreateServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_mapper_service_mapper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return file_service_mapper_service_mapper_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServiceRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *CreateServiceRequest) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type CreateServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TenantId string  `protobuf:"bytes,3,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Version  *string `protobuf:"bytes,4,opt,name=version,proto3,oneof" json:"version,omitempty"`
}

func (x *CreateServiceResponse) Reset() {
	*x = CreateServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mapper_service_mapper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceResponse) ProtoMessage() {}

func (x *CreateServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_mapper_service_mapper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceResponse.ProtoReflect.Descriptor instead.
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return file_service_mapper_service_mapper_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServiceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateServiceResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServiceResponse) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *CreateServiceResponse) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

var File_service_mapper_service_mapper_proto protoreflect.FileDescriptor

var file_service_mapper_service_mapper_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x6d, 0x0a, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x5c, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x42, 0x12, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x18, 0x6e, 0x6f, 0x72, 0x64, 0x69, 0x63, 0x68, 0x69, 0x6c, 0x6c, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_mapper_service_mapper_proto_rawDescOnce sync.Once
	file_service_mapper_service_mapper_proto_rawDescData = file_service_mapper_service_mapper_proto_rawDesc
)

func file_service_mapper_service_mapper_proto_rawDescGZIP() []byte {
	file_service_mapper_service_mapper_proto_rawDescOnce.Do(func() {
		file_service_mapper_service_mapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_mapper_service_mapper_proto_rawDescData)
	})
	return file_service_mapper_service_mapper_proto_rawDescData
}

var file_service_mapper_service_mapper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_mapper_service_mapper_proto_goTypes = []interface{}{
	(*CreateServiceRequest)(nil),  // 0: serviceMapper.CreateServiceRequest
	(*CreateServiceResponse)(nil), // 1: serviceMapper.CreateServiceResponse
}
var file_service_mapper_service_mapper_proto_depIdxs = []int32{
	0, // 0: serviceMapper.ServiceMapper.CreateService:input_type -> serviceMapper.CreateServiceRequest
	1, // 1: serviceMapper.ServiceMapper.CreateService:output_type -> serviceMapper.CreateServiceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_mapper_service_mapper_proto_init() }
func file_service_mapper_service_mapper_proto_init() {
	if File_service_mapper_service_mapper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_mapper_service_mapper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceRequest); i {
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
		file_service_mapper_service_mapper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceResponse); i {
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
	file_service_mapper_service_mapper_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_service_mapper_service_mapper_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_mapper_service_mapper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_mapper_service_mapper_proto_goTypes,
		DependencyIndexes: file_service_mapper_service_mapper_proto_depIdxs,
		MessageInfos:      file_service_mapper_service_mapper_proto_msgTypes,
	}.Build()
	File_service_mapper_service_mapper_proto = out.File
	file_service_mapper_service_mapper_proto_rawDesc = nil
	file_service_mapper_service_mapper_proto_goTypes = nil
	file_service_mapper_service_mapper_proto_depIdxs = nil
}
