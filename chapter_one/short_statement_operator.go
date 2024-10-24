package chapter_one

import "fmt"

func ShortStatementOperator() {
	// Short statement operator
	// This operator is used to declare and initialize variables
	// It is used to declare and initialize variables in a single line
	// Syntax
	// variable_name := value
	// infer the type of the variable from the value
	x := 10
	y := 20
	fmt.Printf("the: x = %d, y = %d, type of x = %T, type of y = %T\n", x, y, x, y)

	// Keywords in Go, not allowed to use as variable names
	// ref: https://go.dev/ref/spec#Keywords
	// break        default      func         interface    select
	// case         defer        go           map          struct
	// chan         else         goto         package      switch
	// const        fallthrough  if           range        type
	// continue     for          import       return       var
}
