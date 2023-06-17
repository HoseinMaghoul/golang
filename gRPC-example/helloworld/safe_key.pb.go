// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: helloworld/safe_key.proto

package safekey

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

type AgentProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName      string                     `protobuf:"bytes,1,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	IpAddress     string                     `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	MacAddress    string                     `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	AgentMetadata *AgentMetadataProtoMessage `protobuf:"bytes,4,opt,name=agent_metadata,json=agentMetadata,proto3" json:"agent_metadata,omitempty"`
}

func (x *AgentProtoMessage) Reset() {
	*x = AgentProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_safe_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentProtoMessage) ProtoMessage() {}

func (x *AgentProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_safe_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentProtoMessage.ProtoReflect.Descriptor instead.
func (*AgentProtoMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_safe_key_proto_rawDescGZIP(), []int{0}
}

func (x *AgentProtoMessage) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *AgentProtoMessage) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *AgentProtoMessage) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *AgentProtoMessage) GetAgentMetadata() *AgentMetadataProtoMessage {
	if x != nil {
		return x.AgentMetadata
	}
	return nil
}

type AgentMetadataProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAiTrained    bool   `protobuf:"varint,1,opt,name=is_ai_trained,json=isAiTrained,proto3" json:"is_ai_trained,omitempty"`
	RegisteredTime string `protobuf:"bytes,2,opt,name=registered_time,json=registeredTime,proto3" json:"registered_time,omitempty"`
	IsRegistered   bool   `protobuf:"varint,3,opt,name=is_registered,json=isRegistered,proto3" json:"is_registered,omitempty"`
}

func (x *AgentMetadataProtoMessage) Reset() {
	*x = AgentMetadataProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_safe_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentMetadataProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMetadataProtoMessage) ProtoMessage() {}

func (x *AgentMetadataProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_safe_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMetadataProtoMessage.ProtoReflect.Descriptor instead.
func (*AgentMetadataProtoMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_safe_key_proto_rawDescGZIP(), []int{1}
}

func (x *AgentMetadataProtoMessage) GetIsAiTrained() bool {
	if x != nil {
		return x.IsAiTrained
	}
	return false
}

func (x *AgentMetadataProtoMessage) GetRegisteredTime() string {
	if x != nil {
		return x.RegisteredTime
	}
	return ""
}

func (x *AgentMetadataProtoMessage) GetIsRegistered() bool {
	if x != nil {
		return x.IsRegistered
	}
	return false
}

type SafeKeyProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string             `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	HashKey     string             `protobuf:"bytes,2,opt,name=hash_key,json=hashKey,proto3" json:"hash_key,omitempty"`
	Name        string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ExpiresDate string             `protobuf:"bytes,4,opt,name=expires_date,json=expiresDate,proto3" json:"expires_date,omitempty"`
	Agent       *AgentProtoMessage `protobuf:"bytes,5,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *SafeKeyProtoMessage) Reset() {
	*x = SafeKeyProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_safe_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SafeKeyProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SafeKeyProtoMessage) ProtoMessage() {}

func (x *SafeKeyProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_safe_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SafeKeyProtoMessage.ProtoReflect.Descriptor instead.
func (*SafeKeyProtoMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_safe_key_proto_rawDescGZIP(), []int{2}
}

func (x *SafeKeyProtoMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SafeKeyProtoMessage) GetHashKey() string {
	if x != nil {
		return x.HashKey
	}
	return ""
}

func (x *SafeKeyProtoMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SafeKeyProtoMessage) GetExpiresDate() string {
	if x != nil {
		return x.ExpiresDate
	}
	return ""
}

func (x *SafeKeyProtoMessage) GetAgent() *AgentProtoMessage {
	if x != nil {
		return x.Agent
	}
	return nil
}

type SafeKeyNameProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SafeKeyNameProtoMessage) Reset() {
	*x = SafeKeyNameProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_safe_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SafeKeyNameProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SafeKeyNameProtoMessage) ProtoMessage() {}

func (x *SafeKeyNameProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_safe_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SafeKeyNameProtoMessage.ProtoReflect.Descriptor instead.
func (*SafeKeyNameProtoMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_safe_key_proto_rawDescGZIP(), []int{3}
}

func (x *SafeKeyNameProtoMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HashKeyProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Hash    string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *HashKeyProtoMessage) Reset() {
	*x = HashKeyProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_safe_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashKeyProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashKeyProtoMessage) ProtoMessage() {}

func (x *HashKeyProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_safe_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashKeyProtoMessage.ProtoReflect.Descriptor instead.
func (*HashKeyProtoMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_safe_key_proto_rawDescGZIP(), []int{4}
}

func (x *HashKeyProtoMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HashKeyProtoMessage) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_helloworld_safe_key_proto protoreflect.FileDescriptor

var file_helloworld_safe_key_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x73, 0x61, 0x66,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x11,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x41,
	0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x8d, 0x01, 0x0a, 0x19, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x61, 0x69, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x69, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x22, 0xab, 0x01, 0x0a, 0x13, 0x53, 0x61, 0x66, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22,
	0x2d, 0x0a, 0x17, 0x53, 0x61, 0x66, 0x65, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43,
	0x0a, 0x13, 0x48, 0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x32, 0x9a, 0x01, 0x0a, 0x13, 0x53, 0x61, 0x66, 0x65, 0x4b, 0x65, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x15, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x61, 0x66, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x4b, 0x65, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x61, 0x66, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x53, 0x61, 0x66,
	0x65, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x6b, 0x65, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_safe_key_proto_rawDescOnce sync.Once
	file_helloworld_safe_key_proto_rawDescData = file_helloworld_safe_key_proto_rawDesc
)

func file_helloworld_safe_key_proto_rawDescGZIP() []byte {
	file_helloworld_safe_key_proto_rawDescOnce.Do(func() {
		file_helloworld_safe_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_safe_key_proto_rawDescData)
	})
	return file_helloworld_safe_key_proto_rawDescData
}

var file_helloworld_safe_key_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_helloworld_safe_key_proto_goTypes = []interface{}{
	(*AgentProtoMessage)(nil),         // 0: AgentProtoMessage
	(*AgentMetadataProtoMessage)(nil), // 1: AgentMetadataProtoMessage
	(*SafeKeyProtoMessage)(nil),       // 2: SafeKeyProtoMessage
	(*SafeKeyNameProtoMessage)(nil),   // 3: SafeKeyNameProtoMessage
	(*HashKeyProtoMessage)(nil),       // 4: HashKeyProtoMessage
}
var file_helloworld_safe_key_proto_depIdxs = []int32{
	1, // 0: AgentProtoMessage.agent_metadata:type_name -> AgentMetadataProtoMessage
	0, // 1: SafeKeyProtoMessage.agent:type_name -> AgentProtoMessage
	3, // 2: SafeKeyProtoService.generate_new_safe_key:input_type -> SafeKeyNameProtoMessage
	4, // 3: SafeKeyProtoService.get_safe_key:input_type -> HashKeyProtoMessage
	4, // 4: SafeKeyProtoService.generate_new_safe_key:output_type -> HashKeyProtoMessage
	2, // 5: SafeKeyProtoService.get_safe_key:output_type -> SafeKeyProtoMessage
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_helloworld_safe_key_proto_init() }
func file_helloworld_safe_key_proto_init() {
	if File_helloworld_safe_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_safe_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentProtoMessage); i {
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
		file_helloworld_safe_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentMetadataProtoMessage); i {
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
		file_helloworld_safe_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SafeKeyProtoMessage); i {
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
		file_helloworld_safe_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SafeKeyNameProtoMessage); i {
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
		file_helloworld_safe_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashKeyProtoMessage); i {
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
			RawDescriptor: file_helloworld_safe_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_safe_key_proto_goTypes,
		DependencyIndexes: file_helloworld_safe_key_proto_depIdxs,
		MessageInfos:      file_helloworld_safe_key_proto_msgTypes,
	}.Build()
	File_helloworld_safe_key_proto = out.File
	file_helloworld_safe_key_proto_rawDesc = nil
	file_helloworld_safe_key_proto_goTypes = nil
	file_helloworld_safe_key_proto_depIdxs = nil
}
