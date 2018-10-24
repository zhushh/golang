package main

import (
	"fmt"
	"os"
)

func GetCmdName() string {
	return os.Args[0]
}

func Test() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	cmdName := GetCmdName()
	fmt.Println(cmdName)
	fmt.Println("Arguments: ", os.Args[1:])
}

func name() {

}
