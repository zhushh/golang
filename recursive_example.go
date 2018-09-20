package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n > 2 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return 1
}

func main() {
	fmt.Printf("fibonacci(%d) = %d\n", 1, fibonacci(1))
	fmt.Printf("fibonacci(%d) = %d\n", 2, fibonacci(2))
	fmt.Printf("fibonacci(%d) = %d\n", 3, fibonacci(3))
	fmt.Printf("fibonacci(%d) = %d\n", 10, fibonacci(10))
}
