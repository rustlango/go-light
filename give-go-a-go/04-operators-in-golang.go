// Operators in golang

// There three main categories of operators in goalng
// 1. Arithmetic Operators
// 2. Relational Operators
// 3. Logical Operators

package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20
	fmt.Println(a / b)  // will be zero because the datat types are integers not float
	fmt.Println(a > b)  // false
	fmt.Println(a == b) //  false: 10 is ot equal to b
	fmt.Println(a != b) // true: a is not equal to b
	var c float32 = 10
	var d float32 = 20
	fmt.Println(c / d)  // returns 0.5 because flota32 type not ineteger
	fmt.Println(c < d)  // true
	fmt.Println(a == b) //  false: c is not equal to d
	fmt.Println(a != b) // true: c is not equal to d
}
