package chapter_two

import "fmt"

// TypeNumeric is a function that demonstrates the use of the numeric types in Go.
func TypeNumeric() {
	typeInt := 10

	fmt.Println("Decimal - Type - Binary - Hexadecimal")
	fmt.Printf("%d - %T - %b - %#x\n", typeInt, typeInt, typeInt, typeInt)
}
