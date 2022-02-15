package usecases

import "github.com/aleesilva/hash-challenger-back-end/integrations"

func ApplyDiscount(cartProduct CartProduct) int {
	discount, err := integrations.GetDiscount(cartProduct.Id)
	if err == nil {
		return int(cartProduct.TotalAmount / 100 * discount)
	}
	return 0
}
