// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: proto3opt/opt.proto

package proto3opt

import (
	binary "encoding/binary"
	fmt "fmt"
	protohelpers "github.com/DataDog/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	math "math"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *OptionalFieldInProto3) CloneVT() *OptionalFieldInProto3 {
	if m == nil {
		return (*OptionalFieldInProto3)(nil)
	}
	r := new(OptionalFieldInProto3)
	if rhs := m.OptionalInt32; rhs != nil {
		tmpVal := *rhs
		r.OptionalInt32 = &tmpVal
	}
	if rhs := m.OptionalInt64; rhs != nil {
		tmpVal := *rhs
		r.OptionalInt64 = &tmpVal
	}
	if rhs := m.OptionalUint32; rhs != nil {
		tmpVal := *rhs
		r.OptionalUint32 = &tmpVal
	}
	if rhs := m.OptionalUint64; rhs != nil {
		tmpVal := *rhs
		r.OptionalUint64 = &tmpVal
	}
	if rhs := m.OptionalSint32; rhs != nil {
		tmpVal := *rhs
		r.OptionalSint32 = &tmpVal
	}
	if rhs := m.OptionalSint64; rhs != nil {
		tmpVal := *rhs
		r.OptionalSint64 = &tmpVal
	}
	if rhs := m.OptionalFixed32; rhs != nil {
		tmpVal := *rhs
		r.OptionalFixed32 = &tmpVal
	}
	if rhs := m.OptionalFixed64; rhs != nil {
		tmpVal := *rhs
		r.OptionalFixed64 = &tmpVal
	}
	if rhs := m.OptionalSfixed32; rhs != nil {
		tmpVal := *rhs
		r.OptionalSfixed32 = &tmpVal
	}
	if rhs := m.OptionalSfixed64; rhs != nil {
		tmpVal := *rhs
		r.OptionalSfixed64 = &tmpVal
	}
	if rhs := m.OptionalFloat; rhs != nil {
		tmpVal := *rhs
		r.OptionalFloat = &tmpVal
	}
	if rhs := m.OptionalDouble; rhs != nil {
		tmpVal := *rhs
		r.OptionalDouble = &tmpVal
	}
	if rhs := m.OptionalBool; rhs != nil {
		tmpVal := *rhs
		r.OptionalBool = &tmpVal
	}
	if rhs := m.OptionalString; rhs != nil {
		tmpVal := *rhs
		r.OptionalString = &tmpVal
	}
	if rhs := m.OptionalBytes; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.OptionalBytes = tmpBytes
	}
	if rhs := m.OptionalEnum; rhs != nil {
		tmpVal := *rhs
		r.OptionalEnum = &tmpVal
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OptionalFieldInProto3) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *OptionalFieldInProto3) EqualVT(that *OptionalFieldInProto3) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if p, q := this.OptionalInt32, that.OptionalInt32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalInt64, that.OptionalInt64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalUint32, that.OptionalUint32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalUint64, that.OptionalUint64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSint32, that.OptionalSint32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSint64, that.OptionalSint64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalFixed32, that.OptionalFixed32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalFixed64, that.OptionalFixed64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSfixed32, that.OptionalSfixed32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSfixed64, that.OptionalSfixed64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalFloat, that.OptionalFloat; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalDouble, that.OptionalDouble; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalBool, that.OptionalBool; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalString, that.OptionalString; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalBytes, that.OptionalBytes; (p == nil && q != nil) || (p != nil && q == nil) || string(p) != string(q) {
		return false
	}
	if p, q := this.OptionalEnum, that.OptionalEnum; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OptionalFieldInProto3) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*OptionalFieldInProto3)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *OptionalFieldInProto3) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OptionalFieldInProto3) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OptionalFieldInProto3) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.OptionalEnum != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalEnum))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.OptionalBytes != nil {
		i -= len(m.OptionalBytes)
		copy(dAtA[i:], m.OptionalBytes)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OptionalBytes)))
		i--
		dAtA[i] = 0x7a
	}
	if m.OptionalString != nil {
		i -= len(*m.OptionalString)
		copy(dAtA[i:], *m.OptionalString)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(*m.OptionalString)))
		i--
		dAtA[i] = 0x72
	}
	if m.OptionalBool != nil {
		i--
		if *m.OptionalBool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.OptionalDouble != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.OptionalDouble))))
		i--
		dAtA[i] = 0x61
	}
	if m.OptionalFloat != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(*m.OptionalFloat))))
		i--
		dAtA[i] = 0x5d
	}
	if m.OptionalSfixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalSfixed64))
		i--
		dAtA[i] = 0x51
	}
	if m.OptionalSfixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalSfixed32))
		i--
		dAtA[i] = 0x4d
	}
	if m.OptionalFixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalFixed64))
		i--
		dAtA[i] = 0x41
	}
	if m.OptionalFixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalFixed32))
		i--
		dAtA[i] = 0x3d
	}
	if m.OptionalSint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint64(*m.OptionalSint64)<<1)^uint64((*m.OptionalSint64>>63))))
		i--
		dAtA[i] = 0x30
	}
	if m.OptionalSint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint32(*m.OptionalSint32)<<1)^uint32((*m.OptionalSint32>>31))))
		i--
		dAtA[i] = 0x28
	}
	if m.OptionalUint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint64))
		i--
		dAtA[i] = 0x20
	}
	if m.OptionalUint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint32))
		i--
		dAtA[i] = 0x18
	}
	if m.OptionalInt64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt64))
		i--
		dAtA[i] = 0x10
	}
	if m.OptionalInt32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt32))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OptionalFieldInProto3) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OptionalFieldInProto3) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OptionalFieldInProto3) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.OptionalEnum != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalEnum))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.OptionalBytes != nil {
		i -= len(m.OptionalBytes)
		copy(dAtA[i:], m.OptionalBytes)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OptionalBytes)))
		i--
		dAtA[i] = 0x7a
	}
	if m.OptionalString != nil {
		i -= len(*m.OptionalString)
		copy(dAtA[i:], *m.OptionalString)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(*m.OptionalString)))
		i--
		dAtA[i] = 0x72
	}
	if m.OptionalBool != nil {
		i--
		if *m.OptionalBool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.OptionalDouble != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.OptionalDouble))))
		i--
		dAtA[i] = 0x61
	}
	if m.OptionalFloat != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(*m.OptionalFloat))))
		i--
		dAtA[i] = 0x5d
	}
	if m.OptionalSfixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalSfixed64))
		i--
		dAtA[i] = 0x51
	}
	if m.OptionalSfixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalSfixed32))
		i--
		dAtA[i] = 0x4d
	}
	if m.OptionalFixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalFixed64))
		i--
		dAtA[i] = 0x41
	}
	if m.OptionalFixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalFixed32))
		i--
		dAtA[i] = 0x3d
	}
	if m.OptionalSint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint64(*m.OptionalSint64)<<1)^uint64((*m.OptionalSint64>>63))))
		i--
		dAtA[i] = 0x30
	}
	if m.OptionalSint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint32(*m.OptionalSint32)<<1)^uint32((*m.OptionalSint32>>31))))
		i--
		dAtA[i] = 0x28
	}
	if m.OptionalUint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint64))
		i--
		dAtA[i] = 0x20
	}
	if m.OptionalUint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint32))
		i--
		dAtA[i] = 0x18
	}
	if m.OptionalInt64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt64))
		i--
		dAtA[i] = 0x10
	}
	if m.OptionalInt32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt32))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OptionalFieldInProto3) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OptionalInt32 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalInt32))
	}
	if m.OptionalInt64 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalInt64))
	}
	if m.OptionalUint32 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalUint32))
	}
	if m.OptionalUint64 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalUint64))
	}
	if m.OptionalSint32 != nil {
		n += 1 + protohelpers.SizeOfZigzag(uint64(*m.OptionalSint32))
	}
	if m.OptionalSint64 != nil {
		n += 1 + protohelpers.SizeOfZigzag(uint64(*m.OptionalSint64))
	}
	if m.OptionalFixed32 != nil {
		n += 5
	}
	if m.OptionalFixed64 != nil {
		n += 9
	}
	if m.OptionalSfixed32 != nil {
		n += 5
	}
	if m.OptionalSfixed64 != nil {
		n += 9
	}
	if m.OptionalFloat != nil {
		n += 5
	}
	if m.OptionalDouble != nil {
		n += 9
	}
	if m.OptionalBool != nil {
		n += 2
	}
	if m.OptionalString != nil {
		l = len(*m.OptionalString)
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.OptionalBytes != nil {
		l = len(m.OptionalBytes)
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.OptionalEnum != nil {
		n += 2 + protohelpers.SizeOfVarint(uint64(*m.OptionalEnum))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OptionalFieldInProto3) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: OptionalFieldInProto3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OptionalFieldInProto3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalInt32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalInt32 = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalInt64", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
			m.OptionalInt64 = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalUint32", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint32 = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalUint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint64 = &v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSint32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.OptionalSint32 = &v
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			v2 := int64(v)
			m.OptionalSint64 = &v2
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalFixed32", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalFixed32 = &v
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalFixed64", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalFixed64 = &v
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSfixed32", wireType)
			}
			var v int32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalSfixed32 = &v
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSfixed64", wireType)
			}
			var v int64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalSfixed64 = &v
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalFloat", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			v2 := float32(math.Float32frombits(v))
			m.OptionalFloat = &v2
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalDouble", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.OptionalDouble = &v2
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalBool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
			m.OptionalBool = &b
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.OptionalString = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OptionalBytes = append(m.OptionalBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.OptionalBytes == nil {
				m.OptionalBytes = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalEnum", wireType)
			}
			var v SimpleEnum
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= SimpleEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalEnum = &v
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OptionalFieldInProto3) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: OptionalFieldInProto3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OptionalFieldInProto3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalInt32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalInt32 = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalInt64", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
			m.OptionalInt64 = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalUint32", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint32 = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalUint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint64 = &v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSint32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.OptionalSint32 = &v
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			v2 := int64(v)
			m.OptionalSint64 = &v2
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalFixed32", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalFixed32 = &v
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalFixed64", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalFixed64 = &v
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSfixed32", wireType)
			}
			var v int32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalSfixed32 = &v
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalSfixed64", wireType)
			}
			var v int64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalSfixed64 = &v
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalFloat", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			v2 := float32(math.Float32frombits(v))
			m.OptionalFloat = &v2
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalDouble", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.OptionalDouble = &v2
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalBool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
			m.OptionalBool = &b
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := unsafe.String(&dAtA[iNdEx], intStringLen)
			m.OptionalString = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OptionalBytes = dAtA[iNdEx:postIndex]
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalEnum", wireType)
			}
			var v SimpleEnum
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= SimpleEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalEnum = &v
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
