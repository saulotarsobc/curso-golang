package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tablea Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numero reais ... float32 float64
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola dev"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho de s1 é", len(s1))

	// string com multiplas linhas
	s2 := `Ola
	dev
	aqui`
	fmt.Println("O tamanho de s2 é", len(s2))

	// char??? (o tipoo de char em go é um int32)
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))

	// comparações
	fmt.Println("1 == 1", 1 == 1)
	fmt.Println("1 == 2", 1 == 2)
	fmt.Println("1 != 2", 1 != 2)
	fmt.Println("1 < 2", 1 < 2)
	fmt.Println("1 > 2", 1 > 2)
	fmt.Println("1 <= 2", 1 <= 2)
	fmt.Println("1 >= 2", 1 >= 2)

	// operadores logicos
	fmt.Println("v && f", true && false)
	fmt.Println("v || f", true || false)
	fmt.Println("!v", !true)

	// operadores unarios
	x++
	fmt.Println("x++", x)
	x--
	fmt.Println("x--", x)
}
