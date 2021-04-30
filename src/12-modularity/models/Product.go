package models

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func Sum(x, y int) int {
	return x + y
}

/* Methods for Stringer interface of the 'fmt' package */
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, cost = %v, units = %d, category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

// only memebers start with upper case are visible outside the package when imported
