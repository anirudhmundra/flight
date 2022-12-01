// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	entities "sahaj/flight/entities"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomValidator is a mock of CustomValidator interface.
type MockCustomValidator struct {
	ctrl     *gomock.Controller
	recorder *MockCustomValidatorMockRecorder
}

// MockCustomValidatorMockRecorder is the mock recorder for MockCustomValidator.
type MockCustomValidatorMockRecorder struct {
	mock *MockCustomValidator
}

// NewMockCustomValidator creates a new mock instance.
func NewMockCustomValidator(ctrl *gomock.Controller) *MockCustomValidator {
	mock := &MockCustomValidator{ctrl: ctrl}
	mock.recorder = &MockCustomValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomValidator) EXPECT() *MockCustomValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockCustomValidator) Validate(ticket entities.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ticket)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockCustomValidatorMockRecorder) Validate(ticket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockCustomValidator)(nil).Validate), ticket)
}
