package main

import "fmt"

func main() {
	//var x, y int = 10, 20
	x, y := 10, 20
	sum := x + y
	//fmt.Println("Sum = ", sum)
	fmt.Printf("Sum of %d(%T) and %d(%T) is %d(%T)\n", x, x, y, y, sum, sum)

	n1 := 10
	n2 := float64(n1)
	n3 := 1.2
	product := n2 * n3

	fmt.Printf("Product of %f(%T) and %f(%T) is %f(%T)\n", n2, n2, n3, n3, product, product)
}
