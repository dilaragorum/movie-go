// Code generated by MockGen. DO NOT EDIT.
// Source: service/movie_service_interface.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	model "github.com/dilaragorum/movie-go/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIMovieService is a mock of IMovieService interface.
type MockIMovieService struct {
	ctrl     *gomock.Controller
	recorder *MockIMovieServiceMockRecorder
}

// MockIMovieServiceMockRecorder is the mock recorder for MockIMovieService.
type MockIMovieServiceMockRecorder struct {
	mock *MockIMovieService
}

// NewMockIMovieService creates a new mock instance.
func NewMockIMovieService(ctrl *gomock.Controller) *MockIMovieService {
	mock := &MockIMovieService{ctrl: ctrl}
	mock.recorder = &MockIMovieServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMovieService) EXPECT() *MockIMovieServiceMockRecorder {
	return m.recorder
}

// CreateMovie mocks base method.
func (m *MockIMovieService) CreateMovie(movie model.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovie", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMovie indicates an expected call of CreateMovie.
func (mr *MockIMovieServiceMockRecorder) CreateMovie(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovie", reflect.TypeOf((*MockIMovieService)(nil).CreateMovie), movie)
}

// DeleteAllMovie mocks base method.
func (m *MockIMovieService) DeleteAllMovie() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllMovie")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllMovie indicates an expected call of DeleteAllMovie.
func (mr *MockIMovieServiceMockRecorder) DeleteAllMovie() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllMovie", reflect.TypeOf((*MockIMovieService)(nil).DeleteAllMovie))
}

// DeleteMovie mocks base method.
func (m *MockIMovieService) DeleteMovie(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockIMovieServiceMockRecorder) DeleteMovie(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockIMovieService)(nil).DeleteMovie), id)
}

// GetMovie mocks base method.
func (m *MockIMovieService) GetMovie(id int) (model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", id)
	ret0, _ := ret[0].(model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockIMovieServiceMockRecorder) GetMovie(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockIMovieService)(nil).GetMovie), id)
}

// GetMovies mocks base method.
func (m *MockIMovieService) GetMovies() ([]model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies")
	ret0, _ := ret[0].([]model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockIMovieServiceMockRecorder) GetMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockIMovieService)(nil).GetMovies))
}

// UpdateMovie mocks base method.
func (m *MockIMovieService) UpdateMovie(id int, movie model.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", id, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockIMovieServiceMockRecorder) UpdateMovie(id, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockIMovieService)(nil).UpdateMovie), id, movie)
}
