package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Kayan"
	idade := 28
	var versao float32 = 1.1

	fmt.Println("Olá, Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Endereço da variavel comando na memoria é", &comando)
}
