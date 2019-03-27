package main

import "fmt"

func main() {
    const (
        LENGTH = 10
        WIDTH = 20
    )

    area := LENGTH * WIDTH
    fmt.Println("Length = ", LENGTH)
    fmt.Println("Width = ", WIDTH)
    fmt.Println("area = ", area)
}
