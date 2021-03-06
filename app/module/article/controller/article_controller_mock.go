// Code generated by MockGen. DO NOT EDIT.
// Source: article_controller.go

// Package mock_controller is a generated GoMock package.
package controller

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockIArticleController is a mock of IArticleController interface.
type MockIArticleController struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleControllerMockRecorder
}

// MockIArticleControllerMockRecorder is the mock recorder for MockIArticleController.
type MockIArticleControllerMockRecorder struct {
	mock *MockIArticleController
}

// NewMockIArticleController creates a new mock instance.
func NewMockIArticleController(ctrl *gomock.Controller) *MockIArticleController {
	mock := &MockIArticleController{ctrl: ctrl}
	mock.recorder = &MockIArticleControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleController) EXPECT() *MockIArticleControllerMockRecorder {
	return m.recorder
}

// Destroy mocks base method.
func (m *MockIArticleController) Destroy(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockIArticleControllerMockRecorder) Destroy(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockIArticleController)(nil).Destroy), c)
}

// Index mocks base method.
func (m *MockIArticleController) Index(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *MockIArticleControllerMockRecorder) Index(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockIArticleController)(nil).Index), c)
}

// Show mocks base method.
func (m *MockIArticleController) Show(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Show indicates an expected call of Show.
func (mr *MockIArticleControllerMockRecorder) Show(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockIArticleController)(nil).Show), c)
}

// Store mocks base method.
func (m *MockIArticleController) Store(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockIArticleControllerMockRecorder) Store(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockIArticleController)(nil).Store), c)
}

// Update mocks base method.
func (m *MockIArticleController) Update(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIArticleControllerMockRecorder) Update(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIArticleController)(nil).Update), c)
}
