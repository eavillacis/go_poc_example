package service

import (
	"errors"
	"fmt"
	"go-poc-example/internal/domain"
	"go-poc-example/internal/repository/memdb" // Updated import path
)

// CouponService defines the interface for the coupon service, including business logic.
type CouponService interface {
    ApplyDiscount(code string, basketValue int) (int, error)
	ListAllCoupons() ([]domain.Coupon, error)
    SaveCoupon(coupon domain.Coupon) (*domain.Coupon, error)
}

// service implements the CouponService interface.
type service struct {
    repo repository.CouponRepository // Use the CouponRepository interface
}

// NewService creates a new instance of the coupon service with the necessary dependencies.
func NewService(repo repository.CouponRepository) CouponService { // Accept CouponRepository interface
    return &service{repo: repo}
}

// ListAllCoupons returns a slice of all coupons from the repository.
func (s *service) ListAllCoupons() ([]domain.Coupon, error) {
    coupons, err := s.repo.FindAll()
    if err != nil {
        return nil, err
    }
    return coupons, nil
}

// CreateCoupon creates a new coupon with the given parameters.
func (s *service) SaveCoupon(coupon domain.Coupon) (*domain.Coupon, error) {    
    couponFound, err := s.repo.FindByID(coupon.ID)
	if err != nil && err.Error() != "coupon not found"{
		return nil, err
	}
	if couponFound != nil {
		return nil, errors.New("coupon already exists")
	}

	if err := s.repo.Save(coupon); err != nil {
        return nil, err
    }

    return &coupon, nil
}

// ApplyDiscount applies a discount to a basket if the coupon code is valid.
func (s *service) ApplyDiscount(code string, basketValue int) (int, error) {
    coupon, err := s.repo.FindByCode(code)
    if err != nil {
        return 0, err
    }

    if basketValue < coupon.MinBasketValue {
        return 0, fmt.Errorf("basket value too low for coupon, should be at least %d", coupon.MinBasketValue)
    }

    return coupon.Discount, nil
}
