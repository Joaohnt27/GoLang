package main

import "fmt"

func main() {

	// strings
	var nome1 string = "Real"
	var nome2 = "Madrid"
	var nome3 string

	fmt.Println(nome1, nome2, nome3)

	nome1 = "Atletico"
	nome3 = "0 Champions League"

	fmt.Println(nome1, nome2, nome3)

	nome4 := "Nanicos..."
	fmt.Println(nome4)

	// ints
	var idade1 int = 20
	var idade2 = 30
	idade3 := 40

	fmt.Println(idade1, idade2, idade3)

	// bits e memória
	var num1 int8 = 20 // Significa que a variável num1 terá 8 bits de memória
	var num2 int8 = -128
	var num3 uint = 25 // Significa que o número incluído na variável não pode ser negativo

	// floats
	var nota1 float32 = 3.32
	var nota2 float64 = 5.75
	nota 3 := 8.5
}
