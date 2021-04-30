package main

import (
	"fmt"
	mod "modularity/models"

	"github.com/fatih/color"
)

func main() {
	products := mod.Products{
		{Id: 100, Name: "Pen", Cost: 10, Units: 10, Category: "stationary"},
		{Id: 106, Name: "Den", Cost: 16, Units: 50, Category: "stationary"},
		{Id: 107, Name: "Ken", Cost: 12, Units: 20, Category: "utencil"},
		{Id: 102, Name: "Zen", Cost: 18, Units: 70, Category: "stationary"},
		{Id: 104, Name: "Ten", Cost: 15, Units: 30, Category: "utencil"},
		{Id: 103, Name: "Len", Cost: 14, Units: 50, Category: "stationary"},
	}
	color.Yellow("Initial List")
	fmt.Println(products)

	color.Blue("Sort By id")
	products.Sort("id")
	fmt.Println(products)

	color.Magenta("Sort By id - descending")
	products.SortDesc("id")
	fmt.Println(products)

	fmt.Println("Sort By name")
	products.Sort("name")
	fmt.Println(products)

	fmt.Println("Sort By cost")
	products.Sort("cost")
	fmt.Println(products)

	fmt.Println("Sort By units")
	products.Sort("units")
	fmt.Println(products)

	fmt.Println("Sort By category")
	products.Sort("category")
	fmt.Println(products)
}
