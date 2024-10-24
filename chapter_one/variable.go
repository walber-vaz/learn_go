package chapter_one

import "fmt"

// Package level scope
var y = 20

// Variable is a function that demonstrates variable scope
func Variable() {
	x := 10
	show(x)
}

// Function written in lower case is only available
// within the package it is defined
func show(param int) {
	// x is not accessible here
	fmt.Println(y)
	fmt.Println(param)
}
