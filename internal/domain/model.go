package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Coupon represents the domain model for a coupon.
type Coupon struct {
    ID             string `json:"id" validate:"required,uuid"`
    Discount       int    `json:"discount" validate:"gte=0,lte=100"`
    Code           string `json:"code" validate:"required,alphanum"`
    MinBasketValue int    `json:"minBasketValue" validate:"gt=0"`
}

type CouponDiscount struct {
    Code        string `json:"code" validate:"required,alphanum"`
    BasketValue int    `json:"basketValue" validate:"gt=0"`
}

func NewCoupon(discount int, code string, minBasketValue int) *Coupon {
    return &Coupon{
        ID:             uuid.NewString(),
        Discount:       discount,
        Code:           code,
        MinBasketValue: minBasketValue,
    }
}

func (c *Coupon) Validate() error {
	validate := validator.New()
    return validate.Struct(c)
}

func (c *CouponDiscount) ValidateDiscount() error {
	validate := validator.New()
    return validate.Struct(c)
}
