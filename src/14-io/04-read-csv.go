package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	data := `
firstname, lastname, age
magesh, kuppan, 42
suresh, kannan, 40
ramesh, jayaraman, 44
`
	reader := csv.NewReader(strings.NewReader(data))
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
