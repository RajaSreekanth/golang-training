package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, _ := os.Open("emp.csv")
	reader := csv.NewReader(fileHandle)
	header := true
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		if header {
			header = false
		} else {
			fmt.Println(line[0], line[1], line[2])
		}
	}

}
