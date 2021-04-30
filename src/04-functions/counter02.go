package main

import "fmt"

// import math.rand

func main() {
	increment, decrement := getCounters()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
	fmt.Println(decrement())

	main2()
	fmt.Println("back from main1")
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
	fmt.Println(decrement())
	main2()

}

func main2() {
	fmt.Println("from main2")
	increment, decrement := getCounters()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
}

/* multiple return function type 1
func getCounters() (func() int, func() int) {
	var count int = 0
	var func1 = func() int {
		count++
		return count
	}

	var func2 = func() int {
		count--
		return count
	}

	return func1, func2
}
*/

/* multiple return function type 1 */
func getCounters() (increment func() int, decrement func() int) {
	var count int = 0
	increment = func() int {
		count++
		return count
	}

	decrement = func() int {
		count--
		return count
	}

	return
}
