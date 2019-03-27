package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	count := make(map[string]int)
	if len(files) == 0 {
		countlines(os.Stdin, count)
	} else {
		for _, fileName := range files {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open file fail. err = %v\n", err)
				continue
			}
			countlines(file, count)
			file.Close()
		}
	}

	for key, val := range count {
		if val > 1 {
			fmt.Println(key, val)
		}
	}
}

func countlines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
