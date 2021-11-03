package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insufiente"
	}
}

func main() {

	contadoKayan := ContaCorrente{titular: "Kayan",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	fmt.Println(contadoKayan)

	contaDaSilvia := ContaCorrente{"Silvia", 589, 1234987, 500}

	fmt.Println(contaDaSilvia)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	fmt.Println(*contaDaCris)

	contadoKayan2 := ContaCorrente{titular: "Kayan",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	//É igual pelo fato de está apontando para o mesmo loca na memoria
	fmt.Println(contadoKayan2 == contadoKayan)

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"

	//É diferente pelo fato de apontar para locais da memoria diferente
	fmt.Println(*contaDaCris == *contaDaCris2)

	fmt.Println(contaDaSilvia.Sacar(200))
}
