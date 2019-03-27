package main

import "fmt"

func main() {
    var b int = 15
    var a int 

    numbers := [6]int{1, 2, 3, 5}

    for a := 0; a < 10; a++ {
        fmt.Printf("a = %d\n", a)
    }

    for a < b {
        fmt.Printf("a = %d\n", a)
        a++
    }

    for i,x := range numbers {
        fmt.Printf("The %dth x = %d\n", i, x)
    }
}
