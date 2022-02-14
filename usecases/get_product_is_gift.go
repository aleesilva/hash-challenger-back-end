package usecases

import (
	"github.com/aleesilva/hash-challenger-back-end/repository"
)

func GetProductGift(cartProduct []CartProduct) (*CartProduct, bool) {
	var giftProductId int
	newGiftProduct := CartProduct{}
	product := repository.GetProductGift()
	var productAlreadyInCart bool

	for _, value := range cartProduct {
		if value.Id == product.Id {
			productAlreadyInCart = true
		}
		productAlreadyInCart = false
	}

	if productAlreadyInCart {
		return &CartProduct{}, false
	}

	newGiftProduct = CartProduct{
		Id:          giftProductId,
		Quantity:    1,
		UnitAmount:  0,
		Discount:    0,
		TotalAmount: 0,
		IsGift:      true,
	}
	return &newGiftProduct, true
}
