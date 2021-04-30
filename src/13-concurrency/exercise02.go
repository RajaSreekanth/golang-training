package main

import (
	"fmt"
	"time"
)

func add(data []int, resultCh chan int) {
	result := 0
	for _, v := range data {
		result += v
	}
	resultCh <- result
}

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	start := time.Now()
	oddData := []int{}
	evenData := []int{}
	for idx, v := range data {
		if idx%2 == 0 {
			evenData = append(evenData, v)
		} else {
			oddData = append(oddData, v)
		}
	}
	sum1 := make(chan int)
	sum2 := make(chan int)
	go add(evenData, sum1)
	go add(oddData, sum2)
	fmt.Printf("total is %d\n", <-sum1+<-sum2)
	elapsed := time.Since(start)
	fmt.Printf("time taken is %v\n", elapsed)
}
