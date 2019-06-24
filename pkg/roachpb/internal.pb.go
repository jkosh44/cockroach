// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roachpb/internal.proto

package roachpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import encoding_binary "encoding/binary"

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

// InternalTimeSeriesData is a collection of data samples for some
// measurable value, where each sample is taken over a uniform time
// interval.
//
// The collection itself contains a start timestamp (in seconds since the unix
// epoch) and a sample duration (in milliseconds). Each sample in the collection
// will contain a positive integer offset that indicates the length of time
// between the start_timestamp of the collection and the time when the sample
// began, expressed as an whole number of sample intervals. For example, if the
// sample duration is 60000 (indicating 1 minute), then a contained sample with
// an offset value of 5 begins (5*60000ms = 300000ms = 5 minutes) after the
// start timestamp of this data.
//
// This is meant to be an efficient internal representation of time series data,
// ensuring that very little redundant data is stored on disk. With this goal in
// mind, this message does not identify the variable which is actually being
// measured; that information is expected be encoded in the key where this
// message is stored.
//
// The actual samples can be stored in one of two formats: a Row-based format in
// the "samples" repeated field, or a columnar format spread across several
// different repeated columns. The row-based format will eventually be
// deprecated, but is maintained for backwards compatibility. There is no flag
// that indicates whether the data is stored as rows or columns; columnar data
// is indicated by the presence of a non-zero-length "offset" collection, while
// row data is indicated by a non-zero-length "samples" collection. Each data
// message must have all of its data either row format or column format.
//
// One feature of the columnar layout is that it is "sparse", and columns
// without useful information are elided. Specifically, the "offset" and "last"
// columns will always be populated, but the other columns are only populated
// for resolutions which contain detailed "rollup" information about long sample
// periods. In the case of non-rollup data there is only one measurement per
// sample period, and the value of all optional columns can be directly inferred
// from the "last" column. Eliding those columns represents a significant memory
// and on-disk savings for our highest resolution data.
type InternalTimeSeriesData struct {
	// Holds a wall time, expressed as a unix epoch time in nanoseconds. This
	// represents the earliest possible timestamp for a sample within the
	// collection.
	StartTimestampNanos int64 `protobuf:"varint,1,opt,name=start_timestamp_nanos,json=startTimestampNanos" json:"start_timestamp_nanos"`
	// The duration of each sample interval, expressed in nanoseconds.
	SampleDurationNanos int64 `protobuf:"varint,2,opt,name=sample_duration_nanos,json=sampleDurationNanos" json:"sample_duration_nanos"`
	// The data samples for this metric if this data was written in the old
	// row format.
	Samples []InternalTimeSeriesSample `protobuf:"bytes,3,rep,name=samples" json:"samples"` // Deprecated: Do not use.
	// Columnar array containing the ordered offsets of the samples in this
	// data set.
	Offset []int32 `protobuf:"varint,4,rep,packed,name=offset" json:"offset,omitempty"`
	// Columnar array containing the last value of the samples in this data set;
	// the "last" value is the most recent individual measurement during a sample
	// period.
	Last []float64 `protobuf:"fixed64,5,rep,packed,name=last" json:"last,omitempty"`
	// Columnar array containing the total number of measurements that were taken
	// during this sample period.
	Count []uint32 `protobuf:"varint,6,rep,packed,name=count" json:"count,omitempty"`
	// Columnar array containing the sum of measurements that were taken during
	// this sample period. If this column is elided, its value for all samples is
	// 1.
	Sum []float64 `protobuf:"fixed64,7,rep,packed,name=sum" json:"sum,omitempty"`
	// Columnar array containing the maximum value of any single measurement taken
	// during this sample period. If this column is elided, its value for all
	// samples is equal to "last".
	Max []float64 `protobuf:"fixed64,8,rep,packed,name=max" json:"max,omitempty"`
	// Columnar array containing the minimum value of any single measurements
	// taken during this sample period. If this column is elided, its value for
	// all samples is equal to "last".
	Min []float64 `protobuf:"fixed64,9,rep,packed,name=min" json:"min,omitempty"`
	// Columnar array containing the first value of the samples in this data set;
	// the "first" value is the earliest individual measurement during a sample
	// period. If this column is elided, its value for all samples is equal to
	// "last".
	First []float64 `protobuf:"fixed64,10,rep,packed,name=first" json:"first,omitempty"`
	// Columnar array containing the variance of measurements that were taken
	// during this sample period. If this column is elided, its value for all
	// samples is zero.
	Variance             []float64 `protobuf:"fixed64,11,rep,packed,name=variance" json:"variance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InternalTimeSeriesData) Reset()         { *m = InternalTimeSeriesData{} }
func (m *InternalTimeSeriesData) String() string { return proto.CompactTextString(m) }
func (*InternalTimeSeriesData) ProtoMessage()    {}
func (*InternalTimeSeriesData) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_c54ce32009eb2046, []int{0}
}
func (m *InternalTimeSeriesData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InternalTimeSeriesData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *InternalTimeSeriesData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTimeSeriesData.Merge(dst, src)
}
func (m *InternalTimeSeriesData) XXX_Size() int {
	return m.Size()
}
func (m *InternalTimeSeriesData) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTimeSeriesData.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTimeSeriesData proto.InternalMessageInfo

// A InternalTimeSeriesSample represents data gathered from multiple
// measurements of a variable value over a given period of time. The
// length of that period of time is stored in an
// InternalTimeSeriesData message; a sample cannot be interpreted
// correctly without a start timestamp and sample duration.
//
// Each sample may contain data gathered from multiple measurements of the same
// variable, as long as all of those measurements occurred within the sample
// period. The sample stores several aggregated values from these measurements:
// - The sum of all measured values
// - A count of all measurements taken
// - The maximum individual measurement seen
// - The minimum individual measurement seen
//
// If zero measurements are present in a sample, then it should be omitted
// entirely from any collection it would be a part of.
//
// If the count of measurements is 1, then max and min fields may be omitted
// and assumed equal to the sum field.
type InternalTimeSeriesSample struct {
	// Temporal offset from the "start_timestamp" of the InternalTimeSeriesData
	// collection this data point is part in. The units of this value are
	// determined by the value of the "sample_duration_milliseconds" field of
	// the TimeSeriesData collection.
	Offset int32 `protobuf:"varint,1,opt,name=offset" json:"offset"`
	// Sum of all measurements.
	Sum float64 `protobuf:"fixed64,7,opt,name=sum" json:"sum"`
	// Count of measurements taken within this sample.
	Count uint32 `protobuf:"varint,6,opt,name=count" json:"count"`
	// Maximum encountered measurement in this sample.
	Max *float64 `protobuf:"fixed64,8,opt,name=max" json:"max,omitempty"`
	// Minimum encountered measurement in this sample.
	Min                  *float64 `protobuf:"fixed64,9,opt,name=min" json:"min,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTimeSeriesSample) Reset()         { *m = InternalTimeSeriesSample{} }
func (m *InternalTimeSeriesSample) String() string { return proto.CompactTextString(m) }
func (*InternalTimeSeriesSample) ProtoMessage()    {}
func (*InternalTimeSeriesSample) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_c54ce32009eb2046, []int{1}
}
func (m *InternalTimeSeriesSample) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InternalTimeSeriesSample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *InternalTimeSeriesSample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTimeSeriesSample.Merge(dst, src)
}
func (m *InternalTimeSeriesSample) XXX_Size() int {
	return m.Size()
}
func (m *InternalTimeSeriesSample) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTimeSeriesSample.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTimeSeriesSample proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InternalTimeSeriesData)(nil), "cockroach.roachpb.InternalTimeSeriesData")
	proto.RegisterType((*InternalTimeSeriesSample)(nil), "cockroach.roachpb.InternalTimeSeriesSample")
}
func (m *InternalTimeSeriesData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalTimeSeriesData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintInternal(dAtA, i, uint64(m.StartTimestampNanos))
	dAtA[i] = 0x10
	i++
	i = encodeVarintInternal(dAtA, i, uint64(m.SampleDurationNanos))
	if len(m.Samples) > 0 {
		for _, msg := range m.Samples {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintInternal(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Offset) > 0 {
		dAtA2 := make([]byte, len(m.Offset)*10)
		var j1 int
		for _, num1 := range m.Offset {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintInternal(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.Last) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Last)*8))
		for _, num := range m.Last {
			f3 := math.Float64bits(float64(num))
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f3))
			i += 8
		}
	}
	if len(m.Count) > 0 {
		dAtA5 := make([]byte, len(m.Count)*10)
		var j4 int
		for _, num := range m.Count {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x32
		i++
		i = encodeVarintInternal(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	if len(m.Sum) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Sum)*8))
		for _, num := range m.Sum {
			f6 := math.Float64bits(float64(num))
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f6))
			i += 8
		}
	}
	if len(m.Max) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Max)*8))
		for _, num := range m.Max {
			f7 := math.Float64bits(float64(num))
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f7))
			i += 8
		}
	}
	if len(m.Min) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Min)*8))
		for _, num := range m.Min {
			f8 := math.Float64bits(float64(num))
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f8))
			i += 8
		}
	}
	if len(m.First) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.First)*8))
		for _, num := range m.First {
			f9 := math.Float64bits(float64(num))
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f9))
			i += 8
		}
	}
	if len(m.Variance) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Variance)*8))
		for _, num := range m.Variance {
			f10 := math.Float64bits(float64(num))
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f10))
			i += 8
		}
	}
	return i, nil
}

func (m *InternalTimeSeriesSample) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalTimeSeriesSample) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintInternal(dAtA, i, uint64(m.Offset))
	dAtA[i] = 0x30
	i++
	i = encodeVarintInternal(dAtA, i, uint64(m.Count))
	dAtA[i] = 0x39
	i++
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Sum))))
	i += 8
	if m.Max != nil {
		dAtA[i] = 0x41
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.Max))))
		i += 8
	}
	if m.Min != nil {
		dAtA[i] = 0x49
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.Min))))
		i += 8
	}
	return i, nil
}

func encodeVarintInternal(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InternalTimeSeriesData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.StartTimestampNanos))
	n += 1 + sovInternal(uint64(m.SampleDurationNanos))
	if len(m.Samples) > 0 {
		for _, e := range m.Samples {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	if len(m.Offset) > 0 {
		l = 0
		for _, e := range m.Offset {
			l += sovInternal(uint64(e))
		}
		n += 1 + sovInternal(uint64(l)) + l
	}
	if len(m.Last) > 0 {
		n += 1 + sovInternal(uint64(len(m.Last)*8)) + len(m.Last)*8
	}
	if len(m.Count) > 0 {
		l = 0
		for _, e := range m.Count {
			l += sovInternal(uint64(e))
		}
		n += 1 + sovInternal(uint64(l)) + l
	}
	if len(m.Sum) > 0 {
		n += 1 + sovInternal(uint64(len(m.Sum)*8)) + len(m.Sum)*8
	}
	if len(m.Max) > 0 {
		n += 1 + sovInternal(uint64(len(m.Max)*8)) + len(m.Max)*8
	}
	if len(m.Min) > 0 {
		n += 1 + sovInternal(uint64(len(m.Min)*8)) + len(m.Min)*8
	}
	if len(m.First) > 0 {
		n += 1 + sovInternal(uint64(len(m.First)*8)) + len(m.First)*8
	}
	if len(m.Variance) > 0 {
		n += 1 + sovInternal(uint64(len(m.Variance)*8)) + len(m.Variance)*8
	}
	return n
}

func (m *InternalTimeSeriesSample) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.Offset))
	n += 1 + sovInternal(uint64(m.Count))
	n += 9
	if m.Max != nil {
		n += 9
	}
	if m.Min != nil {
		n += 9
	}
	return n
}

func sovInternal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternal(x uint64) (n int) {
	return sovInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InternalTimeSeriesData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: InternalTimeSeriesData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalTimeSeriesData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimestampNanos", wireType)
			}
			m.StartTimestampNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTimestampNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SampleDurationNanos", wireType)
			}
			m.SampleDurationNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SampleDurationNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Samples", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Samples = append(m.Samples, InternalTimeSeriesSample{})
			if err := m.Samples[len(m.Samples)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Offset = append(m.Offset, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
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
				if elementCount != 0 && len(m.Offset) == 0 {
					m.Offset = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowInternal
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Offset = append(m.Offset, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
		case 5:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Last = append(m.Last, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.Last) == 0 {
					m.Last = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Last = append(m.Last, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Last", wireType)
			}
		case 6:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Count = append(m.Count, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
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
				if elementCount != 0 && len(m.Count) == 0 {
					m.Count = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowInternal
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Count = append(m.Count, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
		case 7:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Sum = append(m.Sum, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.Sum) == 0 {
					m.Sum = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Sum = append(m.Sum, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
		case 8:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Max = append(m.Max, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.Max) == 0 {
					m.Max = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Max = append(m.Max, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
		case 9:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Min = append(m.Min, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.Min) == 0 {
					m.Min = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Min = append(m.Min, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
		case 10:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.First = append(m.First, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.First) == 0 {
					m.First = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.First = append(m.First, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field First", wireType)
			}
		case 11:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Variance = append(m.Variance, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowInternal
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
					return ErrInvalidLengthInternal
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.Variance) == 0 {
					m.Variance = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Variance = append(m.Variance, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Variance", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *InternalTimeSeriesSample) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: InternalTimeSeriesSample: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalTimeSeriesSample: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Sum = float64(math.Float64frombits(v))
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.Max = &v2
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.Min = &v2
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func skipInternal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
				return 0, ErrInvalidLengthInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInternal
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
				next, err := skipInternal(dAtA[start:])
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
	ErrInvalidLengthInternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternal   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("roachpb/internal.proto", fileDescriptor_internal_c54ce32009eb2046) }

var fileDescriptor_internal_c54ce32009eb2046 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x8e, 0xda, 0x30,
	0x00, 0x86, 0x31, 0x26, 0x40, 0x8d, 0x90, 0xa8, 0x4b, 0x23, 0x0b, 0x55, 0x69, 0xca, 0x64, 0xa9,
	0x52, 0x90, 0x3a, 0x75, 0x8e, 0x58, 0x3a, 0xb4, 0x03, 0x30, 0x75, 0x41, 0x6e, 0x6a, 0xa8, 0x55,
	0x62, 0x47, 0xb6, 0xa9, 0x78, 0x0c, 0x5e, 0xa2, 0xef, 0xc2, 0xd8, 0xb1, 0x53, 0xd5, 0xcb, 0xbd,
	0xc8, 0x29, 0x8e, 0x13, 0xee, 0x84, 0x6e, 0xc3, 0xdf, 0xff, 0xff, 0x16, 0x5f, 0x8c, 0x42, 0xad,
	0x58, 0xf6, 0xa3, 0xf8, 0xb6, 0x10, 0xd2, 0x72, 0x2d, 0xd9, 0x21, 0x29, 0xb4, 0xb2, 0x0a, 0xbf,
	0xcc, 0x54, 0xf6, 0xd3, 0x65, 0x89, 0x6f, 0xcc, 0xa6, 0x7b, 0xb5, 0x57, 0x2e, 0x5d, 0x54, 0xbf,
	0xea, 0xe2, 0xfc, 0x37, 0x44, 0xe1, 0x27, 0xbf, 0xdd, 0x88, 0x9c, 0xaf, 0xb9, 0x16, 0xdc, 0x2c,
	0x99, 0x65, 0xf8, 0x23, 0x7a, 0x6d, 0x2c, 0xd3, 0x76, 0x6b, 0x45, 0xce, 0x8d, 0x65, 0x79, 0xb1,
	0x95, 0x4c, 0x2a, 0x43, 0x40, 0x0c, 0x28, 0x4c, 0x7b, 0x97, 0x7f, 0x6f, 0x3b, 0xab, 0x57, 0xae,
	0xb2, 0x69, 0x1a, 0x5f, 0xaa, 0x82, 0x5b, 0xb2, 0xbc, 0x38, 0xf0, 0xed, 0xf7, 0xa3, 0x66, 0x56,
	0x28, 0xe9, 0x97, 0xdd, 0x27, 0x4b, 0x57, 0x59, 0xfa, 0x46, 0xbd, 0xfc, 0x8c, 0x06, 0x35, 0x36,
	0x04, 0xc6, 0x90, 0x8e, 0x3e, 0xbc, 0x4f, 0x6e, 0x4c, 0x92, 0xdb, 0xff, 0xbb, 0x76, 0x9b, 0xb4,
	0x5f, 0x5d, 0x4c, 0xc0, 0xaa, 0xb9, 0x03, 0xcf, 0x50, 0x5f, 0xed, 0x76, 0x86, 0x5b, 0xd2, 0x8b,
	0x21, 0x0d, 0xd2, 0xee, 0x04, 0xac, 0x3c, 0xc1, 0x21, 0xea, 0x1d, 0x98, 0xb1, 0x24, 0x88, 0x21,
	0x05, 0x2e, 0x71, 0x67, 0x4c, 0x50, 0x90, 0xa9, 0xa3, 0xb4, 0xa4, 0x1f, 0x43, 0x3a, 0x76, 0x41,
	0x0d, 0xf0, 0x14, 0x41, 0x73, 0xcc, 0xc9, 0xa0, 0x1d, 0x54, 0xc7, 0x8a, 0xe6, 0xec, 0x44, 0x86,
	0x57, 0x9a, 0xb3, 0x93, 0xa3, 0x42, 0x92, 0x17, 0x8f, 0xa8, 0x90, 0xd5, 0xdd, 0x3b, 0xa1, 0x8d,
	0x25, 0xa8, 0xe5, 0x35, 0xc0, 0x11, 0x1a, 0xfe, 0x62, 0x5a, 0x30, 0x99, 0x71, 0x32, 0x6a, 0xc3,
	0x96, 0xcd, 0xcf, 0x00, 0x91, 0xe7, 0xbc, 0xf1, 0x9b, 0x56, 0xb3, 0x7a, 0x9a, 0xc0, 0x7f, 0xe0,
	0x46, 0x74, 0x76, 0x15, 0x02, 0x74, 0xec, 0x43, 0xaf, 0x14, 0x36, 0x4a, 0x80, 0x02, 0x9f, 0x38,
	0xa9, 0x49, 0x23, 0x05, 0xa8, 0x17, 0x9a, 0x34, 0x42, 0x35, 0x11, 0x32, 0x7d, 0x77, 0xb9, 0x8b,
	0x3a, 0x97, 0x32, 0x02, 0x7f, 0xca, 0x08, 0xfc, 0x2d, 0x23, 0xf0, 0xbf, 0x8c, 0xc0, 0xf9, 0x3e,
	0xea, 0x7c, 0x1d, 0xf8, 0x97, 0x7a, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x47, 0x27, 0x4b, 0x9f,
	0x02, 0x00, 0x00,
}
