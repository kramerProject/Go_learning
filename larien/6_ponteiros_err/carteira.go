package main

import "fmt"

var ErroSaldoInsuficiente = fmt.Errorf("não é possível retirar, saldo insuficiente")

type Stringer interface {
	String() string
}
type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {

	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= quantidade
	return nil
}
