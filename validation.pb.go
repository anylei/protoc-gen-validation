// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validation.proto

package validation

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FieldValidation struct {
	NotEmptyString       *bool    `protobuf:"varint,1,opt,name=not_empty_string,json=notEmptyString" json:"not_empty_string,omitempty"`
	Matches              *string  `protobuf:"bytes,2,opt,name=matches" json:"matches,omitempty"`
	Contains             *string  `protobuf:"bytes,3,opt,name=contains" json:"contains,omitempty"`
	Regex                *string  `protobuf:"bytes,4,opt,name=regex" json:"regex,omitempty"`
	IntLte               *int64   `protobuf:"varint,5,opt,name=int_lte,json=intLte" json:"int_lte,omitempty"`
	IntGte               *int64   `protobuf:"varint,6,opt,name=int_gte,json=intGte" json:"int_gte,omitempty"`
	IntEq                *int64   `protobuf:"varint,7,opt,name=int_eq,json=intEq" json:"int_eq,omitempty"`
	FloatLte             *float64 `protobuf:"fixed64,8,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	FloatGte             *float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	FloatEq              *float64 `protobuf:"fixed64,10,opt,name=float_eq,json=floatEq" json:"float_eq,omitempty"`
	MinLen               *int64   `protobuf:"varint,11,opt,name=min_len,json=minLen" json:"min_len,omitempty"`
	MaxLen               *int64   `protobuf:"varint,12,opt,name=max_len,json=maxLen" json:"max_len,omitempty"`
	EqLen                *int64   `protobuf:"varint,13,opt,name=eq_len,json=eqLen" json:"eq_len,omitempty"`
	Error                *string  `protobuf:"bytes,14,opt,name=error" json:"error,omitempty"`
	IsUuid               *bool    `protobuf:"varint,15,opt,name=is_uuid,json=isUuid" json:"is_uuid,omitempty"`
	IsEmail              *bool    `protobuf:"varint,16,opt,name=is_email,json=isEmail" json:"is_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldValidation) Reset()         { *m = FieldValidation{} }
func (m *FieldValidation) String() string { return proto.CompactTextString(m) }
func (*FieldValidation) ProtoMessage()    {}
func (*FieldValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{0}
}
func (m *FieldValidation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FieldValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FieldValidation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FieldValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValidation.Merge(m, src)
}
func (m *FieldValidation) XXX_Size() int {
	return m.Size()
}
func (m *FieldValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValidation.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValidation proto.InternalMessageInfo

func (m *FieldValidation) GetNotEmptyString() bool {
	if m != nil && m.NotEmptyString != nil {
		return *m.NotEmptyString
	}
	return false
}

func (m *FieldValidation) GetMatches() string {
	if m != nil && m.Matches != nil {
		return *m.Matches
	}
	return ""
}

func (m *FieldValidation) GetContains() string {
	if m != nil && m.Contains != nil {
		return *m.Contains
	}
	return ""
}

func (m *FieldValidation) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidation) GetIntLte() int64 {
	if m != nil && m.IntLte != nil {
		return *m.IntLte
	}
	return 0
}

func (m *FieldValidation) GetIntGte() int64 {
	if m != nil && m.IntGte != nil {
		return *m.IntGte
	}
	return 0
}

func (m *FieldValidation) GetIntEq() int64 {
	if m != nil && m.IntEq != nil {
		return *m.IntEq
	}
	return 0
}

func (m *FieldValidation) GetFloatLte() float64 {
	if m != nil && m.FloatLte != nil {
		return *m.FloatLte
	}
	return 0
}

func (m *FieldValidation) GetFloatGte() float64 {
	if m != nil && m.FloatGte != nil {
		return *m.FloatGte
	}
	return 0
}

func (m *FieldValidation) GetFloatEq() float64 {
	if m != nil && m.FloatEq != nil {
		return *m.FloatEq
	}
	return 0
}

func (m *FieldValidation) GetMinLen() int64 {
	if m != nil && m.MinLen != nil {
		return *m.MinLen
	}
	return 0
}

func (m *FieldValidation) GetMaxLen() int64 {
	if m != nil && m.MaxLen != nil {
		return *m.MaxLen
	}
	return 0
}

func (m *FieldValidation) GetEqLen() int64 {
	if m != nil && m.EqLen != nil {
		return *m.EqLen
	}
	return 0
}

func (m *FieldValidation) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *FieldValidation) GetIsUuid() bool {
	if m != nil && m.IsUuid != nil {
		return *m.IsUuid
	}
	return false
}

func (m *FieldValidation) GetIsEmail() bool {
	if m != nil && m.IsEmail != nil {
		return *m.IsEmail
	}
	return false
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldValidation)(nil),
	Field:         61032,
	Name:          "validation.field",
	Tag:           "bytes,61032,opt,name=field",
	Filename:      "validation.proto",
}

func init() {
	proto.RegisterType((*FieldValidation)(nil), "validation.FieldValidation")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("validation.proto", fileDescriptor_bfc2ab0b60b7792f) }

var fileDescriptor_bfc2ab0b60b7792f = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x65, 0x4a, 0x9a, 0xd6, 0x0b, 0xbb, 0x95, 0xb5, 0x2b, 0xcc, 0xae, 0xa8, 0x22, 0x4e,
	0x39, 0x65, 0x25, 0x8e, 0x1c, 0x91, 0x0a, 0x97, 0x4a, 0x48, 0x41, 0x70, 0xe0, 0x12, 0x85, 0x66,
	0x1a, 0x46, 0x4a, 0xec, 0xc6, 0x9e, 0xa2, 0xf2, 0x66, 0x3c, 0x02, 0x47, 0x1e, 0x80, 0x03, 0xea,
	0x89, 0xc7, 0x40, 0x1e, 0xd3, 0x50, 0x71, 0xf3, 0xff, 0x7f, 0xc9, 0xaf, 0xb1, 0xff, 0x91, 0x8b,
	0x2f, 0x75, 0x87, 0x4d, 0x4d, 0x68, 0x4d, 0xb1, 0x73, 0x96, 0xac, 0x92, 0xff, 0x9c, 0xdb, 0xac,
	0xb5, 0xb6, 0xed, 0xe0, 0x9e, 0xc9, 0xa7, 0xfd, 0xf6, 0xbe, 0x01, 0xbf, 0x71, 0xb8, 0x23, 0xeb,
	0xe2, 0xd7, 0xcf, 0xbf, 0x4d, 0xe4, 0xd5, 0x6b, 0x84, 0xae, 0xf9, 0x30, 0xfe, 0xa5, 0x72, 0xb9,
	0x30, 0x96, 0x2a, 0xe8, 0x77, 0xf4, 0xb5, 0xf2, 0xe4, 0xd0, 0xb4, 0x5a, 0x64, 0x22, 0x9f, 0x95,
	0x97, 0xc6, 0xd2, 0x2a, 0xd8, 0xef, 0xd8, 0x55, 0x5a, 0xa6, 0x7d, 0x4d, 0x9b, 0xcf, 0xe0, 0xf5,
	0x83, 0x4c, 0xe4, 0xf3, 0xf2, 0x24, 0xd5, 0xad, 0x9c, 0x6d, 0xac, 0xa1, 0x1a, 0x8d, 0xd7, 0x13,
	0x46, 0xa3, 0x56, 0xd7, 0x32, 0x71, 0xd0, 0xc2, 0x41, 0x3f, 0x64, 0x10, 0x85, 0x7a, 0x22, 0x53,
	0x34, 0x54, 0x75, 0x04, 0x3a, 0xc9, 0x44, 0x3e, 0x29, 0xa7, 0x68, 0x68, 0x4d, 0x70, 0x02, 0x2d,
	0x81, 0x9e, 0x8e, 0xe0, 0x0d, 0x81, 0xba, 0x91, 0xe1, 0x54, 0xc1, 0xa0, 0x53, 0xf6, 0x13, 0x34,
	0xb4, 0x1a, 0xd4, 0x9d, 0x9c, 0x6f, 0x3b, 0x5b, 0xc7, 0xa8, 0x59, 0x26, 0x72, 0x51, 0xce, 0xd8,
	0x08, 0x61, 0x23, 0x0c, 0x71, 0xf3, 0x33, 0x18, 0x02, 0x9f, 0xca, 0x78, 0x0e, 0x91, 0x92, 0x59,
	0xca, 0x7a, 0x35, 0x84, 0x21, 0x7a, 0x34, 0x55, 0x07, 0x46, 0x5f, 0xc4, 0x21, 0x7a, 0x34, 0x6b,
	0x30, 0x0c, 0xea, 0x03, 0x83, 0x47, 0x7f, 0x41, 0x7d, 0x08, 0xe0, 0x46, 0x4e, 0x61, 0x60, 0xff,
	0x71, 0x9c, 0x0e, 0x86, 0x60, 0x5f, 0xcb, 0x04, 0x9c, 0xb3, 0x4e, 0x5f, 0xc6, 0xcb, 0xb3, 0xe0,
	0x3b, 0xfa, 0x6a, 0xbf, 0xc7, 0x46, 0x5f, 0xf1, 0x4b, 0x4f, 0xd1, 0xbf, 0xdf, 0x63, 0x13, 0x46,
	0x42, 0x5f, 0x41, 0x5f, 0x63, 0xa7, 0x17, 0x4c, 0x52, 0xf4, 0xab, 0x20, 0x5f, 0x96, 0x32, 0xd9,
	0x86, 0xe6, 0xd4, 0xb3, 0x22, 0xd6, 0x5c, 0x9c, 0x6a, 0x2e, 0xb8, 0xd1, 0xb7, 0xbb, 0xd0, 0xa6,
	0xd7, 0xbf, 0x7f, 0x86, 0x02, 0x2e, 0x5e, 0xdc, 0x15, 0x67, 0xbb, 0xf2, 0x5f, 0xe7, 0x65, 0x8c,
	0x7a, 0xa5, 0xbf, 0x1f, 0x97, 0xe2, 0xc7, 0x71, 0x29, 0x7e, 0x1d, 0x97, 0xe2, 0xe3, 0xd9, 0x2a,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xa1, 0x48, 0x5f, 0x69, 0x02, 0x00, 0x00,
}

func (m *FieldValidation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldValidation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FieldValidation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsEmail != nil {
		i--
		if *m.IsEmail {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.IsUuid != nil {
		i--
		if *m.IsUuid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if m.Error != nil {
		i -= len(*m.Error)
		copy(dAtA[i:], *m.Error)
		i = encodeVarintValidation(dAtA, i, uint64(len(*m.Error)))
		i--
		dAtA[i] = 0x72
	}
	if m.EqLen != nil {
		i = encodeVarintValidation(dAtA, i, uint64(*m.EqLen))
		i--
		dAtA[i] = 0x68
	}
	if m.MaxLen != nil {
		i = encodeVarintValidation(dAtA, i, uint64(*m.MaxLen))
		i--
		dAtA[i] = 0x60
	}
	if m.MinLen != nil {
		i = encodeVarintValidation(dAtA, i, uint64(*m.MinLen))
		i--
		dAtA[i] = 0x58
	}
	if m.FloatEq != nil {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.FloatEq))))
		i--
		dAtA[i] = 0x51
	}
	if m.FloatGte != nil {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.FloatGte))))
		i--
		dAtA[i] = 0x49
	}
	if m.FloatLte != nil {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.FloatLte))))
		i--
		dAtA[i] = 0x41
	}
	if m.IntEq != nil {
		i = encodeVarintValidation(dAtA, i, uint64(*m.IntEq))
		i--
		dAtA[i] = 0x38
	}
	if m.IntGte != nil {
		i = encodeVarintValidation(dAtA, i, uint64(*m.IntGte))
		i--
		dAtA[i] = 0x30
	}
	if m.IntLte != nil {
		i = encodeVarintValidation(dAtA, i, uint64(*m.IntLte))
		i--
		dAtA[i] = 0x28
	}
	if m.Regex != nil {
		i -= len(*m.Regex)
		copy(dAtA[i:], *m.Regex)
		i = encodeVarintValidation(dAtA, i, uint64(len(*m.Regex)))
		i--
		dAtA[i] = 0x22
	}
	if m.Contains != nil {
		i -= len(*m.Contains)
		copy(dAtA[i:], *m.Contains)
		i = encodeVarintValidation(dAtA, i, uint64(len(*m.Contains)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Matches != nil {
		i -= len(*m.Matches)
		copy(dAtA[i:], *m.Matches)
		i = encodeVarintValidation(dAtA, i, uint64(len(*m.Matches)))
		i--
		dAtA[i] = 0x12
	}
	if m.NotEmptyString != nil {
		i--
		if *m.NotEmptyString {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidation(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FieldValidation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotEmptyString != nil {
		n += 2
	}
	if m.Matches != nil {
		l = len(*m.Matches)
		n += 1 + l + sovValidation(uint64(l))
	}
	if m.Contains != nil {
		l = len(*m.Contains)
		n += 1 + l + sovValidation(uint64(l))
	}
	if m.Regex != nil {
		l = len(*m.Regex)
		n += 1 + l + sovValidation(uint64(l))
	}
	if m.IntLte != nil {
		n += 1 + sovValidation(uint64(*m.IntLte))
	}
	if m.IntGte != nil {
		n += 1 + sovValidation(uint64(*m.IntGte))
	}
	if m.IntEq != nil {
		n += 1 + sovValidation(uint64(*m.IntEq))
	}
	if m.FloatLte != nil {
		n += 9
	}
	if m.FloatGte != nil {
		n += 9
	}
	if m.FloatEq != nil {
		n += 9
	}
	if m.MinLen != nil {
		n += 1 + sovValidation(uint64(*m.MinLen))
	}
	if m.MaxLen != nil {
		n += 1 + sovValidation(uint64(*m.MaxLen))
	}
	if m.EqLen != nil {
		n += 1 + sovValidation(uint64(*m.EqLen))
	}
	if m.Error != nil {
		l = len(*m.Error)
		n += 1 + l + sovValidation(uint64(l))
	}
	if m.IsUuid != nil {
		n += 2
	}
	if m.IsEmail != nil {
		n += 3
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovValidation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidation(x uint64) (n int) {
	return sovValidation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FieldValidation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidation
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
			return fmt.Errorf("proto: FieldValidation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldValidation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotEmptyString", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.NotEmptyString = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matches", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
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
				return ErrInvalidLengthValidation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Matches = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
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
				return ErrInvalidLengthValidation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Contains = &s
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
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
				return ErrInvalidLengthValidation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Regex = &s
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntLte", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IntLte = &v
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntGte", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IntGte = &v
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntEq", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IntEq = &v
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatLte", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.FloatLte = &v2
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatGte", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.FloatGte = &v2
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatEq", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.FloatEq = &v2
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinLen", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MinLen = &v
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLen", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MaxLen = &v
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EqLen", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EqLen = &v
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
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
				return ErrInvalidLengthValidation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Error = &s
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsUuid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.IsUuid = &b
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsEmail", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.IsEmail = &b
		default:
			iNdEx = preIndex
			skippy, err := skipValidation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValidation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValidation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipValidation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidation
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
					return 0, ErrIntOverflowValidation
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowValidation
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
				return 0, ErrInvalidLengthValidation
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthValidation
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowValidation
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipValidation(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthValidation
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthValidation = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidation   = fmt.Errorf("proto: integer overflow")
)
