package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	endCh := make(chan string)
	go func() {
		fmt.Println("\tAttempting to write the data into the channel")
		ch <- "Hello"
		fmt.Println("\tAttempt to write the data into the channel completed")
		fmt.Println("\tAttempting to end the go routine")
		<-endCh
		fmt.Println("\tAttempting to end the go routine completed")
	}()
	fmt.Println("Attempting to read the data from the channel")
	fmt.Printf("Data from channel = %v\n", <-ch)
	fmt.Println("Attempt to read the data from the channel completed")
	fmt.Println("Delaying the ending of the goroutine by 5 seconds")
	time.Sleep(5 * time.Second)
	endCh <- "stop"
	fmt.Println("press ENTER to shutdown!")
	var input string
	fmt.Scanln(&input)
}
