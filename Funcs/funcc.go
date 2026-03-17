package main

import "fmt"

func add(num int , num1 int) {
	sum := num + num1
	fmt.Println(sum)
}

func min(num int , num1 int) {
	sum := num - num1
	fmt.Println(sum)
}

func getNumbers(num int , num2 int) (int,int) {
	sum := num + num2
	mul := num * num2

	return sum , mul
}

func main() {
	
a := 10
b := 20

add(a,b)

min(a,b)
 
getNumbers(a,b)

p, q := getNumbers(a,b)

fmt.Println(p)
fmt.Println(q)

	

}