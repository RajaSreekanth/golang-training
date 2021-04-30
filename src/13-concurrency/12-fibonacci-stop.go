package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int, quit chan string) {

	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(2 * time.Second)
			x, y = y, x+y
		case <-quit:
			break
		}
	}
	fmt.Println("Closing the channel")
	close(ch)
}

func main() {
	fibCh := make(chan int, 1)
	quit := make(chan string)
	go fibonacci(fibCh, quit)
	go func() {
		for fibNo := range fibCh {
			fmt.Println(fibNo)
		}
	}()
	fmt.Println("press ENTER to stop")
	var input string
	fmt.Scanln(&input)
	quit <- "done"
	fmt.Println("Done")
}
