package main

import "fmt"

// import math.rand

func main() {
	count := getCounter()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	main2()
	fmt.Println("back from main1")
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	main2()

}

func main2() {
	fmt.Println("from main2")
	count := getCounter()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}
func getCounter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}
