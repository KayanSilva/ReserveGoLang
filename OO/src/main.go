package main

import (
	"fmt"

	"github.com/KayanSilva/ReserveGo/OO/clientes"
	"github.com/KayanSilva/ReserveGo/OO/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contadoKayan := contas.ContaCorrente{Titular: clientes.Titular{
		Nome: "Kayan", CPF: "111.222.333.12", Profissao: "Autônomo"},
		NumeroAgencia: 589, NumeroConta: 123456, saldo: 12}

	fmt.Println(contadoKayan)

	clientSilvia := clientes.Titular{Nome: "Silvia", CPF: "111.222.333.22", Profissao: "Arquiteta"}
	contaDaSilvia := contas.ContaCorrente{clientSilvia, 589, 1234987, 500}

	fmt.Println(contaDaSilvia)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = clientes.Titular{Nome: "Cris", CPF: "111.222.333.33", Profissao: "Agricultora"}

	fmt.Println(*contaDaCris)

	contadoKayan2 := contas.ContaCorrente{Titular: "Kayan",
		NumeroAgencia: 589, NumeroConta: 123456, saldo: 125.5}

	//É igual pelo fato de está apontando para o mesmo loca na memoria
	fmt.Println(contadoKayan2 == contadoKayan)

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular = "Cris"

	//É diferente pelo fato de apontar para locais da memoria diferente
	fmt.Println(*contaDaCris == *contaDaCris2)

	fmt.Println(contaDaSilvia.Sacar(200))

	//Testando deposito
	fmt.Println(contaDaSilvia.Depositar(500))

	//Transferencia entre contas
	contaDaSilvia.Transferir(100, &contadoKayan)
	fmt.Println(contadoKayan.getSaldo())

	PagarBoleto(&contadoKayan, 100)
	fmt.Println(contadoKayan.getSaldo())
}
