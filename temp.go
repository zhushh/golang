package main

import (
	"fmt"
)

func main() {
	var vals []int
	for i := 0; i < 10; i++ {
		vals = append(vals, i)
	}

	fmt.Println(vals[:0])
	fmt.Println(vals)
}
