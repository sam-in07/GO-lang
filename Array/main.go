package main

import "fmt"

func main() {
    var numbers [5]int // array of 5 integers

    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30

    fmt.Println(numbers)     // [10 20 30 0 0]
    fmt.Println(numbers[2])  // 30
    fmt.Println(len(numbers)) // 5
}