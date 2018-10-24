package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	//"strconv"
)

type A struct{}

func (a A) String() string { return "A" }

func inti() {
	fmt.Println("The init function will be called firstly.")
}

func main() {
	//testStruct()

	//testBufio()

	//testHttp()

	//testCountBit()

	//testRawString()

	// GetInput()

	iNow := time.Now().Unix()
	fmt.Println(iNow)

	//arr := new([]int)
	//*arr = append(*arr, 2)
	//*arr = append(*arr, 3)
	//fmt.Println(*arr)
	a := 123
	fmt.Println(a/4)
}

func testSlice(a []int) []int {
	a = append(a, 3)
	a = append(a, 5)
	return a
}

func GetInput() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		fmt.Printf("%s ", word)
	}
}

func testStruct() {
	a := new(A)
	fmt.Printf("%s\n", a)
}

func testBufio() {
	input := "nice    to  meet you"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	data := "nice\nto\nmeet\nyou"
	scanner = bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func testHttp() {
	url := "http://www.qq.com"
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	robots, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)

	fmt.Printf("Status = %d\n", rsp.StatusCode)
	for k, v := range rsp.Header {
		fmt.Printf("Header[%q] = %q\n", k, v)
	}
}

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func testCountBit() {
	x := 12863
	fmt.Printf("%d, %d\n", x, PopCount(uint64(x)))
}

func testRawString() {
	const GoUsage = `go is a tool for managing Go source code.

    Usage:
        go command [argument]
    `

	fmt.Printf("%s", GoUsage)
}
