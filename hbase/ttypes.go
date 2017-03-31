package hbase

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

//Specify type of delete:
// - DELETE_COLUMN means exactly one version will be removed,
// - DELETE_COLUMNS means previous versions will also be removed.
type TDeleteType int64

const (
	TDeleteType_DELETE_COLUMN  TDeleteType = 0
	TDeleteType_DELETE_COLUMNS TDeleteType = 1
)

func (p TDeleteType) String() string {
	switch p {
	case TDeleteType_DELETE_COLUMN:
		return "DELETE_COLUMN"
	case TDeleteType_DELETE_COLUMNS:
		return "DELETE_COLUMNS"
	}
	return "<UNSET>"
}

func TDeleteTypeFromString(s string) (TDeleteType, error) {
	switch s {
	case "DELETE_COLUMN":
		return TDeleteType_DELETE_COLUMN, nil
	case "DELETE_COLUMNS":
		return TDeleteType_DELETE_COLUMNS, nil
	}
	return TDeleteType(0), fmt.Errorf("not a valid TDeleteType string")
}

func TDeleteTypePtr(v TDeleteType) *TDeleteType { return &v }

func (p TDeleteType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TDeleteType) UnmarshalText(text []byte) error {
	q, err := TDeleteTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

//Specify Durability:
// - SKIP_WAL means do not write the Mutation to the WAL.
// - ASYNC_WAL means write the Mutation to the WAL asynchronously,
// - SYNC_WAL means write the Mutation to the WAL synchronously,
// - FSYNC_WAL means Write the Mutation to the WAL synchronously and force the entries to disk.
type TDurability int64

const (
	TDurability_SKIP_WAL  TDurability = 1
	TDurability_ASYNC_WAL TDurability = 2
	TDurability_SYNC_WAL  TDurability = 3
	TDurability_FSYNC_WAL TDurability = 4
)

func (p TDurability) String() string {
	switch p {
	case TDurability_SKIP_WAL:
		return "SKIP_WAL"
	case TDurability_ASYNC_WAL:
		return "ASYNC_WAL"
	case TDurability_SYNC_WAL:
		return "SYNC_WAL"
	case TDurability_FSYNC_WAL:
		return "FSYNC_WAL"
	}
	return "<UNSET>"
}

func TDurabilityFromString(s string) (TDurability, error) {
	switch s {
	case "SKIP_WAL":
		return TDurability_SKIP_WAL, nil
	case "ASYNC_WAL":
		return TDurability_ASYNC_WAL, nil
	case "SYNC_WAL":
		return TDurability_SYNC_WAL, nil
	case "FSYNC_WAL":
		return TDurability_FSYNC_WAL, nil
	}
	return TDurability(0), fmt.Errorf("not a valid TDurability string")
}

func TDurabilityPtr(v TDurability) *TDurability { return &v }

func (p TDurability) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TDurability) UnmarshalText(text []byte) error {
	q, err := TDurabilityFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

//Thrift wrapper around
//org.apache.hadoop.hbase.filter.CompareFilter$CompareOp.
type TCompareOp int64

const (
	TCompareOp_LESS             TCompareOp = 0
	TCompareOp_LESS_OR_EQUAL    TCompareOp = 1
	TCompareOp_EQUAL            TCompareOp = 2
	TCompareOp_NOT_EQUAL        TCompareOp = 3
	TCompareOp_GREATER_OR_EQUAL TCompareOp = 4
	TCompareOp_GREATER          TCompareOp = 5
	TCompareOp_NO_OP            TCompareOp = 6
)

func (p TCompareOp) String() string {
	switch p {
	case TCompareOp_LESS:
		return "LESS"
	case TCompareOp_LESS_OR_EQUAL:
		return "LESS_OR_EQUAL"
	case TCompareOp_EQUAL:
		return "EQUAL"
	case TCompareOp_NOT_EQUAL:
		return "NOT_EQUAL"
	case TCompareOp_GREATER_OR_EQUAL:
		return "GREATER_OR_EQUAL"
	case TCompareOp_GREATER:
		return "GREATER"
	case TCompareOp_NO_OP:
		return "NO_OP"
	}
	return "<UNSET>"
}

func TCompareOpFromString(s string) (TCompareOp, error) {
	switch s {
	case "LESS":
		return TCompareOp_LESS, nil
	case "LESS_OR_EQUAL":
		return TCompareOp_LESS_OR_EQUAL, nil
	case "EQUAL":
		return TCompareOp_EQUAL, nil
	case "NOT_EQUAL":
		return TCompareOp_NOT_EQUAL, nil
	case "GREATER_OR_EQUAL":
		return TCompareOp_GREATER_OR_EQUAL, nil
	case "GREATER":
		return TCompareOp_GREATER, nil
	case "NO_OP":
		return TCompareOp_NO_OP, nil
	}
	return TCompareOp(0), fmt.Errorf("not a valid TCompareOp string")
}

func TCompareOpPtr(v TCompareOp) *TCompareOp { return &v }

func (p TCompareOp) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TCompareOp) UnmarshalText(text []byte) error {
	q, err := TCompareOpFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - MinStamp
//  - MaxStamp
type TTimeRange struct {
	MinStamp int64 `thrift:"minStamp,1,required" json:"minStamp"`
	MaxStamp int64 `thrift:"maxStamp,2,required" json:"maxStamp"`
}

func NewTTimeRange() *TTimeRange {
	return &TTimeRange{}
}

func (p *TTimeRange) GetMinStamp() int64 {
	return p.MinStamp
}

func (p *TTimeRange) GetMaxStamp() int64 {
	return p.MaxStamp
}
func (p *TTimeRange) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetMinStamp bool = false
	var issetMaxStamp bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetMinStamp = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetMaxStamp = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetMinStamp {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field MinStamp is not set"))
	}
	if !issetMaxStamp {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field MaxStamp is not set"))
	}
	return nil
}

func (p *TTimeRange) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.MinStamp = v
	}
	return nil
}

func (p *TTimeRange) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.MaxStamp = v
	}
	return nil
}

func (p *TTimeRange) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TTimeRange"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TTimeRange) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("minStamp", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:minStamp: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.MinStamp)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.minStamp (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:minStamp: ", p), err)
	}
	return err
}

func (p *TTimeRange) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("maxStamp", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:maxStamp: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.MaxStamp)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.maxStamp (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:maxStamp: ", p), err)
	}
	return err
}

func (p *TTimeRange) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TTimeRange(%+v)", *p)
}

// Addresses a single cell or multiple cells
// in a HBase table by column family and optionally
// a column qualifier and timestamp
//
// Attributes:
//  - Family
//  - Qualifier
//  - Timestamp
type TColumn struct {
	Family    []byte `thrift:"family,1,required" json:"family"`
	Qualifier []byte `thrift:"qualifier,2" json:"qualifier,omitempty"`
	Timestamp *int64 `thrift:"timestamp,3" json:"timestamp,omitempty"`
}

func NewTColumn() *TColumn {
	return &TColumn{}
}

func (p *TColumn) GetFamily() []byte {
	return p.Family
}

var TColumn_Qualifier_DEFAULT []byte

func (p *TColumn) GetQualifier() []byte {
	return p.Qualifier
}

var TColumn_Timestamp_DEFAULT int64

func (p *TColumn) GetTimestamp() int64 {
	if !p.IsSetTimestamp() {
		return TColumn_Timestamp_DEFAULT
	}
	return *p.Timestamp
}
func (p *TColumn) IsSetQualifier() bool {
	return p.Qualifier != nil
}

func (p *TColumn) IsSetTimestamp() bool {
	return p.Timestamp != nil
}

func (p *TColumn) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetFamily bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetFamily = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetFamily {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Family is not set"))
	}
	return nil
}

func (p *TColumn) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Family = v
	}
	return nil
}

func (p *TColumn) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Qualifier = v
	}
	return nil
}

func (p *TColumn) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Timestamp = &v
	}
	return nil
}

func (p *TColumn) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TColumn"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TColumn) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("family", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:family: ", p), err)
	}
	if err := oprot.WriteBinary(p.Family); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.family (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:family: ", p), err)
	}
	return err
}

func (p *TColumn) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetQualifier() {
		if err := oprot.WriteFieldBegin("qualifier", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:qualifier: ", p), err)
		}
		if err := oprot.WriteBinary(p.Qualifier); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.qualifier (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:qualifier: ", p), err)
		}
	}
	return err
}

func (p *TColumn) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimestamp() {
		if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timestamp: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timestamp (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timestamp: ", p), err)
		}
	}
	return err
}

func (p *TColumn) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TColumn(%+v)", *p)
}

// Represents a single cell and its value.
//
// Attributes:
//  - Family
//  - Qualifier
//  - Value
//  - Timestamp
//  - Tags
type TColumnValue struct {
	Family    []byte `thrift:"family,1,required" json:"family"`
	Qualifier []byte `thrift:"qualifier,2,required" json:"qualifier"`
	Value     []byte `thrift:"value,3,required" json:"value"`
	Timestamp *int64 `thrift:"timestamp,4" json:"timestamp,omitempty"`
	Tags      []byte `thrift:"tags,5" json:"tags,omitempty"`
}

func NewTColumnValue() *TColumnValue {
	return &TColumnValue{}
}

func (p *TColumnValue) GetFamily() []byte {
	return p.Family
}

func (p *TColumnValue) GetQualifier() []byte {
	return p.Qualifier
}

func (p *TColumnValue) GetValue() []byte {
	return p.Value
}

var TColumnValue_Timestamp_DEFAULT int64

func (p *TColumnValue) GetTimestamp() int64 {
	if !p.IsSetTimestamp() {
		return TColumnValue_Timestamp_DEFAULT
	}
	return *p.Timestamp
}

var TColumnValue_Tags_DEFAULT []byte

func (p *TColumnValue) GetTags() []byte {
	return p.Tags
}
func (p *TColumnValue) IsSetTimestamp() bool {
	return p.Timestamp != nil
}

func (p *TColumnValue) IsSetTags() bool {
	return p.Tags != nil
}

func (p *TColumnValue) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetFamily bool = false
	var issetQualifier bool = false
	var issetValue bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetFamily = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetQualifier = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetValue = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetFamily {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Family is not set"))
	}
	if !issetQualifier {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Qualifier is not set"))
	}
	if !issetValue {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Value is not set"))
	}
	return nil
}

func (p *TColumnValue) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Family = v
	}
	return nil
}

func (p *TColumnValue) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Qualifier = v
	}
	return nil
}

func (p *TColumnValue) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *TColumnValue) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Timestamp = &v
	}
	return nil
}

func (p *TColumnValue) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Tags = v
	}
	return nil
}

func (p *TColumnValue) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TColumnValue"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TColumnValue) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("family", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:family: ", p), err)
	}
	if err := oprot.WriteBinary(p.Family); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.family (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:family: ", p), err)
	}
	return err
}

func (p *TColumnValue) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("qualifier", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:qualifier: ", p), err)
	}
	if err := oprot.WriteBinary(p.Qualifier); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.qualifier (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:qualifier: ", p), err)
	}
	return err
}

func (p *TColumnValue) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:value: ", p), err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:value: ", p), err)
	}
	return err
}

func (p *TColumnValue) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimestamp() {
		if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:timestamp: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timestamp (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:timestamp: ", p), err)
		}
	}
	return err
}

func (p *TColumnValue) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetTags() {
		if err := oprot.WriteFieldBegin("tags", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:tags: ", p), err)
		}
		if err := oprot.WriteBinary(p.Tags); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.tags (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:tags: ", p), err)
		}
	}
	return err
}

func (p *TColumnValue) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TColumnValue(%+v)", *p)
}

// Represents a single cell and the amount to increment it by
//
// Attributes:
//  - Family
//  - Qualifier
//  - Amount
type TColumnIncrement struct {
	Family    []byte `thrift:"family,1,required" json:"family"`
	Qualifier []byte `thrift:"qualifier,2,required" json:"qualifier"`
	Amount    int64  `thrift:"amount,3" json:"amount,omitempty"`
}

func NewTColumnIncrement() *TColumnIncrement {
	return &TColumnIncrement{
		Amount: 1,
	}
}

func (p *TColumnIncrement) GetFamily() []byte {
	return p.Family
}

func (p *TColumnIncrement) GetQualifier() []byte {
	return p.Qualifier
}

var TColumnIncrement_Amount_DEFAULT int64 = 1

func (p *TColumnIncrement) GetAmount() int64 {
	return p.Amount
}
func (p *TColumnIncrement) IsSetAmount() bool {
	return p.Amount != TColumnIncrement_Amount_DEFAULT
}

func (p *TColumnIncrement) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetFamily bool = false
	var issetQualifier bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetFamily = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetQualifier = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetFamily {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Family is not set"))
	}
	if !issetQualifier {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Qualifier is not set"))
	}
	return nil
}

func (p *TColumnIncrement) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Family = v
	}
	return nil
}

func (p *TColumnIncrement) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Qualifier = v
	}
	return nil
}

func (p *TColumnIncrement) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Amount = v
	}
	return nil
}

func (p *TColumnIncrement) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TColumnIncrement"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TColumnIncrement) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("family", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:family: ", p), err)
	}
	if err := oprot.WriteBinary(p.Family); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.family (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:family: ", p), err)
	}
	return err
}

func (p *TColumnIncrement) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("qualifier", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:qualifier: ", p), err)
	}
	if err := oprot.WriteBinary(p.Qualifier); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.qualifier (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:qualifier: ", p), err)
	}
	return err
}

func (p *TColumnIncrement) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetAmount() {
		if err := oprot.WriteFieldBegin("amount", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:amount: ", p), err)
		}
		if err := oprot.WriteI64(int64(p.Amount)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.amount (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:amount: ", p), err)
		}
	}
	return err
}

func (p *TColumnIncrement) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TColumnIncrement(%+v)", *p)
}

// if no Result is found, row and columnValues will not be set.
//
// Attributes:
//  - Row
//  - ColumnValues
type TResult_ struct {
	Row          []byte          `thrift:"row,1" json:"row,omitempty"`
	ColumnValues []*TColumnValue `thrift:"columnValues,2,required" json:"columnValues"`
}

func NewTResult_() *TResult_ {
	return &TResult_{}
}

var TResult__Row_DEFAULT []byte

func (p *TResult_) GetRow() []byte {
	return p.Row
}

func (p *TResult_) GetColumnValues() []*TColumnValue {
	return p.ColumnValues
}
func (p *TResult_) IsSetRow() bool {
	return p.Row != nil
}

func (p *TResult_) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetColumnValues bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetColumnValues = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetColumnValues {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ColumnValues is not set"))
	}
	return nil
}

func (p *TResult_) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TResult_) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumnValue, 0, size)
	p.ColumnValues = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &TColumnValue{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.ColumnValues = append(p.ColumnValues, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TResult_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetRow() {
		if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
		}
		if err := oprot.WriteBinary(p.Row); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
		}
	}
	return err
}

func (p *TResult_) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("columnValues", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:columnValues: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.ColumnValues)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.ColumnValues {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:columnValues: ", p), err)
	}
	return err
}

func (p *TResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TResult_(%+v)", *p)
}

// Attributes:
//  - Labels
type TAuthorization struct {
	Labels []string `thrift:"labels,1" json:"labels,omitempty"`
}

func NewTAuthorization() *TAuthorization {
	return &TAuthorization{}
}

var TAuthorization_Labels_DEFAULT []string

func (p *TAuthorization) GetLabels() []string {
	return p.Labels
}
func (p *TAuthorization) IsSetLabels() bool {
	return p.Labels != nil
}

func (p *TAuthorization) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TAuthorization) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Labels = tSlice
	for i := 0; i < size; i++ {
		var _elem1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.Labels = append(p.Labels, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TAuthorization) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TAuthorization"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TAuthorization) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetLabels() {
		if err := oprot.WriteFieldBegin("labels", thrift.LIST, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:labels: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Labels)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Labels {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:labels: ", p), err)
		}
	}
	return err
}

func (p *TAuthorization) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TAuthorization(%+v)", *p)
}

// Attributes:
//  - Expression
type TCellVisibility struct {
	Expression *string `thrift:"expression,1" json:"expression,omitempty"`
}

func NewTCellVisibility() *TCellVisibility {
	return &TCellVisibility{}
}

var TCellVisibility_Expression_DEFAULT string

func (p *TCellVisibility) GetExpression() string {
	if !p.IsSetExpression() {
		return TCellVisibility_Expression_DEFAULT
	}
	return *p.Expression
}
func (p *TCellVisibility) IsSetExpression() bool {
	return p.Expression != nil
}

func (p *TCellVisibility) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCellVisibility) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Expression = &v
	}
	return nil
}

func (p *TCellVisibility) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TCellVisibility"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCellVisibility) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetExpression() {
		if err := oprot.WriteFieldBegin("expression", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:expression: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Expression)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.expression (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:expression: ", p), err)
		}
	}
	return err
}

func (p *TCellVisibility) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCellVisibility(%+v)", *p)
}

// Used to perform Get operations on a single row.
//
// The scope can be further narrowed down by specifying a list of
// columns or column families.
//
// To get everything for a row, instantiate a Get object with just the row to get.
// To further define the scope of what to get you can add a timestamp or time range
// with an optional maximum number of versions to return.
//
// If you specify a time range and a timestamp the range is ignored.
// Timestamps on TColumns are ignored.
//
// Attributes:
//  - Row
//  - Columns
//  - Timestamp
//  - TimeRange
//  - MaxVersions
//  - FilterString
//  - Attributes
//  - Authorizations
type TGet struct {
	Row            []byte            `thrift:"row,1,required" json:"row"`
	Columns        []*TColumn        `thrift:"columns,2" json:"columns,omitempty"`
	Timestamp      *int64            `thrift:"timestamp,3" json:"timestamp,omitempty"`
	TimeRange      *TTimeRange       `thrift:"timeRange,4" json:"timeRange,omitempty"`
	MaxVersions    *int32            `thrift:"maxVersions,5" json:"maxVersions,omitempty"`
	FilterString   []byte            `thrift:"filterString,6" json:"filterString,omitempty"`
	Attributes     map[string][]byte `thrift:"attributes,7" json:"attributes,omitempty"`
	Authorizations *TAuthorization   `thrift:"authorizations,8" json:"authorizations,omitempty"`
}

func NewTGet() *TGet {
	return &TGet{}
}

func (p *TGet) GetRow() []byte {
	return p.Row
}

var TGet_Columns_DEFAULT []*TColumn

func (p *TGet) GetColumns() []*TColumn {
	return p.Columns
}

var TGet_Timestamp_DEFAULT int64

func (p *TGet) GetTimestamp() int64 {
	if !p.IsSetTimestamp() {
		return TGet_Timestamp_DEFAULT
	}
	return *p.Timestamp
}

var TGet_TimeRange_DEFAULT *TTimeRange

func (p *TGet) GetTimeRange() *TTimeRange {
	if !p.IsSetTimeRange() {
		return TGet_TimeRange_DEFAULT
	}
	return p.TimeRange
}

var TGet_MaxVersions_DEFAULT int32

func (p *TGet) GetMaxVersions() int32 {
	if !p.IsSetMaxVersions() {
		return TGet_MaxVersions_DEFAULT
	}
	return *p.MaxVersions
}

var TGet_FilterString_DEFAULT []byte

func (p *TGet) GetFilterString() []byte {
	return p.FilterString
}

var TGet_Attributes_DEFAULT map[string][]byte

func (p *TGet) GetAttributes() map[string][]byte {
	return p.Attributes
}

var TGet_Authorizations_DEFAULT *TAuthorization

func (p *TGet) GetAuthorizations() *TAuthorization {
	if !p.IsSetAuthorizations() {
		return TGet_Authorizations_DEFAULT
	}
	return p.Authorizations
}
func (p *TGet) IsSetColumns() bool {
	return p.Columns != nil
}

func (p *TGet) IsSetTimestamp() bool {
	return p.Timestamp != nil
}

func (p *TGet) IsSetTimeRange() bool {
	return p.TimeRange != nil
}

func (p *TGet) IsSetMaxVersions() bool {
	return p.MaxVersions != nil
}

func (p *TGet) IsSetFilterString() bool {
	return p.FilterString != nil
}

func (p *TGet) IsSetAttributes() bool {
	return p.Attributes != nil
}

func (p *TGet) IsSetAuthorizations() bool {
	return p.Authorizations != nil
}

func (p *TGet) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRow bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRow = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	return nil
}

func (p *TGet) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TGet) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumn, 0, size)
	p.Columns = tSlice
	for i := 0; i < size; i++ {
		_elem2 := &TColumn{}
		if err := _elem2.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem2), err)
		}
		p.Columns = append(p.Columns, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TGet) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Timestamp = &v
	}
	return nil
}

func (p *TGet) readField4(iprot thrift.TProtocol) error {
	p.TimeRange = &TTimeRange{}
	if err := p.TimeRange.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TimeRange), err)
	}
	return nil
}

func (p *TGet) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.MaxVersions = &v
	}
	return nil
}

func (p *TGet) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.FilterString = v
	}
	return nil
}

func (p *TGet) readField7(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string][]byte, size)
	p.Attributes = tMap
	for i := 0; i < size; i++ {
		var _key3 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key3 = v
		}
		var _val4 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val4 = v
		}
		p.Attributes[_key3] = _val4
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TGet) readField8(iprot thrift.TProtocol) error {
	p.Authorizations = &TAuthorization{}
	if err := p.Authorizations.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Authorizations), err)
	}
	return nil
}

func (p *TGet) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TGet"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TGet) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *TGet) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetColumns() {
		if err := oprot.WriteFieldBegin("columns", thrift.LIST, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:columns: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Columns)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Columns {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:columns: ", p), err)
		}
	}
	return err
}

func (p *TGet) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimestamp() {
		if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timestamp: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timestamp (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timestamp: ", p), err)
		}
	}
	return err
}

func (p *TGet) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimeRange() {
		if err := oprot.WriteFieldBegin("timeRange", thrift.STRUCT, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:timeRange: ", p), err)
		}
		if err := p.TimeRange.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TimeRange), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:timeRange: ", p), err)
		}
	}
	return err
}

func (p *TGet) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetMaxVersions() {
		if err := oprot.WriteFieldBegin("maxVersions", thrift.I32, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:maxVersions: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.MaxVersions)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.maxVersions (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:maxVersions: ", p), err)
		}
	}
	return err
}

func (p *TGet) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetFilterString() {
		if err := oprot.WriteFieldBegin("filterString", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:filterString: ", p), err)
		}
		if err := oprot.WriteBinary(p.FilterString); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.filterString (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:filterString: ", p), err)
		}
	}
	return err
}

func (p *TGet) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetAttributes() {
		if err := oprot.WriteFieldBegin("attributes", thrift.MAP, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:attributes: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Attributes)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Attributes {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:attributes: ", p), err)
		}
	}
	return err
}

func (p *TGet) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetAuthorizations() {
		if err := oprot.WriteFieldBegin("authorizations", thrift.STRUCT, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:authorizations: ", p), err)
		}
		if err := p.Authorizations.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Authorizations), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:authorizations: ", p), err)
		}
	}
	return err
}

func (p *TGet) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TGet(%+v)", *p)
}

// Used to perform Put operations for a single row.
//
// Add column values to this object and they'll be added.
// You can provide a default timestamp if the column values
// don't have one. If you don't provide a default timestamp
// the current time is inserted.
//
// You can specify how this Put should be written to the write-ahead Log (WAL)
// by changing the durability. If you don't provide durability, it defaults to
// column family's default setting for durability.
//
// Attributes:
//  - Row
//  - ColumnValues
//  - Timestamp
//  - Attributes
//  - Durability
//  - CellVisibility
type TPut struct {
	Row          []byte          `thrift:"row,1,required" json:"row"`
	ColumnValues []*TColumnValue `thrift:"columnValues,2,required" json:"columnValues"`
	Timestamp    *int64          `thrift:"timestamp,3" json:"timestamp,omitempty"`
	// unused field # 4
	Attributes     map[string][]byte `thrift:"attributes,5" json:"attributes,omitempty"`
	Durability     *TDurability      `thrift:"durability,6" json:"durability,omitempty"`
	CellVisibility *TCellVisibility  `thrift:"cellVisibility,7" json:"cellVisibility,omitempty"`
}

func NewTPut() *TPut {
	return &TPut{}
}

func (p *TPut) GetRow() []byte {
	return p.Row
}

func (p *TPut) GetColumnValues() []*TColumnValue {
	return p.ColumnValues
}

var TPut_Timestamp_DEFAULT int64

func (p *TPut) GetTimestamp() int64 {
	if !p.IsSetTimestamp() {
		return TPut_Timestamp_DEFAULT
	}
	return *p.Timestamp
}

var TPut_Attributes_DEFAULT map[string][]byte

func (p *TPut) GetAttributes() map[string][]byte {
	return p.Attributes
}

var TPut_Durability_DEFAULT TDurability

func (p *TPut) GetDurability() TDurability {
	if !p.IsSetDurability() {
		return TPut_Durability_DEFAULT
	}
	return *p.Durability
}

var TPut_CellVisibility_DEFAULT *TCellVisibility

func (p *TPut) GetCellVisibility() *TCellVisibility {
	if !p.IsSetCellVisibility() {
		return TPut_CellVisibility_DEFAULT
	}
	return p.CellVisibility
}
func (p *TPut) IsSetTimestamp() bool {
	return p.Timestamp != nil
}

func (p *TPut) IsSetAttributes() bool {
	return p.Attributes != nil
}

func (p *TPut) IsSetDurability() bool {
	return p.Durability != nil
}

func (p *TPut) IsSetCellVisibility() bool {
	return p.CellVisibility != nil
}

func (p *TPut) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRow bool = false
	var issetColumnValues bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRow = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetColumnValues = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	if !issetColumnValues {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ColumnValues is not set"))
	}
	return nil
}

func (p *TPut) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TPut) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumnValue, 0, size)
	p.ColumnValues = tSlice
	for i := 0; i < size; i++ {
		_elem5 := &TColumnValue{}
		if err := _elem5.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem5), err)
		}
		p.ColumnValues = append(p.ColumnValues, _elem5)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TPut) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Timestamp = &v
	}
	return nil
}

func (p *TPut) readField5(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string][]byte, size)
	p.Attributes = tMap
	for i := 0; i < size; i++ {
		var _key6 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key6 = v
		}
		var _val7 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val7 = v
		}
		p.Attributes[_key6] = _val7
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TPut) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		temp := TDurability(v)
		p.Durability = &temp
	}
	return nil
}

func (p *TPut) readField7(iprot thrift.TProtocol) error {
	p.CellVisibility = &TCellVisibility{}
	if err := p.CellVisibility.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.CellVisibility), err)
	}
	return nil
}

func (p *TPut) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TPut"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TPut) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *TPut) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("columnValues", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:columnValues: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.ColumnValues)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.ColumnValues {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:columnValues: ", p), err)
	}
	return err
}

func (p *TPut) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimestamp() {
		if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timestamp: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timestamp (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timestamp: ", p), err)
		}
	}
	return err
}

func (p *TPut) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetAttributes() {
		if err := oprot.WriteFieldBegin("attributes", thrift.MAP, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:attributes: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Attributes)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Attributes {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:attributes: ", p), err)
		}
	}
	return err
}

func (p *TPut) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetDurability() {
		if err := oprot.WriteFieldBegin("durability", thrift.I32, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:durability: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Durability)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.durability (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:durability: ", p), err)
		}
	}
	return err
}

func (p *TPut) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetCellVisibility() {
		if err := oprot.WriteFieldBegin("cellVisibility", thrift.STRUCT, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:cellVisibility: ", p), err)
		}
		if err := p.CellVisibility.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.CellVisibility), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:cellVisibility: ", p), err)
		}
	}
	return err
}

func (p *TPut) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TPut(%+v)", *p)
}

// Used to perform Delete operations on a single row.
//
// The scope can be further narrowed down by specifying a list of
// columns or column families as TColumns.
//
// Specifying only a family in a TColumn will delete the whole family.
// If a timestamp is specified all versions with a timestamp less than
// or equal to this will be deleted. If no timestamp is specified the
// current time will be used.
//
// Specifying a family and a column qualifier in a TColumn will delete only
// this qualifier. If a timestamp is specified only versions equal
// to this timestamp will be deleted. If no timestamp is specified the
// most recent version will be deleted.  To delete all previous versions,
// specify the DELETE_COLUMNS TDeleteType.
//
// The top level timestamp is only used if a complete row should be deleted
// (i.e. no columns are passed) and if it is specified it works the same way
// as if you had added a TColumn for every column family and this timestamp
// (i.e. all versions older than or equal in all column families will be deleted)
//
// You can specify how this Delete should be written to the write-ahead Log (WAL)
// by changing the durability. If you don't provide durability, it defaults to
// column family's default setting for durability.
//
// Attributes:
//  - Row
//  - Columns
//  - Timestamp
//  - DeleteType
//  - Attributes
//  - Durability
type TDelete struct {
	Row        []byte      `thrift:"row,1,required" json:"row"`
	Columns    []*TColumn  `thrift:"columns,2" json:"columns,omitempty"`
	Timestamp  *int64      `thrift:"timestamp,3" json:"timestamp,omitempty"`
	DeleteType TDeleteType `thrift:"deleteType,4" json:"deleteType,omitempty"`
	// unused field # 5
	Attributes map[string][]byte `thrift:"attributes,6" json:"attributes,omitempty"`
	Durability *TDurability      `thrift:"durability,7" json:"durability,omitempty"`
}

func NewTDelete() *TDelete {
	return &TDelete{
		DeleteType: 1,
	}
}

func (p *TDelete) GetRow() []byte {
	return p.Row
}

var TDelete_Columns_DEFAULT []*TColumn

func (p *TDelete) GetColumns() []*TColumn {
	return p.Columns
}

var TDelete_Timestamp_DEFAULT int64

func (p *TDelete) GetTimestamp() int64 {
	if !p.IsSetTimestamp() {
		return TDelete_Timestamp_DEFAULT
	}
	return *p.Timestamp
}

var TDelete_DeleteType_DEFAULT TDeleteType = 1

func (p *TDelete) GetDeleteType() TDeleteType {
	return p.DeleteType
}

var TDelete_Attributes_DEFAULT map[string][]byte

func (p *TDelete) GetAttributes() map[string][]byte {
	return p.Attributes
}

var TDelete_Durability_DEFAULT TDurability

func (p *TDelete) GetDurability() TDurability {
	if !p.IsSetDurability() {
		return TDelete_Durability_DEFAULT
	}
	return *p.Durability
}
func (p *TDelete) IsSetColumns() bool {
	return p.Columns != nil
}

func (p *TDelete) IsSetTimestamp() bool {
	return p.Timestamp != nil
}

func (p *TDelete) IsSetDeleteType() bool {
	return p.DeleteType != TDelete_DeleteType_DEFAULT
}

func (p *TDelete) IsSetAttributes() bool {
	return p.Attributes != nil
}

func (p *TDelete) IsSetDurability() bool {
	return p.Durability != nil
}

func (p *TDelete) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRow bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRow = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	return nil
}

func (p *TDelete) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TDelete) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumn, 0, size)
	p.Columns = tSlice
	for i := 0; i < size; i++ {
		_elem8 := &TColumn{}
		if err := _elem8.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem8), err)
		}
		p.Columns = append(p.Columns, _elem8)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TDelete) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Timestamp = &v
	}
	return nil
}

func (p *TDelete) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		temp := TDeleteType(v)
		p.DeleteType = temp
	}
	return nil
}

func (p *TDelete) readField6(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string][]byte, size)
	p.Attributes = tMap
	for i := 0; i < size; i++ {
		var _key9 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key9 = v
		}
		var _val10 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val10 = v
		}
		p.Attributes[_key9] = _val10
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TDelete) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		temp := TDurability(v)
		p.Durability = &temp
	}
	return nil
}

func (p *TDelete) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TDelete"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TDelete) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *TDelete) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetColumns() {
		if err := oprot.WriteFieldBegin("columns", thrift.LIST, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:columns: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Columns)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Columns {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:columns: ", p), err)
		}
	}
	return err
}

func (p *TDelete) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimestamp() {
		if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timestamp: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timestamp (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timestamp: ", p), err)
		}
	}
	return err
}

func (p *TDelete) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetDeleteType() {
		if err := oprot.WriteFieldBegin("deleteType", thrift.I32, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:deleteType: ", p), err)
		}
		if err := oprot.WriteI32(int32(p.DeleteType)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.deleteType (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:deleteType: ", p), err)
		}
	}
	return err
}

func (p *TDelete) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetAttributes() {
		if err := oprot.WriteFieldBegin("attributes", thrift.MAP, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:attributes: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Attributes)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Attributes {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:attributes: ", p), err)
		}
	}
	return err
}

func (p *TDelete) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetDurability() {
		if err := oprot.WriteFieldBegin("durability", thrift.I32, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:durability: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Durability)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.durability (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:durability: ", p), err)
		}
	}
	return err
}

func (p *TDelete) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDelete(%+v)", *p)
}

// Used to perform Increment operations for a single row.
//
// You can specify how this Increment should be written to the write-ahead Log (WAL)
// by changing the durability. If you don't provide durability, it defaults to
// column family's default setting for durability.
//
// Attributes:
//  - Row
//  - Columns
//  - Attributes
//  - Durability
//  - CellVisibility
type TIncrement struct {
	Row     []byte              `thrift:"row,1,required" json:"row"`
	Columns []*TColumnIncrement `thrift:"columns,2,required" json:"columns"`
	// unused field # 3
	Attributes     map[string][]byte `thrift:"attributes,4" json:"attributes,omitempty"`
	Durability     *TDurability      `thrift:"durability,5" json:"durability,omitempty"`
	CellVisibility *TCellVisibility  `thrift:"cellVisibility,6" json:"cellVisibility,omitempty"`
}

func NewTIncrement() *TIncrement {
	return &TIncrement{}
}

func (p *TIncrement) GetRow() []byte {
	return p.Row
}

func (p *TIncrement) GetColumns() []*TColumnIncrement {
	return p.Columns
}

var TIncrement_Attributes_DEFAULT map[string][]byte

func (p *TIncrement) GetAttributes() map[string][]byte {
	return p.Attributes
}

var TIncrement_Durability_DEFAULT TDurability

func (p *TIncrement) GetDurability() TDurability {
	if !p.IsSetDurability() {
		return TIncrement_Durability_DEFAULT
	}
	return *p.Durability
}

var TIncrement_CellVisibility_DEFAULT *TCellVisibility

func (p *TIncrement) GetCellVisibility() *TCellVisibility {
	if !p.IsSetCellVisibility() {
		return TIncrement_CellVisibility_DEFAULT
	}
	return p.CellVisibility
}
func (p *TIncrement) IsSetAttributes() bool {
	return p.Attributes != nil
}

func (p *TIncrement) IsSetDurability() bool {
	return p.Durability != nil
}

func (p *TIncrement) IsSetCellVisibility() bool {
	return p.CellVisibility != nil
}

func (p *TIncrement) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRow bool = false
	var issetColumns bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRow = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetColumns = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	if !issetColumns {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Columns is not set"))
	}
	return nil
}

func (p *TIncrement) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TIncrement) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumnIncrement, 0, size)
	p.Columns = tSlice
	for i := 0; i < size; i++ {
		_elem11 := &TColumnIncrement{
			Amount: 1,
		}
		if err := _elem11.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem11), err)
		}
		p.Columns = append(p.Columns, _elem11)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TIncrement) readField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string][]byte, size)
	p.Attributes = tMap
	for i := 0; i < size; i++ {
		var _key12 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key12 = v
		}
		var _val13 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val13 = v
		}
		p.Attributes[_key12] = _val13
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TIncrement) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		temp := TDurability(v)
		p.Durability = &temp
	}
	return nil
}

func (p *TIncrement) readField6(iprot thrift.TProtocol) error {
	p.CellVisibility = &TCellVisibility{}
	if err := p.CellVisibility.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.CellVisibility), err)
	}
	return nil
}

func (p *TIncrement) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TIncrement"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TIncrement) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *TIncrement) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("columns", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:columns: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Columns)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Columns {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:columns: ", p), err)
	}
	return err
}

func (p *TIncrement) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetAttributes() {
		if err := oprot.WriteFieldBegin("attributes", thrift.MAP, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:attributes: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Attributes)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Attributes {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:attributes: ", p), err)
		}
	}
	return err
}

func (p *TIncrement) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetDurability() {
		if err := oprot.WriteFieldBegin("durability", thrift.I32, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:durability: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Durability)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.durability (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:durability: ", p), err)
		}
	}
	return err
}

func (p *TIncrement) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetCellVisibility() {
		if err := oprot.WriteFieldBegin("cellVisibility", thrift.STRUCT, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:cellVisibility: ", p), err)
		}
		if err := p.CellVisibility.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.CellVisibility), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:cellVisibility: ", p), err)
		}
	}
	return err
}

func (p *TIncrement) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TIncrement(%+v)", *p)
}

// Attributes:
//  - Row
//  - Columns
//  - Attributes
//  - Durability
//  - CellVisibility
type TAppend struct {
	Row            []byte            `thrift:"row,1,required" json:"row"`
	Columns        []*TColumnValue   `thrift:"columns,2,required" json:"columns"`
	Attributes     map[string][]byte `thrift:"attributes,3" json:"attributes,omitempty"`
	Durability     *TDurability      `thrift:"durability,4" json:"durability,omitempty"`
	CellVisibility *TCellVisibility  `thrift:"cellVisibility,5" json:"cellVisibility,omitempty"`
}

func NewTAppend() *TAppend {
	return &TAppend{}
}

func (p *TAppend) GetRow() []byte {
	return p.Row
}

func (p *TAppend) GetColumns() []*TColumnValue {
	return p.Columns
}

var TAppend_Attributes_DEFAULT map[string][]byte

func (p *TAppend) GetAttributes() map[string][]byte {
	return p.Attributes
}

var TAppend_Durability_DEFAULT TDurability

func (p *TAppend) GetDurability() TDurability {
	if !p.IsSetDurability() {
		return TAppend_Durability_DEFAULT
	}
	return *p.Durability
}

var TAppend_CellVisibility_DEFAULT *TCellVisibility

func (p *TAppend) GetCellVisibility() *TCellVisibility {
	if !p.IsSetCellVisibility() {
		return TAppend_CellVisibility_DEFAULT
	}
	return p.CellVisibility
}
func (p *TAppend) IsSetAttributes() bool {
	return p.Attributes != nil
}

func (p *TAppend) IsSetDurability() bool {
	return p.Durability != nil
}

func (p *TAppend) IsSetCellVisibility() bool {
	return p.CellVisibility != nil
}

func (p *TAppend) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRow bool = false
	var issetColumns bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRow = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetColumns = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	if !issetColumns {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Columns is not set"))
	}
	return nil
}

func (p *TAppend) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TAppend) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumnValue, 0, size)
	p.Columns = tSlice
	for i := 0; i < size; i++ {
		_elem14 := &TColumnValue{}
		if err := _elem14.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem14), err)
		}
		p.Columns = append(p.Columns, _elem14)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TAppend) readField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string][]byte, size)
	p.Attributes = tMap
	for i := 0; i < size; i++ {
		var _key15 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key15 = v
		}
		var _val16 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val16 = v
		}
		p.Attributes[_key15] = _val16
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TAppend) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		temp := TDurability(v)
		p.Durability = &temp
	}
	return nil
}

func (p *TAppend) readField5(iprot thrift.TProtocol) error {
	p.CellVisibility = &TCellVisibility{}
	if err := p.CellVisibility.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.CellVisibility), err)
	}
	return nil
}

func (p *TAppend) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TAppend"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TAppend) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *TAppend) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("columns", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:columns: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Columns)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Columns {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:columns: ", p), err)
	}
	return err
}

func (p *TAppend) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetAttributes() {
		if err := oprot.WriteFieldBegin("attributes", thrift.MAP, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:attributes: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Attributes)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Attributes {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:attributes: ", p), err)
		}
	}
	return err
}

func (p *TAppend) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetDurability() {
		if err := oprot.WriteFieldBegin("durability", thrift.I32, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:durability: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Durability)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.durability (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:durability: ", p), err)
		}
	}
	return err
}

func (p *TAppend) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetCellVisibility() {
		if err := oprot.WriteFieldBegin("cellVisibility", thrift.STRUCT, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:cellVisibility: ", p), err)
		}
		if err := p.CellVisibility.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.CellVisibility), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:cellVisibility: ", p), err)
		}
	}
	return err
}

func (p *TAppend) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TAppend(%+v)", *p)
}

// Any timestamps in the columns are ignored, use timeRange to select by timestamp.
// Max versions defaults to 1.
//
// Attributes:
//  - StartRow
//  - StopRow
//  - Columns
//  - Caching
//  - MaxVersions
//  - TimeRange
//  - FilterString
//  - BatchSize
//  - Attributes
//  - Authorizations
//  - Reversed
//  - CacheBlocks
type TScan struct {
	StartRow       []byte            `thrift:"startRow,1" json:"startRow,omitempty"`
	StopRow        []byte            `thrift:"stopRow,2" json:"stopRow,omitempty"`
	Columns        []*TColumn        `thrift:"columns,3" json:"columns,omitempty"`
	Caching        *int32            `thrift:"caching,4" json:"caching,omitempty"`
	MaxVersions    int32             `thrift:"maxVersions,5" json:"maxVersions,omitempty"`
	TimeRange      *TTimeRange       `thrift:"timeRange,6" json:"timeRange,omitempty"`
	FilterString   []byte            `thrift:"filterString,7" json:"filterString,omitempty"`
	BatchSize      *int32            `thrift:"batchSize,8" json:"batchSize,omitempty"`
	Attributes     map[string][]byte `thrift:"attributes,9" json:"attributes,omitempty"`
	Authorizations *TAuthorization   `thrift:"authorizations,10" json:"authorizations,omitempty"`
	Reversed       *bool             `thrift:"reversed,11" json:"reversed,omitempty"`
	CacheBlocks    *bool             `thrift:"cacheBlocks,12" json:"cacheBlocks,omitempty"`
}

func NewTScan() *TScan {
	return &TScan{
		MaxVersions: 1,
	}
}

var TScan_StartRow_DEFAULT []byte

func (p *TScan) GetStartRow() []byte {
	return p.StartRow
}

var TScan_StopRow_DEFAULT []byte

func (p *TScan) GetStopRow() []byte {
	return p.StopRow
}

var TScan_Columns_DEFAULT []*TColumn

func (p *TScan) GetColumns() []*TColumn {
	return p.Columns
}

var TScan_Caching_DEFAULT int32

func (p *TScan) GetCaching() int32 {
	if !p.IsSetCaching() {
		return TScan_Caching_DEFAULT
	}
	return *p.Caching
}

var TScan_MaxVersions_DEFAULT int32 = 1

func (p *TScan) GetMaxVersions() int32 {
	return p.MaxVersions
}

var TScan_TimeRange_DEFAULT *TTimeRange

func (p *TScan) GetTimeRange() *TTimeRange {
	if !p.IsSetTimeRange() {
		return TScan_TimeRange_DEFAULT
	}
	return p.TimeRange
}

var TScan_FilterString_DEFAULT []byte

func (p *TScan) GetFilterString() []byte {
	return p.FilterString
}

var TScan_BatchSize_DEFAULT int32

func (p *TScan) GetBatchSize() int32 {
	if !p.IsSetBatchSize() {
		return TScan_BatchSize_DEFAULT
	}
	return *p.BatchSize
}

var TScan_Attributes_DEFAULT map[string][]byte

func (p *TScan) GetAttributes() map[string][]byte {
	return p.Attributes
}

var TScan_Authorizations_DEFAULT *TAuthorization

func (p *TScan) GetAuthorizations() *TAuthorization {
	if !p.IsSetAuthorizations() {
		return TScan_Authorizations_DEFAULT
	}
	return p.Authorizations
}

var TScan_Reversed_DEFAULT bool

func (p *TScan) GetReversed() bool {
	if !p.IsSetReversed() {
		return TScan_Reversed_DEFAULT
	}
	return *p.Reversed
}

var TScan_CacheBlocks_DEFAULT bool

func (p *TScan) GetCacheBlocks() bool {
	if !p.IsSetCacheBlocks() {
		return TScan_CacheBlocks_DEFAULT
	}
	return *p.CacheBlocks
}
func (p *TScan) IsSetStartRow() bool {
	return p.StartRow != nil
}

func (p *TScan) IsSetStopRow() bool {
	return p.StopRow != nil
}

func (p *TScan) IsSetColumns() bool {
	return p.Columns != nil
}

func (p *TScan) IsSetCaching() bool {
	return p.Caching != nil
}

func (p *TScan) IsSetMaxVersions() bool {
	return p.MaxVersions != TScan_MaxVersions_DEFAULT
}

func (p *TScan) IsSetTimeRange() bool {
	return p.TimeRange != nil
}

func (p *TScan) IsSetFilterString() bool {
	return p.FilterString != nil
}

func (p *TScan) IsSetBatchSize() bool {
	return p.BatchSize != nil
}

func (p *TScan) IsSetAttributes() bool {
	return p.Attributes != nil
}

func (p *TScan) IsSetAuthorizations() bool {
	return p.Authorizations != nil
}

func (p *TScan) IsSetReversed() bool {
	return p.Reversed != nil
}

func (p *TScan) IsSetCacheBlocks() bool {
	return p.CacheBlocks != nil
}

func (p *TScan) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.readField9(iprot); err != nil {
				return err
			}
		case 10:
			if err := p.readField10(iprot); err != nil {
				return err
			}
		case 11:
			if err := p.readField11(iprot); err != nil {
				return err
			}
		case 12:
			if err := p.readField12(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TScan) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.StartRow = v
	}
	return nil
}

func (p *TScan) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.StopRow = v
	}
	return nil
}

func (p *TScan) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumn, 0, size)
	p.Columns = tSlice
	for i := 0; i < size; i++ {
		_elem17 := &TColumn{}
		if err := _elem17.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem17), err)
		}
		p.Columns = append(p.Columns, _elem17)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TScan) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Caching = &v
	}
	return nil
}

func (p *TScan) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.MaxVersions = v
	}
	return nil
}

func (p *TScan) readField6(iprot thrift.TProtocol) error {
	p.TimeRange = &TTimeRange{}
	if err := p.TimeRange.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TimeRange), err)
	}
	return nil
}

func (p *TScan) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.FilterString = v
	}
	return nil
}

func (p *TScan) readField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.BatchSize = &v
	}
	return nil
}

func (p *TScan) readField9(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string][]byte, size)
	p.Attributes = tMap
	for i := 0; i < size; i++ {
		var _key18 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key18 = v
		}
		var _val19 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val19 = v
		}
		p.Attributes[_key18] = _val19
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TScan) readField10(iprot thrift.TProtocol) error {
	p.Authorizations = &TAuthorization{}
	if err := p.Authorizations.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Authorizations), err)
	}
	return nil
}

func (p *TScan) readField11(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 11: ", err)
	} else {
		p.Reversed = &v
	}
	return nil
}

func (p *TScan) readField12(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 12: ", err)
	} else {
		p.CacheBlocks = &v
	}
	return nil
}

func (p *TScan) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TScan"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := p.writeField10(oprot); err != nil {
		return err
	}
	if err := p.writeField11(oprot); err != nil {
		return err
	}
	if err := p.writeField12(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TScan) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetStartRow() {
		if err := oprot.WriteFieldBegin("startRow", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:startRow: ", p), err)
		}
		if err := oprot.WriteBinary(p.StartRow); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.startRow (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:startRow: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetStopRow() {
		if err := oprot.WriteFieldBegin("stopRow", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:stopRow: ", p), err)
		}
		if err := oprot.WriteBinary(p.StopRow); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.stopRow (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:stopRow: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetColumns() {
		if err := oprot.WriteFieldBegin("columns", thrift.LIST, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:columns: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Columns)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Columns {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:columns: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetCaching() {
		if err := oprot.WriteFieldBegin("caching", thrift.I32, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:caching: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Caching)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.caching (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:caching: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetMaxVersions() {
		if err := oprot.WriteFieldBegin("maxVersions", thrift.I32, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:maxVersions: ", p), err)
		}
		if err := oprot.WriteI32(int32(p.MaxVersions)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.maxVersions (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:maxVersions: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimeRange() {
		if err := oprot.WriteFieldBegin("timeRange", thrift.STRUCT, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:timeRange: ", p), err)
		}
		if err := p.TimeRange.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TimeRange), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:timeRange: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetFilterString() {
		if err := oprot.WriteFieldBegin("filterString", thrift.STRING, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:filterString: ", p), err)
		}
		if err := oprot.WriteBinary(p.FilterString); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.filterString (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:filterString: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetBatchSize() {
		if err := oprot.WriteFieldBegin("batchSize", thrift.I32, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:batchSize: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.BatchSize)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.batchSize (8) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:batchSize: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField9(oprot thrift.TProtocol) (err error) {
	if p.IsSetAttributes() {
		if err := oprot.WriteFieldBegin("attributes", thrift.MAP, 9); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:attributes: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Attributes)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Attributes {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 9:attributes: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField10(oprot thrift.TProtocol) (err error) {
	if p.IsSetAuthorizations() {
		if err := oprot.WriteFieldBegin("authorizations", thrift.STRUCT, 10); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:authorizations: ", p), err)
		}
		if err := p.Authorizations.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Authorizations), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 10:authorizations: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField11(oprot thrift.TProtocol) (err error) {
	if p.IsSetReversed() {
		if err := oprot.WriteFieldBegin("reversed", thrift.BOOL, 11); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:reversed: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Reversed)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.reversed (11) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 11:reversed: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField12(oprot thrift.TProtocol) (err error) {
	if p.IsSetCacheBlocks() {
		if err := oprot.WriteFieldBegin("cacheBlocks", thrift.BOOL, 12); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 12:cacheBlocks: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.CacheBlocks)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.cacheBlocks (12) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 12:cacheBlocks: ", p), err)
		}
	}
	return err
}

func (p *TScan) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TScan(%+v)", *p)
}

// Atomic mutation for the specified row. It can be either Put or Delete.
//
// Attributes:
//  - Put
//  - DeleteSingle
type TMutation struct {
	Put          *TPut    `thrift:"put,1" json:"put,omitempty"`
	DeleteSingle *TDelete `thrift:"deleteSingle,2" json:"deleteSingle,omitempty"`
}

func NewTMutation() *TMutation {
	return &TMutation{}
}

var TMutation_Put_DEFAULT *TPut

func (p *TMutation) GetPut() *TPut {
	if !p.IsSetPut() {
		return TMutation_Put_DEFAULT
	}
	return p.Put
}

var TMutation_DeleteSingle_DEFAULT *TDelete

func (p *TMutation) GetDeleteSingle() *TDelete {
	if !p.IsSetDeleteSingle() {
		return TMutation_DeleteSingle_DEFAULT
	}
	return p.DeleteSingle
}
func (p *TMutation) CountSetFieldsTMutation() int {
	count := 0
	if p.IsSetPut() {
		count++
	}
	if p.IsSetDeleteSingle() {
		count++
	}
	return count

}

func (p *TMutation) IsSetPut() bool {
	return p.Put != nil
}

func (p *TMutation) IsSetDeleteSingle() bool {
	return p.DeleteSingle != nil
}

func (p *TMutation) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TMutation) readField1(iprot thrift.TProtocol) error {
	p.Put = &TPut{}
	if err := p.Put.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Put), err)
	}
	return nil
}

func (p *TMutation) readField2(iprot thrift.TProtocol) error {
	p.DeleteSingle = &TDelete{
		DeleteType: 1,
	}
	if err := p.DeleteSingle.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.DeleteSingle), err)
	}
	return nil
}

func (p *TMutation) Write(oprot thrift.TProtocol) error {
	if c := p.CountSetFieldsTMutation(); c != 1 {
		return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
	}
	if err := oprot.WriteStructBegin("TMutation"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TMutation) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetPut() {
		if err := oprot.WriteFieldBegin("put", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:put: ", p), err)
		}
		if err := p.Put.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Put), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:put: ", p), err)
		}
	}
	return err
}

func (p *TMutation) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetDeleteSingle() {
		if err := oprot.WriteFieldBegin("deleteSingle", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:deleteSingle: ", p), err)
		}
		if err := p.DeleteSingle.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.DeleteSingle), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:deleteSingle: ", p), err)
		}
	}
	return err
}

func (p *TMutation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TMutation(%+v)", *p)
}

// A TRowMutations object is used to apply a number of Mutations to a single row.
//
// Attributes:
//  - Row
//  - Mutations
type TRowMutations struct {
	Row       []byte       `thrift:"row,1,required" json:"row"`
	Mutations []*TMutation `thrift:"mutations,2,required" json:"mutations"`
}

func NewTRowMutations() *TRowMutations {
	return &TRowMutations{}
}

func (p *TRowMutations) GetRow() []byte {
	return p.Row
}

func (p *TRowMutations) GetMutations() []*TMutation {
	return p.Mutations
}
func (p *TRowMutations) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRow bool = false
	var issetMutations bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRow = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetMutations = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	if !issetMutations {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Mutations is not set"))
	}
	return nil
}

func (p *TRowMutations) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TRowMutations) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TMutation, 0, size)
	p.Mutations = tSlice
	for i := 0; i < size; i++ {
		_elem20 := &TMutation{}
		if err := _elem20.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem20), err)
		}
		p.Mutations = append(p.Mutations, _elem20)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TRowMutations) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TRowMutations"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TRowMutations) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *TRowMutations) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mutations", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:mutations: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Mutations)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Mutations {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:mutations: ", p), err)
	}
	return err
}

func (p *TRowMutations) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TRowMutations(%+v)", *p)
}

// Attributes:
//  - RegionId
//  - TableName
//  - StartKey
//  - EndKey
//  - Offline
//  - Split
//  - ReplicaId
type THRegionInfo struct {
	RegionId  int64  `thrift:"regionId,1,required" json:"regionId"`
	TableName []byte `thrift:"tableName,2,required" json:"tableName"`
	StartKey  []byte `thrift:"startKey,3" json:"startKey,omitempty"`
	EndKey    []byte `thrift:"endKey,4" json:"endKey,omitempty"`
	Offline   *bool  `thrift:"offline,5" json:"offline,omitempty"`
	Split     *bool  `thrift:"split,6" json:"split,omitempty"`
	ReplicaId *int32 `thrift:"replicaId,7" json:"replicaId,omitempty"`
}

func NewTHRegionInfo() *THRegionInfo {
	return &THRegionInfo{}
}

func (p *THRegionInfo) GetRegionId() int64 {
	return p.RegionId
}

func (p *THRegionInfo) GetTableName() []byte {
	return p.TableName
}

var THRegionInfo_StartKey_DEFAULT []byte

func (p *THRegionInfo) GetStartKey() []byte {
	return p.StartKey
}

var THRegionInfo_EndKey_DEFAULT []byte

func (p *THRegionInfo) GetEndKey() []byte {
	return p.EndKey
}

var THRegionInfo_Offline_DEFAULT bool

func (p *THRegionInfo) GetOffline() bool {
	if !p.IsSetOffline() {
		return THRegionInfo_Offline_DEFAULT
	}
	return *p.Offline
}

var THRegionInfo_Split_DEFAULT bool

func (p *THRegionInfo) GetSplit() bool {
	if !p.IsSetSplit() {
		return THRegionInfo_Split_DEFAULT
	}
	return *p.Split
}

var THRegionInfo_ReplicaId_DEFAULT int32

func (p *THRegionInfo) GetReplicaId() int32 {
	if !p.IsSetReplicaId() {
		return THRegionInfo_ReplicaId_DEFAULT
	}
	return *p.ReplicaId
}
func (p *THRegionInfo) IsSetStartKey() bool {
	return p.StartKey != nil
}

func (p *THRegionInfo) IsSetEndKey() bool {
	return p.EndKey != nil
}

func (p *THRegionInfo) IsSetOffline() bool {
	return p.Offline != nil
}

func (p *THRegionInfo) IsSetSplit() bool {
	return p.Split != nil
}

func (p *THRegionInfo) IsSetReplicaId() bool {
	return p.ReplicaId != nil
}

func (p *THRegionInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRegionId bool = false
	var issetTableName bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRegionId = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTableName = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRegionId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RegionId is not set"))
	}
	if !issetTableName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TableName is not set"))
	}
	return nil
}

func (p *THRegionInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.RegionId = v
	}
	return nil
}

func (p *THRegionInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.TableName = v
	}
	return nil
}

func (p *THRegionInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.StartKey = v
	}
	return nil
}

func (p *THRegionInfo) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.EndKey = v
	}
	return nil
}

func (p *THRegionInfo) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Offline = &v
	}
	return nil
}

func (p *THRegionInfo) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Split = &v
	}
	return nil
}

func (p *THRegionInfo) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.ReplicaId = &v
	}
	return nil
}

func (p *THRegionInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("THRegionInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *THRegionInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("regionId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:regionId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.RegionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.regionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:regionId: ", p), err)
	}
	return err
}

func (p *THRegionInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tableName", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tableName: ", p), err)
	}
	if err := oprot.WriteBinary(p.TableName); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tableName (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tableName: ", p), err)
	}
	return err
}

func (p *THRegionInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetStartKey() {
		if err := oprot.WriteFieldBegin("startKey", thrift.STRING, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:startKey: ", p), err)
		}
		if err := oprot.WriteBinary(p.StartKey); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.startKey (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:startKey: ", p), err)
		}
	}
	return err
}

func (p *THRegionInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetEndKey() {
		if err := oprot.WriteFieldBegin("endKey", thrift.STRING, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:endKey: ", p), err)
		}
		if err := oprot.WriteBinary(p.EndKey); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.endKey (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:endKey: ", p), err)
		}
	}
	return err
}

func (p *THRegionInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetOffline() {
		if err := oprot.WriteFieldBegin("offline", thrift.BOOL, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:offline: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Offline)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.offline (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:offline: ", p), err)
		}
	}
	return err
}

func (p *THRegionInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetSplit() {
		if err := oprot.WriteFieldBegin("split", thrift.BOOL, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:split: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Split)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.split (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:split: ", p), err)
		}
	}
	return err
}

func (p *THRegionInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetReplicaId() {
		if err := oprot.WriteFieldBegin("replicaId", thrift.I32, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:replicaId: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.ReplicaId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.replicaId (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:replicaId: ", p), err)
		}
	}
	return err
}

func (p *THRegionInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THRegionInfo(%+v)", *p)
}

// Attributes:
//  - HostName
//  - Port
//  - StartCode
type TServerName struct {
	HostName  string `thrift:"hostName,1,required" json:"hostName"`
	Port      *int32 `thrift:"port,2" json:"port,omitempty"`
	StartCode *int64 `thrift:"startCode,3" json:"startCode,omitempty"`
}

func NewTServerName() *TServerName {
	return &TServerName{}
}

func (p *TServerName) GetHostName() string {
	return p.HostName
}

var TServerName_Port_DEFAULT int32

func (p *TServerName) GetPort() int32 {
	if !p.IsSetPort() {
		return TServerName_Port_DEFAULT
	}
	return *p.Port
}

var TServerName_StartCode_DEFAULT int64

func (p *TServerName) GetStartCode() int64 {
	if !p.IsSetStartCode() {
		return TServerName_StartCode_DEFAULT
	}
	return *p.StartCode
}
func (p *TServerName) IsSetPort() bool {
	return p.Port != nil
}

func (p *TServerName) IsSetStartCode() bool {
	return p.StartCode != nil
}

func (p *TServerName) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetHostName bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetHostName = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetHostName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field HostName is not set"))
	}
	return nil
}

func (p *TServerName) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.HostName = v
	}
	return nil
}

func (p *TServerName) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Port = &v
	}
	return nil
}

func (p *TServerName) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.StartCode = &v
	}
	return nil
}

func (p *TServerName) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TServerName"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TServerName) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("hostName", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:hostName: ", p), err)
	}
	if err := oprot.WriteString(string(p.HostName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.hostName (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:hostName: ", p), err)
	}
	return err
}

func (p *TServerName) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetPort() {
		if err := oprot.WriteFieldBegin("port", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:port: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Port)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.port (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:port: ", p), err)
		}
	}
	return err
}

func (p *TServerName) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetStartCode() {
		if err := oprot.WriteFieldBegin("startCode", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:startCode: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.StartCode)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.startCode (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:startCode: ", p), err)
		}
	}
	return err
}

func (p *TServerName) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TServerName(%+v)", *p)
}

// Attributes:
//  - ServerName
//  - RegionInfo
type THRegionLocation struct {
	ServerName *TServerName  `thrift:"serverName,1,required" json:"serverName"`
	RegionInfo *THRegionInfo `thrift:"regionInfo,2,required" json:"regionInfo"`
}

func NewTHRegionLocation() *THRegionLocation {
	return &THRegionLocation{}
}

var THRegionLocation_ServerName_DEFAULT *TServerName

func (p *THRegionLocation) GetServerName() *TServerName {
	if !p.IsSetServerName() {
		return THRegionLocation_ServerName_DEFAULT
	}
	return p.ServerName
}

var THRegionLocation_RegionInfo_DEFAULT *THRegionInfo

func (p *THRegionLocation) GetRegionInfo() *THRegionInfo {
	if !p.IsSetRegionInfo() {
		return THRegionLocation_RegionInfo_DEFAULT
	}
	return p.RegionInfo
}
func (p *THRegionLocation) IsSetServerName() bool {
	return p.ServerName != nil
}

func (p *THRegionLocation) IsSetRegionInfo() bool {
	return p.RegionInfo != nil
}

func (p *THRegionLocation) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetServerName bool = false
	var issetRegionInfo bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetServerName = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetRegionInfo = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetServerName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ServerName is not set"))
	}
	if !issetRegionInfo {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RegionInfo is not set"))
	}
	return nil
}

func (p *THRegionLocation) readField1(iprot thrift.TProtocol) error {
	p.ServerName = &TServerName{}
	if err := p.ServerName.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ServerName), err)
	}
	return nil
}

func (p *THRegionLocation) readField2(iprot thrift.TProtocol) error {
	p.RegionInfo = &THRegionInfo{}
	if err := p.RegionInfo.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RegionInfo), err)
	}
	return nil
}

func (p *THRegionLocation) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("THRegionLocation"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *THRegionLocation) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serverName", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:serverName: ", p), err)
	}
	if err := p.ServerName.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ServerName), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:serverName: ", p), err)
	}
	return err
}

func (p *THRegionLocation) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("regionInfo", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:regionInfo: ", p), err)
	}
	if err := p.RegionInfo.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RegionInfo), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:regionInfo: ", p), err)
	}
	return err
}

func (p *THRegionLocation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THRegionLocation(%+v)", *p)
}

// A TIOError exception signals that an error occurred communicating
// to the HBase master or a HBase region server. Also used to return
// more general HBase error conditions.
//
// Attributes:
//  - Message
type TIOError struct {
	Message *string `thrift:"message,1" json:"message,omitempty"`
}

func NewTIOError() *TIOError {
	return &TIOError{}
}

var TIOError_Message_DEFAULT string

func (p *TIOError) GetMessage() string {
	if !p.IsSetMessage() {
		return TIOError_Message_DEFAULT
	}
	return *p.Message
}
func (p *TIOError) IsSetMessage() bool {
	return p.Message != nil
}

func (p *TIOError) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TIOError) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *TIOError) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TIOError"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TIOError) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *TIOError) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TIOError(%+v)", *p)
}

func (p *TIOError) Error() string {
	return p.String()
}

// A TIllegalArgument exception indicates an illegal or invalid
// argument was passed into a procedure.
//
// Attributes:
//  - Message
type TIllegalArgument struct {
	Message *string `thrift:"message,1" json:"message,omitempty"`
}

func NewTIllegalArgument() *TIllegalArgument {
	return &TIllegalArgument{}
}

var TIllegalArgument_Message_DEFAULT string

func (p *TIllegalArgument) GetMessage() string {
	if !p.IsSetMessage() {
		return TIllegalArgument_Message_DEFAULT
	}
	return *p.Message
}
func (p *TIllegalArgument) IsSetMessage() bool {
	return p.Message != nil
}

func (p *TIllegalArgument) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TIllegalArgument) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *TIllegalArgument) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TIllegalArgument"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TIllegalArgument) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *TIllegalArgument) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TIllegalArgument(%+v)", *p)
}

func (p *TIllegalArgument) Error() string {
	return p.String()
}
