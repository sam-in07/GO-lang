package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of the struct
    p1 := Person{Name: "Alice", Age: 25}

    fmt.Println(p1)
    fmt.Println("Name:", p1.Name)
    fmt.Println("Age:", p1.Age)
}

/*
type Book struct {
    Title  string
    Author string
    Pages  int
    Price  float64
}
*/