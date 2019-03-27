package main

import (
    "fmt"
)

func main() {
    var a int = 100
    var b int = 200
    var ptr *int
    var pPtr **int

    fmt.Println("a =", a, ", b =", b)

    ptr = &a
    pPtr = &ptr
    fmt.Println("address of variable a =", ptr)

    c := (*ptr) + b
    fmt.Println("a + b =", c)

    c = (**pPtr) * b
    fmt.Println("a * b =", c)

    var arr = [3]int{3, 5, 8}
    var pArr [3]*int

    for i := 0; i < len(arr); i++ {
        pArr[i] = &arr[i]
    }

    for i := 0; i < len(pArr); i++ {
        fmt.Printf("%d ", *pArr[i])
    }
    fmt.Printf("\n")
}
