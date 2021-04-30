package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

/*
type PerishableProduct struct {
	product Product
	expiry  int
}
*/

type PerishableProduct struct {
	Product
	expiry int
}

func main() {
	/*
		var p1 Product = Product{}
		p1.id = 100
		p1.name = "Pen"
		p1.cost = 10
		p1.units = 200
		p1.category = "Stationary"
	*/
	/*
		var p1 Product = Product{100, "Pen", 10, 200, "Stationary"}
	*/
	/*
		var p1 Product = Product{name: "Pen", id: 100, units: 200, cost: 10, category: "stationary"}
	*/
	p1 := Product{name: "Pen", id: 100, units: 200, cost: 10, category: "stationary"}
	fmt.Printf("Type of p1 = %T\n", p1)
	fmt.Println("Product p1 = ", p1)
	applyDiscount(&p1, 10)
	fmt.Println("After applying 10% discount p1 = ", p1)

	//grapes := PerishableProduct{product: Product{101, "Grapes", 80, 50, "Fruits"}, expiry: 2}
	//fmt.Println("Perishable Product grapes = ", grapes)

	//grapes := PerishableProduct{Product: Product{101, "Grapes", 80, 50, "Fruits"}, expiry: 2}
	//grapes := PerishableProduct{Product{101, "Grapes", 80, 50, "Fruits"}, 2}
	grapes := NewPerishableProduct(101, "Grapes", 80, 50, "Fruits", 2)
	printPerishableProduct(grapes)
}

/*
func printPerishableProduct(pp PerishableProduct) {
	fmt.Printf("Perishable Product => id=%d, name=%s, cost=%v, units=%d, category=%s, expiry=%d\n", pp.product.id, pp.product.name, pp.product.cost, pp.product.units, pp.product.category, pp.expiry)
}
*/

/*
func printPerishableProduct(pp PerishableProduct) {
	fmt.Printf("Perishable Product => id=%d, name=%s, cost=%v, units=%d, category=%s, expiry=%d\n", pp.Product.id, pp.Product.name, pp.Product.cost, pp.Product.units, pp.Product.category, pp.expiry)
}
*/

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry int) PerishableProduct {
	return PerishableProduct{Product{id, name, cost, units, category}, expiry}
}

func printPerishableProduct(pp PerishableProduct) {
	fmt.Printf("Perishable Product => id=%d, name=%s, cost=%v, units=%d, category=%s, expiry=%d\n", pp.id, pp.name, pp.cost, pp.units, pp.category, pp.expiry)
}

func applyDiscount(productPtr *Product, discount int) /*  */ {
	productPtr.cost = productPtr.cost * ((100 - float32(discount)) / 100)
}
