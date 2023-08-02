// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/mem/v3/cache (interfaces: VictimFinder,Directory)

package cache

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	vm "gitlab.com/akita/mem/v3/vm"
)

// MockVictimFinder is a mock of VictimFinder interface.
type MockVictimFinder struct {
	ctrl     *gomock.Controller
	recorder *MockVictimFinderMockRecorder
}

// MockVictimFinderMockRecorder is the mock recorder for MockVictimFinder.
type MockVictimFinderMockRecorder struct {
	mock *MockVictimFinder
}

// NewMockVictimFinder creates a new mock instance.
func NewMockVictimFinder(ctrl *gomock.Controller) *MockVictimFinder {
	mock := &MockVictimFinder{ctrl: ctrl}
	mock.recorder = &MockVictimFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVictimFinder) EXPECT() *MockVictimFinderMockRecorder {
	return m.recorder
}

// FindVictim mocks base method.
func (m *MockVictimFinder) FindVictim(arg0 *Set) *Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVictim", arg0)
	ret0, _ := ret[0].(*Block)
	return ret0
}

// FindVictim indicates an expected call of FindVictim.
func (mr *MockVictimFinderMockRecorder) FindVictim(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVictim", reflect.TypeOf((*MockVictimFinder)(nil).FindVictim), arg0)
}

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
func (m *MockDirectory) FindVictim(arg0 uint64) *Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVictim", arg0)
	ret0, _ := ret[0].(*Block)
	return ret0
}

// FindVictim indicates an expected call of FindVictim.
func (mr *MockDirectoryMockRecorder) FindVictim(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVictim", reflect.TypeOf((*MockDirectory)(nil).FindVictim), arg0)
}

// GetSets mocks base method.
func (m *MockDirectory) GetSets() []Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSets")
	ret0, _ := ret[0].([]Set)
	return ret0
}

// GetSets indicates an expected call of GetSets.
func (mr *MockDirectoryMockRecorder) GetSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSets", reflect.TypeOf((*MockDirectory)(nil).GetSets))
}

// Lookup mocks base method.
func (m *MockDirectory) Lookup(arg0 vm.PID, arg1 uint64) *Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0, arg1)
	ret0, _ := ret[0].(*Block)
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
func (m *MockDirectory) Visit(arg0 *Block) {
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
