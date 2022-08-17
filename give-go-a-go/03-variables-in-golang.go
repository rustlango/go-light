// Variables in golang

// Data/Values can be stored in temporary storage spacs called variables
// in rust they are called bindings

// main papckage is the entry point of the golang program
package main

// fmt is the standard library f golang
import "fmt"

func main() {
	// how to create a variable in golang
	// explicit variable decalration
	var a int = 10
	var b int = 20
	// implicit variable decalration for type inference
	c := 100
	// Println() is a  method from the standard fmt library to
	// print to the screen
	fmt.Println("Hello, World")
	// %v returns the value and %T returns the type
	fmt.Printf("%v, %T %()", c, c, a+b+c)
	// retruns : 100, int %!{(int=130)}
	// compiler has the ability infer the type for you though implicit variable
	// variable declaration is not always encouraged
}
