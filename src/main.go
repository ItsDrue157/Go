package main

import "fmt"

func main() {
	var nome = "Carlos"

	fmt.Println("Seja Bem-Vindo sr.", nome)

	fmt.Println("1- Iniciar monitoriamento")
	fmt.Println("2- exibir os logs")
	fmt.Println("0- sair do programa")

	var comando int

	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi:", comando)
	
}
