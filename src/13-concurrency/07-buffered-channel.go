package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	go func() {
		fmt.Println("\tAttempting to read from channel")
		<-ch
		fmt.Println("\tAttempt to read from channel completed")
	}()
	// fmt.Println("Attempting to read the data from the channel")
	// fmt.Printf("Data from channel = %v\n", <-ch)
	// fmt.Println("Attempt to read the data from the channel completed")

	fmt.Println("press ENTER to shutdown!")
	var input string
	fmt.Scanln(&input)
}
