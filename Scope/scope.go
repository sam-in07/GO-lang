// package main

// import "fmt"

// var a = 20
// var b = 30

// func add(x int, y int) {
// 	z := x + y
// 	fmt.Println(z)
// }

// func main() {
// 	p := 30
// 	q := 40
//    add(p,q)

//    add(a,b)

//    add(a,p)

//  //  add(b,z)

// }

//coder kono portion a scope/access korte parbo kina

/* local scope and block  */

package main

import (
	"fmt"
	

	"example.com/mathlib"
)

var (
	a = 20
	b = 30
)

/*
1.block => { }
2. package scope : go mod init example.com

*/

// func main(){
// 	x := 18

// 	if x >= 18{
// 		p := 10
// 		fmt.Println("im nega unmature")
// 		fmt.Println("i have", p, "negga gf")
// 	}
// }

//packgae scope

func main() {
	fmt.Println("Showing custome palage")
	mathlib.Add(4,7)
}

//cmd go run scope.go add.go




//go mod init example.com
