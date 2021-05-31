// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: spire/api/agent/privileged/v1/privileged.proto

package privilegedv1

import (
	proto "github.com/golang/protobuf/proto"
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FetchX509SVIDsBySelectorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selectors []*types.Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
}

func (x *FetchX509SVIDsBySelectorsRequest) Reset() {
	*x = FetchX509SVIDsBySelectorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchX509SVIDsBySelectorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchX509SVIDsBySelectorsRequest) ProtoMessage() {}

func (x *FetchX509SVIDsBySelectorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchX509SVIDsBySelectorsRequest.ProtoReflect.Descriptor instead.
func (*FetchX509SVIDsBySelectorsRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_privileged_v1_privileged_proto_rawDescGZIP(), []int{0}
}

func (x *FetchX509SVIDsBySelectorsRequest) GetSelectors() []*types.Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

type FetchX509SVIDsBySelectorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X509Svids        []*X509SVIDWithKey `protobuf:"bytes,1,rep,name=x509_svids,json=x509Svids,proto3" json:"x509_svids,omitempty"`
	Bundle           []byte             `protobuf:"bytes,2,opt,name=bundle,proto3" json:"bundle,omitempty"`
	FederatedBundles map[string][]byte  `protobuf:"bytes,3,rep,name=federated_bundles,json=federatedBundles,proto3" json:"federated_bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FetchX509SVIDsBySelectorsResponse) Reset() {
	*x = FetchX509SVIDsBySelectorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchX509SVIDsBySelectorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchX509SVIDsBySelectorsResponse) ProtoMessage() {}

func (x *FetchX509SVIDsBySelectorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchX509SVIDsBySelectorsResponse.ProtoReflect.Descriptor instead.
func (*FetchX509SVIDsBySelectorsResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_privileged_v1_privileged_proto_rawDescGZIP(), []int{1}
}

func (x *FetchX509SVIDsBySelectorsResponse) GetX509Svids() []*X509SVIDWithKey {
	if x != nil {
		return x.X509Svids
	}
	return nil
}

func (x *FetchX509SVIDsBySelectorsResponse) GetBundle() []byte {
	if x != nil {
		return x.Bundle
	}
	return nil
}

func (x *FetchX509SVIDsBySelectorsResponse) GetFederatedBundles() map[string][]byte {
	if x != nil {
		return x.FederatedBundles
	}
	return nil
}

// X.509 SPIFFE Verifiable Identity Document with the private key.
type X509SVIDWithKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X509Svid    *types.X509SVID `protobuf:"bytes,1,opt,name=x509_svid,json=x509Svid,proto3" json:"x509_svid,omitempty"`
	X509SvidKey []byte          `protobuf:"bytes,2,opt,name=x509_svid_key,json=x509SvidKey,proto3" json:"x509_svid_key,omitempty"`
}

func (x *X509SVIDWithKey) Reset() {
	*x = X509SVIDWithKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509SVIDWithKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509SVIDWithKey) ProtoMessage() {}

func (x *X509SVIDWithKey) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509SVIDWithKey.ProtoReflect.Descriptor instead.
func (*X509SVIDWithKey) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_privileged_v1_privileged_proto_rawDescGZIP(), []int{2}
}

func (x *X509SVIDWithKey) GetX509Svid() *types.X509SVID {
	if x != nil {
		return x.X509Svid
	}
	return nil
}

func (x *X509SVIDWithKey) GetX509SvidKey() []byte {
	if x != nil {
		return x.X509SvidKey
	}
	return nil
}

var File_spire_api_agent_privileged_v1_privileged_proto protoreflect.FileDescriptor

var file_spire_api_agent_privileged_v1_privileged_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a,
	0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x78, 0x35, 0x30, 0x39, 0x73, 0x76, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5b, 0x0a, 0x20, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44,
	0x73, 0x42, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0xd5, 0x02, 0x0a,
	0x21, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x42,
	0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x57,
	0x69, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x11, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53,
	0x56, 0x49, 0x44, 0x73, 0x42, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x1a,
	0x43, 0x0a, 0x15, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x6d, 0x0a, 0x0f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44,
	0x57, 0x69, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x78, 0x35, 0x30, 0x39, 0x5f,
	0x73, 0x76, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x58, 0x35, 0x30,
	0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x08, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x78, 0x35, 0x30, 0x39, 0x53, 0x76, 0x69, 0x64,
	0x4b, 0x65, 0x79, 0x32, 0xaf, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x64, 0x12, 0xa0, 0x01, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39,
	0x53, 0x56, 0x49, 0x44, 0x73, 0x42, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x3f, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x42,
	0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x40, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73,
	0x42, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spire_api_agent_privileged_v1_privileged_proto_rawDescOnce sync.Once
	file_spire_api_agent_privileged_v1_privileged_proto_rawDescData = file_spire_api_agent_privileged_v1_privileged_proto_rawDesc
)

func file_spire_api_agent_privileged_v1_privileged_proto_rawDescGZIP() []byte {
	file_spire_api_agent_privileged_v1_privileged_proto_rawDescOnce.Do(func() {
		file_spire_api_agent_privileged_v1_privileged_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_agent_privileged_v1_privileged_proto_rawDescData)
	})
	return file_spire_api_agent_privileged_v1_privileged_proto_rawDescData
}

var file_spire_api_agent_privileged_v1_privileged_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spire_api_agent_privileged_v1_privileged_proto_goTypes = []interface{}{
	(*FetchX509SVIDsBySelectorsRequest)(nil),  // 0: spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsRequest
	(*FetchX509SVIDsBySelectorsResponse)(nil), // 1: spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsResponse
	(*X509SVIDWithKey)(nil),                   // 2: spire.api.agent.privileged.v1.X509SVIDWithKey
	nil,                                       // 3: spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsResponse.FederatedBundlesEntry
	(*types.Selector)(nil),                    // 4: spire.api.types.Selector
	(*types.X509SVID)(nil),                    // 5: spire.api.types.X509SVID
}
var file_spire_api_agent_privileged_v1_privileged_proto_depIdxs = []int32{
	4, // 0: spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsRequest.selectors:type_name -> spire.api.types.Selector
	2, // 1: spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsResponse.x509_svids:type_name -> spire.api.agent.privileged.v1.X509SVIDWithKey
	3, // 2: spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsResponse.federated_bundles:type_name -> spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsResponse.FederatedBundlesEntry
	5, // 3: spire.api.agent.privileged.v1.X509SVIDWithKey.x509_svid:type_name -> spire.api.types.X509SVID
	0, // 4: spire.api.agent.privileged.v1.Privileged.FetchX509SVIDsBySelectors:input_type -> spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsRequest
	1, // 5: spire.api.agent.privileged.v1.Privileged.FetchX509SVIDsBySelectors:output_type -> spire.api.agent.privileged.v1.FetchX509SVIDsBySelectorsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spire_api_agent_privileged_v1_privileged_proto_init() }
func file_spire_api_agent_privileged_v1_privileged_proto_init() {
	if File_spire_api_agent_privileged_v1_privileged_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchX509SVIDsBySelectorsRequest); i {
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
		file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchX509SVIDsBySelectorsResponse); i {
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
		file_spire_api_agent_privileged_v1_privileged_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509SVIDWithKey); i {
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
			RawDescriptor: file_spire_api_agent_privileged_v1_privileged_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_api_agent_privileged_v1_privileged_proto_goTypes,
		DependencyIndexes: file_spire_api_agent_privileged_v1_privileged_proto_depIdxs,
		MessageInfos:      file_spire_api_agent_privileged_v1_privileged_proto_msgTypes,
	}.Build()
	File_spire_api_agent_privileged_v1_privileged_proto = out.File
	file_spire_api_agent_privileged_v1_privileged_proto_rawDesc = nil
	file_spire_api_agent_privileged_v1_privileged_proto_goTypes = nil
	file_spire_api_agent_privileged_v1_privileged_proto_depIdxs = nil
}
