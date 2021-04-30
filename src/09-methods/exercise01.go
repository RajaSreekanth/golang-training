package main

import (
	"fmt"
	"reflect"
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

	i := products.Index(Product{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"})
	fmt.Println("index is", i)
	exists := products.Include(Product{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"})
	fmt.Println("exists:", exists)

	anyMatch := products.Any1(MatchCriteria{"Zen", "stationary"})
	fmt.Println("any match", anyMatch)

	anyMatch = products.Any1(MatchCriteria{"Zen", "unknown"})
	fmt.Println("any match2", anyMatch)

	allMatch := products.All1(MatchCriteria{"Pen", "stationary"})
	fmt.Println("all match", allMatch)

	allMatch = products.All1(MatchCriteria{"Pens", "station"})
	fmt.Println("all match2", allMatch)

	filteredProducts := products.Filter1(MatchCriteria{"Pen", "stationary"})
	fmt.Println("Filtered products", filteredProducts)

	stationaryProductCriteria := func(product Product) bool {
		return product.category == "stationary"
	}

	anyStationaryProducts := products.Any(stationaryProductCriteria)

	fmt.Println("Any stationary products ? : ", anyStationaryProducts)

	anyGroceryProducts := products.Any(func(product Product) bool {
		return product.category == "grocery"
	})

	fmt.Println("Any grocery products ? : ", anyGroceryProducts)

	allStationaryProducts := products.All(stationaryProductCriteria)

	fmt.Println("all stationary products?", allStationaryProducts)

	products2 := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "stationary"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
		{id: 108, name: "Pen", cost: 20, units: 30, category: "stationary"},
	}
	allStationaryProducts = products2.All(stationaryProductCriteria)

	fmt.Println("all stationary products?", allStationaryProducts)

}

func (products Products) Index(p Product) int {
	for i := 0; i < len(products); i++ {
		if p.id == products[i].id {
			return i

		}
	}
	return -1
}

func (products Products) Include(p Product) bool {
	exists := false
	for i := 0; i < len(products); i++ {
		if reflect.DeepEqual(p, products[i]) {
			exists = true
			break
		}
	}
	return exists
}

func (products Products) Any1(m MatchCriteria) bool {
	match := false
	for i := 0; i < len(products); i++ {
		if m.name == products[i].name || m.category == products[i].category {
			match = true
			break
		}
	}
	return match
}

func (products Products) All1(m MatchCriteria) bool {
	match := false
	for i := 0; i < len(products); i++ {
		if m.name == products[i].name && m.category == products[i].category {
			match = true
			break
		}
	}
	return match
}

func (products Products) Filter1(m MatchCriteria) []Product {
	var matchedProducts = []Product{}
	for i := 0; i < len(products); i++ {
		if m.name == products[i].name && m.category == products[i].category {
			matchedProducts = append(matchedProducts, products[i])
		}
	}
	return matchedProducts
}

func (products Products) Any(criteria func(product Product) bool) bool {
	for _, product := range products {
		if criteria(product) {
			return true
		}
	}
	return false
}

func (products Products) All(criteria func(product Product) bool) bool {
	matches := true
	for _, product := range products {
		if !criteria(product) {
			return false
		}
	}
	return matches
}
