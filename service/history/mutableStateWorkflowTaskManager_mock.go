// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: mutableStateWorkflowTaskManager.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	enums "go.temporal.io/api/enums/v1"
	failure "go.temporal.io/api/failure/v1"
	history "go.temporal.io/api/history/v1"
	workflowservice "go.temporal.io/api/workflowservice/v1"
	reflect "reflect"
)

// MockmutableStateWorkflowTaskManager is a mock of mutableStateWorkflowTaskManager interface.
type MockmutableStateWorkflowTaskManager struct {
	ctrl     *gomock.Controller
	recorder *MockmutableStateWorkflowTaskManagerMockRecorder
}

// MockmutableStateWorkflowTaskManagerMockRecorder is the mock recorder for MockmutableStateWorkflowTaskManager.
type MockmutableStateWorkflowTaskManagerMockRecorder struct {
	mock *MockmutableStateWorkflowTaskManager
}

// NewMockmutableStateWorkflowTaskManager creates a new mock instance.
func NewMockmutableStateWorkflowTaskManager(ctrl *gomock.Controller) *MockmutableStateWorkflowTaskManager {
	mock := &MockmutableStateWorkflowTaskManager{ctrl: ctrl}
	mock.recorder = &MockmutableStateWorkflowTaskManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmutableStateWorkflowTaskManager) EXPECT() *MockmutableStateWorkflowTaskManagerMockRecorder {
	return m.recorder
}

// ReplicateWorkflowTaskScheduledEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) ReplicateWorkflowTaskScheduledEvent(version, scheduleID int64, taskQueue string, startToCloseTimeoutSeconds int32, attempt, scheduleTimestamp, originalScheduledTimestamp int64) (*decisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateWorkflowTaskScheduledEvent", version, scheduleID, taskQueue, startToCloseTimeoutSeconds, attempt, scheduleTimestamp, originalScheduledTimestamp)
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplicateWorkflowTaskScheduledEvent indicates an expected call of ReplicateWorkflowTaskScheduledEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) ReplicateWorkflowTaskScheduledEvent(version, scheduleID, taskQueue, startToCloseTimeoutSeconds, attempt, scheduleTimestamp, originalScheduledTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateWorkflowTaskScheduledEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).ReplicateWorkflowTaskScheduledEvent), version, scheduleID, taskQueue, startToCloseTimeoutSeconds, attempt, scheduleTimestamp, originalScheduledTimestamp)
}

// ReplicateTransientWorkflowTaskScheduled mocks base method.
func (m *MockmutableStateWorkflowTaskManager) ReplicateTransientWorkflowTaskScheduled() (*decisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateTransientWorkflowTaskScheduled")
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplicateTransientWorkflowTaskScheduled indicates an expected call of ReplicateTransientWorkflowTaskScheduled.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) ReplicateTransientWorkflowTaskScheduled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateTransientWorkflowTaskScheduled", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).ReplicateTransientWorkflowTaskScheduled))
}

// ReplicateWorkflowTaskStartedEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) ReplicateWorkflowTaskStartedEvent(decision *decisionInfo, version, scheduleID, startedID int64, requestID string, timestamp int64) (*decisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateWorkflowTaskStartedEvent", decision, version, scheduleID, startedID, requestID, timestamp)
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplicateWorkflowTaskStartedEvent indicates an expected call of ReplicateWorkflowTaskStartedEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) ReplicateWorkflowTaskStartedEvent(decision, version, scheduleID, startedID, requestID, timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateWorkflowTaskStartedEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).ReplicateWorkflowTaskStartedEvent), decision, version, scheduleID, startedID, requestID, timestamp)
}

// ReplicateWorkflowTaskCompletedEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) ReplicateWorkflowTaskCompletedEvent(event *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateWorkflowTaskCompletedEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateWorkflowTaskCompletedEvent indicates an expected call of ReplicateWorkflowTaskCompletedEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) ReplicateWorkflowTaskCompletedEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateWorkflowTaskCompletedEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).ReplicateWorkflowTaskCompletedEvent), event)
}

// ReplicateWorkflowTaskFailedEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) ReplicateWorkflowTaskFailedEvent() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateWorkflowTaskFailedEvent")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateWorkflowTaskFailedEvent indicates an expected call of ReplicateWorkflowTaskFailedEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) ReplicateWorkflowTaskFailedEvent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateWorkflowTaskFailedEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).ReplicateWorkflowTaskFailedEvent))
}

// ReplicateWorkflowTaskTimedOutEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) ReplicateWorkflowTaskTimedOutEvent(timeoutType enums.TimeoutType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateWorkflowTaskTimedOutEvent", timeoutType)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateWorkflowTaskTimedOutEvent indicates an expected call of ReplicateWorkflowTaskTimedOutEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) ReplicateWorkflowTaskTimedOutEvent(timeoutType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateWorkflowTaskTimedOutEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).ReplicateWorkflowTaskTimedOutEvent), timeoutType)
}

// AddWorkflowTaskScheduleToStartTimeoutEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddWorkflowTaskScheduleToStartTimeoutEvent(scheduleEventID int64) (*history.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkflowTaskScheduleToStartTimeoutEvent", scheduleEventID)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWorkflowTaskScheduleToStartTimeoutEvent indicates an expected call of AddWorkflowTaskScheduleToStartTimeoutEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddWorkflowTaskScheduleToStartTimeoutEvent(scheduleEventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkflowTaskScheduleToStartTimeoutEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddWorkflowTaskScheduleToStartTimeoutEvent), scheduleEventID)
}

// AddWorkflowTaskScheduledEventAsHeartbeat mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddWorkflowTaskScheduledEventAsHeartbeat(bypassTaskGeneration bool, originalScheduledTimestamp int64) (*decisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkflowTaskScheduledEventAsHeartbeat", bypassTaskGeneration, originalScheduledTimestamp)
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWorkflowTaskScheduledEventAsHeartbeat indicates an expected call of AddWorkflowTaskScheduledEventAsHeartbeat.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddWorkflowTaskScheduledEventAsHeartbeat(bypassTaskGeneration, originalScheduledTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkflowTaskScheduledEventAsHeartbeat", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddWorkflowTaskScheduledEventAsHeartbeat), bypassTaskGeneration, originalScheduledTimestamp)
}

// AddWorkflowTaskScheduledEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddWorkflowTaskScheduledEvent(bypassTaskGeneration bool) (*decisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkflowTaskScheduledEvent", bypassTaskGeneration)
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWorkflowTaskScheduledEvent indicates an expected call of AddWorkflowTaskScheduledEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddWorkflowTaskScheduledEvent(bypassTaskGeneration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkflowTaskScheduledEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddWorkflowTaskScheduledEvent), bypassTaskGeneration)
}

// AddFirstWorkflowTaskScheduled mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddFirstWorkflowTaskScheduled(startEvent *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFirstWorkflowTaskScheduled", startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFirstWorkflowTaskScheduled indicates an expected call of AddFirstWorkflowTaskScheduled.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddFirstWorkflowTaskScheduled(startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFirstWorkflowTaskScheduled", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddFirstWorkflowTaskScheduled), startEvent)
}

// AddWorkflowTaskStartedEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddWorkflowTaskStartedEvent(scheduleEventID int64, requestID string, request *workflowservice.PollWorkflowTaskQueueRequest) (*history.HistoryEvent, *decisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkflowTaskStartedEvent", scheduleEventID, requestID, request)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(*decisionInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddWorkflowTaskStartedEvent indicates an expected call of AddWorkflowTaskStartedEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddWorkflowTaskStartedEvent(scheduleEventID, requestID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkflowTaskStartedEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddWorkflowTaskStartedEvent), scheduleEventID, requestID, request)
}

// AddWorkflowTaskCompletedEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddWorkflowTaskCompletedEvent(scheduleEventID, startedEventID int64, request *workflowservice.RespondWorkflowTaskCompletedRequest, maxResetPoints int) (*history.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkflowTaskCompletedEvent", scheduleEventID, startedEventID, request, maxResetPoints)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWorkflowTaskCompletedEvent indicates an expected call of AddWorkflowTaskCompletedEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddWorkflowTaskCompletedEvent(scheduleEventID, startedEventID, request, maxResetPoints interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkflowTaskCompletedEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddWorkflowTaskCompletedEvent), scheduleEventID, startedEventID, request, maxResetPoints)
}

// AddWorkflowTaskFailedEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddWorkflowTaskFailedEvent(scheduleEventID, startedEventID int64, cause enums.WorkflowTaskFailedCause, failure *failure.Failure, identity, binChecksum, baseRunID, newRunID string, forkEventVersion int64) (*history.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkflowTaskFailedEvent", scheduleEventID, startedEventID, cause, failure, identity, binChecksum, baseRunID, newRunID, forkEventVersion)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWorkflowTaskFailedEvent indicates an expected call of AddWorkflowTaskFailedEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddWorkflowTaskFailedEvent(scheduleEventID, startedEventID, cause, failure, identity, binChecksum, baseRunID, newRunID, forkEventVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkflowTaskFailedEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddWorkflowTaskFailedEvent), scheduleEventID, startedEventID, cause, failure, identity, binChecksum, baseRunID, newRunID, forkEventVersion)
}

// AddWorkflowTaskTimedOutEvent mocks base method.
func (m *MockmutableStateWorkflowTaskManager) AddWorkflowTaskTimedOutEvent(scheduleEventID, startedEventID int64) (*history.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkflowTaskTimedOutEvent", scheduleEventID, startedEventID)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWorkflowTaskTimedOutEvent indicates an expected call of AddWorkflowTaskTimedOutEvent.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) AddWorkflowTaskTimedOutEvent(scheduleEventID, startedEventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkflowTaskTimedOutEvent", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).AddWorkflowTaskTimedOutEvent), scheduleEventID, startedEventID)
}

// FailDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) FailDecision(incrementAttempt bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailDecision", incrementAttempt)
}

// FailDecision indicates an expected call of FailDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) FailDecision(incrementAttempt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).FailDecision), incrementAttempt)
}

// DeleteDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) DeleteDecision() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteDecision")
}

// DeleteDecision indicates an expected call of DeleteDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) DeleteDecision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).DeleteDecision))
}

// UpdateDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) UpdateDecision(decision *decisionInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDecision", decision)
}

// UpdateDecision indicates an expected call of UpdateDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) UpdateDecision(decision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).UpdateDecision), decision)
}

// HasPendingDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) HasPendingDecision() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPendingDecision")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPendingDecision indicates an expected call of HasPendingDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) HasPendingDecision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPendingDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).HasPendingDecision))
}

// GetPendingDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) GetPendingDecision() (*decisionInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingDecision")
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPendingDecision indicates an expected call of GetPendingDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) GetPendingDecision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).GetPendingDecision))
}

// HasInFlightDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) HasInFlightDecision() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasInFlightDecision")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasInFlightDecision indicates an expected call of HasInFlightDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) HasInFlightDecision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasInFlightDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).HasInFlightDecision))
}

// GetInFlightDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) GetInFlightDecision() (*decisionInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInFlightDecision")
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetInFlightDecision indicates an expected call of GetInFlightDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) GetInFlightDecision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInFlightDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).GetInFlightDecision))
}

// HasProcessedOrPendingDecision mocks base method.
func (m *MockmutableStateWorkflowTaskManager) HasProcessedOrPendingDecision() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasProcessedOrPendingDecision")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasProcessedOrPendingDecision indicates an expected call of HasProcessedOrPendingDecision.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) HasProcessedOrPendingDecision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasProcessedOrPendingDecision", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).HasProcessedOrPendingDecision))
}

// GetDecisionInfo mocks base method.
func (m *MockmutableStateWorkflowTaskManager) GetDecisionInfo(scheduleEventID int64) (*decisionInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDecisionInfo", scheduleEventID)
	ret0, _ := ret[0].(*decisionInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDecisionInfo indicates an expected call of GetDecisionInfo.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) GetDecisionInfo(scheduleEventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDecisionInfo", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).GetDecisionInfo), scheduleEventID)
}

// CreateTransientDecisionEvents mocks base method.
func (m *MockmutableStateWorkflowTaskManager) CreateTransientDecisionEvents(decision *decisionInfo, identity string) (*history.HistoryEvent, *history.HistoryEvent) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransientDecisionEvents", decision, identity)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(*history.HistoryEvent)
	return ret0, ret1
}

// CreateTransientDecisionEvents indicates an expected call of CreateTransientDecisionEvents.
func (mr *MockmutableStateWorkflowTaskManagerMockRecorder) CreateTransientDecisionEvents(decision, identity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransientDecisionEvents", reflect.TypeOf((*MockmutableStateWorkflowTaskManager)(nil).CreateTransientDecisionEvents), decision, identity)
}