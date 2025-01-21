package chapter_two

import "fmt"

// TypeString is a function that demonstrates the use of the string type in Go.
func TypeString() {
	str := "Hello, World! ↓ こんにちは、世界！"

	// Usando range para iterar sobre a string ele imprime o valor do caractere
	for _, v := range str {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	// Usando for estilo c para iterar sobre a string ele imprime o valor do byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", str[i], str[i], str[i], str[i])
	}
}
