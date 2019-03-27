package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	start := time.Now()
	sum := 0
	for i := 0; i < count; i++ {
		sum += i
	}
	sec := time.Since(start).Nanoseconds()
	fmt.Printf("calculate i+...+%d = %d, cost %d\n", count, sum, sec)
}
