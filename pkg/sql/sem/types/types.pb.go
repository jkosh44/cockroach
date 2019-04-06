// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/sem/types/types.proto

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_lib_pq_oid "github.com/lib/pq/oid"

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

// See the comment header for the T.SemanticType method for more details.
type SemanticType int32

const (
	// BOOL is the family of boolean true/false types.
	//
	//   Canonical: types.Bool
	//   Oid      : T_bool
	//
	// Examples:
	//   BOOL
	//
	BOOL SemanticType = 0
	// INT is the family of signed integer types.
	//
	//   Canonical: types.Int
	//   Oid      : T_int8, T_int4, T_int2
	//   Width    : 64, 32, 16
	//
	// Examples:
	//   INT
	//   INT8
	//   INT4
	//
	INT SemanticType = 1
	// FLOAT is the family of base-2 floating-point types (IEEE 754).
	//
	//   Canonical: types.Float
	//   Oid      : T_float8, T_float4
	//   Width    : 64, 32
	//
	// Examples:
	//   FLOAT8
	//   FLOAT4
	//
	FLOAT SemanticType = 2
	// DECIMAL is the family of base-10 floating and fixed point types.
	//
	//   Canonical    : types.Decimal
	//   Oid          : T_numeric
	//   Precision    : max # decimal digits (0 = no specified limit)
	//   Width (Scale): # digits after decimal point (0 = no specified limit)
	//
	// Examples:
	//   DECIMAL
	//   DECIMAL(10)
	//   DECIMAL(10,3)
	//
	DECIMAL SemanticType = 3
	// DATE is the family of date types that store only year/month/day with no
	// time component.
	//
	//   Canonical: types.Date
	//   Oid      : T_date
	//
	// Examples:
	//   DATE
	//
	DATE SemanticType = 4
	// TIMESTAMP is the family of date types that store a year/month/day date
	// component, as well as an hour/minute/second time component. There is no
	// timezone component (see TIMESTAMPTZ). Seconds can have varying precision
	// (defaults to microsecond precision). Currently, only microsecond precision
	// is supported.
	//
	//   Canonical: types.Timestamp
	//   Oid      : T_timestamp
	//   Precision: fractional seconds (3 = ms, 0,6 = us, 9 = ns, etc.)
	//
	// Examples:
	//   TIMESTAMP
	//   TIMESTAMP(6)
	//
	TIMESTAMP SemanticType = 5
	// INTERVAL is the family of types describing a duration of time. Currently,
	// only microsecond precision is supported.
	//
	//   Canonical: types.Interval
	//   Oid      : T_interval
	//
	// Examples:
	//   INTERVAL
	//
	INTERVAL SemanticType = 6
	// STRING is the family of types containing Unicode textual strings. This
	// family includes types constructed by STRING, VARCHAR, CHAR, and "char"
	// column type definitions (CHAR and "char" are distinct PG types). Note
	// that while STRING and VARCHAR have no default width limit, CHAR has a
	// default width of 1.
	// TODO(andyk): "char" should have default width of 1 as well, but doesn't.
	//
	//   Canonical: types.String
	//   Oid      : T_text, T_varchar, T_bpchar, T_char
	//   Width    : max # characters (0 = no specified limit)
	//
	// Examples:
	//   STRING
	//   TEXT
	//   VARCHAR(10)
	//   CHAR
	//
	STRING SemanticType = 7
	// BYTES is the family of types containing a list of raw byte values.
	//
	//   Canonical: types.BYTES
	//   Oid      : T_bytea
	//
	// Examples:
	//   BYTES
	//
	BYTES SemanticType = 8
	// TIMESTAMPTZ is the family of date types that store a year/month/day date
	// component, as well as an hour/minute/second time component, along with a
	// timezone. Seconds can have varying precision (defaults to microsecond
	// precision). Currently, only microsecond precision is supported.
	//
	//   Canonical: types.TimestampTZ
	//   Oid      : T_timestamptz
	//   Precision: fractional seconds (3 = ms, 0,6 = us, 9 = ns, etc.)
	//
	// Examples:
	//   TIMESTAMPTZ
	//   TIMESTAMPTZ(6)
	//
	TIMESTAMPTZ SemanticType = 9
	// COLLATEDSTRING is the family of types containing Unicode textual strings
	// with an associated COLLATE value that specifies the locale used for various
	// character-based operations such as sorting, pattern matching, and builtin
	// functions like lower and upper.
	//
	//   Oid      : T_text, T_varchar, T_bpchar, T_char
	//   Width    : max # characters (0 = no specified limit)
	//   Locale   : name of locale (e.g. EN or DE)
	//
	// Examples:
	//   STRING COLLATE en
	//   VARCHAR(10) COLLATE de
	//
	COLLATEDSTRING SemanticType = 10
	// OID is the family of types containing Postgres Object ID (OID) values.
	// OIDs are integer values that identify some object in the database, like a
	// type, relation, or procedure.
	//
	//   Canonical: types.Oid
	//   Oid      : T_oid, T_regclass, T_regproc, T_regprocedure, T_regtype,
	//              T_regnamespace
	//
	// Examples:
	//   OID
	//   REGCLASS
	//   REGPROC
	//
	// TODO(andyk): OIDs should be part of the INT family, since they are treated
	//              as equivalent to INTs by PG.
	OID SemanticType = 12
	// UNKNOWN is a special type that tags expressions that statically evaluate
	// to NULL. If an expression has type UNKNOWN, then it *must* be NULL. But
	// the inverse is not true, since other types allow NULL values as well. The
	// UNKNOWN type is not supported as a table column type, but can be
	// transferred through distsql streams.
	//
	//   Canonical: types.Unknown
	//   Oid      : T_unknown
	//
	UNKNOWN SemanticType = 13
	// UUID is the family of types containing universally unique identifiers.
	// A UUID is a 128-bit quantity that is very unlikely to ever be generated
	// again, and so can be relied on to be distinct from all other UUID values.
	//
	//   Canonical: types.Uuid
	//   Oid      : T_uuid
	//
	// Examples:
	//   UUID
	//
	UUID SemanticType = 14
	// ARRAY is a family of non-scalar types that contain an ordered list of
	// elements. The elements of an array must all share the same type. Elements
	// can have have any type, including ARRAY. However, while the types package
	// supports nested arrays, other parts of CRDB do not currently support them.
	// Also, the length of array dimension(s) are ignored by PG and CRDB (e.g.
	// an array of length 11 could be inserted into a column declared as INT[11]).
	//
	// Array OID values are special. Rather than having a single T_array OID,
	// Postgres defines a separate OID for each possible array element type.
	// Here are some examples:
	//
	//   T__int8: array of int8 values
	//   T__text: array of text values
	//
	// Notice that each array OID has double underscores to distinguish it from
	// the OID of the scalar type it contains.
	//
	//   Oid          : T__int, T__text, T__numeric, etc.
	//   ArrayContents: types.T of the array element type
	//
	// Examples:
	//   INT[]
	//   VARCHAR(10)[] COLLATE EN
	//   DECIMAL(10,1)[]
	//   TIMESTAMP[5]
	//
	ARRAY SemanticType = 15
	// INET is the family of types containing IPv4 or IPv6 network address
	// identifiers (e.g. 192.168.100.128/25 or FE80:CD00:0:CDE:1257:0:211E:729C).
	//
	//   Canonical: types.INet
	//   Oid      : T_inet
	//
	// Examples:
	//   INET
	//
	INET SemanticType = 16
	// TIME is the family of date types that store only hour/minute/second with
	// no date component. There is no timezone component. Seconds can have
	// varying precision (defaults to microsecond precision). Currently, only
	// microsecond precision is supported.
	//
	//   Canonical: types.Time
	//   Oid      : T_time
	//   Precision: fractional seconds (3 = ms, 0,6 = us, 9 = ns, etc.)
	//
	// Examples:
	//   TIME
	//   TIME(6)
	//
	TIME SemanticType = 17
	// JSON is the family of types containing JavaScript Object Notation (JSON)
	// values. Currently, CRDB only supports JSONB values, which are stored in a
	// decomposed binary format.
	//
	//   Canonical: types.Jsonb
	//   Oid      : T_jsonb
	//
	// Examples:
	//   JSON
	//   JSONB
	//
	JSON SemanticType = 18
	// TUPLE is a family of non-scalar structural types that describes the fields
	// of a row or record. The fields can be of any type, including nested TUPLE
	// types and arrays. Fields can also have optional labels. Currently, CRDB
	// does not support TUPLE types as column types, but it is possible to
	// construct tuples using the ROW function or tuple construction syntax.
	//
	//   Oid          : T_record
	//   TupleContents: []types.T of each tuple field
	//   TupleLabels  : []string of each tuple label
	//
	// Examples:
	//   (1, 'foo')
	//   ((1, 'foo') AS num, str)
	//   ROW(1, 'foo')
	//   (ROW(1, 'foo') AS num, str)
	//
	TUPLE SemanticType = 20
	// BIT is the family of types containing ordered lists of bit values (0 or 1).
	// Note that while VARBIT has no default width limit, BIT has a default width
	// limit of 1.
	//
	//   Canonical: types.VarBit
	//   Oid      : T_varbit, T_bit
	//   Width    : max # of bits (0 = no specified limit)
	//
	// Examples:
	//   VARBIT
	//   VARBIT(10)
	//   BIT
	//   BIT(10)
	//
	BIT SemanticType = 21
	// ANY is a special type used during static analysis as a wildcard type that
	// matches any other type, including scalar, array, and tuple types.
	// Execution-time values should never have this type. As an example of its
	// use, many SQL builtin functions allow an input value to be of any type,
	// and so use this type in their static definitions.
	//
	//   Canonical: types.Any
	//   Oid      : T_anyelement
	//
	ANY SemanticType = 100
)

var SemanticType_name = map[int32]string{
	0:   "BOOL",
	1:   "INT",
	2:   "FLOAT",
	3:   "DECIMAL",
	4:   "DATE",
	5:   "TIMESTAMP",
	6:   "INTERVAL",
	7:   "STRING",
	8:   "BYTES",
	9:   "TIMESTAMPTZ",
	10:  "COLLATEDSTRING",
	12:  "OID",
	13:  "UNKNOWN",
	14:  "UUID",
	15:  "ARRAY",
	16:  "INET",
	17:  "TIME",
	18:  "JSON",
	20:  "TUPLE",
	21:  "BIT",
	100: "ANY",
}
var SemanticType_value = map[string]int32{
	"BOOL":           0,
	"INT":            1,
	"FLOAT":          2,
	"DECIMAL":        3,
	"DATE":           4,
	"TIMESTAMP":      5,
	"INTERVAL":       6,
	"STRING":         7,
	"BYTES":          8,
	"TIMESTAMPTZ":    9,
	"COLLATEDSTRING": 10,
	"OID":            12,
	"UNKNOWN":        13,
	"UUID":           14,
	"ARRAY":          15,
	"INET":           16,
	"TIME":           17,
	"JSON":           18,
	"TUPLE":          20,
	"BIT":            21,
	"ANY":            100,
}

func (x SemanticType) Enum() *SemanticType {
	p := new(SemanticType)
	*p = x
	return p
}
func (x SemanticType) String() string {
	return proto.EnumName(SemanticType_name, int32(x))
}
func (x *SemanticType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SemanticType_value, data, "SemanticType")
	if err != nil {
		return err
	}
	*x = SemanticType(value)
	return nil
}
func (SemanticType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_590ae0e7819e34f1, []int{0}
}

// InternalType is the protobuf encoding for SQL types. It is always wrapped by
// a T struct, and should never be used directly by outside packages. See the
// comment header for the T struct for more details.
type InternalType struct {
	// SemanticType specifies a group of types that are compatible with one
	// another. See the header for the T.SemanticType method for more details.
	SemanticType SemanticType `protobuf:"varint,1,opt,name=semantic_type,json=semanticType,enum=cockroach.sql.sem.types.SemanticType" json:"semantic_type"`
	// Width is the size or scale of the type, such as number of bits or
	// characters. See the T.Width method for more details.
	Width int32 `protobuf:"varint,2,opt,name=width" json:"width"`
	// Precision is the accuracy of the data type. See the T.Precision method for
	// more details. This field was also by FLOAT pre-2.1 (this was incorrect.)
	Precision int32 `protobuf:"varint,3,opt,name=precision" json:"precision"`
	// ArrayDimensions is deprecated in 19.2, since it was never used. It
	// previously contained the length of each dimension in the array. A
	// dimension of -1 meant that no bound was specified for that dimension. If
	// arrayDimensions was nil, then the array had one unbounded dimension.
	ArrayDimensions []int32 `protobuf:"varint,4,rep,name=array_dimensions,json=arrayDimensions" json:"array_dimensions,omitempty"`
	// Locale identifies a specific geographical, political, or cultural region that
	// impacts various character-based operations such as sorting, pattern matching,
	// and builtin functions like lower and upper. See the T.Locale method for
	// more details.
	Locale *string `protobuf:"bytes,5,opt,name=locale" json:"locale,omitempty"`
	// VisibleType is deprecated in 19.2, since it is now superseded by the Oid
	// field. It previously contained an alias for any types where our internal
	// representation is different than the user specification. Examples are INT4,
	// FLOAT4, etc. Mostly for Postgres compatibility.
	VisibleType int32 `protobuf:"varint,6,opt,name=visible_type,json=visibleType" json:"visible_type"`
	// ArrayElemType is deprecated in 19.2, since it is now superseded by the
	// ArrayContents field. It previously contained the semantic type of array
	// elements. The other array fields (width/precision/locale/etc) were used
	// to store the other attributes of the array's element type.
	ArrayElemType *SemanticType `protobuf:"varint,7,opt,name=array_elem_type,json=arrayElemType,enum=cockroach.sql.sem.types.SemanticType" json:"array_elem_type,omitempty"`
	// TupleContents returns a slice containing the type of each tuple field. This
	// is nil for non-TUPLE types.
	TupleContents []T `protobuf:"bytes,8,rep,name=tuple_contents,json=tupleContents,customtype=T" json:"tuple_contents"`
	// TupleLabels returns a slice containing the labels of each tuple field. This
	// is nil for non-TUPLE types, or if the TUPLE type does not specify labels.
	TupleLabels []string `protobuf:"bytes,9,rep,name=tuple_labels,json=tupleLabels" json:"tuple_labels,omitempty"`
	// Oid returns the type's Postgres Object ID. See the header for the T.Oid
	// method for more details.
	Oid github_com_lib_pq_oid.Oid `protobuf:"varint,10,opt,name=oid,customtype=github.com/lib/pq/oid.Oid" json:"oid"`
	// ArrayContents returns the type of array elements. This is nil for non-ARRAY
	// types.
	ArrayContents        *T       `protobuf:"bytes,11,opt,name=array_contents,json=arrayContents,customtype=T" json:"array_contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalType) Reset()         { *m = InternalType{} }
func (m *InternalType) String() string { return proto.CompactTextString(m) }
func (*InternalType) ProtoMessage()    {}
func (*InternalType) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_590ae0e7819e34f1, []int{0}
}
func (m *InternalType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InternalType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *InternalType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalType.Merge(dst, src)
}
func (m *InternalType) XXX_Size() int {
	return m.Size()
}
func (m *InternalType) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalType.DiscardUnknown(m)
}

var xxx_messageInfo_InternalType proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InternalType)(nil), "cockroach.sql.sem.types.InternalType")
	proto.RegisterEnum("cockroach.sql.sem.types.SemanticType", SemanticType_name, SemanticType_value)
}
func (m *InternalType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.SemanticType))
	dAtA[i] = 0x10
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.Width))
	dAtA[i] = 0x18
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.Precision))
	if len(m.ArrayDimensions) > 0 {
		for _, num := range m.ArrayDimensions {
			dAtA[i] = 0x20
			i++
			i = encodeVarintTypes(dAtA, i, uint64(num))
		}
	}
	if m.Locale != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(*m.Locale)))
		i += copy(dAtA[i:], *m.Locale)
	}
	dAtA[i] = 0x30
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.VisibleType))
	if m.ArrayElemType != nil {
		dAtA[i] = 0x38
		i++
		i = encodeVarintTypes(dAtA, i, uint64(*m.ArrayElemType))
	}
	if len(m.TupleContents) > 0 {
		for _, msg := range m.TupleContents {
			dAtA[i] = 0x42
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.TupleLabels) > 0 {
		for _, s := range m.TupleLabels {
			dAtA[i] = 0x4a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	dAtA[i] = 0x50
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.Oid))
	if m.ArrayContents != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.ArrayContents.Size()))
		n1, err := m.ArrayContents.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InternalType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovTypes(uint64(m.SemanticType))
	n += 1 + sovTypes(uint64(m.Width))
	n += 1 + sovTypes(uint64(m.Precision))
	if len(m.ArrayDimensions) > 0 {
		for _, e := range m.ArrayDimensions {
			n += 1 + sovTypes(uint64(e))
		}
	}
	if m.Locale != nil {
		l = len(*m.Locale)
		n += 1 + l + sovTypes(uint64(l))
	}
	n += 1 + sovTypes(uint64(m.VisibleType))
	if m.ArrayElemType != nil {
		n += 1 + sovTypes(uint64(*m.ArrayElemType))
	}
	if len(m.TupleContents) > 0 {
		for _, e := range m.TupleContents {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.TupleLabels) > 0 {
		for _, s := range m.TupleLabels {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	n += 1 + sovTypes(uint64(m.Oid))
	if m.ArrayContents != nil {
		l = m.ArrayContents.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InternalType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: InternalType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SemanticType", wireType)
			}
			m.SemanticType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SemanticType |= (SemanticType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				m.ArrayDimensions = append(m.ArrayDimensions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
					return ErrInvalidLengthTypes
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
				if elementCount != 0 && len(m.ArrayDimensions) == 0 {
					m.ArrayDimensions = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
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
					m.ArrayDimensions = append(m.ArrayDimensions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrayDimensions", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locale", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Locale = &s
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VisibleType", wireType)
			}
			m.VisibleType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VisibleType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrayElemType", wireType)
			}
			var v SemanticType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (SemanticType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ArrayElemType = &v
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TupleContents", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v T
			m.TupleContents = append(m.TupleContents, v)
			if err := m.TupleContents[len(m.TupleContents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TupleLabels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TupleLabels = append(m.TupleLabels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oid", wireType)
			}
			m.Oid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Oid |= (github_com_lib_pq_oid.Oid(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrayContents", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v T
			m.ArrayContents = &v
			if err := m.ArrayContents.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
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
				next, err := skipTypes(dAtA[start:])
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
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sql/sem/types/types.proto", fileDescriptor_types_590ae0e7819e34f1) }

var fileDescriptor_types_590ae0e7819e34f1 = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xce, 0xd4, 0x76, 0x62, 0x4f, 0x9c, 0x74, 0xee, 0xb4, 0xf7, 0x5e, 0xb7, 0x8b, 0xc4, 0x44,
	0x42, 0x18, 0x84, 0x1c, 0xc4, 0x92, 0x9d, 0xdd, 0x18, 0x64, 0x70, 0xec, 0xca, 0x99, 0x80, 0xda,
	0x4d, 0x94, 0xd8, 0xa3, 0xd6, 0xc2, 0x3f, 0x69, 0xec, 0x82, 0xfa, 0x06, 0x2c, 0x59, 0xb3, 0x85,
	0x87, 0x49, 0x77, 0x2c, 0x11, 0x8b, 0x0a, 0xc2, 0x8b, 0xa0, 0xb1, 0x4d, 0xc9, 0x86, 0x05, 0x1b,
	0xeb, 0xf8, 0xfb, 0x3b, 0xdf, 0x91, 0x06, 0x1e, 0xe4, 0x17, 0xf1, 0x30, 0xa7, 0xc9, 0xb0, 0xb8,
	0x5a, 0xd2, 0xbc, 0xfa, 0xea, 0xcb, 0x55, 0x56, 0x64, 0xf8, 0xff, 0x20, 0x0b, 0x5e, 0xaf, 0xb2,
	0x79, 0x70, 0xae, 0xe7, 0x17, 0xb1, 0x9e, 0xd3, 0x44, 0x2f, 0xe9, 0xc3, 0xfd, 0xb3, 0xec, 0x2c,
	0x2b, 0x35, 0x43, 0x36, 0x55, 0xf2, 0xc1, 0x07, 0x1e, 0xca, 0x76, 0x5a, 0xd0, 0x55, 0x3a, 0x8f,
	0xc9, 0xd5, 0x92, 0xe2, 0x63, 0xd8, 0xc9, 0x69, 0x32, 0x4f, 0x8b, 0x28, 0x98, 0x31, 0xa3, 0x02,
	0x54, 0xa0, 0x75, 0x1f, 0xdf, 0xd5, 0xff, 0x90, 0xab, 0x4f, 0x6a, 0x35, 0x73, 0x9b, 0xfc, 0xfa,
	0xa6, 0xdf, 0xf0, 0xe5, 0x7c, 0x0b, 0xc3, 0x87, 0x50, 0x78, 0x1b, 0x85, 0xc5, 0xb9, 0xb2, 0xa3,
	0x02, 0x4d, 0xa8, 0x25, 0x15, 0x84, 0x07, 0x50, 0x5a, 0xae, 0x68, 0x10, 0xe5, 0x51, 0x96, 0x2a,
	0xdc, 0x16, 0xff, 0x1b, 0xc6, 0xf7, 0x21, 0x9a, 0xaf, 0x56, 0xf3, 0xab, 0x59, 0x18, 0x25, 0x34,
	0x65, 0x50, 0xae, 0xf0, 0x2a, 0xa7, 0x09, 0xfe, 0x6e, 0x89, 0x8f, 0x6e, 0x61, 0xfc, 0x1f, 0x6c,
	0xc6, 0x59, 0x30, 0x8f, 0xa9, 0x22, 0xa8, 0x40, 0x93, 0xfc, 0xfa, 0x0f, 0xdf, 0x83, 0xf2, 0x9b,
	0x28, 0x8f, 0x16, 0x31, 0xad, 0x6e, 0x6a, 0x6e, 0x6d, 0x6a, 0xd7, 0x4c, 0xd9, 0x75, 0x0c, 0xab,
	0xcc, 0x19, 0x8d, 0x69, 0x52, 0x69, 0x5b, 0x7f, 0x71, 0xbf, 0xdf, 0x29, 0xdd, 0x56, 0x4c, 0x93,
	0x32, 0xee, 0x11, 0xec, 0x16, 0x97, 0xcb, 0x98, 0xce, 0x82, 0x2c, 0x2d, 0x68, 0x5a, 0xe4, 0x8a,
	0xa8, 0x72, 0x9a, 0x6c, 0x4a, 0x6c, 0xf3, 0xd7, 0x9b, 0x3e, 0x20, 0x7e, 0xa7, 0x14, 0x1c, 0xd5,
	0x3c, 0xbe, 0x03, 0xe5, 0xca, 0x11, 0xcf, 0x17, 0x34, 0xce, 0x15, 0x49, 0xe5, 0x34, 0xc9, 0x6f,
	0x97, 0x98, 0x53, 0x42, 0xf8, 0x09, 0xe4, 0xb2, 0x28, 0x54, 0xa0, 0x0a, 0xb4, 0x8e, 0xa9, 0xd5,
	0x49, 0x07, 0x67, 0x51, 0x71, 0x7e, 0xb9, 0xd0, 0x83, 0x2c, 0x19, 0xc6, 0xd1, 0x62, 0xb8, 0xbc,
	0x18, 0x66, 0x51, 0xa8, 0x7b, 0x51, 0xb8, 0xb9, 0xe9, 0x73, 0x5e, 0x14, 0xfa, 0xcc, 0x84, 0x1f,
	0xc2, 0x6e, 0x75, 0xdf, 0x6d, 0xa1, 0xb6, 0x0a, 0x34, 0xd9, 0x14, 0xea, 0x32, 0x25, 0xf9, 0xab,
	0xcc, 0x83, 0x4f, 0x3b, 0x50, 0xde, 0x3e, 0x0f, 0x8b, 0x90, 0x37, 0x3d, 0xcf, 0x41, 0x0d, 0xdc,
	0x82, 0x9c, 0xed, 0x12, 0x04, 0xb0, 0x04, 0x85, 0xa7, 0x8e, 0x67, 0x10, 0xb4, 0x83, 0xdb, 0xb0,
	0x35, 0xb2, 0x8e, 0xec, 0xb1, 0xe1, 0x20, 0x8e, 0x49, 0x47, 0x06, 0xb1, 0x10, 0x8f, 0x3b, 0x50,
	0x22, 0xf6, 0xd8, 0x9a, 0x10, 0x63, 0x7c, 0x8c, 0x04, 0x2c, 0x43, 0xd1, 0x76, 0x89, 0xe5, 0xbf,
	0x34, 0x1c, 0xd4, 0xc4, 0x10, 0x36, 0x27, 0xc4, 0xb7, 0xdd, 0x67, 0xa8, 0xc5, 0xa2, 0xcc, 0x13,
	0x62, 0x4d, 0x90, 0x88, 0x77, 0x61, 0xfb, 0xd6, 0x43, 0x4e, 0x91, 0x84, 0x31, 0xec, 0x1e, 0x79,
	0x8e, 0x63, 0x10, 0x6b, 0x54, 0xeb, 0x21, 0xeb, 0xe0, 0xd9, 0x23, 0x24, 0xb3, 0xc5, 0x53, 0xf7,
	0x85, 0xeb, 0xbd, 0x72, 0x51, 0x87, 0x2d, 0x9e, 0x4e, 0xed, 0x11, 0xea, 0xb2, 0x3c, 0xc3, 0xf7,
	0x8d, 0x13, 0xb4, 0xcb, 0x40, 0xdb, 0xb5, 0x08, 0x42, 0x6c, 0x62, 0xc9, 0xe8, 0x1f, 0x36, 0x3d,
	0x9f, 0x78, 0x2e, 0xc2, 0x4c, 0x48, 0xa6, 0xc7, 0x8e, 0x85, 0xf6, 0x59, 0xa6, 0x69, 0x13, 0xf4,
	0x2f, 0x1b, 0x0c, 0xf7, 0x04, 0x85, 0x87, 0xfc, 0xbb, 0x8f, 0xbd, 0xc6, 0x80, 0x17, 0xdb, 0xa8,
	0x3d, 0xe0, 0xc5, 0x3d, 0xb4, 0x37, 0x68, 0x8a, 0x6b, 0x80, 0xd6, 0x60, 0xd0, 0x14, 0xaf, 0x01,
	0xba, 0x06, 0x66, 0x7f, 0xfd, 0xbd, 0xd7, 0x58, 0x6f, 0x7a, 0xe0, 0xf3, 0xa6, 0x07, 0xbe, 0x6c,
	0x7a, 0xe0, 0xdb, 0xa6, 0x07, 0xde, 0xff, 0xe8, 0x35, 0x4e, 0x85, 0xf2, 0x89, 0xfc, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0x52, 0x20, 0x31, 0xac, 0xaf, 0x03, 0x00, 0x00,
}
