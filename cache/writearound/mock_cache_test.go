// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/mem/v3/cache (interfaces: Directory,MSHR)

package writearound

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cache "gitlab.com/akita/mem/v3/cache"
	vm "gitlab.com/akita/mem/v3/vm"
)

// MockDirectory is a mock of Directory interface.
type MockDirectory struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryMockRecorder
}

// MockDirectoryMockRecorder is the mock recorder for MockDirectory.
type MockDirectoryMockRecorder struct {
	mock *MockDirectory
}

// NewMockDirectory creates a new mock instance.
func NewMockDirectory(ctrl *gomock.Controller) *MockDirectory {
	mock := &MockDirectory{ctrl: ctrl}
	mock.recorder = &MockDirectoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectory) EXPECT() *MockDirectoryMockRecorder {
	return m.recorder
}

// FindVictim mocks base method.
func (m *MockDirectory) FindVictim(arg0 uint64) *cache.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVictim", arg0)
	ret0, _ := ret[0].(*cache.Block)
	return ret0
}

// FindVictim indicates an expected call of FindVictim.
func (mr *MockDirectoryMockRecorder) FindVictim(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVictim", reflect.TypeOf((*MockDirectory)(nil).FindVictim), arg0)
}

// GetSets mocks base method.
func (m *MockDirectory) GetSets() []cache.Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSets")
	ret0, _ := ret[0].([]cache.Set)
	return ret0
}

// GetSets indicates an expected call of GetSets.
func (mr *MockDirectoryMockRecorder) GetSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSets", reflect.TypeOf((*MockDirectory)(nil).GetSets))
}

// Lookup mocks base method.
func (m *MockDirectory) Lookup(arg0 vm.PID, arg1 uint64) *cache.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0, arg1)
	ret0, _ := ret[0].(*cache.Block)
	return ret0
}

// Lookup indicates an expected call of Lookup.
func (mr *MockDirectoryMockRecorder) Lookup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockDirectory)(nil).Lookup), arg0, arg1)
}

// Reset mocks base method.
func (m *MockDirectory) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockDirectoryMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockDirectory)(nil).Reset))
}

// TotalSize mocks base method.
func (m *MockDirectory) TotalSize() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalSize")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// TotalSize indicates an expected call of TotalSize.
func (mr *MockDirectoryMockRecorder) TotalSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalSize", reflect.TypeOf((*MockDirectory)(nil).TotalSize))
}

// Visit mocks base method.
func (m *MockDirectory) Visit(arg0 *cache.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Visit", arg0)
}

// Visit indicates an expected call of Visit.
func (mr *MockDirectoryMockRecorder) Visit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Visit", reflect.TypeOf((*MockDirectory)(nil).Visit), arg0)
}

// WayAssociativity mocks base method.
func (m *MockDirectory) WayAssociativity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WayAssociativity")
	ret0, _ := ret[0].(int)
	return ret0
}

// WayAssociativity indicates an expected call of WayAssociativity.
func (mr *MockDirectoryMockRecorder) WayAssociativity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WayAssociativity", reflect.TypeOf((*MockDirectory)(nil).WayAssociativity))
}

// MockMSHR is a mock of MSHR interface.
type MockMSHR struct {
	ctrl     *gomock.Controller
	recorder *MockMSHRMockRecorder
}

// MockMSHRMockRecorder is the mock recorder for MockMSHR.
type MockMSHRMockRecorder struct {
	mock *MockMSHR
}

// NewMockMSHR creates a new mock instance.
func NewMockMSHR(ctrl *gomock.Controller) *MockMSHR {
	mock := &MockMSHR{ctrl: ctrl}
	mock.recorder = &MockMSHRMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMSHR) EXPECT() *MockMSHRMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMSHR) Add(arg0 vm.PID, arg1 uint64) *cache.MSHREntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(*cache.MSHREntry)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMSHRMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMSHR)(nil).Add), arg0, arg1)
}

// AllEntries mocks base method.
func (m *MockMSHR) AllEntries() []*cache.MSHREntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllEntries")
	ret0, _ := ret[0].([]*cache.MSHREntry)
	return ret0
}

// AllEntries indicates an expected call of AllEntries.
func (mr *MockMSHRMockRecorder) AllEntries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllEntries", reflect.TypeOf((*MockMSHR)(nil).AllEntries))
}

// IsFull mocks base method.
func (m *MockMSHR) IsFull() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFull")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFull indicates an expected call of IsFull.
func (mr *MockMSHRMockRecorder) IsFull() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFull", reflect.TypeOf((*MockMSHR)(nil).IsFull))
}

// Query mocks base method.
func (m *MockMSHR) Query(arg0 vm.PID, arg1 uint64) *cache.MSHREntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].(*cache.MSHREntry)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockMSHRMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockMSHR)(nil).Query), arg0, arg1)
}

// Remove mocks base method.
func (m *MockMSHR) Remove(arg0 vm.PID, arg1 uint64) *cache.MSHREntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(*cache.MSHREntry)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockMSHRMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockMSHR)(nil).Remove), arg0, arg1)
}

// Reset mocks base method.
func (m *MockMSHR) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockMSHRMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockMSHR)(nil).Reset))
}
