package main

import "fmt"

func main() {
	defer onMainExit()
	fmt.Println("Main is executed")
}

func onMainExit() {
	fmt.Println("Main is exiting")
}
