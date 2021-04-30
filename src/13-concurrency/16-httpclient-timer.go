package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getResponse(ch chan *http.Response) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalf("Something went wrong : %v", err)
	}
	ch <- res
}

func main() {
	resCh := make(chan *http.Response)
	go getResponse(resCh)
	timeoutCh := time.After(50 * time.Millisecond)
	for {
		select {
		case res := <-resCh:
			io.Copy(os.Stdout, res.Body)
			return
		case <-timeoutCh:
			fmt.Println("Timeout occured")
			return
		}
	}
}
