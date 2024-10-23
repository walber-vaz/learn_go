package chapter_one

import "fmt"

// HelloWorld prints "Hello, World!" to the console.
func HelloWorld() {
	// function Println from the fmt package is used to print to the console
	// https://pkg.go.dev/fmt#Println

	// Println formats using the default formats for its operands and writes to standard output.
	// Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.

	// ... is a variadic parameter that allows the function to accept any number of arguments
	// https://golang.org/ref/spec#Passing_arguments_to_..._parameters

	// Any types implement interface{} which is the empty interface
	// https://golang.org/ref/spec#Interface_types

	// Not variable usage _ is used to ignore the return value
	// https://golang.org/ref/spec#Blank_identifier
	anyBytesWritten, _ := fmt.Println("Hello, World!")
	fmt.Println(anyBytesWritten)
}
