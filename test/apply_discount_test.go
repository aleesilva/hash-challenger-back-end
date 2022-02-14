package test

import (
	"testing"

	"github.com/aleesilva/hash-challenger-back-end/usecases"
)

var MockProduct = &usecases.CartProduct{
	Id:          1,
	Quantity:    2,
	UnitAmount:  10,
	TotalAmount: 15157,
	Discount:    0,
	IsGift:      false,
}

func TestApplyDiscount(t *testing.T) {
	result := usecases.ApplyDiscount(*MockProduct)
	if result == 0 {
		t.Log("discount not apply service discount is not online")
	}
	if result == 7 {
		t.Logf("your discount was : %d", result)
	}
}
