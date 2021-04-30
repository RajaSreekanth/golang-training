package main

import "fmt"

// import math.rand

func main() {
	// name := "Sree"
	// greeting := greet(name)
	// fmt.Println(greeting)
	for {
		operation := getOperation()
		if operation == 5 {
			break
		}
		fmt.Println("Enter first operand")
		var a, b int32
		fmt.Scanf("%d", &a)
		fmt.Println("Enter second operand")
		fmt.Scanf("%d", &b)
		operations := map[int]func(int32, int32) (int32, error){
			1: add,
			2: multiply,
			3: sub,
			4: divide,
		}

	}
}

func getOperation() int {
	var op int
	fmt.Println()
	fmt.Println("select an operation")
	fmt.Println("1 - add")
	fmt.Println("2 - multiply")
	fmt.Println("3 - substract")
	fmt.Println("4 - divide")
	fmt.Println("5 - Exit")
	fmt.Scanf("%d\n", &op)
	return op
}

func display(result int32) {
	fmt.Println("Result is %d\n", result)
}
func greet(val string) string {
	return fmt.Sprintf("Hi %s, have a nice day!", val)
}

func add(a, b int32) (int32, error) {
	return a + b
}

func sub(a, b int32) (int32, error) {
	return a - b
}

func multiply(a, b int32) int32 {
	return a * b
}

func divide(a, b int32) (int32, error) {
	return a / b
}
