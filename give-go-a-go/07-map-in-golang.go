// Maps in golang
// maps are key value pairs
package main

import "fmt"

func main() {
	// This is one way of declaring or creating a map in go using a map literal
	crypto_wallet_balances := map[string]float32{
		"Etheurem": 23.77792,
		"Bitcoin":  212.288004322,
		"Atom":     123.3221,
		"DOT":      2399.999,
	}
	// it will print the map KV properties i alphabetical order
	fmt.Println("Crypto balances in wallet are a follows: ", crypto_wallet_balances)
}
