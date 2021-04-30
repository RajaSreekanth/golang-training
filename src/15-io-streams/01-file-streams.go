package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	data1Ch := make(chan string)
	data2Ch := make(chan string)

	go readFile("data1.dat", data1Ch)
	go readFile("data2.dat", data2Ch)

	oddCh := make(chan int)
	evenCh := make(chan int)

	go splitter(data1Ch, oddCh, evenCh)
	go splitter(data2Ch, oddCh, evenCh)
	quit := make(chan string)
	oddSum := make(chan int)
	evenSum := make(chan int)
	go aggregate(oddCh, evenCh, oddSum, evenSum, quit)

	fmt.Println("press ENTER to stop")
	var input string
	fmt.Scanln(&input)
	quit <- "stop"
	go writeResults("result.txt", oddSum, evenSum, quit)
}

func readFile(fileName string, ch chan string) {
	fmt.Println("reading", fileName)
	fileHandle, fileError := os.Open(fileName)
	defer fileHandle.Close()
	if fileError != nil {
		log.Fatalln(fileError)
	}
	inputReader := bufio.NewReader(fileHandle)
	for {
		line, err := inputReader.ReadString('\n')
		fmt.Println("read", line)
		ch <- line
		if err == io.EOF {
			break
		}
	}
}

func splitter(inCh chan string, odd chan int, even chan int) {
	fmt.Println("Splitting off/even")
	for str := range inCh {
		x, err := strconv.Atoi(str)
		if err != nil {
			if x%2 == 0 {
				fmt.Println("even", x)
				even <- x
			} else {
				fmt.Println("odd", x)
				odd <- x
			}

		}
	}
	close(odd)
	close(even)
}

func aggregate(oddCh, evenCh, oddSum, evenSum chan int, quit chan string) {
	fmt.Println("aggregating odd and even values")
	o, e := 0, 0
	oddSum1, evenSum1 := 0, 0
	select {
	case o = <-oddCh:
		fmt.Println("adding to oddSum", o)
		oddSum1 += o
	case e = <-evenCh:
		fmt.Println("adding to evenSum", e)
		evenSum1 += e
	case <-quit:
		oddSum <- oddSum1
		evenSum <- evenSum1
	}
	return
}

func writeResults(fileName string, odd, even chan int, quit chan string) {
	<-quit
	result := fmt.Sprintf("odd = %d\n even = %d\n", <-odd, <-even)
	err := ioutil.WriteFile("result.txt", []byte(result), 0x777)
	if err != nil {
		log.Fatalf("Something went wrong : %v\n", err)
	}
}
