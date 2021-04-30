package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int) {

	x, y := 0, 1
	timerCh := time.After(15 * time.Second)
	for {
		select {
		case ch <- x:
			time.Sleep(1 * time.Second)
			x, y = y, x+y
		case <-timerCh:
			fmt.Println("quitting")
			return
		}
	}

}

func main() {
	fibCh := make(chan int)
	//quit := make(chan string)
	go fibonacci(fibCh)
	go func() {
		for {
			fmt.Println(<-fibCh)
		}
		fmt.Println("Done")
	}()
	fmt.Println("press ENTER to stop")
	var input string
	fmt.Scanln(&input)
	/* time.Sleep(15 * time.Second)
	quit <- "done" */

}
