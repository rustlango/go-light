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
	// golang supports interpolation so use %d
	fmt.Printf(" c is = to %v, of type %T, a+b+c = %d\n", c, c, a+b+c)
	// retruns : 100, int %!{(int=130)}
	// compiler has the ability infer the type for you though implicit variable
	// variable declaration is not always encouraged

	// In go you can decalre multiple variables in a variable block as shown
	// below
	// this variable block syntax is for variabbles that are always use together
	var (
		stu_name  string  = "Lemuel"
		stu_marks float32 = 98.07
		stu_id    int     = 14
	)

	fmt.Printf("%v %T\n", stu_name, stu_name)
	fmt.Printf("%v %T\n", stu_marks, stu_marks)
	fmt.Printf("%v %T\n", stu_id, stu_id)
}
