package main

import (
	"fmt"
)

// 数组声明如下
// var array_name [array_size] array_type

// 多维数组使用如下
// var array_name [SIZE1][SIZE2]...[SIZEN] array_type

func arrayTest() {
	// 数组的声明必须指定长度
	var arr [6]int
	var data = [6]int{2, 4, 6, 8, 10}

	for i, item := range data {
		arr[i] = item
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
}

func multiArrayTest() {
	// 此处初始化不能直接换行
	var a = [5][2]int{{0, 0}, {1, 2}, {4, 6}, {3, 5}, {4, 8}}

	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func main() {
	arrayTest()
	multiArrayTest()
}
