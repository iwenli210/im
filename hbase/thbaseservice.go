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

type THBaseService interface {
	// Test for the existence of columns in the table, as specified in the TGet.
	//
	// @return true if the specified TGet matches one or more keys, false if not
	//
	// Parameters:
	//  - Table: the table to check on
	//  - Tget: the TGet to check for
	Exists(table []byte, tget *TGet) (r bool, err error)
	// Method for getting data from a row.
	//
	// If the row cannot be found an empty Result is returned.
	// This can be checked by the empty field of the TResult
	//
	// @return the result
	//
	// Parameters:
	//  - Table: the table to get from
	//  - Tget: the TGet to fetch
	Get(table []byte, tget *TGet) (r *TResult_, err error)
	// Method for getting multiple rows.
	//
	// If a row cannot be found there will be a null
	// value in the result list for that TGet at the
	// same position.
	//
	// So the Results are in the same order as the TGets.
	//
	// Parameters:
	//  - Table: the table to get from
	//  - Tgets: a list of TGets to fetch, the Result list
	// will have the Results at corresponding positions
	// or null if there was an error
	GetMultiple(table []byte, tgets []*TGet) (r []*TResult_, err error)
	// Commit a TPut to a table.
	//
	// Parameters:
	//  - Table: the table to put data in
	//  - Tput: the TPut to put
	Put(table []byte, tput *TPut) (err error)
	// Atomically checks if a row/family/qualifier value matches the expected
	// value. If it does, it adds the TPut.
	//
	// @return true if the new put was executed, false otherwise
	//
	// Parameters:
	//  - Table: to check in and put to
	//  - Row: row to check
	//  - Family: column family to check
	//  - Qualifier: column qualifier to check
	//  - Value: the expected value, if not provided the
	// check is for the non-existence of the
	// column in question
	//  - Tput: the TPut to put if the check succeeds
	CheckAndPut(table []byte, row []byte, family []byte, qualifier []byte, value []byte, tput *TPut) (r bool, err error)
	// Commit a List of Puts to the table.
	//
	// Parameters:
	//  - Table: the table to put data in
	//  - Tputs: a list of TPuts to commit
	PutMultiple(table []byte, tputs []*TPut) (err error)
	// Deletes as specified by the TDelete.
	//
	// Note: "delete" is a reserved keyword and cannot be used in Thrift
	// thus the inconsistent naming scheme from the other functions.
	//
	// Parameters:
	//  - Table: the table to delete from
	//  - Tdelete: the TDelete to delete
	DeleteSingle(table []byte, tdelete *TDelete) (err error)
	// Bulk commit a List of TDeletes to the table.
	//
	// Throws a TIOError if any of the deletes fail.
	//
	// Always returns an empty list for backwards compatibility.
	//
	// Parameters:
	//  - Table: the table to delete from
	//  - Tdeletes: list of TDeletes to delete
	DeleteMultiple(table []byte, tdeletes []*TDelete) (r []*TDelete, err error)
	// Atomically checks if a row/family/qualifier value matches the expected
	// value. If it does, it adds the delete.
	//
	// @return true if the new delete was executed, false otherwise
	//
	// Parameters:
	//  - Table: to check in and delete from
	//  - Row: row to check
	//  - Family: column family to check
	//  - Qualifier: column qualifier to check
	//  - Value: the expected value, if not provided the
	// check is for the non-existence of the
	// column in question
	//  - Tdelete: the TDelete to execute if the check succeeds
	CheckAndDelete(table []byte, row []byte, family []byte, qualifier []byte, value []byte, tdelete *TDelete) (r bool, err error)
	// Parameters:
	//  - Table: the table to increment the value on
	//  - Tincrement: the TIncrement to increment
	Increment(table []byte, tincrement *TIncrement) (r *TResult_, err error)
	// Parameters:
	//  - Table: the table to append the value on
	//  - Tappend: the TAppend to append
	Append(table []byte, tappend *TAppend) (r *TResult_, err error)
	// Get a Scanner for the provided TScan object.
	//
	// @return Scanner Id to be used with other scanner procedures
	//
	// Parameters:
	//  - Table: the table to get the Scanner for
	//  - Tscan: the scan object to get a Scanner for
	OpenScanner(table []byte, tscan *TScan) (r int32, err error)
	// Grabs multiple rows from a Scanner.
	//
	// @return Between zero and numRows TResults
	//
	// Parameters:
	//  - ScannerId: the Id of the Scanner to return rows from. This is an Id returned from the openScanner function.
	//  - NumRows: number of rows to return
	GetScannerRows(scannerId int32, numRows int32) (r []*TResult_, err error)
	// Closes the scanner. Should be called to free server side resources timely.
	// Typically close once the scanner is not needed anymore, i.e. after looping
	// over it to get all the required rows.
	//
	// Parameters:
	//  - ScannerId: the Id of the Scanner to close *
	CloseScanner(scannerId int32) (err error)
	// mutateRow performs multiple mutations atomically on a single row.
	//
	// Parameters:
	//  - Table: table to apply the mutations
	//  - TrowMutations: mutations to apply
	MutateRow(table []byte, trowMutations *TRowMutations) (err error)
	// Get results for the provided TScan object.
	// This helper function opens a scanner, get the results and close the scanner.
	//
	// @return between zero and numRows TResults
	//
	// Parameters:
	//  - Table: the table to get the Scanner for
	//  - Tscan: the scan object to get a Scanner for
	//  - NumRows: number of rows to return
	GetScannerResults(table []byte, tscan *TScan, numRows int32) (r []*TResult_, err error)
	// Given a table and a row get the location of the region that
	// would contain the given row key.
	//
	// reload = true means the cache will be cleared and the location
	// will be fetched from meta.
	//
	// Parameters:
	//  - Table
	//  - Row
	//  - Reload
	GetRegionLocation(table []byte, row []byte, reload bool) (r *THRegionLocation, err error)
	// Get all of the region locations for a given table.
	//
	//
	// Parameters:
	//  - Table
	GetAllRegionLocations(table []byte) (r []*THRegionLocation, err error)
	// Atomically checks if a row/family/qualifier value matches the expected
	// value. If it does, it mutates the row.
	//
	// @return true if the row was mutated, false otherwise
	//
	// Parameters:
	//  - Table: to check in and delete from
	//  - Row: row to check
	//  - Family: column family to check
	//  - Qualifier: column qualifier to check
	//  - CompareOp: comparison to make on the value
	//  - Value: the expected value to be compared against, if not provided the
	// check is for the non-existence of the column in question
	//  - RowMutations: row mutations to execute if the value matches
	CheckAndMutate(table []byte, row []byte, family []byte, qualifier []byte, compareOp TCompareOp, value []byte, rowMutations *TRowMutations) (r bool, err error)
}

type THBaseServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewTHBaseServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *THBaseServiceClient {
	return &THBaseServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewTHBaseServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *THBaseServiceClient {
	return &THBaseServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Test for the existence of columns in the table, as specified in the TGet.
//
// @return true if the specified TGet matches one or more keys, false if not
//
// Parameters:
//  - Table: the table to check on
//  - Tget: the TGet to check for
func (p *THBaseServiceClient) Exists(table []byte, tget *TGet) (r bool, err error) {
	if err = p.sendExists(table, tget); err != nil {
		return
	}
	return p.recvExists()
}

func (p *THBaseServiceClient) sendExists(table []byte, tget *TGet) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("exists", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceExistsArgs{
		Table: table,
		Tget:  tget,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvExists() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "exists" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "exists failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "exists failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error21 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error22 error
		error22, err = error21.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error22
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "exists failed: invalid message type")
		return
	}
	result := THBaseServiceExistsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Method for getting data from a row.
//
// If the row cannot be found an empty Result is returned.
// This can be checked by the empty field of the TResult
//
// @return the result
//
// Parameters:
//  - Table: the table to get from
//  - Tget: the TGet to fetch
func (p *THBaseServiceClient) Get(table []byte, tget *TGet) (r *TResult_, err error) {
	if err = p.sendGet(table, tget); err != nil {
		return
	}
	return p.recvGet()
}

func (p *THBaseServiceClient) sendGet(table []byte, tget *TGet) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("get", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceGetArgs{
		Table: table,
		Tget:  tget,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvGet() (value *TResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "get" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "get failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "get failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error23 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error24 error
		error24, err = error23.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error24
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "get failed: invalid message type")
		return
	}
	result := THBaseServiceGetResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Method for getting multiple rows.
//
// If a row cannot be found there will be a null
// value in the result list for that TGet at the
// same position.
//
// So the Results are in the same order as the TGets.
//
// Parameters:
//  - Table: the table to get from
//  - Tgets: a list of TGets to fetch, the Result list
// will have the Results at corresponding positions
// or null if there was an error
func (p *THBaseServiceClient) GetMultiple(table []byte, tgets []*TGet) (r []*TResult_, err error) {
	if err = p.sendGetMultiple(table, tgets); err != nil {
		return
	}
	return p.recvGetMultiple()
}

func (p *THBaseServiceClient) sendGetMultiple(table []byte, tgets []*TGet) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getMultiple", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceGetMultipleArgs{
		Table: table,
		Tgets: tgets,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvGetMultiple() (value []*TResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getMultiple" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getMultiple failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getMultiple failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error25 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error26 error
		error26, err = error25.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error26
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getMultiple failed: invalid message type")
		return
	}
	result := THBaseServiceGetMultipleResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Commit a TPut to a table.
//
// Parameters:
//  - Table: the table to put data in
//  - Tput: the TPut to put
func (p *THBaseServiceClient) Put(table []byte, tput *TPut) (err error) {
	if err = p.sendPut(table, tput); err != nil {
		return
	}
	return p.recvPut()
}

func (p *THBaseServiceClient) sendPut(table []byte, tput *TPut) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("put", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServicePutArgs{
		Table: table,
		Tput:  tput,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvPut() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "put" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "put failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "put failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error27 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error28 error
		error28, err = error27.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error28
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "put failed: invalid message type")
		return
	}
	result := THBaseServicePutResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	return
}

// Atomically checks if a row/family/qualifier value matches the expected
// value. If it does, it adds the TPut.
//
// @return true if the new put was executed, false otherwise
//
// Parameters:
//  - Table: to check in and put to
//  - Row: row to check
//  - Family: column family to check
//  - Qualifier: column qualifier to check
//  - Value: the expected value, if not provided the
// check is for the non-existence of the
// column in question
//  - Tput: the TPut to put if the check succeeds
func (p *THBaseServiceClient) CheckAndPut(table []byte, row []byte, family []byte, qualifier []byte, value []byte, tput *TPut) (r bool, err error) {
	if err = p.sendCheckAndPut(table, row, family, qualifier, value, tput); err != nil {
		return
	}
	return p.recvCheckAndPut()
}

func (p *THBaseServiceClient) sendCheckAndPut(table []byte, row []byte, family []byte, qualifier []byte, value []byte, tput *TPut) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("checkAndPut", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceCheckAndPutArgs{
		Table:     table,
		Row:       row,
		Family:    family,
		Qualifier: qualifier,
		Value:     value,
		Tput:      tput,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvCheckAndPut() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "checkAndPut" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "checkAndPut failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "checkAndPut failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error29 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error30 error
		error30, err = error29.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error30
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "checkAndPut failed: invalid message type")
		return
	}
	result := THBaseServiceCheckAndPutResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Commit a List of Puts to the table.
//
// Parameters:
//  - Table: the table to put data in
//  - Tputs: a list of TPuts to commit
func (p *THBaseServiceClient) PutMultiple(table []byte, tputs []*TPut) (err error) {
	if err = p.sendPutMultiple(table, tputs); err != nil {
		return
	}
	return p.recvPutMultiple()
}

func (p *THBaseServiceClient) sendPutMultiple(table []byte, tputs []*TPut) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("putMultiple", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServicePutMultipleArgs{
		Table: table,
		Tputs: tputs,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvPutMultiple() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "putMultiple" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "putMultiple failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "putMultiple failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error31 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error32 error
		error32, err = error31.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error32
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "putMultiple failed: invalid message type")
		return
	}
	result := THBaseServicePutMultipleResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	return
}

// Deletes as specified by the TDelete.
//
// Note: "delete" is a reserved keyword and cannot be used in Thrift
// thus the inconsistent naming scheme from the other functions.
//
// Parameters:
//  - Table: the table to delete from
//  - Tdelete: the TDelete to delete
func (p *THBaseServiceClient) DeleteSingle(table []byte, tdelete *TDelete) (err error) {
	if err = p.sendDeleteSingle(table, tdelete); err != nil {
		return
	}
	return p.recvDeleteSingle()
}

func (p *THBaseServiceClient) sendDeleteSingle(table []byte, tdelete *TDelete) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("deleteSingle", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceDeleteSingleArgs{
		Table:   table,
		Tdelete: tdelete,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvDeleteSingle() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "deleteSingle" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "deleteSingle failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "deleteSingle failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error33 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error34 error
		error34, err = error33.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error34
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "deleteSingle failed: invalid message type")
		return
	}
	result := THBaseServiceDeleteSingleResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	return
}

// Bulk commit a List of TDeletes to the table.
//
// Throws a TIOError if any of the deletes fail.
//
// Always returns an empty list for backwards compatibility.
//
// Parameters:
//  - Table: the table to delete from
//  - Tdeletes: list of TDeletes to delete
func (p *THBaseServiceClient) DeleteMultiple(table []byte, tdeletes []*TDelete) (r []*TDelete, err error) {
	if err = p.sendDeleteMultiple(table, tdeletes); err != nil {
		return
	}
	return p.recvDeleteMultiple()
}

func (p *THBaseServiceClient) sendDeleteMultiple(table []byte, tdeletes []*TDelete) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("deleteMultiple", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceDeleteMultipleArgs{
		Table:    table,
		Tdeletes: tdeletes,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvDeleteMultiple() (value []*TDelete, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "deleteMultiple" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "deleteMultiple failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "deleteMultiple failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error35 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error36 error
		error36, err = error35.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error36
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "deleteMultiple failed: invalid message type")
		return
	}
	result := THBaseServiceDeleteMultipleResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Atomically checks if a row/family/qualifier value matches the expected
// value. If it does, it adds the delete.
//
// @return true if the new delete was executed, false otherwise
//
// Parameters:
//  - Table: to check in and delete from
//  - Row: row to check
//  - Family: column family to check
//  - Qualifier: column qualifier to check
//  - Value: the expected value, if not provided the
// check is for the non-existence of the
// column in question
//  - Tdelete: the TDelete to execute if the check succeeds
func (p *THBaseServiceClient) CheckAndDelete(table []byte, row []byte, family []byte, qualifier []byte, value []byte, tdelete *TDelete) (r bool, err error) {
	if err = p.sendCheckAndDelete(table, row, family, qualifier, value, tdelete); err != nil {
		return
	}
	return p.recvCheckAndDelete()
}

func (p *THBaseServiceClient) sendCheckAndDelete(table []byte, row []byte, family []byte, qualifier []byte, value []byte, tdelete *TDelete) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("checkAndDelete", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceCheckAndDeleteArgs{
		Table:     table,
		Row:       row,
		Family:    family,
		Qualifier: qualifier,
		Value:     value,
		Tdelete:   tdelete,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvCheckAndDelete() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "checkAndDelete" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "checkAndDelete failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "checkAndDelete failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error37 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error38 error
		error38, err = error37.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error38
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "checkAndDelete failed: invalid message type")
		return
	}
	result := THBaseServiceCheckAndDeleteResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Table: the table to increment the value on
//  - Tincrement: the TIncrement to increment
func (p *THBaseServiceClient) Increment(table []byte, tincrement *TIncrement) (r *TResult_, err error) {
	if err = p.sendIncrement(table, tincrement); err != nil {
		return
	}
	return p.recvIncrement()
}

func (p *THBaseServiceClient) sendIncrement(table []byte, tincrement *TIncrement) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("increment", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceIncrementArgs{
		Table:      table,
		Tincrement: tincrement,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvIncrement() (value *TResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "increment" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "increment failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "increment failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error39 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error40 error
		error40, err = error39.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error40
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "increment failed: invalid message type")
		return
	}
	result := THBaseServiceIncrementResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Table: the table to append the value on
//  - Tappend: the TAppend to append
func (p *THBaseServiceClient) Append(table []byte, tappend *TAppend) (r *TResult_, err error) {
	if err = p.sendAppend(table, tappend); err != nil {
		return
	}
	return p.recvAppend()
}

func (p *THBaseServiceClient) sendAppend(table []byte, tappend *TAppend) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("append", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceAppendArgs{
		Table:   table,
		Tappend: tappend,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvAppend() (value *TResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "append" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "append failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "append failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error41 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error42 error
		error42, err = error41.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error42
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "append failed: invalid message type")
		return
	}
	result := THBaseServiceAppendResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Get a Scanner for the provided TScan object.
//
// @return Scanner Id to be used with other scanner procedures
//
// Parameters:
//  - Table: the table to get the Scanner for
//  - Tscan: the scan object to get a Scanner for
func (p *THBaseServiceClient) OpenScanner(table []byte, tscan *TScan) (r int32, err error) {
	if err = p.sendOpenScanner(table, tscan); err != nil {
		return
	}
	return p.recvOpenScanner()
}

func (p *THBaseServiceClient) sendOpenScanner(table []byte, tscan *TScan) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("openScanner", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceOpenScannerArgs{
		Table: table,
		Tscan: tscan,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvOpenScanner() (value int32, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "openScanner" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "openScanner failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "openScanner failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error43 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error44 error
		error44, err = error43.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error44
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "openScanner failed: invalid message type")
		return
	}
	result := THBaseServiceOpenScannerResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Grabs multiple rows from a Scanner.
//
// @return Between zero and numRows TResults
//
// Parameters:
//  - ScannerId: the Id of the Scanner to return rows from. This is an Id returned from the openScanner function.
//  - NumRows: number of rows to return
func (p *THBaseServiceClient) GetScannerRows(scannerId int32, numRows int32) (r []*TResult_, err error) {
	if err = p.sendGetScannerRows(scannerId, numRows); err != nil {
		return
	}
	return p.recvGetScannerRows()
}

func (p *THBaseServiceClient) sendGetScannerRows(scannerId int32, numRows int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getScannerRows", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceGetScannerRowsArgs{
		ScannerId: scannerId,
		NumRows:   numRows,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvGetScannerRows() (value []*TResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getScannerRows" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getScannerRows failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getScannerRows failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error45 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error46 error
		error46, err = error45.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error46
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getScannerRows failed: invalid message type")
		return
	}
	result := THBaseServiceGetScannerRowsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	} else if result.Ia != nil {
		err = result.Ia
		return
	}
	value = result.GetSuccess()
	return
}

// Closes the scanner. Should be called to free server side resources timely.
// Typically close once the scanner is not needed anymore, i.e. after looping
// over it to get all the required rows.
//
// Parameters:
//  - ScannerId: the Id of the Scanner to close *
func (p *THBaseServiceClient) CloseScanner(scannerId int32) (err error) {
	if err = p.sendCloseScanner(scannerId); err != nil {
		return
	}
	return p.recvCloseScanner()
}

func (p *THBaseServiceClient) sendCloseScanner(scannerId int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("closeScanner", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceCloseScannerArgs{
		ScannerId: scannerId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvCloseScanner() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "closeScanner" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "closeScanner failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "closeScanner failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error47 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error48 error
		error48, err = error47.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error48
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "closeScanner failed: invalid message type")
		return
	}
	result := THBaseServiceCloseScannerResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	} else if result.Ia != nil {
		err = result.Ia
		return
	}
	return
}

// mutateRow performs multiple mutations atomically on a single row.
//
// Parameters:
//  - Table: table to apply the mutations
//  - TrowMutations: mutations to apply
func (p *THBaseServiceClient) MutateRow(table []byte, trowMutations *TRowMutations) (err error) {
	if err = p.sendMutateRow(table, trowMutations); err != nil {
		return
	}
	return p.recvMutateRow()
}

func (p *THBaseServiceClient) sendMutateRow(table []byte, trowMutations *TRowMutations) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("mutateRow", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceMutateRowArgs{
		Table:         table,
		TrowMutations: trowMutations,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvMutateRow() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "mutateRow" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "mutateRow failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "mutateRow failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error49 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error50 error
		error50, err = error49.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error50
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "mutateRow failed: invalid message type")
		return
	}
	result := THBaseServiceMutateRowResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	return
}

// Get results for the provided TScan object.
// This helper function opens a scanner, get the results and close the scanner.
//
// @return between zero and numRows TResults
//
// Parameters:
//  - Table: the table to get the Scanner for
//  - Tscan: the scan object to get a Scanner for
//  - NumRows: number of rows to return
func (p *THBaseServiceClient) GetScannerResults(table []byte, tscan *TScan, numRows int32) (r []*TResult_, err error) {
	if err = p.sendGetScannerResults(table, tscan, numRows); err != nil {
		return
	}
	return p.recvGetScannerResults()
}

func (p *THBaseServiceClient) sendGetScannerResults(table []byte, tscan *TScan, numRows int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getScannerResults", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceGetScannerResultsArgs{
		Table:   table,
		Tscan:   tscan,
		NumRows: numRows,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvGetScannerResults() (value []*TResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getScannerResults" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getScannerResults failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getScannerResults failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error51 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error52 error
		error52, err = error51.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error52
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getScannerResults failed: invalid message type")
		return
	}
	result := THBaseServiceGetScannerResultsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Given a table and a row get the location of the region that
// would contain the given row key.
//
// reload = true means the cache will be cleared and the location
// will be fetched from meta.
//
// Parameters:
//  - Table
//  - Row
//  - Reload
func (p *THBaseServiceClient) GetRegionLocation(table []byte, row []byte, reload bool) (r *THRegionLocation, err error) {
	if err = p.sendGetRegionLocation(table, row, reload); err != nil {
		return
	}
	return p.recvGetRegionLocation()
}

func (p *THBaseServiceClient) sendGetRegionLocation(table []byte, row []byte, reload bool) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getRegionLocation", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceGetRegionLocationArgs{
		Table:  table,
		Row:    row,
		Reload: reload,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvGetRegionLocation() (value *THRegionLocation, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getRegionLocation" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getRegionLocation failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getRegionLocation failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error53 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error54 error
		error54, err = error53.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error54
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getRegionLocation failed: invalid message type")
		return
	}
	result := THBaseServiceGetRegionLocationResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Get all of the region locations for a given table.
//
//
// Parameters:
//  - Table
func (p *THBaseServiceClient) GetAllRegionLocations(table []byte) (r []*THRegionLocation, err error) {
	if err = p.sendGetAllRegionLocations(table); err != nil {
		return
	}
	return p.recvGetAllRegionLocations()
}

func (p *THBaseServiceClient) sendGetAllRegionLocations(table []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getAllRegionLocations", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceGetAllRegionLocationsArgs{
		Table: table,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvGetAllRegionLocations() (value []*THRegionLocation, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getAllRegionLocations" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getAllRegionLocations failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getAllRegionLocations failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error55 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error56 error
		error56, err = error55.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error56
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getAllRegionLocations failed: invalid message type")
		return
	}
	result := THBaseServiceGetAllRegionLocationsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

// Atomically checks if a row/family/qualifier value matches the expected
// value. If it does, it mutates the row.
//
// @return true if the row was mutated, false otherwise
//
// Parameters:
//  - Table: to check in and delete from
//  - Row: row to check
//  - Family: column family to check
//  - Qualifier: column qualifier to check
//  - CompareOp: comparison to make on the value
//  - Value: the expected value to be compared against, if not provided the
// check is for the non-existence of the column in question
//  - RowMutations: row mutations to execute if the value matches
func (p *THBaseServiceClient) CheckAndMutate(table []byte, row []byte, family []byte, qualifier []byte, compareOp TCompareOp, value []byte, rowMutations *TRowMutations) (r bool, err error) {
	if err = p.sendCheckAndMutate(table, row, family, qualifier, compareOp, value, rowMutations); err != nil {
		return
	}
	return p.recvCheckAndMutate()
}

func (p *THBaseServiceClient) sendCheckAndMutate(table []byte, row []byte, family []byte, qualifier []byte, compareOp TCompareOp, value []byte, rowMutations *TRowMutations) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("checkAndMutate", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := THBaseServiceCheckAndMutateArgs{
		Table:        table,
		Row:          row,
		Family:       family,
		Qualifier:    qualifier,
		CompareOp:    compareOp,
		Value:        value,
		RowMutations: rowMutations,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *THBaseServiceClient) recvCheckAndMutate() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "checkAndMutate" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "checkAndMutate failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "checkAndMutate failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error57 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error58 error
		error58, err = error57.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error58
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "checkAndMutate failed: invalid message type")
		return
	}
	result := THBaseServiceCheckAndMutateResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Io != nil {
		err = result.Io
		return
	}
	value = result.GetSuccess()
	return
}

type THBaseServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      THBaseService
}

func (p *THBaseServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *THBaseServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *THBaseServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTHBaseServiceProcessor(handler THBaseService) *THBaseServiceProcessor {

	self59 := &THBaseServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self59.processorMap["exists"] = &tHBaseServiceProcessorExists{handler: handler}
	self59.processorMap["get"] = &tHBaseServiceProcessorGet{handler: handler}
	self59.processorMap["getMultiple"] = &tHBaseServiceProcessorGetMultiple{handler: handler}
	self59.processorMap["put"] = &tHBaseServiceProcessorPut{handler: handler}
	self59.processorMap["checkAndPut"] = &tHBaseServiceProcessorCheckAndPut{handler: handler}
	self59.processorMap["putMultiple"] = &tHBaseServiceProcessorPutMultiple{handler: handler}
	self59.processorMap["deleteSingle"] = &tHBaseServiceProcessorDeleteSingle{handler: handler}
	self59.processorMap["deleteMultiple"] = &tHBaseServiceProcessorDeleteMultiple{handler: handler}
	self59.processorMap["checkAndDelete"] = &tHBaseServiceProcessorCheckAndDelete{handler: handler}
	self59.processorMap["increment"] = &tHBaseServiceProcessorIncrement{handler: handler}
	self59.processorMap["append"] = &tHBaseServiceProcessorAppend{handler: handler}
	self59.processorMap["openScanner"] = &tHBaseServiceProcessorOpenScanner{handler: handler}
	self59.processorMap["getScannerRows"] = &tHBaseServiceProcessorGetScannerRows{handler: handler}
	self59.processorMap["closeScanner"] = &tHBaseServiceProcessorCloseScanner{handler: handler}
	self59.processorMap["mutateRow"] = &tHBaseServiceProcessorMutateRow{handler: handler}
	self59.processorMap["getScannerResults"] = &tHBaseServiceProcessorGetScannerResults{handler: handler}
	self59.processorMap["getRegionLocation"] = &tHBaseServiceProcessorGetRegionLocation{handler: handler}
	self59.processorMap["getAllRegionLocations"] = &tHBaseServiceProcessorGetAllRegionLocations{handler: handler}
	self59.processorMap["checkAndMutate"] = &tHBaseServiceProcessorCheckAndMutate{handler: handler}
	return self59
}

func (p *THBaseServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x60 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x60.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x60

}

type tHBaseServiceProcessorExists struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorExists) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceExistsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("exists", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceExistsResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.Exists(args.Table, args.Tget); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing exists: "+err2.Error())
			oprot.WriteMessageBegin("exists", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("exists", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorGet struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorGet) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceGetArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceGetResult{}
	var retval *TResult_
	var err2 error
	if retval, err2 = p.handler.Get(args.Table, args.Tget); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get: "+err2.Error())
			oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("get", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorGetMultiple struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorGetMultiple) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceGetMultipleArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getMultiple", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceGetMultipleResult{}
	var retval []*TResult_
	var err2 error
	if retval, err2 = p.handler.GetMultiple(args.Table, args.Tgets); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getMultiple: "+err2.Error())
			oprot.WriteMessageBegin("getMultiple", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getMultiple", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorPut struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorPut) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServicePutArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("put", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServicePutResult{}
	var err2 error
	if err2 = p.handler.Put(args.Table, args.Tput); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing put: "+err2.Error())
			oprot.WriteMessageBegin("put", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("put", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorCheckAndPut struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorCheckAndPut) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceCheckAndPutArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("checkAndPut", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceCheckAndPutResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.CheckAndPut(args.Table, args.Row, args.Family, args.Qualifier, args.Value, args.Tput); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing checkAndPut: "+err2.Error())
			oprot.WriteMessageBegin("checkAndPut", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("checkAndPut", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorPutMultiple struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorPutMultiple) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServicePutMultipleArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("putMultiple", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServicePutMultipleResult{}
	var err2 error
	if err2 = p.handler.PutMultiple(args.Table, args.Tputs); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing putMultiple: "+err2.Error())
			oprot.WriteMessageBegin("putMultiple", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("putMultiple", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorDeleteSingle struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorDeleteSingle) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceDeleteSingleArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("deleteSingle", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceDeleteSingleResult{}
	var err2 error
	if err2 = p.handler.DeleteSingle(args.Table, args.Tdelete); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteSingle: "+err2.Error())
			oprot.WriteMessageBegin("deleteSingle", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("deleteSingle", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorDeleteMultiple struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorDeleteMultiple) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceDeleteMultipleArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("deleteMultiple", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceDeleteMultipleResult{}
	var retval []*TDelete
	var err2 error
	if retval, err2 = p.handler.DeleteMultiple(args.Table, args.Tdeletes); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteMultiple: "+err2.Error())
			oprot.WriteMessageBegin("deleteMultiple", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("deleteMultiple", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorCheckAndDelete struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorCheckAndDelete) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceCheckAndDeleteArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("checkAndDelete", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceCheckAndDeleteResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.CheckAndDelete(args.Table, args.Row, args.Family, args.Qualifier, args.Value, args.Tdelete); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing checkAndDelete: "+err2.Error())
			oprot.WriteMessageBegin("checkAndDelete", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("checkAndDelete", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorIncrement struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorIncrement) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceIncrementArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("increment", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceIncrementResult{}
	var retval *TResult_
	var err2 error
	if retval, err2 = p.handler.Increment(args.Table, args.Tincrement); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing increment: "+err2.Error())
			oprot.WriteMessageBegin("increment", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("increment", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorAppend struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorAppend) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceAppendArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceAppendResult{}
	var retval *TResult_
	var err2 error
	if retval, err2 = p.handler.Append(args.Table, args.Tappend); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing append: "+err2.Error())
			oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("append", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorOpenScanner struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorOpenScanner) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceOpenScannerArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("openScanner", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceOpenScannerResult{}
	var retval int32
	var err2 error
	if retval, err2 = p.handler.OpenScanner(args.Table, args.Tscan); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing openScanner: "+err2.Error())
			oprot.WriteMessageBegin("openScanner", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("openScanner", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorGetScannerRows struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorGetScannerRows) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceGetScannerRowsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getScannerRows", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceGetScannerRowsResult{}
	var retval []*TResult_
	var err2 error
	if retval, err2 = p.handler.GetScannerRows(args.ScannerId, args.NumRows); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		case *TIllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getScannerRows: "+err2.Error())
			oprot.WriteMessageBegin("getScannerRows", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getScannerRows", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorCloseScanner struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorCloseScanner) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceCloseScannerArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("closeScanner", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceCloseScannerResult{}
	var err2 error
	if err2 = p.handler.CloseScanner(args.ScannerId); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		case *TIllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing closeScanner: "+err2.Error())
			oprot.WriteMessageBegin("closeScanner", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("closeScanner", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorMutateRow struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorMutateRow) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceMutateRowArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("mutateRow", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceMutateRowResult{}
	var err2 error
	if err2 = p.handler.MutateRow(args.Table, args.TrowMutations); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing mutateRow: "+err2.Error())
			oprot.WriteMessageBegin("mutateRow", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("mutateRow", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorGetScannerResults struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorGetScannerResults) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceGetScannerResultsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getScannerResults", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceGetScannerResultsResult{}
	var retval []*TResult_
	var err2 error
	if retval, err2 = p.handler.GetScannerResults(args.Table, args.Tscan, args.NumRows); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getScannerResults: "+err2.Error())
			oprot.WriteMessageBegin("getScannerResults", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getScannerResults", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorGetRegionLocation struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorGetRegionLocation) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceGetRegionLocationArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRegionLocation", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceGetRegionLocationResult{}
	var retval *THRegionLocation
	var err2 error
	if retval, err2 = p.handler.GetRegionLocation(args.Table, args.Row, args.Reload); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRegionLocation: "+err2.Error())
			oprot.WriteMessageBegin("getRegionLocation", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRegionLocation", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorGetAllRegionLocations struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorGetAllRegionLocations) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceGetAllRegionLocationsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getAllRegionLocations", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceGetAllRegionLocationsResult{}
	var retval []*THRegionLocation
	var err2 error
	if retval, err2 = p.handler.GetAllRegionLocations(args.Table); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getAllRegionLocations: "+err2.Error())
			oprot.WriteMessageBegin("getAllRegionLocations", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getAllRegionLocations", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tHBaseServiceProcessorCheckAndMutate struct {
	handler THBaseService
}

func (p *tHBaseServiceProcessorCheckAndMutate) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := THBaseServiceCheckAndMutateArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("checkAndMutate", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := THBaseServiceCheckAndMutateResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.CheckAndMutate(args.Table, args.Row, args.Family, args.Qualifier, args.CompareOp, args.Value, args.RowMutations); err2 != nil {
		switch v := err2.(type) {
		case *TIOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing checkAndMutate: "+err2.Error())
			oprot.WriteMessageBegin("checkAndMutate", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("checkAndMutate", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Table: the table to check on
//  - Tget: the TGet to check for
type THBaseServiceExistsArgs struct {
	Table []byte `thrift:"table,1,required" json:"table"`
	Tget  *TGet  `thrift:"tget,2,required" json:"tget"`
}

func NewTHBaseServiceExistsArgs() *THBaseServiceExistsArgs {
	return &THBaseServiceExistsArgs{}
}

func (p *THBaseServiceExistsArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceExistsArgs_Tget_DEFAULT *TGet

func (p *THBaseServiceExistsArgs) GetTget() *TGet {
	if !p.IsSetTget() {
		return THBaseServiceExistsArgs_Tget_DEFAULT
	}
	return p.Tget
}
func (p *THBaseServiceExistsArgs) IsSetTget() bool {
	return p.Tget != nil
}

func (p *THBaseServiceExistsArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTget bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTget = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTget {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tget is not set"))
	}
	return nil
}

func (p *THBaseServiceExistsArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceExistsArgs) readField2(iprot thrift.TProtocol) error {
	p.Tget = &TGet{}
	if err := p.Tget.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tget), err)
	}
	return nil
}

func (p *THBaseServiceExistsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("exists_args"); err != nil {
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

func (p *THBaseServiceExistsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceExistsArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tget", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tget: ", p), err)
	}
	if err := p.Tget.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tget), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tget: ", p), err)
	}
	return err
}

func (p *THBaseServiceExistsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceExistsArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceExistsResult struct {
	Success *bool     `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceExistsResult() *THBaseServiceExistsResult {
	return &THBaseServiceExistsResult{}
}

var THBaseServiceExistsResult_Success_DEFAULT bool

func (p *THBaseServiceExistsResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return THBaseServiceExistsResult_Success_DEFAULT
	}
	return *p.Success
}

var THBaseServiceExistsResult_Io_DEFAULT *TIOError

func (p *THBaseServiceExistsResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceExistsResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceExistsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceExistsResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceExistsResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceExistsResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *THBaseServiceExistsResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceExistsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("exists_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceExistsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceExistsResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceExistsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceExistsResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to get from
//  - Tget: the TGet to fetch
type THBaseServiceGetArgs struct {
	Table []byte `thrift:"table,1,required" json:"table"`
	Tget  *TGet  `thrift:"tget,2,required" json:"tget"`
}

func NewTHBaseServiceGetArgs() *THBaseServiceGetArgs {
	return &THBaseServiceGetArgs{}
}

func (p *THBaseServiceGetArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceGetArgs_Tget_DEFAULT *TGet

func (p *THBaseServiceGetArgs) GetTget() *TGet {
	if !p.IsSetTget() {
		return THBaseServiceGetArgs_Tget_DEFAULT
	}
	return p.Tget
}
func (p *THBaseServiceGetArgs) IsSetTget() bool {
	return p.Tget != nil
}

func (p *THBaseServiceGetArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTget bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTget = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTget {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tget is not set"))
	}
	return nil
}

func (p *THBaseServiceGetArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceGetArgs) readField2(iprot thrift.TProtocol) error {
	p.Tget = &TGet{}
	if err := p.Tget.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tget), err)
	}
	return nil
}

func (p *THBaseServiceGetArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("get_args"); err != nil {
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

func (p *THBaseServiceGetArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tget", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tget: ", p), err)
	}
	if err := p.Tget.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tget), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tget: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceGetResult struct {
	Success *TResult_ `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceGetResult() *THBaseServiceGetResult {
	return &THBaseServiceGetResult{}
}

var THBaseServiceGetResult_Success_DEFAULT *TResult_

func (p *THBaseServiceGetResult) GetSuccess() *TResult_ {
	if !p.IsSetSuccess() {
		return THBaseServiceGetResult_Success_DEFAULT
	}
	return p.Success
}

var THBaseServiceGetResult_Io_DEFAULT *TIOError

func (p *THBaseServiceGetResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceGetResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceGetResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceGetResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceGetResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceGetResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &TResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *THBaseServiceGetResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceGetResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("get_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceGetResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to get from
//  - Tgets: a list of TGets to fetch, the Result list
// will have the Results at corresponding positions
// or null if there was an error
type THBaseServiceGetMultipleArgs struct {
	Table []byte  `thrift:"table,1,required" json:"table"`
	Tgets []*TGet `thrift:"tgets,2,required" json:"tgets"`
}

func NewTHBaseServiceGetMultipleArgs() *THBaseServiceGetMultipleArgs {
	return &THBaseServiceGetMultipleArgs{}
}

func (p *THBaseServiceGetMultipleArgs) GetTable() []byte {
	return p.Table
}

func (p *THBaseServiceGetMultipleArgs) GetTgets() []*TGet {
	return p.Tgets
}
func (p *THBaseServiceGetMultipleArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTgets bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTgets = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTgets {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tgets is not set"))
	}
	return nil
}

func (p *THBaseServiceGetMultipleArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceGetMultipleArgs) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TGet, 0, size)
	p.Tgets = tSlice
	for i := 0; i < size; i++ {
		_elem61 := &TGet{}
		if err := _elem61.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem61), err)
		}
		p.Tgets = append(p.Tgets, _elem61)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServiceGetMultipleArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getMultiple_args"); err != nil {
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

func (p *THBaseServiceGetMultipleArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetMultipleArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tgets", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tgets: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Tgets)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Tgets {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tgets: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetMultipleArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetMultipleArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceGetMultipleResult struct {
	Success []*TResult_ `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError   `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceGetMultipleResult() *THBaseServiceGetMultipleResult {
	return &THBaseServiceGetMultipleResult{}
}

var THBaseServiceGetMultipleResult_Success_DEFAULT []*TResult_

func (p *THBaseServiceGetMultipleResult) GetSuccess() []*TResult_ {
	return p.Success
}

var THBaseServiceGetMultipleResult_Io_DEFAULT *TIOError

func (p *THBaseServiceGetMultipleResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceGetMultipleResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceGetMultipleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceGetMultipleResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceGetMultipleResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceGetMultipleResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TResult_, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem62 := &TResult_{}
		if err := _elem62.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem62), err)
		}
		p.Success = append(p.Success, _elem62)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServiceGetMultipleResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceGetMultipleResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getMultiple_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceGetMultipleResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetMultipleResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetMultipleResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetMultipleResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to put data in
//  - Tput: the TPut to put
type THBaseServicePutArgs struct {
	Table []byte `thrift:"table,1,required" json:"table"`
	Tput  *TPut  `thrift:"tput,2,required" json:"tput"`
}

func NewTHBaseServicePutArgs() *THBaseServicePutArgs {
	return &THBaseServicePutArgs{}
}

func (p *THBaseServicePutArgs) GetTable() []byte {
	return p.Table
}

var THBaseServicePutArgs_Tput_DEFAULT *TPut

func (p *THBaseServicePutArgs) GetTput() *TPut {
	if !p.IsSetTput() {
		return THBaseServicePutArgs_Tput_DEFAULT
	}
	return p.Tput
}
func (p *THBaseServicePutArgs) IsSetTput() bool {
	return p.Tput != nil
}

func (p *THBaseServicePutArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTput bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTput = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTput {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tput is not set"))
	}
	return nil
}

func (p *THBaseServicePutArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServicePutArgs) readField2(iprot thrift.TProtocol) error {
	p.Tput = &TPut{}
	if err := p.Tput.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tput), err)
	}
	return nil
}

func (p *THBaseServicePutArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("put_args"); err != nil {
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

func (p *THBaseServicePutArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServicePutArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tput", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tput: ", p), err)
	}
	if err := p.Tput.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tput), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tput: ", p), err)
	}
	return err
}

func (p *THBaseServicePutArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServicePutArgs(%+v)", *p)
}

// Attributes:
//  - Io
type THBaseServicePutResult struct {
	Io *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServicePutResult() *THBaseServicePutResult {
	return &THBaseServicePutResult{}
}

var THBaseServicePutResult_Io_DEFAULT *TIOError

func (p *THBaseServicePutResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServicePutResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServicePutResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServicePutResult) Read(iprot thrift.TProtocol) error {
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

func (p *THBaseServicePutResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServicePutResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("put_result"); err != nil {
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

func (p *THBaseServicePutResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServicePutResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServicePutResult(%+v)", *p)
}

// Attributes:
//  - Table: to check in and put to
//  - Row: row to check
//  - Family: column family to check
//  - Qualifier: column qualifier to check
//  - Value: the expected value, if not provided the
// check is for the non-existence of the
// column in question
//  - Tput: the TPut to put if the check succeeds
type THBaseServiceCheckAndPutArgs struct {
	Table     []byte `thrift:"table,1,required" json:"table"`
	Row       []byte `thrift:"row,2,required" json:"row"`
	Family    []byte `thrift:"family,3,required" json:"family"`
	Qualifier []byte `thrift:"qualifier,4,required" json:"qualifier"`
	Value     []byte `thrift:"value,5" json:"value"`
	Tput      *TPut  `thrift:"tput,6,required" json:"tput"`
}

func NewTHBaseServiceCheckAndPutArgs() *THBaseServiceCheckAndPutArgs {
	return &THBaseServiceCheckAndPutArgs{}
}

func (p *THBaseServiceCheckAndPutArgs) GetTable() []byte {
	return p.Table
}

func (p *THBaseServiceCheckAndPutArgs) GetRow() []byte {
	return p.Row
}

func (p *THBaseServiceCheckAndPutArgs) GetFamily() []byte {
	return p.Family
}

func (p *THBaseServiceCheckAndPutArgs) GetQualifier() []byte {
	return p.Qualifier
}

func (p *THBaseServiceCheckAndPutArgs) GetValue() []byte {
	return p.Value
}

var THBaseServiceCheckAndPutArgs_Tput_DEFAULT *TPut

func (p *THBaseServiceCheckAndPutArgs) GetTput() *TPut {
	if !p.IsSetTput() {
		return THBaseServiceCheckAndPutArgs_Tput_DEFAULT
	}
	return p.Tput
}
func (p *THBaseServiceCheckAndPutArgs) IsSetTput() bool {
	return p.Tput != nil
}

func (p *THBaseServiceCheckAndPutArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetRow bool = false
	var issetFamily bool = false
	var issetQualifier bool = false
	var issetTput bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetRow = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetFamily = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
			issetQualifier = true
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
			issetTput = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	if !issetFamily {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Family is not set"))
	}
	if !issetQualifier {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Qualifier is not set"))
	}
	if !issetTput {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tput is not set"))
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Family = v
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Qualifier = v
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) readField6(iprot thrift.TProtocol) error {
	p.Tput = &TPut{}
	if err := p.Tput.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tput), err)
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("checkAndPut_args"); err != nil {
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
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *THBaseServiceCheckAndPutArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndPutArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:row: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndPutArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("family", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:family: ", p), err)
	}
	if err := oprot.WriteBinary(p.Family); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.family (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:family: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndPutArgs) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("qualifier", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:qualifier: ", p), err)
	}
	if err := oprot.WriteBinary(p.Qualifier); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.qualifier (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:qualifier: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndPutArgs) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:value: ", p), err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:value: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndPutArgs) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tput", thrift.STRUCT, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:tput: ", p), err)
	}
	if err := p.Tput.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tput), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:tput: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndPutArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCheckAndPutArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceCheckAndPutResult struct {
	Success *bool     `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceCheckAndPutResult() *THBaseServiceCheckAndPutResult {
	return &THBaseServiceCheckAndPutResult{}
}

var THBaseServiceCheckAndPutResult_Success_DEFAULT bool

func (p *THBaseServiceCheckAndPutResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return THBaseServiceCheckAndPutResult_Success_DEFAULT
	}
	return *p.Success
}

var THBaseServiceCheckAndPutResult_Io_DEFAULT *TIOError

func (p *THBaseServiceCheckAndPutResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceCheckAndPutResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceCheckAndPutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceCheckAndPutResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceCheckAndPutResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceCheckAndPutResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *THBaseServiceCheckAndPutResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceCheckAndPutResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("checkAndPut_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceCheckAndPutResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCheckAndPutResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCheckAndPutResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCheckAndPutResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to put data in
//  - Tputs: a list of TPuts to commit
type THBaseServicePutMultipleArgs struct {
	Table []byte  `thrift:"table,1,required" json:"table"`
	Tputs []*TPut `thrift:"tputs,2,required" json:"tputs"`
}

func NewTHBaseServicePutMultipleArgs() *THBaseServicePutMultipleArgs {
	return &THBaseServicePutMultipleArgs{}
}

func (p *THBaseServicePutMultipleArgs) GetTable() []byte {
	return p.Table
}

func (p *THBaseServicePutMultipleArgs) GetTputs() []*TPut {
	return p.Tputs
}
func (p *THBaseServicePutMultipleArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTputs bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTputs = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTputs {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tputs is not set"))
	}
	return nil
}

func (p *THBaseServicePutMultipleArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServicePutMultipleArgs) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TPut, 0, size)
	p.Tputs = tSlice
	for i := 0; i < size; i++ {
		_elem63 := &TPut{}
		if err := _elem63.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem63), err)
		}
		p.Tputs = append(p.Tputs, _elem63)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServicePutMultipleArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("putMultiple_args"); err != nil {
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

func (p *THBaseServicePutMultipleArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServicePutMultipleArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tputs", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tputs: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Tputs)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Tputs {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tputs: ", p), err)
	}
	return err
}

func (p *THBaseServicePutMultipleArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServicePutMultipleArgs(%+v)", *p)
}

// Attributes:
//  - Io
type THBaseServicePutMultipleResult struct {
	Io *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServicePutMultipleResult() *THBaseServicePutMultipleResult {
	return &THBaseServicePutMultipleResult{}
}

var THBaseServicePutMultipleResult_Io_DEFAULT *TIOError

func (p *THBaseServicePutMultipleResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServicePutMultipleResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServicePutMultipleResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServicePutMultipleResult) Read(iprot thrift.TProtocol) error {
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

func (p *THBaseServicePutMultipleResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServicePutMultipleResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("putMultiple_result"); err != nil {
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

func (p *THBaseServicePutMultipleResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServicePutMultipleResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServicePutMultipleResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to delete from
//  - Tdelete: the TDelete to delete
type THBaseServiceDeleteSingleArgs struct {
	Table   []byte   `thrift:"table,1,required" json:"table"`
	Tdelete *TDelete `thrift:"tdelete,2,required" json:"tdelete"`
}

func NewTHBaseServiceDeleteSingleArgs() *THBaseServiceDeleteSingleArgs {
	return &THBaseServiceDeleteSingleArgs{}
}

func (p *THBaseServiceDeleteSingleArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceDeleteSingleArgs_Tdelete_DEFAULT *TDelete

func (p *THBaseServiceDeleteSingleArgs) GetTdelete() *TDelete {
	if !p.IsSetTdelete() {
		return THBaseServiceDeleteSingleArgs_Tdelete_DEFAULT
	}
	return p.Tdelete
}
func (p *THBaseServiceDeleteSingleArgs) IsSetTdelete() bool {
	return p.Tdelete != nil
}

func (p *THBaseServiceDeleteSingleArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTdelete bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTdelete = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTdelete {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tdelete is not set"))
	}
	return nil
}

func (p *THBaseServiceDeleteSingleArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceDeleteSingleArgs) readField2(iprot thrift.TProtocol) error {
	p.Tdelete = &TDelete{
		DeleteType: 1,
	}
	if err := p.Tdelete.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tdelete), err)
	}
	return nil
}

func (p *THBaseServiceDeleteSingleArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("deleteSingle_args"); err != nil {
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

func (p *THBaseServiceDeleteSingleArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceDeleteSingleArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tdelete", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tdelete: ", p), err)
	}
	if err := p.Tdelete.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tdelete), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tdelete: ", p), err)
	}
	return err
}

func (p *THBaseServiceDeleteSingleArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceDeleteSingleArgs(%+v)", *p)
}

// Attributes:
//  - Io
type THBaseServiceDeleteSingleResult struct {
	Io *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceDeleteSingleResult() *THBaseServiceDeleteSingleResult {
	return &THBaseServiceDeleteSingleResult{}
}

var THBaseServiceDeleteSingleResult_Io_DEFAULT *TIOError

func (p *THBaseServiceDeleteSingleResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceDeleteSingleResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceDeleteSingleResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceDeleteSingleResult) Read(iprot thrift.TProtocol) error {
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

func (p *THBaseServiceDeleteSingleResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceDeleteSingleResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("deleteSingle_result"); err != nil {
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

func (p *THBaseServiceDeleteSingleResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceDeleteSingleResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceDeleteSingleResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to delete from
//  - Tdeletes: list of TDeletes to delete
type THBaseServiceDeleteMultipleArgs struct {
	Table    []byte     `thrift:"table,1,required" json:"table"`
	Tdeletes []*TDelete `thrift:"tdeletes,2,required" json:"tdeletes"`
}

func NewTHBaseServiceDeleteMultipleArgs() *THBaseServiceDeleteMultipleArgs {
	return &THBaseServiceDeleteMultipleArgs{}
}

func (p *THBaseServiceDeleteMultipleArgs) GetTable() []byte {
	return p.Table
}

func (p *THBaseServiceDeleteMultipleArgs) GetTdeletes() []*TDelete {
	return p.Tdeletes
}
func (p *THBaseServiceDeleteMultipleArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTdeletes bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTdeletes = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTdeletes {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tdeletes is not set"))
	}
	return nil
}

func (p *THBaseServiceDeleteMultipleArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceDeleteMultipleArgs) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TDelete, 0, size)
	p.Tdeletes = tSlice
	for i := 0; i < size; i++ {
		_elem64 := &TDelete{
			DeleteType: 1,
		}
		if err := _elem64.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem64), err)
		}
		p.Tdeletes = append(p.Tdeletes, _elem64)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServiceDeleteMultipleArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("deleteMultiple_args"); err != nil {
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

func (p *THBaseServiceDeleteMultipleArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceDeleteMultipleArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tdeletes", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tdeletes: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Tdeletes)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Tdeletes {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tdeletes: ", p), err)
	}
	return err
}

func (p *THBaseServiceDeleteMultipleArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceDeleteMultipleArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceDeleteMultipleResult struct {
	Success []*TDelete `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError  `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceDeleteMultipleResult() *THBaseServiceDeleteMultipleResult {
	return &THBaseServiceDeleteMultipleResult{}
}

var THBaseServiceDeleteMultipleResult_Success_DEFAULT []*TDelete

func (p *THBaseServiceDeleteMultipleResult) GetSuccess() []*TDelete {
	return p.Success
}

var THBaseServiceDeleteMultipleResult_Io_DEFAULT *TIOError

func (p *THBaseServiceDeleteMultipleResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceDeleteMultipleResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceDeleteMultipleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceDeleteMultipleResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceDeleteMultipleResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceDeleteMultipleResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TDelete, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem65 := &TDelete{
			DeleteType: 1,
		}
		if err := _elem65.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem65), err)
		}
		p.Success = append(p.Success, _elem65)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServiceDeleteMultipleResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceDeleteMultipleResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("deleteMultiple_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceDeleteMultipleResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceDeleteMultipleResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceDeleteMultipleResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceDeleteMultipleResult(%+v)", *p)
}

// Attributes:
//  - Table: to check in and delete from
//  - Row: row to check
//  - Family: column family to check
//  - Qualifier: column qualifier to check
//  - Value: the expected value, if not provided the
// check is for the non-existence of the
// column in question
//  - Tdelete: the TDelete to execute if the check succeeds
type THBaseServiceCheckAndDeleteArgs struct {
	Table     []byte   `thrift:"table,1,required" json:"table"`
	Row       []byte   `thrift:"row,2,required" json:"row"`
	Family    []byte   `thrift:"family,3,required" json:"family"`
	Qualifier []byte   `thrift:"qualifier,4,required" json:"qualifier"`
	Value     []byte   `thrift:"value,5" json:"value"`
	Tdelete   *TDelete `thrift:"tdelete,6,required" json:"tdelete"`
}

func NewTHBaseServiceCheckAndDeleteArgs() *THBaseServiceCheckAndDeleteArgs {
	return &THBaseServiceCheckAndDeleteArgs{}
}

func (p *THBaseServiceCheckAndDeleteArgs) GetTable() []byte {
	return p.Table
}

func (p *THBaseServiceCheckAndDeleteArgs) GetRow() []byte {
	return p.Row
}

func (p *THBaseServiceCheckAndDeleteArgs) GetFamily() []byte {
	return p.Family
}

func (p *THBaseServiceCheckAndDeleteArgs) GetQualifier() []byte {
	return p.Qualifier
}

func (p *THBaseServiceCheckAndDeleteArgs) GetValue() []byte {
	return p.Value
}

var THBaseServiceCheckAndDeleteArgs_Tdelete_DEFAULT *TDelete

func (p *THBaseServiceCheckAndDeleteArgs) GetTdelete() *TDelete {
	if !p.IsSetTdelete() {
		return THBaseServiceCheckAndDeleteArgs_Tdelete_DEFAULT
	}
	return p.Tdelete
}
func (p *THBaseServiceCheckAndDeleteArgs) IsSetTdelete() bool {
	return p.Tdelete != nil
}

func (p *THBaseServiceCheckAndDeleteArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetRow bool = false
	var issetFamily bool = false
	var issetQualifier bool = false
	var issetTdelete bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetRow = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetFamily = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
			issetQualifier = true
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
			issetTdelete = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	if !issetFamily {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Family is not set"))
	}
	if !issetQualifier {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Qualifier is not set"))
	}
	if !issetTdelete {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tdelete is not set"))
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Family = v
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Qualifier = v
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) readField6(iprot thrift.TProtocol) error {
	p.Tdelete = &TDelete{
		DeleteType: 1,
	}
	if err := p.Tdelete.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tdelete), err)
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("checkAndDelete_args"); err != nil {
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
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:row: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("family", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:family: ", p), err)
	}
	if err := oprot.WriteBinary(p.Family); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.family (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:family: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteArgs) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("qualifier", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:qualifier: ", p), err)
	}
	if err := oprot.WriteBinary(p.Qualifier); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.qualifier (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:qualifier: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteArgs) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:value: ", p), err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:value: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteArgs) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tdelete", thrift.STRUCT, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:tdelete: ", p), err)
	}
	if err := p.Tdelete.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tdelete), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:tdelete: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCheckAndDeleteArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceCheckAndDeleteResult struct {
	Success *bool     `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceCheckAndDeleteResult() *THBaseServiceCheckAndDeleteResult {
	return &THBaseServiceCheckAndDeleteResult{}
}

var THBaseServiceCheckAndDeleteResult_Success_DEFAULT bool

func (p *THBaseServiceCheckAndDeleteResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return THBaseServiceCheckAndDeleteResult_Success_DEFAULT
	}
	return *p.Success
}

var THBaseServiceCheckAndDeleteResult_Io_DEFAULT *TIOError

func (p *THBaseServiceCheckAndDeleteResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceCheckAndDeleteResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceCheckAndDeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceCheckAndDeleteResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceCheckAndDeleteResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceCheckAndDeleteResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceCheckAndDeleteResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("checkAndDelete_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceCheckAndDeleteResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCheckAndDeleteResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCheckAndDeleteResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to increment the value on
//  - Tincrement: the TIncrement to increment
type THBaseServiceIncrementArgs struct {
	Table      []byte      `thrift:"table,1,required" json:"table"`
	Tincrement *TIncrement `thrift:"tincrement,2,required" json:"tincrement"`
}

func NewTHBaseServiceIncrementArgs() *THBaseServiceIncrementArgs {
	return &THBaseServiceIncrementArgs{}
}

func (p *THBaseServiceIncrementArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceIncrementArgs_Tincrement_DEFAULT *TIncrement

func (p *THBaseServiceIncrementArgs) GetTincrement() *TIncrement {
	if !p.IsSetTincrement() {
		return THBaseServiceIncrementArgs_Tincrement_DEFAULT
	}
	return p.Tincrement
}
func (p *THBaseServiceIncrementArgs) IsSetTincrement() bool {
	return p.Tincrement != nil
}

func (p *THBaseServiceIncrementArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTincrement bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTincrement = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTincrement {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tincrement is not set"))
	}
	return nil
}

func (p *THBaseServiceIncrementArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceIncrementArgs) readField2(iprot thrift.TProtocol) error {
	p.Tincrement = &TIncrement{}
	if err := p.Tincrement.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tincrement), err)
	}
	return nil
}

func (p *THBaseServiceIncrementArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("increment_args"); err != nil {
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

func (p *THBaseServiceIncrementArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceIncrementArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tincrement", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tincrement: ", p), err)
	}
	if err := p.Tincrement.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tincrement), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tincrement: ", p), err)
	}
	return err
}

func (p *THBaseServiceIncrementArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceIncrementArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceIncrementResult struct {
	Success *TResult_ `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceIncrementResult() *THBaseServiceIncrementResult {
	return &THBaseServiceIncrementResult{}
}

var THBaseServiceIncrementResult_Success_DEFAULT *TResult_

func (p *THBaseServiceIncrementResult) GetSuccess() *TResult_ {
	if !p.IsSetSuccess() {
		return THBaseServiceIncrementResult_Success_DEFAULT
	}
	return p.Success
}

var THBaseServiceIncrementResult_Io_DEFAULT *TIOError

func (p *THBaseServiceIncrementResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceIncrementResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceIncrementResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceIncrementResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceIncrementResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceIncrementResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &TResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *THBaseServiceIncrementResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceIncrementResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("increment_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceIncrementResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceIncrementResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceIncrementResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceIncrementResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to append the value on
//  - Tappend: the TAppend to append
type THBaseServiceAppendArgs struct {
	Table   []byte   `thrift:"table,1,required" json:"table"`
	Tappend *TAppend `thrift:"tappend,2,required" json:"tappend"`
}

func NewTHBaseServiceAppendArgs() *THBaseServiceAppendArgs {
	return &THBaseServiceAppendArgs{}
}

func (p *THBaseServiceAppendArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceAppendArgs_Tappend_DEFAULT *TAppend

func (p *THBaseServiceAppendArgs) GetTappend() *TAppend {
	if !p.IsSetTappend() {
		return THBaseServiceAppendArgs_Tappend_DEFAULT
	}
	return p.Tappend
}
func (p *THBaseServiceAppendArgs) IsSetTappend() bool {
	return p.Tappend != nil
}

func (p *THBaseServiceAppendArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTappend bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTappend = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTappend {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tappend is not set"))
	}
	return nil
}

func (p *THBaseServiceAppendArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceAppendArgs) readField2(iprot thrift.TProtocol) error {
	p.Tappend = &TAppend{}
	if err := p.Tappend.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tappend), err)
	}
	return nil
}

func (p *THBaseServiceAppendArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("append_args"); err != nil {
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

func (p *THBaseServiceAppendArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceAppendArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tappend", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tappend: ", p), err)
	}
	if err := p.Tappend.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tappend), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tappend: ", p), err)
	}
	return err
}

func (p *THBaseServiceAppendArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceAppendArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceAppendResult struct {
	Success *TResult_ `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceAppendResult() *THBaseServiceAppendResult {
	return &THBaseServiceAppendResult{}
}

var THBaseServiceAppendResult_Success_DEFAULT *TResult_

func (p *THBaseServiceAppendResult) GetSuccess() *TResult_ {
	if !p.IsSetSuccess() {
		return THBaseServiceAppendResult_Success_DEFAULT
	}
	return p.Success
}

var THBaseServiceAppendResult_Io_DEFAULT *TIOError

func (p *THBaseServiceAppendResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceAppendResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceAppendResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceAppendResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceAppendResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceAppendResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &TResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *THBaseServiceAppendResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceAppendResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("append_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceAppendResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceAppendResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceAppendResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceAppendResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to get the Scanner for
//  - Tscan: the scan object to get a Scanner for
type THBaseServiceOpenScannerArgs struct {
	Table []byte `thrift:"table,1,required" json:"table"`
	Tscan *TScan `thrift:"tscan,2,required" json:"tscan"`
}

func NewTHBaseServiceOpenScannerArgs() *THBaseServiceOpenScannerArgs {
	return &THBaseServiceOpenScannerArgs{}
}

func (p *THBaseServiceOpenScannerArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceOpenScannerArgs_Tscan_DEFAULT *TScan

func (p *THBaseServiceOpenScannerArgs) GetTscan() *TScan {
	if !p.IsSetTscan() {
		return THBaseServiceOpenScannerArgs_Tscan_DEFAULT
	}
	return p.Tscan
}
func (p *THBaseServiceOpenScannerArgs) IsSetTscan() bool {
	return p.Tscan != nil
}

func (p *THBaseServiceOpenScannerArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTscan bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTscan = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTscan {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tscan is not set"))
	}
	return nil
}

func (p *THBaseServiceOpenScannerArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceOpenScannerArgs) readField2(iprot thrift.TProtocol) error {
	p.Tscan = &TScan{
		MaxVersions: 1,
	}
	if err := p.Tscan.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tscan), err)
	}
	return nil
}

func (p *THBaseServiceOpenScannerArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("openScanner_args"); err != nil {
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

func (p *THBaseServiceOpenScannerArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceOpenScannerArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tscan", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tscan: ", p), err)
	}
	if err := p.Tscan.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tscan), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tscan: ", p), err)
	}
	return err
}

func (p *THBaseServiceOpenScannerArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceOpenScannerArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceOpenScannerResult struct {
	Success *int32    `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceOpenScannerResult() *THBaseServiceOpenScannerResult {
	return &THBaseServiceOpenScannerResult{}
}

var THBaseServiceOpenScannerResult_Success_DEFAULT int32

func (p *THBaseServiceOpenScannerResult) GetSuccess() int32 {
	if !p.IsSetSuccess() {
		return THBaseServiceOpenScannerResult_Success_DEFAULT
	}
	return *p.Success
}

var THBaseServiceOpenScannerResult_Io_DEFAULT *TIOError

func (p *THBaseServiceOpenScannerResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceOpenScannerResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceOpenScannerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceOpenScannerResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceOpenScannerResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceOpenScannerResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *THBaseServiceOpenScannerResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceOpenScannerResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("openScanner_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceOpenScannerResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceOpenScannerResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceOpenScannerResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceOpenScannerResult(%+v)", *p)
}

// Attributes:
//  - ScannerId: the Id of the Scanner to return rows from. This is an Id returned from the openScanner function.
//  - NumRows: number of rows to return
type THBaseServiceGetScannerRowsArgs struct {
	ScannerId int32 `thrift:"scannerId,1,required" json:"scannerId"`
	NumRows   int32 `thrift:"numRows,2" json:"numRows"`
}

func NewTHBaseServiceGetScannerRowsArgs() *THBaseServiceGetScannerRowsArgs {
	return &THBaseServiceGetScannerRowsArgs{
		NumRows: 1,
	}
}

func (p *THBaseServiceGetScannerRowsArgs) GetScannerId() int32 {
	return p.ScannerId
}

func (p *THBaseServiceGetScannerRowsArgs) GetNumRows() int32 {
	return p.NumRows
}
func (p *THBaseServiceGetScannerRowsArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetScannerId bool = false

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
			issetScannerId = true
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
	if !issetScannerId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ScannerId is not set"))
	}
	return nil
}

func (p *THBaseServiceGetScannerRowsArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ScannerId = v
	}
	return nil
}

func (p *THBaseServiceGetScannerRowsArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.NumRows = v
	}
	return nil
}

func (p *THBaseServiceGetScannerRowsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getScannerRows_args"); err != nil {
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

func (p *THBaseServiceGetScannerRowsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("scannerId", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:scannerId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ScannerId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.scannerId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:scannerId: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetScannerRowsArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("numRows", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:numRows: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.NumRows)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.numRows (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:numRows: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetScannerRowsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetScannerRowsArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
//  - Ia: if the scannerId is invalid
type THBaseServiceGetScannerRowsResult struct {
	Success []*TResult_       `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError         `thrift:"io,1" json:"io,omitempty"`
	Ia      *TIllegalArgument `thrift:"ia,2" json:"ia,omitempty"`
}

func NewTHBaseServiceGetScannerRowsResult() *THBaseServiceGetScannerRowsResult {
	return &THBaseServiceGetScannerRowsResult{}
}

var THBaseServiceGetScannerRowsResult_Success_DEFAULT []*TResult_

func (p *THBaseServiceGetScannerRowsResult) GetSuccess() []*TResult_ {
	return p.Success
}

var THBaseServiceGetScannerRowsResult_Io_DEFAULT *TIOError

func (p *THBaseServiceGetScannerRowsResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceGetScannerRowsResult_Io_DEFAULT
	}
	return p.Io
}

var THBaseServiceGetScannerRowsResult_Ia_DEFAULT *TIllegalArgument

func (p *THBaseServiceGetScannerRowsResult) GetIa() *TIllegalArgument {
	if !p.IsSetIa() {
		return THBaseServiceGetScannerRowsResult_Ia_DEFAULT
	}
	return p.Ia
}
func (p *THBaseServiceGetScannerRowsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceGetScannerRowsResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceGetScannerRowsResult) IsSetIa() bool {
	return p.Ia != nil
}

func (p *THBaseServiceGetScannerRowsResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceGetScannerRowsResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TResult_, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem66 := &TResult_{}
		if err := _elem66.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem66), err)
		}
		p.Success = append(p.Success, _elem66)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServiceGetScannerRowsResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceGetScannerRowsResult) readField2(iprot thrift.TProtocol) error {
	p.Ia = &TIllegalArgument{}
	if err := p.Ia.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Ia), err)
	}
	return nil
}

func (p *THBaseServiceGetScannerRowsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getScannerRows_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceGetScannerRowsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetScannerRowsResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetScannerRowsResult) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetIa() {
		if err := oprot.WriteFieldBegin("ia", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ia: ", p), err)
		}
		if err := p.Ia.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Ia), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ia: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetScannerRowsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetScannerRowsResult(%+v)", *p)
}

// Attributes:
//  - ScannerId: the Id of the Scanner to close *
type THBaseServiceCloseScannerArgs struct {
	ScannerId int32 `thrift:"scannerId,1,required" json:"scannerId"`
}

func NewTHBaseServiceCloseScannerArgs() *THBaseServiceCloseScannerArgs {
	return &THBaseServiceCloseScannerArgs{}
}

func (p *THBaseServiceCloseScannerArgs) GetScannerId() int32 {
	return p.ScannerId
}
func (p *THBaseServiceCloseScannerArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetScannerId bool = false

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
			issetScannerId = true
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
	if !issetScannerId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ScannerId is not set"))
	}
	return nil
}

func (p *THBaseServiceCloseScannerArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ScannerId = v
	}
	return nil
}

func (p *THBaseServiceCloseScannerArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("closeScanner_args"); err != nil {
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

func (p *THBaseServiceCloseScannerArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("scannerId", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:scannerId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ScannerId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.scannerId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:scannerId: ", p), err)
	}
	return err
}

func (p *THBaseServiceCloseScannerArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCloseScannerArgs(%+v)", *p)
}

// Attributes:
//  - Io
//  - Ia: if the scannerId is invalid
type THBaseServiceCloseScannerResult struct {
	Io *TIOError         `thrift:"io,1" json:"io,omitempty"`
	Ia *TIllegalArgument `thrift:"ia,2" json:"ia,omitempty"`
}

func NewTHBaseServiceCloseScannerResult() *THBaseServiceCloseScannerResult {
	return &THBaseServiceCloseScannerResult{}
}

var THBaseServiceCloseScannerResult_Io_DEFAULT *TIOError

func (p *THBaseServiceCloseScannerResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceCloseScannerResult_Io_DEFAULT
	}
	return p.Io
}

var THBaseServiceCloseScannerResult_Ia_DEFAULT *TIllegalArgument

func (p *THBaseServiceCloseScannerResult) GetIa() *TIllegalArgument {
	if !p.IsSetIa() {
		return THBaseServiceCloseScannerResult_Ia_DEFAULT
	}
	return p.Ia
}
func (p *THBaseServiceCloseScannerResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceCloseScannerResult) IsSetIa() bool {
	return p.Ia != nil
}

func (p *THBaseServiceCloseScannerResult) Read(iprot thrift.TProtocol) error {
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

func (p *THBaseServiceCloseScannerResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceCloseScannerResult) readField2(iprot thrift.TProtocol) error {
	p.Ia = &TIllegalArgument{}
	if err := p.Ia.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Ia), err)
	}
	return nil
}

func (p *THBaseServiceCloseScannerResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("closeScanner_result"); err != nil {
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

func (p *THBaseServiceCloseScannerResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCloseScannerResult) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetIa() {
		if err := oprot.WriteFieldBegin("ia", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ia: ", p), err)
		}
		if err := p.Ia.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Ia), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ia: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCloseScannerResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCloseScannerResult(%+v)", *p)
}

// Attributes:
//  - Table: table to apply the mutations
//  - TrowMutations: mutations to apply
type THBaseServiceMutateRowArgs struct {
	Table         []byte         `thrift:"table,1,required" json:"table"`
	TrowMutations *TRowMutations `thrift:"trowMutations,2,required" json:"trowMutations"`
}

func NewTHBaseServiceMutateRowArgs() *THBaseServiceMutateRowArgs {
	return &THBaseServiceMutateRowArgs{}
}

func (p *THBaseServiceMutateRowArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceMutateRowArgs_TrowMutations_DEFAULT *TRowMutations

func (p *THBaseServiceMutateRowArgs) GetTrowMutations() *TRowMutations {
	if !p.IsSetTrowMutations() {
		return THBaseServiceMutateRowArgs_TrowMutations_DEFAULT
	}
	return p.TrowMutations
}
func (p *THBaseServiceMutateRowArgs) IsSetTrowMutations() bool {
	return p.TrowMutations != nil
}

func (p *THBaseServiceMutateRowArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTrowMutations bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTrowMutations = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTrowMutations {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TrowMutations is not set"))
	}
	return nil
}

func (p *THBaseServiceMutateRowArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceMutateRowArgs) readField2(iprot thrift.TProtocol) error {
	p.TrowMutations = &TRowMutations{}
	if err := p.TrowMutations.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TrowMutations), err)
	}
	return nil
}

func (p *THBaseServiceMutateRowArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("mutateRow_args"); err != nil {
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

func (p *THBaseServiceMutateRowArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceMutateRowArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("trowMutations", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:trowMutations: ", p), err)
	}
	if err := p.TrowMutations.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TrowMutations), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:trowMutations: ", p), err)
	}
	return err
}

func (p *THBaseServiceMutateRowArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceMutateRowArgs(%+v)", *p)
}

// Attributes:
//  - Io
type THBaseServiceMutateRowResult struct {
	Io *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceMutateRowResult() *THBaseServiceMutateRowResult {
	return &THBaseServiceMutateRowResult{}
}

var THBaseServiceMutateRowResult_Io_DEFAULT *TIOError

func (p *THBaseServiceMutateRowResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceMutateRowResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceMutateRowResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceMutateRowResult) Read(iprot thrift.TProtocol) error {
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

func (p *THBaseServiceMutateRowResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceMutateRowResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("mutateRow_result"); err != nil {
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

func (p *THBaseServiceMutateRowResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceMutateRowResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceMutateRowResult(%+v)", *p)
}

// Attributes:
//  - Table: the table to get the Scanner for
//  - Tscan: the scan object to get a Scanner for
//  - NumRows: number of rows to return
type THBaseServiceGetScannerResultsArgs struct {
	Table   []byte `thrift:"table,1,required" json:"table"`
	Tscan   *TScan `thrift:"tscan,2,required" json:"tscan"`
	NumRows int32  `thrift:"numRows,3" json:"numRows"`
}

func NewTHBaseServiceGetScannerResultsArgs() *THBaseServiceGetScannerResultsArgs {
	return &THBaseServiceGetScannerResultsArgs{
		NumRows: 1,
	}
}

func (p *THBaseServiceGetScannerResultsArgs) GetTable() []byte {
	return p.Table
}

var THBaseServiceGetScannerResultsArgs_Tscan_DEFAULT *TScan

func (p *THBaseServiceGetScannerResultsArgs) GetTscan() *TScan {
	if !p.IsSetTscan() {
		return THBaseServiceGetScannerResultsArgs_Tscan_DEFAULT
	}
	return p.Tscan
}

func (p *THBaseServiceGetScannerResultsArgs) GetNumRows() int32 {
	return p.NumRows
}
func (p *THBaseServiceGetScannerResultsArgs) IsSetTscan() bool {
	return p.Tscan != nil
}

func (p *THBaseServiceGetScannerResultsArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetTscan bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTscan = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetTscan {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tscan is not set"))
	}
	return nil
}

func (p *THBaseServiceGetScannerResultsArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceGetScannerResultsArgs) readField2(iprot thrift.TProtocol) error {
	p.Tscan = &TScan{
		MaxVersions: 1,
	}
	if err := p.Tscan.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tscan), err)
	}
	return nil
}

func (p *THBaseServiceGetScannerResultsArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.NumRows = v
	}
	return nil
}

func (p *THBaseServiceGetScannerResultsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getScannerResults_args"); err != nil {
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

func (p *THBaseServiceGetScannerResultsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetScannerResultsArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tscan", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tscan: ", p), err)
	}
	if err := p.Tscan.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tscan), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tscan: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetScannerResultsArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("numRows", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:numRows: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.NumRows)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.numRows (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:numRows: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetScannerResultsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetScannerResultsArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceGetScannerResultsResult struct {
	Success []*TResult_ `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError   `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceGetScannerResultsResult() *THBaseServiceGetScannerResultsResult {
	return &THBaseServiceGetScannerResultsResult{}
}

var THBaseServiceGetScannerResultsResult_Success_DEFAULT []*TResult_

func (p *THBaseServiceGetScannerResultsResult) GetSuccess() []*TResult_ {
	return p.Success
}

var THBaseServiceGetScannerResultsResult_Io_DEFAULT *TIOError

func (p *THBaseServiceGetScannerResultsResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceGetScannerResultsResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceGetScannerResultsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceGetScannerResultsResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceGetScannerResultsResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceGetScannerResultsResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TResult_, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem67 := &TResult_{}
		if err := _elem67.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem67), err)
		}
		p.Success = append(p.Success, _elem67)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServiceGetScannerResultsResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceGetScannerResultsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getScannerResults_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceGetScannerResultsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetScannerResultsResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetScannerResultsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetScannerResultsResult(%+v)", *p)
}

// Attributes:
//  - Table
//  - Row
//  - Reload
type THBaseServiceGetRegionLocationArgs struct {
	Table  []byte `thrift:"table,1,required" json:"table"`
	Row    []byte `thrift:"row,2,required" json:"row"`
	Reload bool   `thrift:"reload,3" json:"reload"`
}

func NewTHBaseServiceGetRegionLocationArgs() *THBaseServiceGetRegionLocationArgs {
	return &THBaseServiceGetRegionLocationArgs{}
}

func (p *THBaseServiceGetRegionLocationArgs) GetTable() []byte {
	return p.Table
}

func (p *THBaseServiceGetRegionLocationArgs) GetRow() []byte {
	return p.Row
}

func (p *THBaseServiceGetRegionLocationArgs) GetReload() bool {
	return p.Reload
}
func (p *THBaseServiceGetRegionLocationArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetRow = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	return nil
}

func (p *THBaseServiceGetRegionLocationArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceGetRegionLocationArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *THBaseServiceGetRegionLocationArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Reload = v
	}
	return nil
}

func (p *THBaseServiceGetRegionLocationArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getRegionLocation_args"); err != nil {
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

func (p *THBaseServiceGetRegionLocationArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetRegionLocationArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:row: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetRegionLocationArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("reload", thrift.BOOL, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:reload: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.Reload)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.reload (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:reload: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetRegionLocationArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetRegionLocationArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceGetRegionLocationResult struct {
	Success *THRegionLocation `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError         `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceGetRegionLocationResult() *THBaseServiceGetRegionLocationResult {
	return &THBaseServiceGetRegionLocationResult{}
}

var THBaseServiceGetRegionLocationResult_Success_DEFAULT *THRegionLocation

func (p *THBaseServiceGetRegionLocationResult) GetSuccess() *THRegionLocation {
	if !p.IsSetSuccess() {
		return THBaseServiceGetRegionLocationResult_Success_DEFAULT
	}
	return p.Success
}

var THBaseServiceGetRegionLocationResult_Io_DEFAULT *TIOError

func (p *THBaseServiceGetRegionLocationResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceGetRegionLocationResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceGetRegionLocationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceGetRegionLocationResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceGetRegionLocationResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceGetRegionLocationResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &THRegionLocation{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *THBaseServiceGetRegionLocationResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceGetRegionLocationResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getRegionLocation_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceGetRegionLocationResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetRegionLocationResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetRegionLocationResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetRegionLocationResult(%+v)", *p)
}

// Attributes:
//  - Table
type THBaseServiceGetAllRegionLocationsArgs struct {
	Table []byte `thrift:"table,1,required" json:"table"`
}

func NewTHBaseServiceGetAllRegionLocationsArgs() *THBaseServiceGetAllRegionLocationsArgs {
	return &THBaseServiceGetAllRegionLocationsArgs{}
}

func (p *THBaseServiceGetAllRegionLocationsArgs) GetTable() []byte {
	return p.Table
}
func (p *THBaseServiceGetAllRegionLocationsArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false

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
			issetTable = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	return nil
}

func (p *THBaseServiceGetAllRegionLocationsArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceGetAllRegionLocationsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllRegionLocations_args"); err != nil {
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

func (p *THBaseServiceGetAllRegionLocationsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceGetAllRegionLocationsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetAllRegionLocationsArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceGetAllRegionLocationsResult struct {
	Success []*THRegionLocation `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError           `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceGetAllRegionLocationsResult() *THBaseServiceGetAllRegionLocationsResult {
	return &THBaseServiceGetAllRegionLocationsResult{}
}

var THBaseServiceGetAllRegionLocationsResult_Success_DEFAULT []*THRegionLocation

func (p *THBaseServiceGetAllRegionLocationsResult) GetSuccess() []*THRegionLocation {
	return p.Success
}

var THBaseServiceGetAllRegionLocationsResult_Io_DEFAULT *TIOError

func (p *THBaseServiceGetAllRegionLocationsResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceGetAllRegionLocationsResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceGetAllRegionLocationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceGetAllRegionLocationsResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceGetAllRegionLocationsResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceGetAllRegionLocationsResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*THRegionLocation, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem68 := &THRegionLocation{}
		if err := _elem68.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem68), err)
		}
		p.Success = append(p.Success, _elem68)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *THBaseServiceGetAllRegionLocationsResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceGetAllRegionLocationsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllRegionLocations_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceGetAllRegionLocationsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetAllRegionLocationsResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceGetAllRegionLocationsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceGetAllRegionLocationsResult(%+v)", *p)
}

// Attributes:
//  - Table: to check in and delete from
//  - Row: row to check
//  - Family: column family to check
//  - Qualifier: column qualifier to check
//  - CompareOp: comparison to make on the value
//  - Value: the expected value to be compared against, if not provided the
// check is for the non-existence of the column in question
//  - RowMutations: row mutations to execute if the value matches
type THBaseServiceCheckAndMutateArgs struct {
	Table        []byte         `thrift:"table,1,required" json:"table"`
	Row          []byte         `thrift:"row,2,required" json:"row"`
	Family       []byte         `thrift:"family,3,required" json:"family"`
	Qualifier    []byte         `thrift:"qualifier,4,required" json:"qualifier"`
	CompareOp    TCompareOp     `thrift:"compareOp,5,required" json:"compareOp"`
	Value        []byte         `thrift:"value,6" json:"value"`
	RowMutations *TRowMutations `thrift:"rowMutations,7,required" json:"rowMutations"`
}

func NewTHBaseServiceCheckAndMutateArgs() *THBaseServiceCheckAndMutateArgs {
	return &THBaseServiceCheckAndMutateArgs{}
}

func (p *THBaseServiceCheckAndMutateArgs) GetTable() []byte {
	return p.Table
}

func (p *THBaseServiceCheckAndMutateArgs) GetRow() []byte {
	return p.Row
}

func (p *THBaseServiceCheckAndMutateArgs) GetFamily() []byte {
	return p.Family
}

func (p *THBaseServiceCheckAndMutateArgs) GetQualifier() []byte {
	return p.Qualifier
}

func (p *THBaseServiceCheckAndMutateArgs) GetCompareOp() TCompareOp {
	return p.CompareOp
}

func (p *THBaseServiceCheckAndMutateArgs) GetValue() []byte {
	return p.Value
}

var THBaseServiceCheckAndMutateArgs_RowMutations_DEFAULT *TRowMutations

func (p *THBaseServiceCheckAndMutateArgs) GetRowMutations() *TRowMutations {
	if !p.IsSetRowMutations() {
		return THBaseServiceCheckAndMutateArgs_RowMutations_DEFAULT
	}
	return p.RowMutations
}
func (p *THBaseServiceCheckAndMutateArgs) IsSetRowMutations() bool {
	return p.RowMutations != nil
}

func (p *THBaseServiceCheckAndMutateArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTable bool = false
	var issetRow bool = false
	var issetFamily bool = false
	var issetQualifier bool = false
	var issetCompareOp bool = false
	var issetRowMutations bool = false

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
			issetTable = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetRow = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetFamily = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
			issetQualifier = true
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
			issetCompareOp = true
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
			issetRowMutations = true
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
	if !issetTable {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Table is not set"))
	}
	if !issetRow {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Row is not set"))
	}
	if !issetFamily {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Family is not set"))
	}
	if !issetQualifier {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Qualifier is not set"))
	}
	if !issetCompareOp {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field CompareOp is not set"))
	}
	if !issetRowMutations {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RowMutations is not set"))
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Family = v
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Qualifier = v
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		temp := TCompareOp(v)
		p.CompareOp = temp
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) readField7(iprot thrift.TProtocol) error {
	p.RowMutations = &TRowMutations{}
	if err := p.RowMutations.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RowMutations), err)
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("checkAndMutate_args"); err != nil {
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

func (p *THBaseServiceCheckAndMutateArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndMutateArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:row: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndMutateArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("family", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:family: ", p), err)
	}
	if err := oprot.WriteBinary(p.Family); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.family (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:family: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndMutateArgs) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("qualifier", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:qualifier: ", p), err)
	}
	if err := oprot.WriteBinary(p.Qualifier); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.qualifier (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:qualifier: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndMutateArgs) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("compareOp", thrift.I32, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:compareOp: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.CompareOp)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.compareOp (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:compareOp: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndMutateArgs) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:value: ", p), err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:value: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndMutateArgs) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("rowMutations", thrift.STRUCT, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:rowMutations: ", p), err)
	}
	if err := p.RowMutations.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RowMutations), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:rowMutations: ", p), err)
	}
	return err
}

func (p *THBaseServiceCheckAndMutateArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCheckAndMutateArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Io
type THBaseServiceCheckAndMutateResult struct {
	Success *bool     `thrift:"success,0" json:"success,omitempty"`
	Io      *TIOError `thrift:"io,1" json:"io,omitempty"`
}

func NewTHBaseServiceCheckAndMutateResult() *THBaseServiceCheckAndMutateResult {
	return &THBaseServiceCheckAndMutateResult{}
}

var THBaseServiceCheckAndMutateResult_Success_DEFAULT bool

func (p *THBaseServiceCheckAndMutateResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return THBaseServiceCheckAndMutateResult_Success_DEFAULT
	}
	return *p.Success
}

var THBaseServiceCheckAndMutateResult_Io_DEFAULT *TIOError

func (p *THBaseServiceCheckAndMutateResult) GetIo() *TIOError {
	if !p.IsSetIo() {
		return THBaseServiceCheckAndMutateResult_Io_DEFAULT
	}
	return p.Io
}
func (p *THBaseServiceCheckAndMutateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *THBaseServiceCheckAndMutateResult) IsSetIo() bool {
	return p.Io != nil
}

func (p *THBaseServiceCheckAndMutateResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *THBaseServiceCheckAndMutateResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateResult) readField1(iprot thrift.TProtocol) error {
	p.Io = &TIOError{}
	if err := p.Io.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Io), err)
	}
	return nil
}

func (p *THBaseServiceCheckAndMutateResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("checkAndMutate_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *THBaseServiceCheckAndMutateResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCheckAndMutateResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetIo() {
		if err := oprot.WriteFieldBegin("io", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:io: ", p), err)
		}
		if err := p.Io.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Io), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:io: ", p), err)
		}
	}
	return err
}

func (p *THBaseServiceCheckAndMutateResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("THBaseServiceCheckAndMutateResult(%+v)", *p)
}
