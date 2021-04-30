package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.After(5 * time.Second)
	fmt.Println(<-x)
}
