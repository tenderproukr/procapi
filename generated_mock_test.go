// Code generated by MockGen. DO NOT EDIT.
// Source: procapi.go

// Package procapi is a generated GoMock package.
package procapi

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockResult is a mock of Result interface
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *MockResultMockRecorder
}

// MockResultMockRecorder is the mock recorder for MockResult
type MockResultMockRecorder struct {
	mock *MockResult
}

// NewMockResult creates a new mock instance
func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &MockResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResult) EXPECT() *MockResultMockRecorder {
	return m.recorder
}

// RowsAffected mocks base method
func (m *MockResult) RowsAffected() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowsAffected")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RowsAffected indicates an expected call of RowsAffected
func (mr *MockResultMockRecorder) RowsAffected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowsAffected", reflect.TypeOf((*MockResult)(nil).RowsAffected))
}

// MockRows is a mock of Rows interface
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// StructScan mocks base method
func (m *MockRows) StructScan(dest interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StructScan", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// StructScan indicates an expected call of StructScan
func (mr *MockRowsMockRecorder) StructScan(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StructScan", reflect.TypeOf((*MockRows)(nil).StructScan), dest)
}

// MapScan mocks base method
func (m *MockRows) MapScan(dest map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapScan", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// MapScan indicates an expected call of MapScan
func (mr *MockRowsMockRecorder) MapScan(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapScan", reflect.TypeOf((*MockRows)(nil).MapScan), dest)
}

// Scan mocks base method
func (m *MockRows) Scan(dest ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *MockRowsMockRecorder) Scan(dest ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRows)(nil).Scan), dest...)
}

// Close mocks base method
func (m *MockRows) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRowsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRows)(nil).Close))
}

// Next mocks base method
func (m *MockRows) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockRowsMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRows)(nil).Next))
}

// MockTx is a mock of Tx interface
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx
type MockTxMockRecorder struct {
	mock *MockTx
}

// NewMockTx creates a new mock instance
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &MockTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// Queryx mocks base method
func (m *MockTx) Queryx(query string, args ...interface{}) (Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Queryx", varargs...)
	ret0, _ := ret[0].(Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Queryx indicates an expected call of Queryx
func (mr *MockTxMockRecorder) Queryx(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Queryx", reflect.TypeOf((*MockTx)(nil).Queryx), varargs...)
}

// Exec mocks base method
func (m *MockTx) Exec(query string, args ...interface{}) (Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec
func (mr *MockTxMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockTx)(nil).Exec), varargs...)
}

// Rollback mocks base method
func (m *MockTx) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockTxMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTx)(nil).Rollback))
}

// Commit mocks base method
func (m *MockTx) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockTxMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTx)(nil).Commit))
}

// MockDB is a mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// Beginx mocks base method
func (m *MockDB) Beginx() (Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Beginx")
	ret0, _ := ret[0].(Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Beginx indicates an expected call of Beginx
func (mr *MockDBMockRecorder) Beginx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Beginx", reflect.TypeOf((*MockDB)(nil).Beginx))
}

// Close mocks base method
func (m *MockDB) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDB)(nil).Close))
}

// MockMarshaller is a mock of Marshaller interface
type MockMarshaller struct {
	ctrl     *gomock.Controller
	recorder *MockMarshallerMockRecorder
}

// MockMarshallerMockRecorder is the mock recorder for MockMarshaller
type MockMarshallerMockRecorder struct {
	mock *MockMarshaller
}

// NewMockMarshaller creates a new mock instance
func NewMockMarshaller(ctrl *gomock.Controller) *MockMarshaller {
	mock := &MockMarshaller{ctrl: ctrl}
	mock.recorder = &MockMarshallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarshaller) EXPECT() *MockMarshallerMockRecorder {
	return m.recorder
}

// Marshal mocks base method
func (m *MockMarshaller) Marshal(typ string, v interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal", typ, v)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal
func (mr *MockMarshallerMockRecorder) Marshal(typ, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockMarshaller)(nil).Marshal), typ, v)
}

// Unmarshal mocks base method
func (m *MockMarshaller) Unmarshal(typ string, data interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", typ, data)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockMarshallerMockRecorder) Unmarshal(typ, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockMarshaller)(nil).Unmarshal), typ, data)
}
