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

func main() {
	products := []Product{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
		{id: 108, name: "Pen", cost: 20, units: 30, category: "stationary"},
	}

	i := Index(Product{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"}, products)
	fmt.Println("index is", i)
	exists := Include(Product{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"}, products)
	fmt.Println("exists:", exists)

	anyMatch := Any1(MatchCriteria{"Zen", "stationary"}, products)
	fmt.Println("any match", anyMatch)

	anyMatch = Any1(MatchCriteria{"Zen", "unknown"}, products)
	fmt.Println("any match2", anyMatch)

	allMatch := All1(MatchCriteria{"Pen", "stationary"}, products)
	fmt.Println("all match", allMatch)

	allMatch = All1(MatchCriteria{"Pens", "station"}, products)
	fmt.Println("all match2", allMatch)

	filteredProducts := Filter1(MatchCriteria{"Pen", "stationary"}, products)
	fmt.Println("Filtered products", filteredProducts)

	stationaryProductCriteria := func(product Product) bool {
		return product.category == "stationary"
	}

	anyStationaryProducts := Any(products, stationaryProductCriteria)

	fmt.Println("Any stationary products ? : ", anyStationaryProducts)

	anyGroceryProducts := Any(products, func(product Product) bool {
		return product.category == "grocery"
	})

	fmt.Println("Any grocery products ? : ", anyGroceryProducts)

	allStationaryProducts := All(products, stationaryProductCriteria)

	fmt.Println("all stationary products?", allStationaryProducts)

	products2 := []Product{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "stationary"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
		{id: 108, name: "Pen", cost: 20, units: 30, category: "stationary"},
	}
	allStationaryProducts = All(products2, stationaryProductCriteria)

	fmt.Println("all stationary products?", allStationaryProducts)

}

func Index(p Product, products []Product) int {
	for i := 0; i < len(products); i++ {
		if p.id == products[i].id {
			return i

		}
	}
	return -1
}

func Include(p Product, products []Product) bool {
	exists := false
	for i := 0; i < len(products); i++ {
		if reflect.DeepEqual(p, products[i]) {
			exists = true
			break
		}
	}
	return exists
}

func Any1(m MatchCriteria, products []Product) bool {
	match := false
	for i := 0; i < len(products); i++ {
		if m.name == products[i].name || m.category == products[i].category {
			match = true
			break
		}
	}
	return match
}

func All1(m MatchCriteria, products []Product) bool {
	match := false
	for i := 0; i < len(products); i++ {
		if m.name == products[i].name && m.category == products[i].category {
			match = true
			break
		}
	}
	return match
}

func Filter1(m MatchCriteria, products []Product) []Product {
	var matchedProducts = []Product{}
	for i := 0; i < len(products); i++ {
		if m.name == products[i].name && m.category == products[i].category {
			matchedProducts = append(matchedProducts, products[i])
		}
	}
	return matchedProducts
}

func Any(products []Product, criteria func(product Product) bool) bool {
	for _, product := range products {
		if criteria(product) {
			return true
		}
	}
	return false
}

func All(products []Product, criteria func(product Product) bool) bool {
	matches := true
	for _, product := range products {
		if !criteria(product) {
			return false
		}
	}
	return matches
}
