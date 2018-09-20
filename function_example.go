package main 

import (
    "fmt"
    "os"
    "strconv"
)

func MaxNum(num1 int, num2 int) int {
    var result int

    if (num1 < num2) {
        result = num2
    } else {
        result = num1
    }

    return result
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: cmd number1 number2")
    } else {
        var a int 
        var b int
        var ret int

        a, _ = strconv.Atoi(os.Args[1])
        b, _ = strconv.Atoi(os.Args[2])

        ret = MaxNum(a, b)
        fmt.Printf("Max number: %d\n", ret)
    }
}
