// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: beacon/core/types/v1/deposit.proto

package typesv1

import (
	_ "github.com/prysmaticlabs/prysm/v5/proto/eth/ext"
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

// Deposit into the consensus layer from the deposit contract in the execution layer.
type Deposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key of the validator, which is compatible to the implementations
	// of the PubKey interface in Cosmos SDK. 32-byte ed25519 public key is preferred.
	ValidatorPubkey []byte `protobuf:"bytes,1,opt,name=validator_pubkey,json=validatorPubkey,proto3" json:"validator_pubkey,omitempty" spec-name:"pubkey" ssz-max:"48"`
	// A staking credentials with
	// 1 byte prefix + 11 bytes padding + 20 bytes address = 32 bytes.
	StakingCredentials []byte `protobuf:"bytes,2,opt,name=staking_credentials,json=stakingCredentials,proto3" json:"staking_credentials,omitempty" ssz-size:"32"`
	// Deposit amount in gwei.
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// Signature of the deposit data.
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty" ssz-max:"96"`
}

func (x *Deposit) Reset() {
	*x = Deposit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beacon_core_types_v1_deposit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deposit) ProtoMessage() {}

func (x *Deposit) ProtoReflect() protoreflect.Message {
	mi := &file_beacon_core_types_v1_deposit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deposit.ProtoReflect.Descriptor instead.
func (*Deposit) Descriptor() ([]byte, []int) {
	return file_beacon_core_types_v1_deposit_proto_rawDescGZIP(), []int{0}
}

func (x *Deposit) GetValidatorPubkey() []byte {
	if x != nil {
		return x.ValidatorPubkey
	}
	return nil
}

func (x *Deposit) GetStakingCredentials() []byte {
	if x != nil {
		return x.StakingCredentials
	}
	return nil
}

func (x *Deposit) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Deposit) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_beacon_core_types_v1_deposit_proto protoreflect.FileDescriptor

var file_beacon_core_types_v1_deposit_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x07, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x10, 0x92, 0xb5, 0x18, 0x02, 0x34, 0x38, 0x9a, 0xb5, 0x18, 0x06, 0x70, 0x75, 0x62, 0x6b,
	0x65, 0x79, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x12, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x92, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x72, 0x61, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x62, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_beacon_core_types_v1_deposit_proto_rawDescOnce sync.Once
	file_beacon_core_types_v1_deposit_proto_rawDescData = file_beacon_core_types_v1_deposit_proto_rawDesc
)

func file_beacon_core_types_v1_deposit_proto_rawDescGZIP() []byte {
	file_beacon_core_types_v1_deposit_proto_rawDescOnce.Do(func() {
		file_beacon_core_types_v1_deposit_proto_rawDescData = protoimpl.X.CompressGZIP(file_beacon_core_types_v1_deposit_proto_rawDescData)
	})
	return file_beacon_core_types_v1_deposit_proto_rawDescData
}

var file_beacon_core_types_v1_deposit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beacon_core_types_v1_deposit_proto_goTypes = []interface{}{
	(*Deposit)(nil), // 0: beacon.core.types.v1.Deposit
}
var file_beacon_core_types_v1_deposit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_beacon_core_types_v1_deposit_proto_init() }
func file_beacon_core_types_v1_deposit_proto_init() {
	if File_beacon_core_types_v1_deposit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beacon_core_types_v1_deposit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deposit); i {
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
			RawDescriptor: file_beacon_core_types_v1_deposit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beacon_core_types_v1_deposit_proto_goTypes,
		DependencyIndexes: file_beacon_core_types_v1_deposit_proto_depIdxs,
		MessageInfos:      file_beacon_core_types_v1_deposit_proto_msgTypes,
	}.Build()
	File_beacon_core_types_v1_deposit_proto = out.File
	file_beacon_core_types_v1_deposit_proto_rawDesc = nil
	file_beacon_core_types_v1_deposit_proto_goTypes = nil
	file_beacon_core_types_v1_deposit_proto_depIdxs = nil
}
