package chapter_one

import "fmt"

// Type Primitives
var typeInt = 10
var typeFloat = 10.10
var typeString = "Hello, World!"
var typeBool = true

// Types in Go
// Go is language with static types, which means that the type of a variable is known
// at compile time.
func Types() {
	fmt.Printf("Type of typeInt: %T\n", typeInt)
	fmt.Printf("Type of typeFloat: %T\n", typeFloat)
	fmt.Printf("Type of typeString: %T\n", typeString)
	fmt.Printf("Type of typeBool: %T\n", typeBool)
}
