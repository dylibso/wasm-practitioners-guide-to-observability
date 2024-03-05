package main

import (
	"encoding/json"
	"fmt"

	observe "github.com/dylibso/observe-sdk/observe-api/go"
)

type Data struct {
	Name        string
	ProductItem Product
}

const productJson = `{"id":1,"title":"iPhone 9","description":"An apple mobile which is nothing like apple","price":549,"discountPercentage":12.96,"rating":4.69,"stock":94,"brand":"Apple","category":"smartphones","thumbnail":"https://cdn.dummyjson.com/product-images/1/thumbnail.jpg","images":["https://cdn.dummyjson.com/product-images/1/1.jpg","https://cdn.dummyjson.com/product-images/1/2.jpg","https://cdn.dummyjson.com/product-images/1/3.jpg","https://cdn.dummyjson.com/product-images/1/4.jpg","https://cdn.dummyjson.com/product-images/1/thumbnail.jpg"]}`

type Product struct {
	Id                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

func (d *Data) SetName(name string) {
	d.Name = name
}

func (d *Data) SetProduct(input string) {
	var product Product
	err := json.Unmarshal([]byte(input), &product)
	if err != nil {
		fmt.Println("failed to unmarshal json:", err)
		return
	}

	if product.Brand != "Apple" {
		observe.SpanTags([]string{"brand:unknown"})
	}
	observe.SpanTags([]string{
		fmt.Sprintf("price:%d", product.Price),
		fmt.Sprintf("discountPercentage:%f", product.DiscountPercentage),
	})
	d.ProductItem = product
}

func main() {
	data := &Data{}
	data.SetName("NewProduct")
	data.SetProduct(productJson)

	fmt.Println(data)
}
