// Code generated by MockGen. DO NOT EDIT.
// Source: services/transaction_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"
	models "transactions-routine/models"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionServiceI is a mock of TransactionServiceI interface.
type MockTransactionServiceI struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceIMockRecorder
}

// MockTransactionServiceIMockRecorder is the mock recorder for MockTransactionServiceI.
type MockTransactionServiceIMockRecorder struct {
	mock *MockTransactionServiceI
}

// NewMockTransactionServiceI creates a new mock instance.
func NewMockTransactionServiceI(ctrl *gomock.Controller) *MockTransactionServiceI {
	mock := &MockTransactionServiceI{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionServiceI) EXPECT() *MockTransactionServiceIMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionServiceI) CreateTransaction(accountID, operationTypeID int64, amount float64, eventDate time.Time) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", accountID, operationTypeID, amount, eventDate)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionServiceIMockRecorder) CreateTransaction(accountID, operationTypeID, amount, eventDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionServiceI)(nil).CreateTransaction), accountID, operationTypeID, amount, eventDate)
}
