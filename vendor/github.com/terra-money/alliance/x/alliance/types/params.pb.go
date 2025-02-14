// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alliance/alliance/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	RewardDelayTime time.Duration `protobuf:"bytes,1,opt,name=reward_delay_time,json=rewardDelayTime,proto3,stdduration" json:"reward_delay_time"`
	// Time interval between consecutive applications of `take_rate`
	TakeRateClaimInterval time.Duration `protobuf:"bytes,2,opt,name=take_rate_claim_interval,json=takeRateClaimInterval,proto3,stdduration" json:"take_rate_claim_interval"`
	// Last application of `take_rate` on assets
	LastTakeRateClaimTime time.Time `protobuf:"bytes,3,opt,name=last_take_rate_claim_time,json=lastTakeRateClaimTime,proto3,stdtime" json:"last_take_rate_claim_time"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_245e75f0d2ae44fd, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetRewardDelayTime() time.Duration {
	if m != nil {
		return m.RewardDelayTime
	}
	return 0
}

func (m *Params) GetTakeRateClaimInterval() time.Duration {
	if m != nil {
		return m.TakeRateClaimInterval
	}
	return 0
}

func (m *Params) GetLastTakeRateClaimTime() time.Time {
	if m != nil {
		return m.LastTakeRateClaimTime
	}
	return time.Time{}
}

type RewardHistory struct {
	Denom    string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Index    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=index,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"index"`
	Alliance string                                 `protobuf:"bytes,3,opt,name=alliance,proto3" json:"alliance,omitempty"`
}

func (m *RewardHistory) Reset()         { *m = RewardHistory{} }
func (m *RewardHistory) String() string { return proto.CompactTextString(m) }
func (*RewardHistory) ProtoMessage()    {}
func (*RewardHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_245e75f0d2ae44fd, []int{1}
}
func (m *RewardHistory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardHistory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardHistory.Merge(m, src)
}
func (m *RewardHistory) XXX_Size() int {
	return m.Size()
}
func (m *RewardHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardHistory.DiscardUnknown(m)
}

var xxx_messageInfo_RewardHistory proto.InternalMessageInfo

func (m *RewardHistory) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *RewardHistory) GetAlliance() string {
	if m != nil {
		return m.Alliance
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "alliance.alliance.Params")
	proto.RegisterType((*RewardHistory)(nil), "alliance.alliance.RewardHistory")
}

func init() { proto.RegisterFile("alliance/alliance/params.proto", fileDescriptor_245e75f0d2ae44fd) }

var fileDescriptor_245e75f0d2ae44fd = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x83, 0x3b, 0x5d, 0x8d, 0x10, 0xba, 0xe8, 0x90, 0xda, 0x0c, 0x0e, 0xba, 0x01,
	0xb1, 0x34, 0x41, 0xb0, 0x21, 0xa6, 0x92, 0x01, 0x58, 0x40, 0x51, 0x27, 0x84, 0x88, 0xdc, 0xe4,
	0x11, 0xac, 0x8b, 0xe3, 0xc8, 0x76, 0xe1, 0xf2, 0x2d, 0x6e, 0x44, 0x42, 0x48, 0x7c, 0x08, 0x3e,
	0xc4, 0x8d, 0x15, 0x13, 0x62, 0x28, 0xa8, 0x5d, 0xf8, 0x18, 0xc8, 0x76, 0xd2, 0x56, 0x65, 0xb9,
	0x29, 0xef, 0xe5, 0xff, 0xfc, 0x7b, 0xef, 0x6f, 0x3f, 0x4c, 0x68, 0x55, 0x31, 0x5a, 0xe7, 0x10,
	0x6f, 0x82, 0x86, 0x4a, 0xca, 0x55, 0xd4, 0x48, 0xa1, 0x85, 0x7f, 0xd2, 0xff, 0x8e, 0xfa, 0x20,
	0x38, 0x2d, 0x45, 0x29, 0xac, 0x1a, 0x9b, 0xc8, 0x15, 0x06, 0xa3, 0x5c, 0x28, 0x2e, 0x54, 0xe6,
	0x04, 0x97, 0x74, 0x12, 0x29, 0x85, 0x28, 0x2b, 0x88, 0x6d, 0x36, 0x9b, 0xbf, 0x8f, 0x8b, 0xb9,
	0xa4, 0x9a, 0x89, 0xba, 0xd3, 0xc3, 0x7d, 0x5d, 0x33, 0x0e, 0x4a, 0x53, 0xde, 0xb8, 0x82, 0xb3,
	0xaf, 0x07, 0xf8, 0xe8, 0xb5, 0x9d, 0xca, 0x7f, 0x85, 0x4f, 0x24, 0x7c, 0xa2, 0xb2, 0xc8, 0x0a,
	0xa8, 0x68, 0x9b, 0x99, 0xd2, 0x21, 0xba, 0x87, 0x1e, 0xdc, 0x7a, 0x34, 0x8a, 0x1c, 0x27, 0xea,
	0x39, 0x51, 0xd2, 0xf5, 0x99, 0x1c, 0x5f, 0x2d, 0x43, 0xef, 0xf3, 0xef, 0x10, 0xa5, 0x77, 0xdc,
	0xe9, 0xc4, 0x1c, 0x9e, 0x32, 0x0e, 0xfe, 0x5b, 0x3c, 0xd4, 0xf4, 0x1c, 0x32, 0x49, 0x35, 0x64,
	0x79, 0x45, 0x19, 0xcf, 0x58, 0xad, 0x41, 0x7e, 0xa4, 0xd5, 0xf0, 0xe0, 0xfa, 0xdc, 0xbb, 0x06,
	0x92, 0x52, 0x0d, 0xcf, 0x0c, 0xe2, 0x45, 0x47, 0xf0, 0xdf, 0xe1, 0x51, 0x45, 0x95, 0xce, 0xf6,
	0x5b, 0xd8, 0xb1, 0x6f, 0x58, 0x7c, 0xf0, 0x1f, 0x7e, 0xda, 0xdb, 0x77, 0xfc, 0x4b, 0xcb, 0x37,
	0x98, 0xe9, 0x6e, 0x0f, 0x53, 0xf5, 0xe4, 0xe6, 0xdf, 0x6f, 0x21, 0x3a, 0xfb, 0x82, 0xf0, 0xed,
	0xd4, 0xfa, 0x7a, 0xce, 0x94, 0x16, 0xb2, 0xf5, 0x4f, 0xf1, 0x61, 0x01, 0xb5, 0xe0, 0xf6, 0x6a,
	0x06, 0xa9, 0x4b, 0xfc, 0x14, 0x1f, 0xb2, 0xba, 0x80, 0x0b, 0x6b, 0x6c, 0x30, 0x79, 0x6a, 0xe8,
	0xbf, 0x96, 0xe1, 0xfd, 0x92, 0xe9, 0x0f, 0xf3, 0x59, 0x94, 0x0b, 0xde, 0x3d, 0x5c, 0xf7, 0x19,
	0xab, 0xe2, 0x3c, 0xd6, 0x6d, 0x03, 0x2a, 0x4a, 0x20, 0xff, 0xf1, 0x7d, 0x8c, 0xbb, 0x77, 0x4d,
	0x20, 0x4f, 0x1d, 0xca, 0x0f, 0xf0, 0x71, 0xbf, 0x19, 0xd6, 0xd0, 0x20, 0xdd, 0xe4, 0x6e, 0xba,
	0xc9, 0xcb, 0xab, 0x15, 0x41, 0x8b, 0x15, 0x41, 0x7f, 0x56, 0x04, 0x5d, 0xae, 0x89, 0xb7, 0x58,
	0x13, 0xef, 0xe7, 0x9a, 0x78, 0x6f, 0x1e, 0xee, 0x34, 0xd6, 0x20, 0x25, 0x1d, 0x73, 0x51, 0x43,
	0xbb, 0x5d, 0xc5, 0x8b, 0x6d, 0x68, 0xc7, 0x98, 0x1d, 0xd9, 0x4b, 0x7a, 0xfc, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x61, 0xb2, 0x67, 0xb7, 0xb7, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RewardDelayTime != that1.RewardDelayTime {
		return false
	}
	if this.TakeRateClaimInterval != that1.TakeRateClaimInterval {
		return false
	}
	if !this.LastTakeRateClaimTime.Equal(that1.LastTakeRateClaimTime) {
		return false
	}
	return true
}
func (this *RewardHistory) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RewardHistory)
	if !ok {
		that2, ok := that.(RewardHistory)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if !this.Index.Equal(that1.Index) {
		return false
	}
	if this.Alliance != that1.Alliance {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastTakeRateClaimTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastTakeRateClaimTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.TakeRateClaimInterval, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.TakeRateClaimInterval):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.RewardDelayTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.RewardDelayTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RewardHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardHistory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardHistory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Alliance) > 0 {
		i -= len(m.Alliance)
		copy(dAtA[i:], m.Alliance)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Alliance)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Index.Size()
		i -= size
		if _, err := m.Index.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.RewardDelayTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.TakeRateClaimInterval)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastTakeRateClaimTime)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *RewardHistory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.Index.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.Alliance)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDelayTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.RewardDelayTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakeRateClaimInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.TakeRateClaimInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTakeRateClaimTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastTakeRateClaimTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *RewardHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: RewardHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Index.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alliance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alliance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
