package main

import "fmt"

func main() {
	a := 3
	b := 2

	fmt.Println("Soma: ", a+b)
	fmt.Println("Subtração: ", a-b)
	fmt.Println("Divisão: ", a/b)
	fmt.Println("Multiplicação: ", a*b)
	fmt.Println("Módulo: ", a%b)

	// bitwise
	fmt.Println("E: ", a&b)    // AND - 11 & 10 = 10
	fmt.Println("Ou: ", a|b)   // OR - 11 | 10 = 11
	fmt.Println("Xor: ", a^b)  // XOR - 11 ^ 10 = 01
	fmt.Println("Shl: ", a<<b) // SHL - 11 << 1 = 110
	fmt.Println("Shr: ", a>>b) // SHR - 11 >> 1 = 01

	// outas operações aritimeticas do go usando o math
	c := 3.0
	d := 2.0
	fmt.Println("Maior: ", c > d)
	fmt.Println("Menor: ", c < d)
	fmt.Println("Maior ou igual: ", c >= d)
	fmt.Println("Menor ou igual: ", c <= d)
	fmt.Println("Diferente: ", c != d)
	fmt.Println("Igual: ", c == d)
}
