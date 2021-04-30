// implement sort method on Products type
// sort by name, category, cost, units
package main

import (
	"fmt"
	"sort"
)

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

type MatchCriteria struct {
	name     string
	category string
}

type Products []Product

func main() {
	products := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
		{id: 108, name: "Pen", cost: 20, units: 30, category: "stationary"},
	}
	for _, product := range products {
		fmt.Println(product)
	}

	fmt.Println("After sorting by ID...")
	products.Sort("id", true)
	fmt.Println(products)

	fmt.Println("After sorting by Name...")
	products.Sort("name", true)
	fmt.Println(products)

	fmt.Println("After sorting by Cost in descending order...")
	products.Sort("cost", false)
	fmt.Println(products)

	fmt.Println("After sorting by Units in descending order...")
	products.Sort("units", false)
	fmt.Println(products)

	fmt.Println("After sorting by Category in descending order...")
	products.Sort("category", false)
	fmt.Println(products)
}

func (products Products) Sort(criterion string, ascending bool) {
	/*
		compareByFunc := compareBy[criterion]

		sort.Slice(products, func(i, j int) bool {
			return compareByFunc(products[i], products[j])
		})
	*/
	currentProductComparer = compareBy[criterion]
	if ascending {
		sort.Sort(products)
	} else {
		sort.Sort(sort.Reverse(products))
	}
}

func sortID(p1 Product, p2 Product) bool {
	return p1.id < p2.id
}

func sortName(p1 Product, p2 Product) bool {
	return p1.name < p2.name
}

func sortCost(p1 Product, p2 Product) bool {
	return p1.cost < p2.cost
}

func sortUnits(p1 Product, p2 Product) bool {
	return p1.units < p2.units
}
func sortCategory(p1 Product, p2 Product) bool {
	return p1.category < p2.category
}

var currentProductComparer func(p1, p2 Product) bool

var compareBy map[string]func(p1, p2 Product) bool = map[string]func(p1, p2 Product) bool{
	"id":       sortID,
	"name":     sortName,
	"cost":     sortCost,
	"units":    sortUnits,
	"category": sortCategory,
}

/* Methods for 'Interface' interface of 'sort' package  */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return currentProductComparer(products[i], products[j])
}

/* Methods for Stringer interface of the 'fmt' package */
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, cost = %v, units = %d, category = %s\n", product.id, product.name, product.cost, product.units, product.category)
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprint(product)
	}
	return result
}
