package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contadoKayan := ContaCorrente{titular: "Kayan",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	fmt.Println(contadoKayan)

	contaDaSilvia := ContaCorrente{"Silvia", 589, 1234987, 125.5}

	fmt.Println(contaDaSilvia)
}
