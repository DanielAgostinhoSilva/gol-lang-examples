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

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa...")
	} else {
		fmt.Println("Nao conheco este comando")
	}
}
