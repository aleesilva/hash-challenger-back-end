package usecases

import (
	"errors"
	"strconv"

	"github.com/aleesilva/hash-challenger-back-end/repository"
)

type ProductPayload struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}

type CartProduct struct {
	Id          int     `json:"id"`
	UnitAmount  float32 `json:"unit_amount"`
	TotalAmount float32 `json:"total_amount"`
	Discount    float32 `json:"discount"`
	IsGift      bool    `json:"is_gift"`
	Quantity    int     `json:"quantity"`
}

func ProductDetails(productsPayload ProductPayload) (*CartProduct, error) {
	cartProduct := CartProduct{}
	product, exists := repository.GetProductById(productsPayload.Id)

	if !exists {
		return nil, errors.New("Product not found")
	}

	if product.IsGift {
		return nil, errors.New("product with id: " + strconv.Itoa(productsPayload.Id) + " cannot be added to cart")
	}

	cartProduct = CartProduct{
		Id:          product.Id,
		Quantity:    productsPayload.Quantity,
		UnitAmount:  float32(product.Amount),
		TotalAmount: float32(product.Amount) * float32(productsPayload.Quantity),
		Discount:    0,
		IsGift:      product.IsGift,
	}

	return &cartProduct, nil
}
