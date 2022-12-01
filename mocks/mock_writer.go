// Code generated by MockGen. DO NOT EDIT.
// Source: writer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	dto "sahaj/flight/dto"

	gomock "github.com/golang/mock/gomock"
)

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockWriter) Write(processedTickets []dto.ValidTicketCSV, failedTickets []dto.InvalidTicketCSV) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", processedTickets, failedTickets)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockWriterMockRecorder) Write(processedTickets, failedTickets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), processedTickets, failedTickets)
}