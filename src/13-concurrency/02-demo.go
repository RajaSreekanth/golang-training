package main

import (
	"fmt"
	"time"
)

func print(message string, ch chan string) {
	time.Sleep(3 * time.Second)
	fmt.Println(message)
	ch <- "done"
}

func main() {
	ch := make(chan string)
	go print("Hello", ch)
	<-ch
	fmt.Println("Done")
}

/*
	for writing data into the channel
		ch <- data

	for reading data from the channel
		x := <- ch

*/
