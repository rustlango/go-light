// Operators in golang

// There are three main categories of operators in golang
// 1. Arithmetic Operators
// 2. Relational Operators
// 3. Logical Operators
// && is an AND logical operator and || is the OR operator

package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20
	fmt.Println("what is c int divided by d int?", a/b) // will be zero because the data types are integers not float
	fmt.Println("is a greater than b?", a > b)        // false
	fmt.Println("is a equal to b?", a == b)           //  false: 10 is ot equal to b
	fmt.Println("is a not equal to b?", a != b)       // true: a is not equal to b
	var c float32 = 10
	var d float32 = 20
	fmt.Println("what is c float32 / d float32?", c/d) // returns 0.5 because flota32 type not ineteger
	fmt.Println("is c less than d?", c < d)            // true
	fmt.Println("is c = to d?", c == d)                //  false: c is not equal to d
	fmt.Println("is c not = to d?", c != d)            // true: c is not equal to d
	// && is an AND logical operator and || is the OR operator
	var e bool = false
	var f bool = true
	fmt.Println("e AND f", e && f) // returns false
	fmt.Println("e OR f", e || f)  // returns true
}
