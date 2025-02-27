// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	domain "go-poc-example/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// NewService is an autogenerated mock type for the CouponService type
type NewService struct {
	mock.Mock
}

// ApplyDiscount provides a mock function with given fields: code, basketValue
func (_m *NewService) ApplyDiscount(code string, basketValue int) (int, error) {
	ret := _m.Called(code, basketValue)

	if len(ret) == 0 {
		panic("no return value specified for ApplyDiscount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (int, error)); ok {
		return rf(code, basketValue)
	}
	if rf, ok := ret.Get(0).(func(string, int) int); ok {
		r0 = rf(code, basketValue)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(code, basketValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllCoupons provides a mock function with given fields:
func (_m *NewService) ListAllCoupons() ([]domain.Coupon, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListAllCoupons")
	}

	var r0 []domain.Coupon
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Coupon, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Coupon); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Coupon)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveCoupon provides a mock function with given fields: coupon
func (_m *NewService) SaveCoupon(coupon domain.Coupon) (*domain.Coupon, error) {
	ret := _m.Called(coupon)

	if len(ret) == 0 {
		panic("no return value specified for SaveCoupon")
	}

	var r0 *domain.Coupon
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Coupon) (*domain.Coupon, error)); ok {
		return rf(coupon)
	}
	if rf, ok := ret.Get(0).(func(domain.Coupon) *domain.Coupon); ok {
		r0 = rf(coupon)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Coupon)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.Coupon) error); ok {
		r1 = rf(coupon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNewService creates a new instance of NewService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NewService {
	mock := &NewService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
