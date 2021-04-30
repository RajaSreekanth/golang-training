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
		switch operation {
		case 1:
			display(log(a, b, add))
		case 2:
			display(log(a, b, multiply))
		case 3:
			display(log(a, b, sub))
		case 4:
			display(log(a, b, divide))
		default:
			fmt.Println("Try again")
		}
	}
}

func log(n1, n2 int32, operation func(int32, int32) int32) int32 {
	fmt.Printf("Processing %d and %d\n", n1, n2)
	return operation(n1, n2)
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
	fmt.Printf("Result is %d\n", result)
}
func greet(val string) string {
	return fmt.Sprintf("Hi %s, have a nice day!", val)
}

func add(a, b int32) int32 {
	return a + b
}

func sub(a, b int32) int32 {
	return a - b
}

func multiply(a, b int32) int32 {
	return a * b
}

func divide(a, b int32) int32 {
	return a / b
}
