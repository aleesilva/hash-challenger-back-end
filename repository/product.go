package repository

import (
	"encoding/json"
	"errors"
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

func GetProductById(productId int) (*Product, error) {
	productArr := []Product{}
	product := Product{}
	file, _ := ioutil.ReadFile("./db/products.json")

	if err := json.Unmarshal([]byte(file), &productArr); err != nil {
		return nil, errors.New("database is not online ")
	}

	for _, value := range productArr {
		if value.Id == productId {
			product = Product{
				Id:          value.Id,
				Title:       value.Title,
				Description: value.Description,
				Amount:      value.Amount,
				IsGift:      value.IsGift,
			}
			return &product, nil
		}

	}
	return nil, errors.New("Product not found")
}
