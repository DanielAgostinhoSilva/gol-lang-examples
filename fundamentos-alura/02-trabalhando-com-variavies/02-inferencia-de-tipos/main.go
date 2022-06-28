package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Daniel"
	idade := 30
	versao := 1.1

	fmt.Println("Ola, sr.", nome, "sua idade e", idade)
	fmt.Println("Este programa esta na versao", versao)

	fmt.Println("O tipo da variavel nome e", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade e", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao e", reflect.TypeOf(versao))
}
