// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: index/index.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Schema struct {
	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty" yaml:"owner_address"`
	SchemaName   string `protobuf:"bytes,2,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty" yaml:"schema_name"`
	Body         string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty" yaml:"body"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_91013751fa82c1bb, []int{0}
}
func (m *Schema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return m.Size()
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *Schema) GetSchemaName() string {
	if m != nil {
		return m.SchemaName
	}
	return ""
}

func (m *Schema) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Doc struct {
	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty" yaml:"owner_address"`
	SchemaName   string `protobuf:"bytes,2,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty" yaml:"schema_name"`
	Body         string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty" yaml:"body"`
}

func (m *Doc) Reset()         { *m = Doc{} }
func (m *Doc) String() string { return proto.CompactTextString(m) }
func (*Doc) ProtoMessage()    {}
func (*Doc) Descriptor() ([]byte, []int) {
	return fileDescriptor_91013751fa82c1bb, []int{1}
}
func (m *Doc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Doc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Doc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Doc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Doc.Merge(m, src)
}
func (m *Doc) XXX_Size() int {
	return m.Size()
}
func (m *Doc) XXX_DiscardUnknown() {
	xxx_messageInfo_Doc.DiscardUnknown(m)
}

var xxx_messageInfo_Doc proto.InternalMessageInfo

func (m *Doc) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *Doc) GetSchemaName() string {
	if m != nil {
		return m.SchemaName
	}
	return ""
}

func (m *Doc) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Schema)(nil), "blockved.glitterchain.index.Schema")
	proto.RegisterType((*Doc)(nil), "blockved.glitterchain.index.Doc")
}

func init() { proto.RegisterFile("index/index.proto", fileDescriptor_91013751fa82c1bb) }

var fileDescriptor_91013751fa82c1bb = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcc, 0x4b, 0x49,
	0xad, 0xd0, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xd2, 0x49, 0x39, 0xf9, 0xc9,
	0xd9, 0x65, 0xa9, 0x29, 0x7a, 0xe9, 0x39, 0x99, 0x25, 0x25, 0xa9, 0x45, 0xc9, 0x19, 0x89, 0x99,
	0x79, 0x7a, 0x60, 0x25, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x75, 0xfa, 0x20, 0x16, 0x44,
	0x8b, 0xd2, 0x52, 0x46, 0x2e, 0xb6, 0xe0, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x21, 0x5b, 0x2e, 0xde,
	0xfc, 0xf2, 0xbc, 0xd4, 0xa2, 0xf8, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x4e, 0x27, 0x89, 0x4f, 0xf7, 0xe4, 0x45, 0x2a, 0x13, 0x73, 0x73, 0xac, 0x94, 0x50,
	0xa4, 0x95, 0x82, 0x78, 0xc0, 0x7c, 0x47, 0x08, 0x57, 0xc8, 0x9c, 0x8b, 0xbb, 0x18, 0x6c, 0x50,
	0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0xb3, 0xd8, 0xa7, 0x7b, 0xf2, 0x42, 0x10, 0xcd,
	0x48, 0x92, 0x4a, 0x41, 0x5c, 0x10, 0x9e, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x32, 0x17, 0x4b, 0x52,
	0x7e, 0x4a, 0xa5, 0x04, 0x33, 0x58, 0x07, 0xff, 0xa7, 0x7b, 0xf2, 0xdc, 0x10, 0x1d, 0x20, 0x51,
	0xa5, 0x20, 0xb0, 0xa4, 0xd2, 0x22, 0x46, 0x2e, 0x66, 0x97, 0xfc, 0xe4, 0x41, 0xed, 0x48, 0x27,
	0xcf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x45, 0x92, 0x3e, 0x34, 0x92, 0x74, 0xc1, 0xb1,
	0xa4, 0x0f, 0x8d, 0x4a, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xf4, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x01, 0xdc, 0xd7, 0x73, 0xe6, 0x01, 0x00, 0x00,
}

func (m *Schema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SchemaName) > 0 {
		i -= len(m.SchemaName)
		copy(dAtA[i:], m.SchemaName)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.SchemaName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Doc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Doc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Doc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SchemaName) > 0 {
		i -= len(m.SchemaName)
		copy(dAtA[i:], m.SchemaName)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.SchemaName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndex(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndex(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Schema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	l = len(m.SchemaName)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	return n
}

func (m *Doc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	l = len(m.SchemaName)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	return n
}

func sovIndex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndex(x uint64) (n int) {
	return sovIndex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Schema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: Schema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndex
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
func (m *Doc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: Doc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Doc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndex
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
func skipIndex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndex
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
					return 0, ErrIntOverflowIndex
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
					return 0, ErrIntOverflowIndex
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
				return 0, ErrInvalidLengthIndex
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndex
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndex
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndex        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndex          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndex = fmt.Errorf("proto: unexpected end of group")
)