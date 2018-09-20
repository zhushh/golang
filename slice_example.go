package main

import (
	"fmt"
)

func main() {
	// 通过make方式创建切片
	var numbers = make([]int, 5, 7)

	// 中括号中没有指定长度，表示声明的是slice，带长度时为数组
	data := []int{3, 5, 7, 1}

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(data), cap(data), data)

	// 正常数组使用
	for i, num := range data {
		numbers[i] = num
	}
	fmt.Println("numbers =", numbers)

	// 切片使用
	numbers1 := numbers[:2]
	fmt.Println("numbers1 =", numbers1)

	// 切片使用
	numbers2 := numbers[2:5]
	fmt.Println("numbers2 =", numbers2)

	// 切片使用
	numbers3 := numbers[3:]
	fmt.Println("number3 =", numbers3)

	// append
	numbers = append(numbers, 8, 9, 2)
	fmt.Println("numbers =", numbers)

	// copy
	numbers4 := make([]int, len(numbers), (cap(numbers) * 2))
	copy(numbers4, numbers)
	fmt.Println("number4 =", numbers4)
}
