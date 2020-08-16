// Code generated by MockGen. DO NOT EDIT.
// Source: magazine_service.go

// Package mock_service is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mariojuzar/soldier-magazine/api/entity/model"
	request "github.com/mariojuzar/soldier-magazine/api/entity/rest-web/request"
	response "github.com/mariojuzar/soldier-magazine/api/entity/rest-web/response"
	reflect "reflect"
)

// MockMagazineService is a mock of MagazineService interface
type MockMagazineService struct {
	ctrl     *gomock.Controller
	recorder *MockMagazineServiceMockRecorder
}

// MockMagazineServiceMockRecorder is the mock recorder for MockMagazineService
type MockMagazineServiceMockRecorder struct {
	mock *MockMagazineService
}

// NewMockMagazineService creates a new mock instance
func NewMockMagazineService(ctrl *gomock.Controller) *MockMagazineService {
	mock := &MockMagazineService{ctrl: ctrl}
	mock.recorder = &MockMagazineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMagazineService) EXPECT() *MockMagazineServiceMockRecorder {
	return m.recorder
}

// PutMagazine mocks base method
func (m *MockMagazineService) PutMagazine(request request.PutMagazineRequest) (model.Soldier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMagazine", request)
	ret0, _ := ret[0].(model.Soldier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutMagazine indicates an expected call of PutMagazine
func (mr *MockMagazineServiceMockRecorder) PutMagazine(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMagazine", reflect.TypeOf((*MockMagazineService)(nil).PutMagazine), request)
}

// LoadMagazineRandomly mocks base method
func (m *MockMagazineService) LoadMagazineRandomly(request request.LoadMagazineRequest) (response.VerifiedMagazineStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMagazineRandomly", request)
	ret0, _ := ret[0].(response.VerifiedMagazineStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadMagazineRandomly indicates an expected call of LoadMagazineRandomly
func (mr *MockMagazineServiceMockRecorder) LoadMagazineRandomly(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMagazineRandomly", reflect.TypeOf((*MockMagazineService)(nil).LoadMagazineRandomly), request)
}
