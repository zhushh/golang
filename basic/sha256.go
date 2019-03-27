package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "niceashaiaukAEDO;AFSBKLG"
	cs := sha256.Sum256([]byte(s))
	fmt.Printf("%x\n", cs)

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n", c1, c2)
}
