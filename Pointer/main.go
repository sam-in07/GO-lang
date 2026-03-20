package main

import "fmt"

func main() {
    x := 42
    var p *int = &x // p stores the address of x

    fmt.Println("x =", x)      // 42
    fmt.Println("p =", p)      // memory address of x
    fmt.Println("*p =", *p)    // 42 (value at address)

    *p = 100 // change x through the pointer
    fmt.Println("x after *p =", x) // 100
}