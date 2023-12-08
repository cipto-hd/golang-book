package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductData struct {
	Products []Product `json:"products"`
}

func main() {
	fRead, fErr := os.ReadFile("products.json")
	if fErr != nil {
		log.Println(fErr)
	}

	productData := ProductData{}

	_ = json.Unmarshal([]byte(fRead), &productData)

	var nextProductId int

	if len(productData.Products) == 0 {
		nextProductId = 1
	} else {
		nextProductId = productData.Products[len(productData.Products)-1].Id + 1
	}

	product := Product{Id: nextProductId, Name: "product"}

	products := append(productData.Products, product)

	for j := 0; j < len(products); j++ {
		item := products[j]
		fmt.Println("Item id: ", item.Id, "Item name: ", item.Name)
	}

	productsStr, _ := json.Marshal(products)
	productsJson := fmt.Sprintf("{\"products\": %s}", productsStr)

	f, err := os.OpenFile("products.json", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}

	f.WriteString(productsJson)

	f.Close()

}
