package repository

import (
	"encoding/json"
	"io/ioutil"
)

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}

func ListProducts() []Product {
	var product []Product
	file, _ := ioutil.ReadFile("./db/products.json")
	if err := json.Unmarshal([]byte(file), &product); err != nil {
		panic(err)
	}

	return product
}

func GetProductById(productId int) (Product, bool) {
	var product []Product
	file, _ := ioutil.ReadFile("./db/products.json")
	if err := json.Unmarshal([]byte(file), &product); err != nil {
		panic(err)
	}
	for _, value := range product {
		if value.Id == productId {
			return value, true
		}

	}
	var noP Product
	return noP, false
}
