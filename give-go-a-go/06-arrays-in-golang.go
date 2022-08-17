// Arrays in golang
// Arrays are homogenous ontainers used to store multple elements
// of the same data type

package main

import "fmt"

func main() {

	var fruits [4]string

	fruits[0] = "Orange"
	fruits[1] = "Apple"
	fruits[2] = "Stawberry"
	fruits[3] = "Litchi"
	fmt.Print("Fruits array contains:", fruits)
}
