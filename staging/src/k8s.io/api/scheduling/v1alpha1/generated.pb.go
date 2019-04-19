/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1/generated.proto

package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *PriorityClass) Reset()      { *m = PriorityClass{} }
func (*PriorityClass) ProtoMessage() {}
func (*PriorityClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9f04948bacd88e32, []int{0}
}
func (m *PriorityClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *PriorityClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClass.Merge(dst, src)
}
func (m *PriorityClass) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClass) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClass.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClass proto.InternalMessageInfo

func (m *PriorityClassList) Reset()      { *m = PriorityClassList{} }
func (*PriorityClassList) ProtoMessage() {}
func (*PriorityClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9f04948bacd88e32, []int{1}
}
func (m *PriorityClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *PriorityClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClassList.Merge(dst, src)
}
func (m *PriorityClassList) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClassList.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClassList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PriorityClass)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClass")
	proto.RegisterType((*PriorityClassList)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClassList")
}
func (m *PriorityClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClass) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x10
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Value))
	dAtA[i] = 0x18
	i++
	if m.GlobalDefault {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Description)))
	i += copy(dAtA[i:], m.Description)
	return i, nil
}

func (m *PriorityClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClassList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n2, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PriorityClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Value))
	n += 2
	l = len(m.Description)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PriorityClassList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PriorityClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`GlobalDefault:` + fmt.Sprintf("%v", this.GlobalDefault) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PriorityClassList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PriorityClass", "PriorityClass", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PriorityClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PriorityClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalDefault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.GlobalDefault = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PriorityClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PriorityClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, PriorityClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1/generated.proto", fileDescriptor_generated_9f04948bacd88e32)
}

var fileDescriptor_generated_9f04948bacd88e32 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0x5d, 0x0b, 0x75, 0x4a, 0x41, 0x23, 0x42, 0xe8, 0x61, 0x36, 0xac, 0x97, 0x5c,
	0x76, 0xc6, 0x2e, 0x2a, 0x82, 0xb7, 0xb8, 0xb0, 0x08, 0x8a, 0x92, 0x83, 0x07, 0xf1, 0xe0, 0x24,
	0x79, 0x37, 0x1d, 0x9b, 0x64, 0xc2, 0xcc, 0x24, 0xb0, 0x37, 0xcf, 0x9e, 0xfc, 0x52, 0x42, 0x8f,
	0x7b, 0xdc, 0xd3, 0x62, 0xe3, 0x17, 0x91, 0xa4, 0x69, 0xd3, 0x5a, 0xfc, 0x73, 0xcb, 0x3c, 0xef,
	0xef, 0x79, 0xe6, 0xcd, 0xc3, 0xe0, 0x8b, 0xc5, 0x73, 0x4d, 0x85, 0x64, 0x8b, 0x32, 0x04, 0x95,
	0x83, 0x01, 0xcd, 0x2a, 0xc8, 0x63, 0xa9, 0x58, 0x37, 0xe0, 0x85, 0x60, 0x3a, 0x9a, 0x43, 0x5c,
	0xa6, 0x22, 0x4f, 0x58, 0x35, 0xe3, 0x69, 0x31, 0xe7, 0x33, 0x96, 0x40, 0x0e, 0x8a, 0x1b, 0x88,
	0x69, 0xa1, 0xa4, 0x91, 0x36, 0x59, 0xf3, 0x94, 0x17, 0x82, 0xf6, 0x3c, 0xdd, 0xf0, 0xd3, 0xd3,
	0x44, 0x98, 0x79, 0x19, 0xd2, 0x48, 0x66, 0x2c, 0x91, 0x89, 0x64, 0xad, 0x2d, 0x2c, 0x2f, 0xdb,
	0x53, 0x7b, 0x68, 0xbf, 0xd6, 0x71, 0xd3, 0x27, 0xfd, 0xf5, 0x19, 0x8f, 0xe6, 0x22, 0x07, 0x75,
	0xc5, 0x8a, 0x45, 0xd2, 0x08, 0x9a, 0x65, 0x60, 0x38, 0xab, 0x0e, 0x96, 0x98, 0xb2, 0x3f, 0xb9,
	0x54, 0x99, 0x1b, 0x91, 0xc1, 0x81, 0xe1, 0xd9, 0xbf, 0x0c, 0xcd, 0xaf, 0x64, 0xfc, 0x77, 0xdf,
	0xc9, 0xd7, 0x01, 0x9e, 0xbc, 0x53, 0x42, 0x2a, 0x61, 0xae, 0x5e, 0xa6, 0x5c, 0x6b, 0xfb, 0x13,
	0x1e, 0x35, 0x5b, 0xc5, 0xdc, 0x70, 0x07, 0xb9, 0xc8, 0x1b, 0x9f, 0x3d, 0xa6, 0x7d, 0x25, 0xdb,
	0x70, 0x5a, 0x2c, 0x92, 0x46, 0xd0, 0xb4, 0xa1, 0x69, 0x35, 0xa3, 0x6f, 0xc3, 0xcf, 0x10, 0x99,
	0x37, 0x60, 0xb8, 0x6f, 0x2f, 0x6f, 0x8f, 0xad, 0xfa, 0xf6, 0x18, 0xf7, 0x5a, 0xb0, 0x4d, 0xb5,
	0x1f, 0xe1, 0x61, 0xc5, 0xd3, 0x12, 0x9c, 0x81, 0x8b, 0xbc, 0xa1, 0x3f, 0xe9, 0xe0, 0xe1, 0xfb,
	0x46, 0x0c, 0xd6, 0x33, 0xfb, 0x05, 0x9e, 0x24, 0xa9, 0x0c, 0x79, 0x7a, 0x0e, 0x97, 0xbc, 0x4c,
	0x8d, 0x73, 0xe4, 0x22, 0x6f, 0xe4, 0x3f, 0xec, 0xe0, 0xc9, 0xc5, 0xee, 0x30, 0xd8, 0x67, 0xed,
	0xa7, 0x78, 0x1c, 0x83, 0x8e, 0x94, 0x28, 0x8c, 0x90, 0xb9, 0x73, 0xc7, 0x45, 0xde, 0x5d, 0xff,
	0x41, 0x67, 0x1d, 0x9f, 0xf7, 0xa3, 0x60, 0x97, 0x3b, 0xf9, 0x8e, 0xf0, 0xfd, 0xbd, 0x32, 0x5e,
	0x0b, 0x6d, 0xec, 0x8f, 0x07, 0x85, 0xd0, 0xff, 0x2b, 0xa4, 0x71, 0xb7, 0x75, 0xdc, 0xeb, 0x6e,
	0x1e, 0x6d, 0x94, 0x9d, 0x32, 0x02, 0x3c, 0x14, 0x06, 0x32, 0xed, 0x0c, 0xdc, 0x23, 0x6f, 0x7c,
	0x76, 0x4a, 0xff, 0xfe, 0xfc, 0xe8, 0xde, 0x7e, 0x7d, 0x77, 0xaf, 0x9a, 0x8c, 0x60, 0x1d, 0xe5,
	0xd3, 0xe5, 0x8a, 0x58, 0xd7, 0x2b, 0x62, 0xdd, 0xac, 0x88, 0xf5, 0xa5, 0x26, 0x68, 0x59, 0x13,
	0x74, 0x5d, 0x13, 0x74, 0x53, 0x13, 0xf4, 0xa3, 0x26, 0xe8, 0xdb, 0x4f, 0x62, 0x7d, 0x18, 0x6d,
	0x32, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xab, 0x20, 0x12, 0x63, 0x3c, 0x03, 0x00, 0x00,
}
