package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("sample.txt")
	defer fileHandle.Close()
	if fileError != nil {
		log.Fatalln(fileError)
	}
	inputReader := bufio.NewReader(fileHandle)
	for {
		// to read sentence by sentence use .
		line, err := inputReader.ReadString('.')
		// to read line by line use \n
		// line, err := inputReader.ReadString('\n')
		fmt.Println(line)
		if err == io.EOF {
			break
		}
	}
	fileHandle.Close()
}
