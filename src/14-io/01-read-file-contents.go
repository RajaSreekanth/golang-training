package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("Something went wrong! : %v\n", err)
	}
	fmt.Println(string(fileContents))
}
