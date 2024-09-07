package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) bool {
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

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		fmt.Println("Deposito realizado: ", valorDoDeposito)
		return true
	} else {
		fmt.Println("Deposito nao realizado")
		return false
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
