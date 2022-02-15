package test

import (
	"testing"

	"github.com/aleesilva/hash-challenger-back-end/usecases"
)

func TestProductDetailsError(t *testing.T) {
	_, err := usecases.ProductDetails(usecases.ProductPayload{
		Id:       99,
		Quantity: 3,
	})

	if err.Error() == "product not found" {
		t.Log("product not found")
	}
}

func TestProductDetailsSuccess(t *testing.T) {
	result, _ := usecases.ProductDetails(usecases.ProductPayload{
		Id:       2,
		Quantity: 2,
	})

	t.Log("product details succefuly: ", result)

}
