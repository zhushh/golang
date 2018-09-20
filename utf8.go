package main

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	f.WriteString("Hello, world!\n")
	f.Close()
}
