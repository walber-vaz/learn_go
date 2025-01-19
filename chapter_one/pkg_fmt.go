package chapter_one

import (
	"fmt"
	"strings"
)

// PackageFmt demonstrates the difference between interpreted and raw string literals.
// Interpreted string literals are character sequences between double quotes "".
// Raw string literals are character sequences between backticks “.
func PackageFmt() {
	// Interpreted string literals
	// Interpreted string literals are character sequences between double quotes "".
	interpretedStringLiterals := "Usage Interpreted string literals: Hello, World!\nStudents are learning Go!"

	// Raw string literals
	// Raw string literals are character sequences between backticks ``.
	rawStringLiterals := `Usage Raw string literals: Hello, World!\nStudents are learning Go!`

	fmt.Println(interpretedStringLiterals)
	fmt.Println(rawStringLiterals)

	// Sprint
	// Sprint formats using the default formats for its operands and returns the resulting string.
	usageSprint := "Usage Sprint: Hello, World!\nStudents are learning Go!"
	fmt.Println(usageSprint)

	// Fprint
	// Fprint formats using the default formats for its operands and writes to w.
	var buffer strings.Builder
	_, err := fmt.Fprint(&buffer, "Usage Fprint: Hello, World!\nStudents are learning Go!")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(buffer.String())
	}
}
