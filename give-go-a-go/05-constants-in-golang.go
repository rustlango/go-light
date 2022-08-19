// golang const
// const values cannot be changed
// use the const keyword to declare a constant

package main

import "fmt"

func main() {
	const a int = 10
	// cannot reassign a new value to the constant value a - compiler will cry
	// a = 20
	fmt.Println("constant a equals", a)

	// let us create a const block to contain multple constant values
	// a const block is nice because it reduces the boilerplate
	// of having to write out the const type multple times for each separate
	// declaration if we were to do it the normal way
	const (
		// iota is a counter which starts from 0
		b int = iota
		c int = iota
		d int = iota
	)

	fmt.Println("constant b equals", b)
	fmt.Println("constant c equals", c)
	fmt.Println("constant d equals", d)
}
