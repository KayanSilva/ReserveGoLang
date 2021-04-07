package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	exibiIntroducao()
	exibeMenu()

	// Estrutura de if
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do...")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	comando := leComando()

	// Estrutura com switch
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func devolveNomeEIdade() (string, int, bool) {
	nome := "Kayan"
	idade := 28
	varPraSerIgnorada := true
	return nome, idade, varPraSerIgnorada
}

func exibiIntroducao() {
	nome, idade, _ := devolveNomeEIdade()
	var versao float32 = 1.1

	fmt.Println("Olá, Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Endereço da variavel comando na memoria é", &comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resp.StatusCode)
	}
}
