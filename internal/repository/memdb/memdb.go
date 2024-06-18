package repository

import (
	"errors"
	"sync"

	"go-poc-example/internal/domain"
)

// CouponRepository defines the operations available for a coupon repository.
type CouponRepository interface {
	FindAll() ([]domain.Coupon, error)
	FindByID(ID string) (*domain.Coupon, error)
	FindByCode(code string) (*domain.Coupon, error)
	Save(coupon domain.Coupon) error
}

// MemDB implements the CouponRepository interface for an in-memory storage.
type MemDB struct {
	mu      sync.RWMutex
	coupons map[string]domain.Coupon
}

// NewMemDB creates a new instance of MemDB.
func NewMemDB() *MemDB {
	return &MemDB{
		coupons: make(map[string]domain.Coupon),
	}
}

// FindAll returns a slice of all coupons in the database.
func (db *MemDB) FindAll() ([]domain.Coupon, error) {
    db.mu.RLock()
    defer db.mu.RUnlock()

    var coupons []domain.Coupon
    for _, coupon := range db.coupons {
        coupons = append(coupons, coupon)
    }
    return coupons, nil
}

// FindByID searches for a coupon by its ID and returns the coupon if found.
func (db *MemDB) FindByID(ID string) (*domain.Coupon, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, c := range db.coupons {
		if c.ID == ID {
			return &c, nil
		}
	}
	return nil, errors.New("coupon not found")
}

// FindByCode searches for a coupon by its code and returns the coupon if found.
func (db *MemDB) FindByCode(code string) (*domain.Coupon, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, c := range db.coupons {
		if c.Code == code {
			return &c, nil
		}
	}
	return nil, errors.New("coupon not found")
}

// Save adds a new coupon to the in-memory database or updates an existing one.
func (db *MemDB) Save(coupon domain.Coupon) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.coupons[coupon.ID] = coupon
	return nil
}
