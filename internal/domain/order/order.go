package order

import (
	"github.com/Vagmacker/luzora-api/internal/domain/coupon"
	"github.com/Vagmacker/luzora-api/internal/domain/customer"
)

type Order struct {
	ID       string
	Customer *customer.Customer
	Coupon   *coupon.Coupon
	Items    []*OrderItem
}
