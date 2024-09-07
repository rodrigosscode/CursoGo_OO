package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		fmt.Println("Saque realizado", valorDoSaque)
		return true
	} else {
		fmt.Println("saldo insuficiente")
		return false
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		fmt.Println("Deposito realizado: ", valorDoDeposito)
		return true
	} else {
		fmt.Println("Deposito nao realizado")
		return false
	}
}

func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contaDestino *ContaCorrente) {
	conseguiuSacar := c.Sacar(valorDeTransferencia)
	if conseguiuSacar {
		contaDestino.Depositar(valorDeTransferencia)
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
