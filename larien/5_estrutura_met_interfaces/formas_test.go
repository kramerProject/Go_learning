package main

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retangulo", forma: Retangulo{12, 6}, esperado: 72.0},
		{nome: "Circulo", forma: Circulo{10}, esperado: 314.1592653589793},
		{nome: "Triangulo", forma: Triangulo{12, 6}, esperado: 36},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("%#v resultado %.2f esperado %.2f", tt.forma, resultado, tt.esperado)
		}
	}
}
