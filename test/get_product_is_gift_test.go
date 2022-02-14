package test

import (
	"testing"

	"github.com/aleesilva/hash-challenger-back-end/usecases"
)

func TestGetProductIsGift(t *testing.T) {

	var mock = []usecases.CartProduct{
		{
			Id:          1,
			Quantity:    2,
			UnitAmount:  10028,
			TotalAmount: 12230,
			Discount:    12,
			IsGift:      false,
		},
		{
			Id:          4,
			Quantity:    2,
			UnitAmount:  12822,
			TotalAmount: 130,
			Discount:    1244,
			IsGift:      false,
		},
	}

	_, isGift := usecases.GetProductGift(mock)

	if isGift {
		t.Log("success finding product gift")
	}
}
