// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/stashapp/stash/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// GalleryReaderWriter is an autogenerated mock type for the GalleryReaderWriter type
type GalleryReaderWriter struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *GalleryReaderWriter) All() ([]*models.Gallery, error) {
	ret := _m.Called()

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func() []*models.Gallery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBySceneID provides a mock function with given fields: sceneID
func (_m *GalleryReaderWriter) FindBySceneID(sceneID int) (*models.Gallery, error) {
	ret := _m.Called(sceneID)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(int) *models.Gallery); ok {
		r0 = rf(sceneID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(sceneID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ids
func (_m *GalleryReaderWriter) FindMany(ids []int) ([]*models.Gallery, error) {
	ret := _m.Called(ids)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func([]int) []*models.Gallery); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}