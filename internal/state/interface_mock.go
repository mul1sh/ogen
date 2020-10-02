// Code generated by MockGen. DO NOT EDIT.
// Source: internal/state/interface.go

// Package state is a generated GoMock package.
package state

import (
	gomock "github.com/golang/mock/gomock"
	bls "github.com/olympus-protocol/ogen/pkg/bls"
	chainhash "github.com/olympus-protocol/ogen/pkg/chainhash"
	primitives "github.com/olympus-protocol/ogen/pkg/primitives"
	reflect "reflect"
)

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// ProcessSlot mocks base method
func (m *MockState) ProcessSlot(previousBlockRoot chainhash.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessSlot", previousBlockRoot)
}

// ProcessSlot indicates an expected call of ProcessSlot
func (mr *MockStateMockRecorder) ProcessSlot(previousBlockRoot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSlot", reflect.TypeOf((*MockState)(nil).ProcessSlot), previousBlockRoot)
}

// ProcessSlots mocks base method
func (m *MockState) ProcessSlots(requestedSlot uint64, view BlockView) ([]*primitives.EpochReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessSlots", requestedSlot, view)
	ret0, _ := ret[0].([]*primitives.EpochReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessSlots indicates an expected call of ProcessSlots
func (mr *MockStateMockRecorder) ProcessSlots(requestedSlot, view interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSlots", reflect.TypeOf((*MockState)(nil).ProcessSlots), requestedSlot, view)
}

// GetEffectiveBalance mocks base method
func (m *MockState) GetEffectiveBalance(index uint64) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEffectiveBalance", index)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetEffectiveBalance indicates an expected call of GetEffectiveBalance
func (mr *MockStateMockRecorder) GetEffectiveBalance(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEffectiveBalance", reflect.TypeOf((*MockState)(nil).GetEffectiveBalance), index)
}

// ActivateValidator mocks base method
func (m *MockState) ActivateValidator(index uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateValidator", index)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivateValidator indicates an expected call of ActivateValidator
func (mr *MockStateMockRecorder) ActivateValidator(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateValidator", reflect.TypeOf((*MockState)(nil).ActivateValidator), index)
}

// InitiateValidatorExit mocks base method
func (m *MockState) InitiateValidatorExit(index uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiateValidatorExit", index)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitiateValidatorExit indicates an expected call of InitiateValidatorExit
func (mr *MockStateMockRecorder) InitiateValidatorExit(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiateValidatorExit", reflect.TypeOf((*MockState)(nil).InitiateValidatorExit), index)
}

// ExitValidator mocks base method
func (m *MockState) ExitValidator(index, status uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExitValidator", index, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExitValidator indicates an expected call of ExitValidator
func (mr *MockStateMockRecorder) ExitValidator(index, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitValidator", reflect.TypeOf((*MockState)(nil).ExitValidator), index, status)
}

// UpdateValidatorStatus mocks base method
func (m *MockState) UpdateValidatorStatus(index, status uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateValidatorStatus", index, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateValidatorStatus indicates an expected call of UpdateValidatorStatus
func (mr *MockStateMockRecorder) UpdateValidatorStatus(index, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorStatus", reflect.TypeOf((*MockState)(nil).UpdateValidatorStatus), index, status)
}

// GetRecentBlockHash mocks base method
func (m *MockState) GetRecentBlockHash(slotToGet uint64) chainhash.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecentBlockHash", slotToGet)
	ret0, _ := ret[0].(chainhash.Hash)
	return ret0
}

// GetRecentBlockHash indicates an expected call of GetRecentBlockHash
func (mr *MockStateMockRecorder) GetRecentBlockHash(slotToGet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentBlockHash", reflect.TypeOf((*MockState)(nil).GetRecentBlockHash), slotToGet)
}

// GetTotalBalances mocks base method
func (m *MockState) GetTotalBalances() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalBalances")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetTotalBalances indicates an expected call of GetTotalBalances
func (mr *MockStateMockRecorder) GetTotalBalances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalBalances", reflect.TypeOf((*MockState)(nil).GetTotalBalances))
}

// NextVoteEpoch mocks base method
func (m *MockState) NextVoteEpoch(newState uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NextVoteEpoch", newState)
}

// NextVoteEpoch indicates an expected call of NextVoteEpoch
func (mr *MockStateMockRecorder) NextVoteEpoch(newState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextVoteEpoch", reflect.TypeOf((*MockState)(nil).NextVoteEpoch), newState)
}

// CheckForVoteTransitions mocks base method
func (m *MockState) CheckForVoteTransitions() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CheckForVoteTransitions")
}

// CheckForVoteTransitions indicates an expected call of CheckForVoteTransitions
func (mr *MockStateMockRecorder) CheckForVoteTransitions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckForVoteTransitions", reflect.TypeOf((*MockState)(nil).CheckForVoteTransitions))
}

// ProcessEpochTransition mocks base method
func (m *MockState) ProcessEpochTransition() ([]*primitives.EpochReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessEpochTransition")
	ret0, _ := ret[0].([]*primitives.EpochReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessEpochTransition indicates an expected call of ProcessEpochTransition
func (mr *MockStateMockRecorder) ProcessEpochTransition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessEpochTransition", reflect.TypeOf((*MockState)(nil).ProcessEpochTransition))
}

// IsGovernanceVoteValid mocks base method
func (m *MockState) IsGovernanceVoteValid(vote *primitives.GovernanceVote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsGovernanceVoteValid", vote)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsGovernanceVoteValid indicates an expected call of IsGovernanceVoteValid
func (mr *MockStateMockRecorder) IsGovernanceVoteValid(vote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsGovernanceVoteValid", reflect.TypeOf((*MockState)(nil).IsGovernanceVoteValid), vote)
}

// ProcessGovernanceVote mocks base method
func (m *MockState) ProcessGovernanceVote(vote *primitives.GovernanceVote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessGovernanceVote", vote)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessGovernanceVote indicates an expected call of ProcessGovernanceVote
func (mr *MockStateMockRecorder) ProcessGovernanceVote(vote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessGovernanceVote", reflect.TypeOf((*MockState)(nil).ProcessGovernanceVote), vote)
}

// ApplyTransactionSingle mocks base method
func (m *MockState) ApplyTransactionSingle(tx *primitives.Tx, blockWithdrawalAddress [20]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTransactionSingle", tx, blockWithdrawalAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyTransactionSingle indicates an expected call of ApplyTransactionSingle
func (mr *MockStateMockRecorder) ApplyTransactionSingle(tx, blockWithdrawalAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTransactionSingle", reflect.TypeOf((*MockState)(nil).ApplyTransactionSingle), tx, blockWithdrawalAddress)
}

// ApplyTransactionMulti mocks base method
func (m *MockState) ApplyTransactionMulti(tx *primitives.TxMulti, blockWithdrawalAddress [20]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTransactionMulti", tx, blockWithdrawalAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyTransactionMulti indicates an expected call of ApplyTransactionMulti
func (mr *MockStateMockRecorder) ApplyTransactionMulti(tx, blockWithdrawalAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTransactionMulti", reflect.TypeOf((*MockState)(nil).ApplyTransactionMulti), tx, blockWithdrawalAddress)
}

// IsProposerSlashingValid mocks base method
func (m *MockState) IsProposerSlashingValid(ps *primitives.ProposerSlashing) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProposerSlashingValid", ps)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsProposerSlashingValid indicates an expected call of IsProposerSlashingValid
func (mr *MockStateMockRecorder) IsProposerSlashingValid(ps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProposerSlashingValid", reflect.TypeOf((*MockState)(nil).IsProposerSlashingValid), ps)
}

// ApplyProposerSlashing mocks base method
func (m *MockState) ApplyProposerSlashing(ps *primitives.ProposerSlashing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyProposerSlashing", ps)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyProposerSlashing indicates an expected call of ApplyProposerSlashing
func (mr *MockStateMockRecorder) ApplyProposerSlashing(ps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyProposerSlashing", reflect.TypeOf((*MockState)(nil).ApplyProposerSlashing), ps)
}

// IsVoteSlashingValid mocks base method
func (m *MockState) IsVoteSlashingValid(vs *primitives.VoteSlashing) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVoteSlashingValid", vs)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsVoteSlashingValid indicates an expected call of IsVoteSlashingValid
func (mr *MockStateMockRecorder) IsVoteSlashingValid(vs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVoteSlashingValid", reflect.TypeOf((*MockState)(nil).IsVoteSlashingValid), vs)
}

// ApplyVoteSlashing mocks base method
func (m *MockState) ApplyVoteSlashing(vs *primitives.VoteSlashing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyVoteSlashing", vs)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyVoteSlashing indicates an expected call of ApplyVoteSlashing
func (mr *MockStateMockRecorder) ApplyVoteSlashing(vs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyVoteSlashing", reflect.TypeOf((*MockState)(nil).ApplyVoteSlashing), vs)
}

// IsRANDAOSlashingValid mocks base method
func (m *MockState) IsRANDAOSlashingValid(rs *primitives.RANDAOSlashing) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRANDAOSlashingValid", rs)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRANDAOSlashingValid indicates an expected call of IsRANDAOSlashingValid
func (mr *MockStateMockRecorder) IsRANDAOSlashingValid(rs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRANDAOSlashingValid", reflect.TypeOf((*MockState)(nil).IsRANDAOSlashingValid), rs)
}

// ApplyRANDAOSlashing mocks base method
func (m *MockState) ApplyRANDAOSlashing(rs *primitives.RANDAOSlashing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyRANDAOSlashing", rs)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyRANDAOSlashing indicates an expected call of ApplyRANDAOSlashing
func (mr *MockStateMockRecorder) ApplyRANDAOSlashing(rs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyRANDAOSlashing", reflect.TypeOf((*MockState)(nil).ApplyRANDAOSlashing), rs)
}

// GetVoteCommittee mocks base method
func (m *MockState) GetVoteCommittee(slot uint64) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVoteCommittee", slot)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVoteCommittee indicates an expected call of GetVoteCommittee
func (mr *MockStateMockRecorder) GetVoteCommittee(slot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVoteCommittee", reflect.TypeOf((*MockState)(nil).GetVoteCommittee), slot)
}

// IsExitValid mocks base method
func (m *MockState) IsExitValid(exit *primitives.Exit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExitValid", exit)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsExitValid indicates an expected call of IsExitValid
func (mr *MockStateMockRecorder) IsExitValid(exit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExitValid", reflect.TypeOf((*MockState)(nil).IsExitValid), exit)
}

// ApplyExit mocks base method
func (m *MockState) ApplyExit(exit *primitives.Exit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyExit", exit)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyExit indicates an expected call of ApplyExit
func (mr *MockStateMockRecorder) ApplyExit(exit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyExit", reflect.TypeOf((*MockState)(nil).ApplyExit), exit)
}

// IsDepositValid mocks base method
func (m *MockState) IsDepositValid(deposit *primitives.Deposit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDepositValid", deposit)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsDepositValid indicates an expected call of IsDepositValid
func (mr *MockStateMockRecorder) IsDepositValid(deposit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDepositValid", reflect.TypeOf((*MockState)(nil).IsDepositValid), deposit)
}

// ApplyDeposit mocks base method
func (m *MockState) ApplyDeposit(deposit *primitives.Deposit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyDeposit", deposit)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyDeposit indicates an expected call of ApplyDeposit
func (mr *MockStateMockRecorder) ApplyDeposit(deposit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyDeposit", reflect.TypeOf((*MockState)(nil).ApplyDeposit), deposit)
}

// IsVoteValid mocks base method
func (m *MockState) IsVoteValid(v *primitives.MultiValidatorVote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVoteValid", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsVoteValid indicates an expected call of IsVoteValid
func (mr *MockStateMockRecorder) IsVoteValid(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVoteValid", reflect.TypeOf((*MockState)(nil).IsVoteValid), v)
}

// ProcessVote mocks base method
func (m *MockState) ProcessVote(v *primitives.MultiValidatorVote, proposerIndex uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessVote", v, proposerIndex)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessVote indicates an expected call of ProcessVote
func (mr *MockStateMockRecorder) ProcessVote(v, proposerIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessVote", reflect.TypeOf((*MockState)(nil).ProcessVote), v, proposerIndex)
}

// GetProposerPublicKey mocks base method
func (m *MockState) GetProposerPublicKey(b *primitives.Block) (*bls.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposerPublicKey", b)
	ret0, _ := ret[0].(*bls.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposerPublicKey indicates an expected call of GetProposerPublicKey
func (mr *MockStateMockRecorder) GetProposerPublicKey(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposerPublicKey", reflect.TypeOf((*MockState)(nil).GetProposerPublicKey), b)
}

// CheckBlockSignature mocks base method
func (m *MockState) CheckBlockSignature(b *primitives.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBlockSignature", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckBlockSignature indicates an expected call of CheckBlockSignature
func (mr *MockStateMockRecorder) CheckBlockSignature(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBlockSignature", reflect.TypeOf((*MockState)(nil).CheckBlockSignature), b)
}

// ProcessBlock mocks base method
func (m *MockState) ProcessBlock(b *primitives.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBlock", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBlock indicates an expected call of ProcessBlock
func (mr *MockStateMockRecorder) ProcessBlock(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBlock", reflect.TypeOf((*MockState)(nil).ProcessBlock), b)
}

// ToSerializable mocks base method
func (m *MockState) ToSerializable() *primitives.SerializableState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToSerializable")
	ret0, _ := ret[0].(*primitives.SerializableState)
	return ret0
}

// ToSerializable indicates an expected call of ToSerializable
func (mr *MockStateMockRecorder) ToSerializable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToSerializable", reflect.TypeOf((*MockState)(nil).ToSerializable))
}

// FromSerializable mocks base method
func (m *MockState) FromSerializable(ser *primitives.SerializableState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FromSerializable", ser)
}

// FromSerializable indicates an expected call of FromSerializable
func (mr *MockStateMockRecorder) FromSerializable(ser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromSerializable", reflect.TypeOf((*MockState)(nil).FromSerializable), ser)
}

// Marshal mocks base method
func (m *MockState) Marshal() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal
func (mr *MockStateMockRecorder) Marshal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockState)(nil).Marshal))
}

// Unmarshal mocks base method
func (m *MockState) Unmarshal(b []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockStateMockRecorder) Unmarshal(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockState)(nil).Unmarshal), b)
}

// GetValidatorIndicesActiveAt mocks base method
func (m *MockState) GetValidatorIndicesActiveAt(epoch uint64) []uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorIndicesActiveAt", epoch)
	ret0, _ := ret[0].([]uint64)
	return ret0
}

// GetValidatorIndicesActiveAt indicates an expected call of GetValidatorIndicesActiveAt
func (mr *MockStateMockRecorder) GetValidatorIndicesActiveAt(epoch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorIndicesActiveAt", reflect.TypeOf((*MockState)(nil).GetValidatorIndicesActiveAt), epoch)
}

// GetValidators mocks base method
func (m *MockState) GetValidators() ValidatorsInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidators")
	ret0, _ := ret[0].(ValidatorsInfo)
	return ret0
}

// GetValidators indicates an expected call of GetValidators
func (mr *MockStateMockRecorder) GetValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidators", reflect.TypeOf((*MockState)(nil).GetValidators))
}

// GetValidatorsForAccount mocks base method
func (m *MockState) GetValidatorsForAccount(acc []byte) ValidatorsInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorsForAccount", acc)
	ret0, _ := ret[0].(ValidatorsInfo)
	return ret0
}

// GetValidatorsForAccount indicates an expected call of GetValidatorsForAccount
func (mr *MockStateMockRecorder) GetValidatorsForAccount(acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorsForAccount", reflect.TypeOf((*MockState)(nil).GetValidatorsForAccount), acc)
}

// Copy mocks base method
func (m *MockState) Copy() State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(State)
	return ret0
}

// Copy indicates an expected call of Copy
func (mr *MockStateMockRecorder) Copy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockState)(nil).Copy))
}

// GetCoinsState mocks base method
func (m *MockState) GetCoinsState() primitives.CoinsState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoinsState")
	ret0, _ := ret[0].(primitives.CoinsState)
	return ret0
}

// GetCoinsState indicates an expected call of GetCoinsState
func (mr *MockStateMockRecorder) GetCoinsState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoinsState", reflect.TypeOf((*MockState)(nil).GetCoinsState))
}

// GetValidatorRegistry mocks base method
func (m *MockState) GetValidatorRegistry() []*primitives.Validator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorRegistry")
	ret0, _ := ret[0].([]*primitives.Validator)
	return ret0
}

// GetValidatorRegistry indicates an expected call of GetValidatorRegistry
func (mr *MockStateMockRecorder) GetValidatorRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorRegistry", reflect.TypeOf((*MockState)(nil).GetValidatorRegistry))
}

// GetProposerQueue mocks base method
func (m *MockState) GetProposerQueue() []uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposerQueue")
	ret0, _ := ret[0].([]uint64)
	return ret0
}

// GetProposerQueue indicates an expected call of GetProposerQueue
func (mr *MockStateMockRecorder) GetProposerQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposerQueue", reflect.TypeOf((*MockState)(nil).GetProposerQueue))
}

// GetSlot mocks base method
func (m *MockState) GetSlot() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlot")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetSlot indicates an expected call of GetSlot
func (mr *MockStateMockRecorder) GetSlot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlot", reflect.TypeOf((*MockState)(nil).GetSlot))
}

// GetEpochIndex mocks base method
func (m *MockState) GetEpochIndex() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpochIndex")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetEpochIndex indicates an expected call of GetEpochIndex
func (mr *MockStateMockRecorder) GetEpochIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpochIndex", reflect.TypeOf((*MockState)(nil).GetEpochIndex))
}

// GetFinalizedEpoch mocks base method
func (m *MockState) GetFinalizedEpoch() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizedEpoch")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetFinalizedEpoch indicates an expected call of GetFinalizedEpoch
func (mr *MockStateMockRecorder) GetFinalizedEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizedEpoch", reflect.TypeOf((*MockState)(nil).GetFinalizedEpoch))
}

// GetJustifiedEpoch mocks base method
func (m *MockState) GetJustifiedEpoch() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJustifiedEpoch")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetJustifiedEpoch indicates an expected call of GetJustifiedEpoch
func (mr *MockStateMockRecorder) GetJustifiedEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJustifiedEpoch", reflect.TypeOf((*MockState)(nil).GetJustifiedEpoch))
}

// GetJustifiedEpochHash mocks base method
func (m *MockState) GetJustifiedEpochHash() chainhash.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJustifiedEpochHash")
	ret0, _ := ret[0].(chainhash.Hash)
	return ret0
}

// GetJustifiedEpochHash indicates an expected call of GetJustifiedEpochHash
func (mr *MockStateMockRecorder) GetJustifiedEpochHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJustifiedEpochHash", reflect.TypeOf((*MockState)(nil).GetJustifiedEpochHash))
}
