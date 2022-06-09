package main

import "fmt"
import "os"

func main() {
	exibeintroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Nao conheco esse comando")
		os.Exit(-1)

	}
}

func exibeintroducao() {
	var nome = "Carlos"
	var versao = 1.18
	fmt.Println("Seja Bem-Vindo sr.", nome)
	fmt.Println("Esse programa esta na ", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	return comandoLido
}

func exibeMenu() {

	fmt.Println("1- Iniciar monitoriamento")
	fmt.Println("2- exibir os logs")
	fmt.Println("0- sair do programa")

}
