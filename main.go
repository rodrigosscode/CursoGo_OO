package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) bool
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {

	titularRodrigo := clientes.Titular{
		Nome:      "Rodrigo",
		CPF:       "0000000000",
		Profissao: "Dev Android",
	}

	titularAna := clientes.Titular{
		Nome:      "Ana",
		CPF:       "012121200",
		Profissao: "intp libras",
	}

	contaDoRodrigo := contas.ContaCorrente{
		Titular:       titularRodrigo,
		NumeroAgencia: 3,
		NumeroConta:   2,
	}

	contaDaAna := contas.ContaCorrente{
		Titular:       titularAna,
		NumeroAgencia: 3,
		NumeroConta:   2,
	}

	contaDoFulano := contas.ContaPoupanca{}

	contaDoRodrigo.Depositar(6000)

	PagarBoleto(&contaDoRodrigo, 5000)
	PagarBoleto(&contaDoFulano, 5000)

	fmt.Println(contaDoRodrigo)
	fmt.Println(contaDaAna)
	fmt.Println(contaDoFulano)

	createInstancesAndCompare()
}

func createInstancesAndCompare() {
	// contaDoRodrigo := contas.ContaCorrente{
	// 	Titular:       "Rodrigo",
	// 	NumeroAgencia: 3,
	// 	NumeroConta:   2,
	// 	Saldo:         3000,
	// }

	// contaDaAna := contas.ContaCorrente{
	// 	Titular: "Ana", NumeroAgencia: 3, NumeroConta: 2, Saldo: 5000,
	// }

	// contaDaFulano := contas.ContaCorrente{
	// 	Titular: "Fulano",
	// 	Saldo:   2000,
	// }

	// contaDaFulano2 := contas.ContaCorrente{
	// 	Titular: "Fulano",
	// 	Saldo:   2000,
	// }

	// var (
	// 	contaRui  *contas.ContaCorrente
	// 	contaRui2 *contas.ContaCorrente
	// )

	// contaRui = new(contas.ContaCorrente)
	// contaRui.Titular = "Rui"

	// contaRui2 = new(contas.ContaCorrente)
	// contaRui2.Titular = "Rui"

	// fmt.Println(contaRui)
	// fmt.Println(contaDoRodrigo)
	// fmt.Println(contaDaAna)
	// fmt.Println(contaDaFulano)

	// fmt.Println(contaDaFulano == contaDaFulano2)
	// fmt.Println(*contaRui == *contaRui2)
}
