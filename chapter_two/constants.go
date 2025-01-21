package chapter_two

import "fmt"

// Constants are declared like variables, but with the const keyword.
const (
	// Pi is a float64 constant.
	Pi = 3.14159265358979323846
	// E is a float64 constant.
	E = 2.7182818284590452354
)

// Constants can be declared in blocks like variables.
func Constants() {
	fmt.Println("Pi:", Pi)
	fmt.Println("E:", E)
}
