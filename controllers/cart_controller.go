package controllers

import (
	"net/http"

	"github.com/aleesilva/hash-challenger-back-end/helpers"
	"github.com/aleesilva/hash-challenger-back-end/usecases"
	"github.com/gin-gonic/gin"
)

type ProductPayload struct {
	Products []Payload `json:"products"`
}

type Payload struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type Response struct {
	TotalAmount             int                    `json:"total_amount"`
	TotalAmountWithDiscount int                    `json:"total_amount_with_discount"`
	TotalDiscount           int                    `json:"total_discount"`
	Products                []usecases.CartProduct `json:"products"`
}

func Checkout(ctx *gin.Context) {
	payload := ProductPayload{}
	var cartTotalAmount int = 0
	var cartTotalDiscount int = 0
	// var response Response
	var productsCartCheckout []usecases.CartProduct

	if err := ctx.ShouldBindJSON(&payload); err == nil {
		for _, value := range payload.Products {
			// faz a busca dos produtos e retorna um productCheckout
			product, _ := usecases.ProductDetails(usecases.ProductPayload{
				Id:       value.Id,
				Quantity: value.Quantity,
			})

			if err == nil {
				cartTotalAmount += int(product.TotalAmount)
				discount := usecases.ApplyDiscount(*product)
				product.Discount = float32(discount)
				cartTotalDiscount += int(product.Discount)
				productsCartCheckout = append(productsCartCheckout, *product)
			}

		}
		if helpers.VerifyBlackFridayDate() {
			productGift, isGift := usecases.GetProductGift(productsCartCheckout)
			if isGift {
				productsCartCheckout = append(productsCartCheckout, *productGift)
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"response": &Response{
			TotalAmount:             cartTotalAmount,
			TotalAmountWithDiscount: cartTotalAmount - cartTotalDiscount,
			TotalDiscount:           cartTotalDiscount,
			Products:                productsCartCheckout,
		}})
	}
}
