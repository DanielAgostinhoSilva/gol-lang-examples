package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Daniel"
	versao := 1.1

	fmt.Println("Ola, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Porgrama")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comandoLido escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}

	for _, site := range sites {

		response, _ := http.Get(site)

		if response.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
		} else {
			fmt.Println("Site:", site, "esta com problemas. Status code:", response.StatusCode)
		}
	}

}
