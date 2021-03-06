// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: match.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CMD_MATCH_ENUM int32

const (
	CMD_MATCH_INVALID CMD_MATCH_ENUM = 0
	CMD_MATCH_MATCH   CMD_MATCH_ENUM = 1
)

var CMD_MATCH_ENUM_name = map[int32]string{
	0: "INVALID",
	1: "MATCH",
}
var CMD_MATCH_ENUM_value = map[string]int32{
	"INVALID": 0,
	"MATCH":   1,
}

func (x CMD_MATCH_ENUM) String() string {
	return proto.EnumName(CMD_MATCH_ENUM_name, int32(x))
}
func (CMD_MATCH_ENUM) EnumDescriptor() ([]byte, []int) { return fileDescriptorMatch, []int{0, 0} }

type ENUM_MATCH_COMMON_ERROR_ENUM int32

const (
	ENUM_MATCH_COMMON_ERROR_OK           ENUM_MATCH_COMMON_ERROR_ENUM = 0
	ENUM_MATCH_COMMON_ERROR_SYSTEM_ERROR ENUM_MATCH_COMMON_ERROR_ENUM = 1
)

var ENUM_MATCH_COMMON_ERROR_ENUM_name = map[int32]string{
	0: "OK",
	1: "SYSTEM_ERROR",
}
var ENUM_MATCH_COMMON_ERROR_ENUM_value = map[string]int32{
	"OK":           0,
	"SYSTEM_ERROR": 1,
}

func (x ENUM_MATCH_COMMON_ERROR_ENUM) String() string {
	return proto.EnumName(ENUM_MATCH_COMMON_ERROR_ENUM_name, int32(x))
}
func (ENUM_MATCH_COMMON_ERROR_ENUM) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorMatch, []int{1, 0}
}

type CMD_MATCH struct {
}

func (m *CMD_MATCH) Reset()                    { *m = CMD_MATCH{} }
func (m *CMD_MATCH) String() string            { return proto.CompactTextString(m) }
func (*CMD_MATCH) ProtoMessage()               {}
func (*CMD_MATCH) Descriptor() ([]byte, []int) { return fileDescriptorMatch, []int{0} }

type ENUM_MATCH_COMMON_ERROR struct {
}

func (m *ENUM_MATCH_COMMON_ERROR) Reset()                    { *m = ENUM_MATCH_COMMON_ERROR{} }
func (m *ENUM_MATCH_COMMON_ERROR) String() string            { return proto.CompactTextString(m) }
func (*ENUM_MATCH_COMMON_ERROR) ProtoMessage()               {}
func (*ENUM_MATCH_COMMON_ERROR) Descriptor() ([]byte, []int) { return fileDescriptorMatch, []int{1} }

// 请求匹配 ( LOBBY -> MATCH )
type MSG_MATCH_MATCH struct {
	Account string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	RoleID  uint64 `protobuf:"varint,2,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
}

func (m *MSG_MATCH_MATCH) Reset()                    { *m = MSG_MATCH_MATCH{} }
func (m *MSG_MATCH_MATCH) String() string            { return proto.CompactTextString(m) }
func (*MSG_MATCH_MATCH) ProtoMessage()               {}
func (*MSG_MATCH_MATCH) Descriptor() ([]byte, []int) { return fileDescriptorMatch, []int{2} }

func (m *MSG_MATCH_MATCH) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MSG_MATCH_MATCH) GetRoleID() uint64 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

type MSG_MATCH_MATCH_RESULT struct {
	Err     ENUM_MATCH_COMMON_ERROR_ENUM `protobuf:"varint,1,opt,name=Err,proto3,enum=protocol.ENUM_MATCH_COMMON_ERROR_ENUM" json:"Err,omitempty"`
	Account string                       `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	RoleID  uint64                       `protobuf:"varint,3,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	Roles   []*ROLE_BASE_INFO            `protobuf:"bytes,4,rep,name=Roles" json:"Roles,omitempty"`
}

func (m *MSG_MATCH_MATCH_RESULT) Reset()                    { *m = MSG_MATCH_MATCH_RESULT{} }
func (m *MSG_MATCH_MATCH_RESULT) String() string            { return proto.CompactTextString(m) }
func (*MSG_MATCH_MATCH_RESULT) ProtoMessage()               {}
func (*MSG_MATCH_MATCH_RESULT) Descriptor() ([]byte, []int) { return fileDescriptorMatch, []int{3} }

func (m *MSG_MATCH_MATCH_RESULT) GetErr() ENUM_MATCH_COMMON_ERROR_ENUM {
	if m != nil {
		return m.Err
	}
	return ENUM_MATCH_COMMON_ERROR_OK
}

func (m *MSG_MATCH_MATCH_RESULT) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MSG_MATCH_MATCH_RESULT) GetRoleID() uint64 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

func (m *MSG_MATCH_MATCH_RESULT) GetRoles() []*ROLE_BASE_INFO {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*CMD_MATCH)(nil), "protocol.CMD_MATCH")
	proto.RegisterType((*ENUM_MATCH_COMMON_ERROR)(nil), "protocol.ENUM_MATCH_COMMON_ERROR")
	proto.RegisterType((*MSG_MATCH_MATCH)(nil), "protocol.MSG_MATCH_MATCH")
	proto.RegisterType((*MSG_MATCH_MATCH_RESULT)(nil), "protocol.MSG_MATCH_MATCH_RESULT")
	proto.RegisterEnum("protocol.CMD_MATCH_ENUM", CMD_MATCH_ENUM_name, CMD_MATCH_ENUM_value)
	proto.RegisterEnum("protocol.ENUM_MATCH_COMMON_ERROR_ENUM", ENUM_MATCH_COMMON_ERROR_ENUM_name, ENUM_MATCH_COMMON_ERROR_ENUM_value)
}
func (m *CMD_MATCH) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CMD_MATCH) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ENUM_MATCH_COMMON_ERROR) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ENUM_MATCH_COMMON_ERROR) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MSG_MATCH_MATCH) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_MATCH_MATCH) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMatch(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if m.RoleID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMatch(dAtA, i, uint64(m.RoleID))
	}
	return i, nil
}

func (m *MSG_MATCH_MATCH_RESULT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_MATCH_MATCH_RESULT) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Err != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMatch(dAtA, i, uint64(m.Err))
	}
	if len(m.Account) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMatch(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if m.RoleID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMatch(dAtA, i, uint64(m.RoleID))
	}
	if len(m.Roles) > 0 {
		for _, msg := range m.Roles {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMatch(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Match(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Match(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMatch(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CMD_MATCH) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ENUM_MATCH_COMMON_ERROR) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MSG_MATCH_MATCH) Size() (n int) {
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	if m.RoleID != 0 {
		n += 1 + sovMatch(uint64(m.RoleID))
	}
	return n
}

func (m *MSG_MATCH_MATCH_RESULT) Size() (n int) {
	var l int
	_ = l
	if m.Err != 0 {
		n += 1 + sovMatch(uint64(m.Err))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	if m.RoleID != 0 {
		n += 1 + sovMatch(uint64(m.RoleID))
	}
	if len(m.Roles) > 0 {
		for _, e := range m.Roles {
			l = e.Size()
			n += 1 + l + sovMatch(uint64(l))
		}
	}
	return n
}

func sovMatch(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMatch(x uint64) (n int) {
	return sovMatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CMD_MATCH) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CMD_MATCH: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CMD_MATCH: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMatch
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
func (m *ENUM_MATCH_COMMON_ERROR) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ENUM_MATCH_COMMON_ERROR: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ENUM_MATCH_COMMON_ERROR: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMatch
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
func (m *MSG_MATCH_MATCH) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MSG_MATCH_MATCH: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_MATCH_MATCH: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleID", wireType)
			}
			m.RoleID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoleID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMatch
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
func (m *MSG_MATCH_MATCH_RESULT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MSG_MATCH_MATCH_RESULT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_MATCH_MATCH_RESULT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			m.Err = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Err |= (ENUM_MATCH_COMMON_ERROR_ENUM(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleID", wireType)
			}
			m.RoleID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoleID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, &ROLE_BASE_INFO{})
			if err := m.Roles[len(m.Roles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMatch
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
func skipMatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMatch
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMatch
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
				next, err := skipMatch(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthMatch = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMatch   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("match.proto", fileDescriptorMatch) }

var fileDescriptorMatch = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4d, 0x2c, 0x49,
	0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x52, 0x42,
	0x39, 0xf9, 0x49, 0x49, 0x95, 0xf1, 0xc9, 0xa5, 0xc5, 0x25, 0xf9, 0xb9, 0x10, 0x59, 0x25, 0x6d,
	0x2e, 0x4e, 0x67, 0x5f, 0x97, 0x78, 0x5f, 0xc7, 0x10, 0x67, 0x0f, 0x25, 0x39, 0x2e, 0x16, 0x57,
	0xbf, 0x50, 0x5f, 0x21, 0x6e, 0x2e, 0x76, 0x4f, 0xbf, 0x30, 0x47, 0x1f, 0x4f, 0x17, 0x01, 0x06,
	0x21, 0x4e, 0x2e, 0x56, 0xb0, 0xac, 0x00, 0xa3, 0x92, 0x35, 0x97, 0x38, 0x48, 0x1e, 0xa2, 0x3a,
	0xde, 0xd9, 0xdf, 0xd7, 0xd7, 0xdf, 0x2f, 0xde, 0x35, 0x28, 0xc8, 0x3f, 0x48, 0x49, 0x01, 0xaa,
	0x95, 0x8d, 0x8b, 0xc9, 0xdf, 0x5b, 0x80, 0x41, 0x48, 0x80, 0x8b, 0x27, 0x38, 0x32, 0x38, 0xc4,
	0xd5, 0x17, 0x22, 0x2f, 0xc0, 0xa8, 0xe4, 0xcc, 0xc5, 0xef, 0x1b, 0xec, 0x0e, 0xd5, 0x0b, 0x26,
	0x85, 0x24, 0xb8, 0xd8, 0x1d, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x60, 0x5c, 0x21, 0x31, 0x2e, 0xb6, 0xa0, 0xfc, 0x9c, 0x54, 0x4f, 0x17, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0x96, 0x20, 0x28, 0x4f, 0x69, 0x0b, 0x23, 0x97, 0x18, 0x9a, 0x29, 0xf1, 0x41,
	0xae, 0xc1, 0xa1, 0x3e, 0x21, 0x42, 0x16, 0x5c, 0xcc, 0xae, 0x45, 0x45, 0x60, 0x83, 0xf8, 0x8c,
	0xd4, 0xf4, 0x60, 0xbe, 0xd6, 0xc3, 0xe1, 0x62, 0xb0, 0x78, 0x10, 0x48, 0x0b, 0xb2, 0x33, 0x98,
	0x70, 0x39, 0x83, 0x19, 0xd9, 0x19, 0x42, 0x7a, 0x5c, 0xac, 0x20, 0x56, 0xb1, 0x04, 0x8b, 0x02,
	0xb3, 0x06, 0xb7, 0x91, 0x04, 0xc2, 0xb6, 0x20, 0x7f, 0x1f, 0xd7, 0x78, 0x27, 0xc7, 0x60, 0xd7,
	0x78, 0x4f, 0x3f, 0x37, 0xff, 0x20, 0x88, 0x32, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0x0e, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xbf, 0x52, 0xc8, 0xab, 0x01, 0x00, 0x00,
}
