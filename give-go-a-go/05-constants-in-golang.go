// golang const
// values cannot be changed
// use the const keyword to declare a constant

package main

import "fmt"

func main() {
	const a int = 10
	// cannot reassign a new value to the constant value a - compiler will cry
	a = 20
	fmt.Println("constant a equals", a)
}
