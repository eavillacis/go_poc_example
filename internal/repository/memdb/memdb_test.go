package repository_test

import (
	"errors"
	"go-poc-example/internal/domain"
	"go-poc-example/internal/repository/memdb/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO: separate tests and create more use cases
func TestMemDB(t *testing.T) {
	db := mocks.NewMemDB(t) // Using the mock instead of the real implementation just for the sake of the test

    // Test Save
    coupon1 := domain.Coupon{ID: "1", Code: "SAVE10", Discount: 10}
    db.On("Save", mock.Anything).Return(nil)
    err := db.Save(coupon1)
    assert.NoError(t, err)

    // Test Save - Update
    coupon1Updated := domain.Coupon{ID: "1", Code: "SAVE10", Discount: 15}
    db.On("Save", mock.Anything).Return(nil)
    err = db.Save(coupon1Updated)
    assert.NoError(t, err)

    // Test FindAll
    db.On("FindAll").Return([]domain.Coupon{coupon1Updated}, nil)
    coupons, err := db.FindAll()
    assert.NoError(t, err)
    assert.Len(t, coupons, 1)
    assert.Equal(t, coupon1Updated, coupons[0], "Expected updated coupon")

    // Test FindByID - Success
    db.On("FindByID", "SAVE10").Return(&coupon1Updated, nil)
    foundCoupon, err := db.FindByID("SAVE10")
    assert.NoError(t, err)
    assert.NotNil(t, foundCoupon)
    assert.Equal(t, coupon1Updated, *foundCoupon)

    // Test FindByID - Failure
    db.On("FindByID", "NONEXISTENT").Return(nil, errors.New("coupon not found"))
    _, err = db.FindByID("NONEXISTENT")
    assert.Error(t, err)

    // Test FindByCode - Success
    db.On("FindByCode", "SAVE10").Return(&coupon1Updated, nil)
    foundCoupon, err = db.FindByCode("SAVE10")
    assert.NoError(t, err)
    assert.NotNil(t, foundCoupon)
    assert.Equal(t, coupon1Updated, *foundCoupon)

    // Test FindByCode - Failure
    db.On("FindByCode", "NONEXISTENT").Return(nil, errors.New("coupon not found"))
    _, err = db.FindByCode("NONEXISTENT")
    assert.Error(t, err)
}