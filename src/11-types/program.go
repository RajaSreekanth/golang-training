package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x interface{}
	x = "abc"
	x = 100
	x = true

	fmt.Printf("type of x = %T\n", x)
	if value, ok := x.(int); ok {
		fmt.Println(value + 100)
	} else {
		fmt.Printf("%v is not an int\n", x)
	}

	switch v := x.(type) {
	case int:
		fmt.Printf("Twice of %v  is %v\n", v, 2*v)
	case string:
		fmt.Printf("Len of %v is %d\n", v, len(v))
	default:
		fmt.Printf("Unknown type of %v = %T\n", v, v)
	}

	fmt.Printf("sum(10, 20, '30' is %d\n", sum(10, 20, "30"))
	fmt.Printf("sum(10,20,'30',40,50,'abc') is %d\n", sum(10, 20, "30", 40, 50, "abc"))
	fmt.Printf("sum(10,20,[]int{30,40,50}) is %d\n", sum(10, 20, []int{30, 40, 50}))
	fmt.Printf("sum(10,20,[]interface{}{'30',40,50,'abc'}) is %d\n", sum(10, 20, []interface{}{"30", 40, 50, "abc"}))
}

func sum(items ...interface{}) int {
	result := 0
	for i := 0; i < len(items); i++ {
		switch v := items[i].(type) {
		case int:
			result += v
		case string:
			// k, err := strconv.Atoi(v)
			// if err == nil {
			// 	result += k
			// }
			if intVal, ok := strconv.Atoi(v); ok == nil {
				result += intVal
			}
		case []int:
			// fmt.Printf("in []int case: %v of type %T", v, v)
			data := make([]interface{}, len(v))
			for id, value := range v {
				data[id] = value
			}
			result += sum(data...)
		case []interface{}:
			// fmt.Printf("in []interface{} case: %v of type %T", v, v)
			result += sum(v...)
		default:
			break
		}
	}
	return result
}

/*
sum() => 0
sum(10) => 10
sum(10,20) => 30
sum(10,20,30,40,50) => 150
sum(10,20,"30",40,50) => 150
sum(10,20,"30",40,50,"abc") => 150
sum(10,20,[]int{30,40,50}) => 150
sum(10,20,[]interface{}{"30",40,50,"abc"}) => 150
*/
