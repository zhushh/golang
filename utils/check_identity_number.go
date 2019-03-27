package main

import "fmt"

func CheckIdentity(number string) bool {
	l := len(number)
	if l < 18 {
		return false
	}

	iS := 0
	iW := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	verCode := "10X98765432"

	for i := 0; i < 17; i++ {
		iS += int(number[i]-'0') * iW[i]
	}

	iY := iS % 11
	verEndCode := verCode[iY : iY+1]
	endCode := number[17:18]

	if verEndCode == endCode {
		return true
	}

	return false
}

func main() {
	number := ""
	fmt.Scan(&number)
	fmt.Printf("Check result: %v", CheckIdentity(number))
}
