package main

import "fmt"

func main() {
	nome := "João Henrique"
	idade := 19

	fmt.Println("Meu nome é", nome, "e tenho", idade, " anos de idade")

	// Printf (strings formatadas)
	fmt.Printf("Meu nome é %v e tenho %v anos de idade \n", nome, idade)
	fmt.Printf("Meu nome é %q e temho %q anos de idade \n", nome, idade)
	fmt.Printf("Idade é um tipo de: %T \n", idade)
	fmt.Printf("Você tem um score de %f pontos \n", 764.98)
	fmt.Printf("Você tem um score de %0.1f pontos \n", 764.98)

	// Sprintf (salvar strings formatadas)
	var teste string = fmt.Sprintf("Meu nome é %v e tenho %v anos de idade \n", nome, idade)

	fmt.Println("A string formatada é:", teste)
}
