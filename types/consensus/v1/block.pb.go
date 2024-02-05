// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/consensus/v1/block.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BeaconKitBlock represents a generic beacon block that can be used to represent
// any beacon block in the system.
type BeaconKitBlock struct {
	// Beacon chain slot that this block represents.
	Slot github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,1,opt,name=slot,proto3,casttype=github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot" json:"slot,omitempty"`
	// BeaconBlockBody contains the body of the beacon block.
	//
	// Types that are valid to be assigned to Body:
	//
	//	*BeaconKitBlock_BlockBodyGeneric
	Body isBeaconKitBlock_Body `protobuf_oneof:"body"`
	// The payload value of the block.
	PayloadValue uint64 `protobuf:"varint,101,opt,name=payload_value,json=payloadValue,proto3" json:"payload_value,omitempty"`
}

func (m *BeaconKitBlock) Reset()         { *m = BeaconKitBlock{} }
func (m *BeaconKitBlock) String() string { return proto.CompactTextString(m) }
func (*BeaconKitBlock) ProtoMessage()    {}
func (*BeaconKitBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_b29fa1ac1aaec767, []int{0}
}
func (m *BeaconKitBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BeaconKitBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BeaconKitBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BeaconKitBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconKitBlock.Merge(m, src)
}
func (m *BeaconKitBlock) XXX_Size() int {
	return m.Size()
}
func (m *BeaconKitBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconKitBlock.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconKitBlock proto.InternalMessageInfo

type isBeaconKitBlock_Body interface {
	isBeaconKitBlock_Body()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BeaconKitBlock_BlockBodyGeneric struct {
	BlockBodyGeneric *BeaconBlockBody `protobuf:"bytes,2,opt,name=block_body_generic,json=blockBodyGeneric,proto3,oneof" json:"block_body_generic,omitempty"`
}

func (*BeaconKitBlock_BlockBodyGeneric) isBeaconKitBlock_Body() {}

func (m *BeaconKitBlock) GetBody() isBeaconKitBlock_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *BeaconKitBlock) GetSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *BeaconKitBlock) GetBlockBodyGeneric() *BeaconBlockBody {
	if x, ok := m.GetBody().(*BeaconKitBlock_BlockBodyGeneric); ok {
		return x.BlockBodyGeneric
	}
	return nil
}

func (m *BeaconKitBlock) GetPayloadValue() uint64 {
	if m != nil {
		return m.PayloadValue
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BeaconKitBlock) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BeaconKitBlock_BlockBodyGeneric)(nil),
	}
}

// BeaconBlockBody represents the body of a beacon block.
type BeaconBlockBody struct {
	// The validators RANDAO reveal 96 byte value.
	RandaoReveal []byte `protobuf:"bytes,1,opt,name=randao_reveal,json=randaoReveal,proto3" json:"randao_reveal,omitempty"`
	// 32 byte field of arbitrary data. This field may contain any data and
	// is not used for anything other than a fun message.
	Graffiti []byte `protobuf:"bytes,2,opt,name=graffiti,proto3" json:"graffiti,omitempty"`
	// Execution payload from the execution chain. New in Bellatrix network upgrade.
	ExecutionPayload []byte `protobuf:"bytes,3,opt,name=execution_payload,json=executionPayload,proto3" json:"execution_payload,omitempty"`
	// TODO: DEPRECATE WHEN WE BREAK OUT INTO MULTIPLE MESSAGES PER FORK.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *BeaconBlockBody) Reset()         { *m = BeaconBlockBody{} }
func (m *BeaconBlockBody) String() string { return proto.CompactTextString(m) }
func (*BeaconBlockBody) ProtoMessage()    {}
func (*BeaconBlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_b29fa1ac1aaec767, []int{1}
}
func (m *BeaconBlockBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BeaconBlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BeaconBlockBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BeaconBlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconBlockBody.Merge(m, src)
}
func (m *BeaconBlockBody) XXX_Size() int {
	return m.Size()
}
func (m *BeaconBlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconBlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconBlockBody proto.InternalMessageInfo

func (m *BeaconBlockBody) GetRandaoReveal() []byte {
	if m != nil {
		return m.RandaoReveal
	}
	return nil
}

func (m *BeaconBlockBody) GetGraffiti() []byte {
	if m != nil {
		return m.Graffiti
	}
	return nil
}

func (m *BeaconBlockBody) GetExecutionPayload() []byte {
	if m != nil {
		return m.ExecutionPayload
	}
	return nil
}

func (m *BeaconBlockBody) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*BeaconKitBlock)(nil), "types.consensus.v1.BeaconKitBlock")
	proto.RegisterType((*BeaconBlockBody)(nil), "types.consensus.v1.BeaconBlockBody")
}

func init() { proto.RegisterFile("types/consensus/v1/block.proto", fileDescriptor_b29fa1ac1aaec767) }

var fileDescriptor_b29fa1ac1aaec767 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x33, 0xde, 0x70, 0x95, 0x31, 0xea, 0x75, 0x70, 0x11, 0xba, 0x88, 0xa5, 0xdd, 0x14,
	0xc4, 0x8c, 0x55, 0x5f, 0xc0, 0x80, 0x28, 0xba, 0x91, 0x14, 0x04, 0xdd, 0x84, 0x99, 0x64, 0x1a,
	0x07, 0xa7, 0x39, 0x61, 0x66, 0x32, 0x98, 0xb7, 0x70, 0xeb, 0x1b, 0xb9, 0xec, 0xd2, 0x95, 0x48,
	0xbb, 0xf6, 0x05, 0x5c, 0x49, 0x26, 0x6d, 0x95, 0xdb, 0x5d, 0xfe, 0xff, 0x3f, 0xf9, 0xcf, 0x17,
	0x4e, 0x70, 0x62, 0xfb, 0x56, 0x18, 0x5a, 0x42, 0x63, 0x44, 0x63, 0x3a, 0x43, 0xdd, 0x92, 0x72,
	0x05, 0xe5, 0xe7, 0xb4, 0xd5, 0x60, 0x81, 0x10, 0x9f, 0xa7, 0xa7, 0x3c, 0x75, 0xcb, 0xc9, 0x83,
	0x1a, 0x6a, 0xf0, 0x31, 0x1d, 0x9e, 0xc6, 0xc9, 0xd9, 0x6f, 0x84, 0xef, 0x66, 0x82, 0x95, 0xd0,
	0xbc, 0x95, 0x36, 0x1b, 0x2a, 0xc8, 0x07, 0x1c, 0x1a, 0x05, 0x36, 0x46, 0x53, 0xb4, 0x08, 0xb3,
	0x97, 0x7f, 0x7e, 0x3e, 0x7c, 0x51, 0x4b, 0xfb, 0xa9, 0xe3, 0x69, 0x09, 0x1b, 0xda, 0xea, 0xde,
	0x6c, 0x98, 0x95, 0xa5, 0x62, 0xdc, 0x8c, 0x8a, 0xba, 0xe7, 0xff, 0x50, 0x1e, 0x8f, 0x68, 0xad,
	0x96, 0x1b, 0x69, 0xa5, 0x13, 0x26, 0x5d, 0x29, 0xb0, 0xb9, 0xaf, 0x24, 0x2b, 0x4c, 0x3c, 0x66,
	0xc1, 0xa1, 0xea, 0x8b, 0x5a, 0x34, 0x42, 0xcb, 0x32, 0xbe, 0x31, 0x45, 0x8b, 0xdb, 0x4f, 0xe7,
	0xe9, 0x39, 0x74, 0x3a, 0xa2, 0x79, 0xae, 0x0c, 0xaa, 0xfe, 0x75, 0x90, 0x5f, 0xf1, 0xa3, 0x78,
	0x35, 0xbe, 0x4e, 0xe6, 0xf8, 0x4e, 0xcb, 0x7a, 0x05, 0xac, 0x2a, 0x1c, 0x53, 0x9d, 0x88, 0xc5,
	0x00, 0x9e, 0x47, 0x07, 0xf3, 0xfd, 0xe0, 0x65, 0x97, 0x38, 0x1c, 0x76, 0xce, 0xbe, 0x21, 0x7c,
	0xef, 0x5a, 0xe9, 0x50, 0xa0, 0x59, 0x53, 0x31, 0x28, 0xb4, 0x70, 0x82, 0x29, 0xff, 0xe5, 0x51,
	0x1e, 0x8d, 0x66, 0xee, 0x3d, 0x32, 0xc1, 0xb7, 0x6a, 0xcd, 0xd6, 0x6b, 0x69, 0xa5, 0x07, 0x8e,
	0xf2, 0x93, 0x26, 0x8f, 0xf0, 0x7d, 0xf1, 0x45, 0x94, 0x9d, 0x95, 0xd0, 0x14, 0x87, 0xb5, 0xf1,
	0x85, 0x1f, 0xba, 0x3a, 0x05, 0xef, 0x46, 0x9f, 0xc4, 0xf8, 0xa6, 0x13, 0xda, 0x48, 0x68, 0xe2,
	0x70, 0x8a, 0x16, 0x17, 0xf9, 0x51, 0x66, 0x6f, 0xbe, 0xef, 0x12, 0xb4, 0xdd, 0x25, 0xe8, 0xd7,
	0x2e, 0x41, 0x5f, 0xf7, 0x49, 0xb0, 0xdd, 0x27, 0xc1, 0x8f, 0x7d, 0x12, 0x7c, 0x7c, 0xf2, 0xdf,
	0x01, 0xa4, 0x35, 0x95, 0x70, 0x5c, 0x30, 0x4d, 0x39, 0x28, 0xa6, 0xa5, 0xa1, 0xe7, 0x7f, 0x03,
	0xbf, 0xf4, 0xe7, 0x7d, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x79, 0x6b, 0x13, 0x2a, 0x02,
	0x00, 0x00,
}

func (m *BeaconKitBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BeaconKitBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BeaconKitBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PayloadValue != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.PayloadValue))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa8
	}
	if m.Body != nil {
		{
			size := m.Body.Size()
			i -= size
			if _, err := m.Body.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Slot != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BeaconKitBlock_BlockBodyGeneric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BeaconKitBlock_BlockBodyGeneric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BlockBodyGeneric != nil {
		{
			size, err := m.BlockBodyGeneric.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *BeaconBlockBody) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BeaconBlockBody) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BeaconBlockBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ExecutionPayload) > 0 {
		i -= len(m.ExecutionPayload)
		copy(dAtA[i:], m.ExecutionPayload)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.ExecutionPayload)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Graffiti) > 0 {
		i -= len(m.Graffiti)
		copy(dAtA[i:], m.Graffiti)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Graffiti)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RandaoReveal) > 0 {
		i -= len(m.RandaoReveal)
		copy(dAtA[i:], m.RandaoReveal)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.RandaoReveal)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BeaconKitBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Slot != 0 {
		n += 1 + sovBlock(uint64(m.Slot))
	}
	if m.Body != nil {
		n += m.Body.Size()
	}
	if m.PayloadValue != 0 {
		n += 2 + sovBlock(uint64(m.PayloadValue))
	}
	return n
}

func (m *BeaconKitBlock_BlockBodyGeneric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockBodyGeneric != nil {
		l = m.BlockBodyGeneric.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}
func (m *BeaconBlockBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RandaoReveal)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Graffiti)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.ExecutionPayload)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovBlock(uint64(m.Version))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BeaconKitBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BeaconKitBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BeaconKitBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockBodyGeneric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BeaconBlockBody{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Body = &BeaconKitBlock_BlockBodyGeneric{v}
			iNdEx = postIndex
		case 101:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadValue", wireType)
			}
			m.PayloadValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PayloadValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BeaconBlockBody) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BeaconBlockBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BeaconBlockBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandaoReveal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RandaoReveal = append(m.RandaoReveal[:0], dAtA[iNdEx:postIndex]...)
			if m.RandaoReveal == nil {
				m.RandaoReveal = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Graffiti", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Graffiti = append(m.Graffiti[:0], dAtA[iNdEx:postIndex]...)
			if m.Graffiti == nil {
				m.Graffiti = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutionPayload = append(m.ExecutionPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.ExecutionPayload == nil {
				m.ExecutionPayload = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
