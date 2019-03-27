package main

import (
	"fmt"
	"time"
)

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func logTime(f func(), funcName string, times int) {
	start := time.Now()
	for i := 0; i < times; i++ {
		f()
	}
	nansec := time.Since(start).Nanoseconds()
	fmt.Printf("call %s %dtimes, cost time %dnansec\n", funcName, times, nansec)
}

func main() {
	x := 122
	times := 1000000

	logTime(func() { PopCount(uint64(x)) }, "PopCount", times)
	logTime(func() { PopCount2(uint64(x)) }, "PopCount2", times)
}

func PopCount(x uint64) int {
	cnt := 0
	for i := 0; i < 8; i++ {
		cnt += int(pc[byte(x>>uint(i*8))])
	}
	return cnt
}

func PopCount2(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
