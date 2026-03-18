package main

import "fmt"

var (
	a = 20
	b = 30
)

func printNum(num int) {
	fmt.Println(num)
}

func add(x int , y int)  {
	res := x + y
	println(res)
}

func main()  {
	add(a,b)
}