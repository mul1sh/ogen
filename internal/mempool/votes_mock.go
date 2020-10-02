// Code generated by MockGen. DO NOT EDIT.
// Source: internal/mempool/votes.go

// Package mempool is a generated GoMock package.
package mempool

import (
	gomock "github.com/golang/mock/gomock"
	state "github.com/olympus-protocol/ogen/internal/state"
	primitives "github.com/olympus-protocol/ogen/pkg/primitives"
	reflect "reflect"
)

// MockVoteMempool is a mock of VoteMempool interface
type MockVoteMempool struct {
	ctrl     *gomock.Controller
	recorder *MockVoteMempoolMockRecorder
}

// MockVoteMempoolMockRecorder is the mock recorder for MockVoteMempool
type MockVoteMempoolMockRecorder struct {
	mock *MockVoteMempool
}

// NewMockVoteMempool creates a new mock instance
func NewMockVoteMempool(ctrl *gomock.Controller) *MockVoteMempool {
	mock := &MockVoteMempool{ctrl: ctrl}
	mock.recorder = &MockVoteMempoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVoteMempool) EXPECT() *MockVoteMempoolMockRecorder {
	return m.recorder
}

// AddValidate mocks base method
func (m *MockVoteMempool) AddValidate(vote *primitives.MultiValidatorVote, state state.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddValidate", vote, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddValidate indicates an expected call of AddValidate
func (mr *MockVoteMempoolMockRecorder) AddValidate(vote, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddValidate", reflect.TypeOf((*MockVoteMempool)(nil).AddValidate), vote, state)
}

// Add mocks base method
func (m *MockVoteMempool) Add(vote *primitives.MultiValidatorVote) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", vote)
}

// Add indicates an expected call of Add
func (mr *MockVoteMempoolMockRecorder) Add(vote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockVoteMempool)(nil).Add), vote)
}

// Get mocks base method
func (m *MockVoteMempool) Get(slot uint64, s state.State, proposerIndex uint64) ([]*primitives.MultiValidatorVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", slot, s, proposerIndex)
	ret0, _ := ret[0].([]*primitives.MultiValidatorVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVoteMempoolMockRecorder) Get(slot, s, proposerIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVoteMempool)(nil).Get), slot, s, proposerIndex)
}

// Remove mocks base method
func (m *MockVoteMempool) Remove(b *primitives.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", b)
}

// Remove indicates an expected call of Remove
func (mr *MockVoteMempoolMockRecorder) Remove(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockVoteMempool)(nil).Remove), b)
}

// Notify mocks base method
func (m *MockVoteMempool) Notify(notifee VoteSlashingNotifee) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", notifee)
}

// Notify indicates an expected call of Notify
func (mr *MockVoteMempoolMockRecorder) Notify(notifee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockVoteMempool)(nil).Notify), notifee)
}

// MockVoteSlashingNotifee is a mock of VoteSlashingNotifee interface
type MockVoteSlashingNotifee struct {
	ctrl     *gomock.Controller
	recorder *MockVoteSlashingNotifeeMockRecorder
}

// MockVoteSlashingNotifeeMockRecorder is the mock recorder for MockVoteSlashingNotifee
type MockVoteSlashingNotifeeMockRecorder struct {
	mock *MockVoteSlashingNotifee
}

// NewMockVoteSlashingNotifee creates a new mock instance
func NewMockVoteSlashingNotifee(ctrl *gomock.Controller) *MockVoteSlashingNotifee {
	mock := &MockVoteSlashingNotifee{ctrl: ctrl}
	mock.recorder = &MockVoteSlashingNotifeeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVoteSlashingNotifee) EXPECT() *MockVoteSlashingNotifeeMockRecorder {
	return m.recorder
}

// NotifyIllegalVotes mocks base method
func (m *MockVoteSlashingNotifee) NotifyIllegalVotes(slashing *primitives.VoteSlashing) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyIllegalVotes", slashing)
}

// NotifyIllegalVotes indicates an expected call of NotifyIllegalVotes
func (mr *MockVoteSlashingNotifeeMockRecorder) NotifyIllegalVotes(slashing interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyIllegalVotes", reflect.TypeOf((*MockVoteSlashingNotifee)(nil).NotifyIllegalVotes), slashing)
}
