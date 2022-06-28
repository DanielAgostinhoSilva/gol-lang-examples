package main

import "fmt"

func main() {
	nome := "Daniel"
	versao := 1.1

	fmt.Println("Ola, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Porgrama")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Nao conheco este comando")
	}
}
