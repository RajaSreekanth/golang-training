package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

type Products []Product

func main() {
	product := Product{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"}
	product.Print()

	products := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
	}

	fmt.Println("Initial list of products")
	//Print(products)
	products.Print()
}

func (product Product) Print() {
	fmt.Printf("Id = %d, Name = %s, cost = %v, units = %d, category = %s\n", product.id, product.name, product.cost, product.units, product.category)
}

func (products Products) Print() {
	for _, product := range products {
		product.Print()
	}
}
