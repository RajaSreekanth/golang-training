package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// initialize a slice
	var mySlice []int
	for i := 0; i < 20; i++ {
		mySlice = append(mySlice, rand.Intn(20))
	}
	fmt.Println("mySlice", mySlice)
	var oddSlice []int
	var evenSlice []int
	for _, value := range mySlice {
		if value%2 == 0 {
			evenSlice = append(evenSlice, value)
		} else {
			oddSlice = append(oddSlice, value)
		}
	}

	fmt.Println("evenSlice", evenSlice)
	fmt.Println("oddSlice", oddSlice)

	// maps

	cityRanks := map[string]int{
		"Udipi":     2,
		"Bangaluru": 4,
		"Mangaluru": 1,
		"Mysuru":    3,
	}

	fmt.Printf("Type of cityRanks %T", cityRanks)
	fmt.Println(cityRanks)
	// add values
	cityRanks["Chennai"] = 5
	fmt.Println(cityRanks)

	for city, rank := range cityRanks {
		fmt.Printf("City = %q, Rank = %d\n", city, rank)
	}

	fmt.Println("Hyderabad rank", cityRanks["Chennai"])

	chennaiRank, exists := cityRanks["Chennai"]
	if exists {
		fmt.Println("rank of Chennai", chennaiRank)
	} else {
		fmt.Println("No rank for Chennai")
	}

	if HydRank, exists := cityRanks["Hyderabad"]; exists {
		fmt.Println("rank of Hyderabad", HydRank)
	} else {
		fmt.Println("Hyderabad is not ranked")
	}

	//remove an key/value pair
	delete(cityRanks, "Chennai")
	fmt.Println("After deleting Chennai")
	if chennaiRank, exists := cityRanks["Chennai"]; exists {
		fmt.Println("Rank of Chennai is ", chennaiRank)
	} else {
		fmt.Println("Chennai is not ranked")
	}
}
