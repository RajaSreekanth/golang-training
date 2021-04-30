package main

import "fmt"

func main() {
	x := 10
	//address of x => &x
	fmt.Printf("value of x is %d\n", x)
	fmt.Printf("address of x is %v\n", &x)

	var addressOfX *int = &x
	fmt.Printf("address of x is %v\n", addressOfX)

	//accessing the value using the address is called "dereferencing"
	fmt.Printf("Value of x using the address of x %d\n", *addressOfX)

	n1 := 10
	n2 := 20
	fmt.Printf("Before swapping n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping n1 = %d and n2 = %d\n", n1, n2)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
