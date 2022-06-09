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

	if comando == 1 {
		fmt.Println("Monitorando...")

	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")

	} else if comando == 0 {
		fmt.Println("Saindo do programa...")

	} else {
		fmt.Println("Nao conheco esse comando")
	}

}
