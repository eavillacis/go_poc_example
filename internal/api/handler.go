package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-poc-example/internal/domain"
	"go-poc-example/internal/service"
)

// Handler holds the dependencies for the HTTP server
type Handler struct {
	svc service.CouponService
}

// NewHandler initializes and returns a new Handler
func NewHandler(svc service.CouponService) *Handler {
	return &Handler{svc: svc}
}

// RegisterRoutes sets up the routes for the HTTP server
func (h *Handler) RegisterRoutes(mux *mux.Router) {
	api := mux.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/", h.HealthCheck).Methods("GET")

    couponRouter := api.PathPrefix("/coupon").Subrouter()
    couponRouter.HandleFunc("/save", h.SaveCoupon).Methods("POST")	
    couponRouter.HandleFunc("/findAll", h.FindAll).Methods("GET")
	couponRouter.HandleFunc("/apply-discount", h.ApplyDiscount).Methods("POST")
}

// healthCheck handles requests to the "/" route
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
    // Write "Service OK" to the response
    w.WriteHeader(http.StatusOK) // Set the status code to 200 OK
    w.Write([]byte("Service OK"))
}

// findAll handles GET requests to list all coupons
func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    coupons, err := h.svc.ListAllCoupons()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(coupons); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// createCoupon handles POST requests to create a new coupon
func (h *Handler) SaveCoupon(w http.ResponseWriter, r *http.Request) {	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
	}

	var coupon domain.Coupon
	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := coupon.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.svc.SaveCoupon(coupon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ApplyDiscount(w http.ResponseWriter, r *http.Request) {
    if r.Body == nil {
        http.Error(w, "Request body is empty", http.StatusBadRequest)
        return
    }

    var couponDiscount domain.CouponDiscount
    if err := json.NewDecoder(r.Body).Decode(&couponDiscount); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err := couponDiscount.ValidateDiscount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    discountedValue, err := h.svc.ApplyDiscount(couponDiscount.Code, couponDiscount.BasketValue)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest) // Use BadRequest assuming most errors are due to invalid input
        return
    }

    response := map[string]int{"discountedValue": discountedValue}
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
