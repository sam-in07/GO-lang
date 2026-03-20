package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Value receiver
func (p Person) CelebrateBirthday() {
    p.Age += 1
    fmt.Println("Happy Birthday,", p.Name, "Age:", p.Age)
}

func main() {
    p1 := Person{Name: "Alice", Age: 25}
    p1.CelebrateBirthday() // Age increases only inside the method
    p2 := Person{Name : "samin" , Age : 24}
	p2.CelebrateBirthday()
	fmt.Println("Outside Age:", p1.Age) // Still 25
}