package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-poc-example/internal/domain"
	"go-poc-example/internal/service/mocks"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO: complete tests for the remaining handler methods
func TestHealthCheck(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(NewHandler(nil).HealthCheck)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Equal(t, "Service OK", rr.Body.String())
}

func TestFindAll(t *testing.T) {
    mockSvc := new(mocks.NewService)
    mockSvc.On("ListAllCoupons").Return([]domain.Coupon{}, nil)

    h := NewHandler(mockSvc)

	prefix := "/api/v1/coupon"
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/findAll", prefix), nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    h.RegisterRoutes(router)
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    // Further assertions can be made on the response body if needed
}

func TestSaveCoupon(t *testing.T) {
    mockSvc := new(mocks.NewService)
    mockSvc.On("SaveCoupon", mock.Anything).Return(nil, nil)

    h := NewHandler(mockSvc)

    coupon := domain.Coupon{ID: "283b8e96-a5f4-4389-a82f-5ec46399de37", Code: "TEST", Discount: 10, MinBasketValue: 50}
    body, err := json.Marshal(coupon)
    assert.NoError(t, err)

	prefix := "/api/v1/coupon"
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/save", prefix), bytes.NewBuffer(body))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    h.RegisterRoutes(router)
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
    // Further assertions can be made on the response body if needed
}
