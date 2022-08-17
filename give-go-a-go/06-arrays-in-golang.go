// Arrays in golang
// Arrays are homogenous containers used to store multiple elements
// of the same data type

package main

import "fmt"

func main() {
	// one way of declaring an array
	// size of array defined as 4. First element starts from 0
	var fruits [4]string

	fruits[0] = "Orange"
	fruits[1] = "Apple"
	fruits[2] = "Stawberry"
	fruits[3] = "Litchi"
	fmt.Println("Fruits array contains:", fruits)

	// one way of declaring an array
	// size of array defined as 4. First element starts from 0
	crypto_balance := [4]float32{12226271.3222, 3272726.3213, 422.322, 43777.4432211}
	fmt.Println("List of crypto balances:", crypto_balance)
}
