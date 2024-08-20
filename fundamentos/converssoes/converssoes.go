package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado... string(123) renderiza um caractere na tabela ASCII
	fmt.Println("Letra a : " + string(97))
	fmt.Println("Chaves : " + string(123))

	// int para string = strconv.Itoa(123) significa converter int para string
	fmt.Println("string para int : " + strconv.Itoa(123))

	// string para int = strconv.Atoi("123") significa converter string para int
	// o primeiro valor retornado é o valor convertido e o segundo o erro
	num, _ := strconv.Atoi("123")
	fmt.Println("int com string (calc) : " + strconv.Itoa(num-122))

	// string para bool = strconv.ParseBool("true") significa converter string para bool
	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}
