package main

import "fmt"

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	add(2, 4)
	//anonymous func
	//IIFE

	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(5, 7)
}

func init() {
	fmt.Println("i.ll be callled first ")
}
