package main

import (
    "fmt"
    "math"
)

// 常量的定义
const PI float64 = 3.14

type Circle struct {
    radius float64
}

// 这是go语言的方法
// 类似C++语言里的类方法
// 需要先定义上面的数据类型 Circle
func (c Circle) getArea() float64 {
    return PI * c.radius * c.radius
}

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    // 多返回值函数的例子
    a, b := swap("Mark", "John")
    fmt.Println(a, b)

    // 函数作为参数的例子
    getSquareRoot := func(x float64) float64 {
        return math.Sqrt(x)
    }
    fmt.Println(getSquareRoot(64))

    // 方法的例子
    var c1 Circle
    c1.radius = 10.00
    fmt.Println("Circle area = ", c1.getArea())
}

