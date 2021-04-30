package models

import (
	"fmt"
	"sort"
)

type Products []Product

var currentProductComparer productsComparer

var productsComparers map[string]productsComparer = map[string]productsComparer{
	"id": func(p1, p2 Product) bool {
		return p1.Id < p2.Id
	},
	"name": func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	},
	"cost": func(p1, p2 Product) bool {
		return p1.Cost < p2.Cost
	},
	"units": func(p1, p2 Product) bool {
		return p1.Units < p2.Units
	},
	"category": func(p1, p2 Product) bool {
		return p1.Category < p2.Category
	},
}

func init() {
	fmt.Println("Package initializer triggered")
	currentProductComparer = productsComparers["id"]
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprint(product)
	}
	return result
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

func (products Products) Sort(attrName string) {
	currentProductComparer = productsComparers[attrName]
	sort.Sort(products)
}

func (products Products) SortDesc(attrName string) {
	currentProductComparer = productsComparers[attrName]
	sort.Sort(sort.Reverse(products))
}
