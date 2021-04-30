package main

import (
	"fmt"
)

func main() {
	// no := 30
	// if no%2 == 0 {
	// 	fmt.Printf("%v is even\n", no)
	// } else {
	// 	fmt.Printf("%v is odd\n", no)
	// }

	// // for construct
	// x := 30
	// for i := 0; i < x; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%v is even\n", i)
	// 	} else {
	// 		fmt.Printf("%v is odd\n", i)
	// 	}
	// }

	// // random values
	// for i := 0; i < 100; i++ {
	// 	mval := rand.Intn(1000)
	// 	if mval%2 == 0 {
	// 		fmt.Printf("%v is even\n", mval)
	// 	} else {
	// 		fmt.Printf("%v is odd\n", mval)
	// 	}
	// }

	// for as while construct

	i := 1
	sum := 0
	// for i < 101 {
	// 	sum += i
	// 	i++
	// }

	for {
		sum += i
		i++
		if i > 100 {
			break
		}
	}
	fmt.Printf("Sum = %v\n", sum)

	b := 4
	remainder := b % 2
	switch remainder {
	case 0:
		fmt.Printf("%v is even\n", b)
	case 1:
		fmt.Printf("%v is odd\n", b)
	}

	var n int
	// switch with fallthrough
	fmt.Println("Enter a value")
	fmt.Scanf("%d", &n)

	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	default:
		fmt.Println("Try again")
	}

	var name string
	fmt.Println("enter a name")
	fmt.Scanf("%s", &name)
	switch name {
	case "sree":
		fmt.Println("Hi sree")
	default:
		fmt.Println("who are you?")
	}
}
