package chapter_one

import "fmt"

// ValuesDefaultTypes demonstrates the default zero values for various Go types.
// It prints the default value and type for int, string, bool, float64, pointer,
// slice, map, channel, and function types. This function helps in understanding
// what values these types hold when they are declared but not explicitly initialized.
func ValuesDefaultTypes() {
	var typeInt int
	var typeString string
	var typeBool bool
	var typeFloat64 float64
	var typePointerInt *int
	var typeArrayInt []int
	var typeMapStringInt map[string]int
	var typeChanInt chan int
	var typeFunc func(string) int

	fmt.Printf("Value: %v, Type: %T\n", typeInt, typeInt)
	fmt.Printf("Value: %v, Type: %T\n", typeString, typeString)
	fmt.Printf("Value: %v, Type: %T\n", typeBool, typeBool)
	fmt.Printf("Value: %v, Type: %T\n", typeFloat64, typeFloat64)
	fmt.Printf("Value: %v, Type: %T\n", typePointerInt, typePointerInt)
	fmt.Printf("Value: %v, Type: %T\n", typeArrayInt, typeArrayInt)
	fmt.Printf("Value: %v, Type: %T\n", typeMapStringInt, typeMapStringInt)
	fmt.Printf("Value: %v, Type: %T\n", typeChanInt, typeChanInt)
	fmt.Printf("Value: %p, Type: %T\n", typeFunc, typeFunc)
}
