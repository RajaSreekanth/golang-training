package main

import (
	"errors"
	"fmt"
)

func main() {
	//success condition
	if res1, err := divide(100, 5); err == nil {
		fmt.Println("Dividing 100 by 5 = ", res1)
	} else {
		fmt.Println("Error : ", err)
	}

	if res2, err := divide(100, 0); err == nil {
		fmt.Println("Dividing 100 by 0 = ", res2)
	} else {
		fmt.Println("Error : ", err)
	}

}

func divide(n1, n2 int) (result int, err error) {
	if n2 == 0 {
		err = errors.New("invalid operand")
		return
	}
	result = n1 / n2
	return
}
