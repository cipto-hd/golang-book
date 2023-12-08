package main

import (
	"fmt"
)

type Row struct {
	Title       string
	Description string
	Quantity    int
	UnitPrice   float32
}

func (r Row) string() string {
	return fmt.Sprintf("%s, %s, %d, %dGBP, %dGBP", r.Title, r.Description, r.Quantity, int(r.UnitPrice), (r.Quantity)*int(r.UnitPrice))
}

func main() {
	row := Row{
		Title:       "LEGO set",
		Description: "4000 pieces",
		Quantity:    1,
		UnitPrice:   600,
	}
	row2 := Row{
		Title:       "Plushy",
		Description: "plush toy",
		Quantity:    3,
		UnitPrice:   5,
	}

	basket := make([]Row, 0)
	basket = append(basket, row)
	basket = append(basket, row2)

	fmt.Println("Title, Description, Quantity, Price per unit, Total")

	var sum int = 0
	for i := 0; i < len(basket); i++ {
		current := basket[i]
		fmt.Println(current.string())
		sum += current.Quantity * int(current.UnitPrice)
	}

	fmt.Println("\nTotal: ", sum, "GBP")
}
