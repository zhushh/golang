package main

import (
	"bufio"
	"fmt"
	"os"
)

const max_n = 10000

var d [max_n][max_n]int
var a, b string

func main() {
	input()
	prepare()
	computeEditLength()
	debug_log()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Edit length =", d[len(a)][len(b)])
}

func input() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	a = input.Text()
	input.Scan()
	b = input.Text()
}

func prepare() {
	for i := 0; i <= len(a); i++ {
		d[i][0] = i
	}
	for i := 0; i <= len(b); i++ {
		d[0][i] = i
	}
}

func computeEditLength() {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				d[i+1][j+1] = min(d[i+1][j]+1, min(d[i][j+1]+1, d[i][j]))
			} else {
				d[i+1][j+1] = min(d[i+1][j]+1, min(d[i][j+1]+1, d[i][j]+1))
			}
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func debug_log() {
	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			fmt.Printf("%d ", d[i][j])
		}
		fmt.Println("")
	}
}
