package main

import (
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("Something went wrong : %v\n", err)
	}
	err = ioutil.WriteFile("sample-copy.txt", fileContents, 0x777)
	if err != nil {
		log.Fatalf("Something went wrong : %v\n", err)
	}
}
