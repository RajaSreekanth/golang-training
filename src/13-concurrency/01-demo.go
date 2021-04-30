package main

import (
	"fmt"
	"time"
)

func print(message string) {
	fmt.Println(message)
}

func main() {
	go print("Hello")
	print("World")
	time.Sleep(1 * time.Second)
}
