// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v2/Anchor.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

type AnchorQueryLiveWearingReq struct {
	// 房管 uid
	UidList []int64 `protobuf:"varint,1,rep,packed,name=uid_list,json=uidList" json:"uid_list"`
}

func (m *AnchorQueryLiveWearingReq) Reset()         { *m = AnchorQueryLiveWearingReq{} }
func (m *AnchorQueryLiveWearingReq) String() string { return proto.CompactTextString(m) }
func (*AnchorQueryLiveWearingReq) ProtoMessage()    {}
func (*AnchorQueryLiveWearingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_Anchor_9efdd09a9d362c87, []int{0}
}
func (m *AnchorQueryLiveWearingReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnchorQueryLiveWearingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnchorQueryLiveWearingReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AnchorQueryLiveWearingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorQueryLiveWearingReq.Merge(dst, src)
}
func (m *AnchorQueryLiveWearingReq) XXX_Size() int {
	return m.Size()
}
func (m *AnchorQueryLiveWearingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorQueryLiveWearingReq.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorQueryLiveWearingReq proto.InternalMessageInfo

func (m *AnchorQueryLiveWearingReq) GetUidList() []int64 {
	if m != nil {
		return m.UidList
	}
	return nil
}

type AnchorQueryLiveWearingResp struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	// 勋章信息 map
	Data map[int64]*AnchorQueryLiveWearingResp_Medal `protobuf:"bytes,3,rep,name=data" json:"data" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AnchorQueryLiveWearingResp) Reset()         { *m = AnchorQueryLiveWearingResp{} }
func (m *AnchorQueryLiveWearingResp) String() string { return proto.CompactTextString(m) }
func (*AnchorQueryLiveWearingResp) ProtoMessage()    {}
func (*AnchorQueryLiveWearingResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_Anchor_9efdd09a9d362c87, []int{1}
}
func (m *AnchorQueryLiveWearingResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnchorQueryLiveWearingResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnchorQueryLiveWearingResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AnchorQueryLiveWearingResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorQueryLiveWearingResp.Merge(dst, src)
}
func (m *AnchorQueryLiveWearingResp) XXX_Size() int {
	return m.Size()
}
func (m *AnchorQueryLiveWearingResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorQueryLiveWearingResp.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorQueryLiveWearingResp proto.InternalMessageInfo

func (m *AnchorQueryLiveWearingResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AnchorQueryLiveWearingResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AnchorQueryLiveWearingResp) GetData() map[int64]*AnchorQueryLiveWearingResp_Medal {
	if m != nil {
		return m.Data
	}
	return nil
}

type AnchorQueryLiveWearingResp_Medal struct {
	//
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	//
	TargetId int64 `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id"`
	//
	MedalId int64 `protobuf:"varint,3,opt,name=medal_id,json=medalId,proto3" json:"medal_id"`
	//
	MedalName string `protobuf:"bytes,4,opt,name=medal_name,json=medalName,proto3" json:"medal_name"`
	//
	Level int64 `protobuf:"varint,5,opt,name=level,proto3" json:"level"`
}

func (m *AnchorQueryLiveWearingResp_Medal) Reset()         { *m = AnchorQueryLiveWearingResp_Medal{} }
func (m *AnchorQueryLiveWearingResp_Medal) String() string { return proto.CompactTextString(m) }
func (*AnchorQueryLiveWearingResp_Medal) ProtoMessage()    {}
func (*AnchorQueryLiveWearingResp_Medal) Descriptor() ([]byte, []int) {
	return fileDescriptor_Anchor_9efdd09a9d362c87, []int{1, 1}
}
func (m *AnchorQueryLiveWearingResp_Medal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnchorQueryLiveWearingResp_Medal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnchorQueryLiveWearingResp_Medal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AnchorQueryLiveWearingResp_Medal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorQueryLiveWearingResp_Medal.Merge(dst, src)
}
func (m *AnchorQueryLiveWearingResp_Medal) XXX_Size() int {
	return m.Size()
}
func (m *AnchorQueryLiveWearingResp_Medal) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorQueryLiveWearingResp_Medal.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorQueryLiveWearingResp_Medal proto.InternalMessageInfo

func (m *AnchorQueryLiveWearingResp_Medal) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AnchorQueryLiveWearingResp_Medal) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *AnchorQueryLiveWearingResp_Medal) GetMedalId() int64 {
	if m != nil {
		return m.MedalId
	}
	return 0
}

func (m *AnchorQueryLiveWearingResp_Medal) GetMedalName() string {
	if m != nil {
		return m.MedalName
	}
	return ""
}

func (m *AnchorQueryLiveWearingResp_Medal) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*AnchorQueryLiveWearingReq)(nil), "fans_medal.v2.AnchorQueryLiveWearingReq")
	proto.RegisterType((*AnchorQueryLiveWearingResp)(nil), "fans_medal.v2.AnchorQueryLiveWearingResp")
	proto.RegisterMapType((map[int64]*AnchorQueryLiveWearingResp_Medal)(nil), "fans_medal.v2.AnchorQueryLiveWearingResp.DataEntry")
	proto.RegisterType((*AnchorQueryLiveWearingResp_Medal)(nil), "fans_medal.v2.AnchorQueryLiveWearingResp.Medal")
}
func (m *AnchorQueryLiveWearingReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnchorQueryLiveWearingReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UidList) > 0 {
		dAtA2 := make([]byte, len(m.UidList)*10)
		var j1 int
		for _, num1 := range m.UidList {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func (m *AnchorQueryLiveWearingResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnchorQueryLiveWearingResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Data) > 0 {
		for k, _ := range m.Data {
			dAtA[i] = 0x1a
			i++
			v := m.Data[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovAnchor(uint64(msgSize))
			}
			mapSize := 1 + sovAnchor(uint64(k)) + msgSize
			i = encodeVarintAnchor(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintAnchor(dAtA, i, uint64(k))
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintAnchor(dAtA, i, uint64(v.Size()))
				n3, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n3
			}
		}
	}
	return i, nil
}

func (m *AnchorQueryLiveWearingResp_Medal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnchorQueryLiveWearingResp_Medal) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(m.Uid))
	}
	if m.TargetId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(m.TargetId))
	}
	if m.MedalId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(m.MedalId))
	}
	if len(m.MedalName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(len(m.MedalName)))
		i += copy(dAtA[i:], m.MedalName)
	}
	if m.Level != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAnchor(dAtA, i, uint64(m.Level))
	}
	return i, nil
}

func encodeVarintAnchor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AnchorQueryLiveWearingReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.UidList) > 0 {
		l = 0
		for _, e := range m.UidList {
			l += sovAnchor(uint64(e))
		}
		n += 1 + sovAnchor(uint64(l)) + l
	}
	return n
}

func (m *AnchorQueryLiveWearingResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovAnchor(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovAnchor(uint64(l))
	}
	if len(m.Data) > 0 {
		for k, v := range m.Data {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovAnchor(uint64(l))
			}
			mapEntrySize := 1 + sovAnchor(uint64(k)) + l
			n += mapEntrySize + 1 + sovAnchor(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *AnchorQueryLiveWearingResp_Medal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovAnchor(uint64(m.Uid))
	}
	if m.TargetId != 0 {
		n += 1 + sovAnchor(uint64(m.TargetId))
	}
	if m.MedalId != 0 {
		n += 1 + sovAnchor(uint64(m.MedalId))
	}
	l = len(m.MedalName)
	if l > 0 {
		n += 1 + l + sovAnchor(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovAnchor(uint64(m.Level))
	}
	return n
}

func sovAnchor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAnchor(x uint64) (n int) {
	return sovAnchor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnchorQueryLiveWearingReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnchor
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
			return fmt.Errorf("proto: AnchorQueryLiveWearingReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnchorQueryLiveWearingReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAnchor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.UidList = append(m.UidList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAnchor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthAnchor
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.UidList) == 0 {
					m.UidList = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAnchor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.UidList = append(m.UidList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field UidList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAnchor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAnchor
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
func (m *AnchorQueryLiveWearingResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnchor
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
			return fmt.Errorf("proto: AnchorQueryLiveWearingResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnchorQueryLiveWearingResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
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
				return ErrInvalidLengthAnchor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
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
				return ErrInvalidLengthAnchor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = make(map[int64]*AnchorQueryLiveWearingResp_Medal)
			}
			var mapkey int64
			var mapvalue *AnchorQueryLiveWearingResp_Medal
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAnchor
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAnchor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAnchor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthAnchor
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthAnchor
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &AnchorQueryLiveWearingResp_Medal{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipAnchor(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthAnchor
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Data[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnchor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAnchor
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
func (m *AnchorQueryLiveWearingResp_Medal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnchor
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
			return fmt.Errorf("proto: Medal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Medal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetId", wireType)
			}
			m.TargetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MedalId", wireType)
			}
			m.MedalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MedalId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MedalName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
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
				return ErrInvalidLengthAnchor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MedalName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnchor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAnchor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAnchor
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
func skipAnchor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnchor
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
					return 0, ErrIntOverflowAnchor
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
					return 0, ErrIntOverflowAnchor
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
				return 0, ErrInvalidLengthAnchor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAnchor
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
				next, err := skipAnchor(dAtA[start:])
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
	ErrInvalidLengthAnchor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnchor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v2/Anchor.proto", fileDescriptor_Anchor_9efdd09a9d362c87) }

var fileDescriptor_Anchor_9efdd09a9d362c87 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xb3, 0x49, 0x13, 0x6f, 0x29, 0xa0, 0x3d, 0xb9, 0x56, 0x65, 0x47, 0xbd, 0x60, 0x90,
	0xea, 0x48, 0xee, 0x05, 0x71, 0xc3, 0x6a, 0x0f, 0x95, 0x0a, 0x52, 0xf7, 0x82, 0xc4, 0xc5, 0xda,
	0x64, 0xb7, 0xce, 0x0a, 0x7f, 0x34, 0xf6, 0xda, 0x28, 0xff, 0x82, 0x7f, 0x05, 0xc7, 0x1e, 0x39,
	0x20, 0x0b, 0x25, 0x37, 0xff, 0x0a, 0xb4, 0xb3, 0x50, 0xe0, 0x50, 0x29, 0x97, 0xf1, 0x9b, 0xe7,
	0xd9, 0x79, 0xf3, 0x66, 0x17, 0x3f, 0x6b, 0xa3, 0xf9, 0xdb, 0x62, 0xb9, 0x2a, 0xab, 0xf0, 0xae,
	0x2a, 0x55, 0x49, 0x8e, 0x6e, 0x59, 0x51, 0x27, 0xb9, 0xe0, 0x2c, 0x0b, 0xdb, 0xc8, 0x3d, 0x4b,
	0xa5, 0x5a, 0x35, 0x8b, 0x70, 0x59, 0xe6, 0xf3, 0xb4, 0x4c, 0xcb, 0x39, 0x54, 0x2d, 0x9a, 0x5b,
	0xc8, 0x20, 0x01, 0x64, 0x4e, 0x9f, 0x5e, 0xe0, 0x63, 0xd3, 0xed, 0xa6, 0x11, 0xd5, 0xe6, 0x5a,
	0xb6, 0xe2, 0x83, 0x60, 0x95, 0x2c, 0x52, 0x2a, 0xd6, 0xe4, 0x05, 0x9e, 0x36, 0x92, 0x27, 0x99,
	0xac, 0x95, 0x63, 0xcd, 0x50, 0x80, 0xe2, 0x27, 0x7d, 0xe7, 0x3f, 0x70, 0x74, 0xd2, 0x48, 0x7e,
	0x2d, 0x6b, 0x75, 0xfa, 0x03, 0x61, 0xf7, 0xb1, 0x36, 0xf5, 0x1d, 0x39, 0xc1, 0xa3, 0x65, 0xc9,
	0x85, 0x63, 0xcd, 0xac, 0x00, 0xc5, 0xd3, 0xbe, 0xf3, 0x21, 0xa7, 0x10, 0xc9, 0x31, 0x46, 0x79,
	0x9d, 0x3a, 0xc3, 0x99, 0x15, 0xd8, 0xf1, 0xa4, 0xef, 0x7c, 0x9d, 0x52, 0x1d, 0xc8, 0x0d, 0x1e,
	0x71, 0xa6, 0x98, 0x83, 0x66, 0x28, 0x38, 0x8c, 0xce, 0xc3, 0xff, 0xac, 0x86, 0x8f, 0x2b, 0x86,
	0x17, 0x4c, 0xb1, 0xcb, 0x42, 0x55, 0x1b, 0xa3, 0xa6, 0x9b, 0x50, 0x88, 0xee, 0x0a, 0xdb, 0x0f,
	0x3f, 0xc9, 0x73, 0x8c, 0x3e, 0x89, 0x8d, 0x99, 0x8b, 0x6a, 0x48, 0x2e, 0xf1, 0xb8, 0x65, 0x59,
	0x23, 0x60, 0x9c, 0xc3, 0x68, 0xbe, 0xbf, 0xe4, 0x3b, 0x5d, 0x43, 0xcd, 0xe9, 0x37, 0xc3, 0xd7,
	0x96, 0xfb, 0xd5, 0xc2, 0x63, 0x20, 0xb5, 0xc3, 0x46, 0xf2, 0xdf, 0xf6, 0xc1, 0x61, 0x23, 0x39,
	0xd5, 0x81, 0xbc, 0xc2, 0xb6, 0x62, 0x55, 0x2a, 0x54, 0x22, 0x39, 0x68, 0xa2, 0xf8, 0xa8, 0xef,
	0xfc, 0xbf, 0x24, 0x9d, 0x1a, 0x78, 0xc5, 0xf5, 0x75, 0xc0, 0x20, 0xba, 0x14, 0x41, 0x29, 0x5c,
	0xc7, 0x1f, 0x8e, 0x4e, 0x00, 0x5d, 0x71, 0x72, 0x86, 0xb1, 0x21, 0x0b, 0x96, 0x0b, 0x67, 0x04,
	0x8b, 0x7d, 0xda, 0x77, 0xfe, 0x3f, 0x2c, 0xb5, 0x01, 0xbf, 0x67, 0xb9, 0x20, 0x3e, 0x1e, 0x67,
	0xa2, 0x15, 0x99, 0x33, 0x86, 0xa6, 0x76, 0xdf, 0xf9, 0x86, 0xa0, 0xe6, 0x13, 0xd5, 0xf8, 0xc0,
	0x18, 0x27, 0x12, 0x93, 0xb5, 0x36, 0x9f, 0x64, 0xb2, 0x15, 0xc9, 0x67, 0x63, 0x9f, 0x04, 0x7b,
	0x6e, 0x69, 0xed, 0xbe, 0xdc, 0x7b, 0x9f, 0xf1, 0xc9, 0xb7, 0xad, 0x67, 0xdd, 0x6f, 0x3d, 0xeb,
	0xe7, 0xd6, 0xb3, 0xbe, 0xec, 0xbc, 0xc1, 0xfd, 0xce, 0x1b, 0x7c, 0xdf, 0x79, 0x83, 0x8f, 0xc3,
	0x36, 0x5a, 0x1c, 0xc0, 0xf3, 0x3d, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x39, 0x27, 0x2a,
	0x0f, 0x03, 0x00, 0x00,
}
