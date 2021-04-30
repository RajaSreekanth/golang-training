package main

import "fmt"

func main() {
	// array
	var nos [5]int
	fmt.Println("nos length =", len(nos))

	for i := 0; i < len(nos); i++ {
		nos[i] = i
	}
	fmt.Println(nos)

	var names [5]string = [5]string{"raja", "sree", "kanth", "nidadavolu", "name"}
	fmt.Printf("type of names %T\n", names)

	// another way of initializing the Array, without explicitly giving the size
	var products = [...]string{"Pen", "Pencil", "Marker", "Scribble Pad", "Mouse"}
	//Iterating using range
	for id, value := range products {
		fmt.Println(id, value)
	}

	for _, value := range products {
		fmt.Println(value)
	}

	//multidimenstional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row-%d*col-%d", i, j)
		}
	}
	fmt.Printf("Matrix - %v\n", matrix)

	// slices
	var nosSlice []int = []int{9, 4, 3, 6, 1}
	fmt.Println(nosSlice)

	nosSlice = append(nosSlice, 1, 2, 5, 3, 6)
	fmt.Println("After appending, nosSlice", nosSlice)

	anotherSlice := []int{300, 400, 500}
	nosSlice = append(nosSlice, anotherSlice...)
	fmt.Println("After appending, nosSlice", nosSlice)

	// slicing a slice
	/*
		nosSlice[lo: hi]
		all starting with lo index up to (hi -1)
		nosSlice[lo : lo] => empty
		nosSlice[lo : lo+1] => item at 'lo'
		nosSlice[lo:] => all items starting from 'lo'
		nosSlice[: hi] => all the values from 0 to hi-1
	*/

	fmt.Println("nosSlice[3:6]", nosSlice[3:6])
	fmt.Println("nosSlice[3:3]", nosSlice[3:3])
	fmt.Println("nosSlice[3:4]", nosSlice[3:4])
	fmt.Println("nosSlice[3:]", nosSlice[3:])
	fmt.Println("nosSlice[:6]", nosSlice[:6])

}
