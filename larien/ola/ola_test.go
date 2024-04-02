package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Kramer")
	esperado := "Olá, Kramer."

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
