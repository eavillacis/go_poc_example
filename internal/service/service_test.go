package service_test

import (
    "errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-poc-example/internal/domain"
	"go-poc-example/internal/service"
	"go-poc-example/internal/repository/memdb/mocks"
	"testing"
)

func TestListAllCoupons(t *testing.T) {
	// Setup
	mockRepo := new(mocks.MemDB)
	expectedCoupons := []domain.Coupon{{ID: "1", Code: "SAVE10", Discount: 10}}
	mockRepo.On("FindAll").Return(expectedCoupons, nil)
	svc := service.NewService(mockRepo)

	// Execute
	coupons, err := svc.ListAllCoupons()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedCoupons, coupons)
	mockRepo.AssertExpectations(t)
}

func TestSaveCoupon(t *testing.T) {
	// Setup
	mockRepo := new(mocks.MemDB)
	inputCoupon := domain.Coupon{Code: "SAVE20", Discount: 20}
	mockRepo.On("FindByID", mock.Anything).Return(nil, nil)
	mockRepo.On("Save", mock.Anything).Return(nil)
	svc := service.NewService(mockRepo)

	// Execute
	savedCoupon, err := svc.SaveCoupon(inputCoupon)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, savedCoupon.ID) // Ensure an ID was assigned
	assert.Equal(t, inputCoupon.Discount, savedCoupon.Discount)
	mockRepo.AssertExpectations(t)
}

func TestSaveCouponFound(t *testing.T) {
	// Setup
	mockRepo := new(mocks.MemDB)
	validCoupon := domain.Coupon{ID: "1", Code: "SAVE10", Discount: 10, MinBasketValue: 50}
	mockRepo.On("FindByID", mock.Anything).Return(&validCoupon, nil)
	mockRepo.On("Save", mock.Anything).Return(nil, errors.New("coupon already exists"))
	svc := service.NewService(mockRepo)

	// Execute
	savedCoupon, err := svc.SaveCoupon(validCoupon)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, savedCoupon) // Ensure an ID was assigned	
}

func TestApplyDiscount(t *testing.T) {
	// Setup for valid coupon
	mockRepo := new(mocks.MemDB)
	validCoupon := domain.Coupon{ID: "1", Code: "SAVE10", Discount: 10, MinBasketValue: 50}
	mockRepo.On("FindByCode", "SAVE10").Return(&validCoupon, nil)
	svc := service.NewService(mockRepo)

	// Execute with valid coupon
	discount, err := svc.ApplyDiscount("SAVE10", 100) // Assuming 100 is the basket value

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 10, discount) // Assuming the discount is a fixed amount and not a percentage
	mockRepo.AssertExpectations(t)

	// Setup for invalid coupon
	mockRepo.On("FindByCode", "INVALID").Return(nil, errors.New("coupon not found"))
	svc = service.NewService(mockRepo)

	// Execute with invalid coupon
	discount, err = svc.ApplyDiscount("INVALID", 100)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0, discount) // Expect no discount for invalid coupon
	mockRepo.AssertExpectations(t)
}
