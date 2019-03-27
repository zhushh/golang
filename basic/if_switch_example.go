package main

import "fmt"

func main() {
    var grade = 80

    if grade > 90 {
        fmt.Println("grade = A")
    } else if grade > 75 {
        fmt.Println("grade = B")
    } else {
        fmt.Println("grade = C")
    }

    var x interface{}
    switch i := x.(type) {
    case nil:
        fmt.Printf("grade 的类型: %T", i)
    case int:
        fmt.Printf("grade 的类型: int")
    default:
        fmt.Printf("未知型")
    }
}

