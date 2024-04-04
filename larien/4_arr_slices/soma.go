package main

func Soma(numeros []int) int {
	var soma int
	for _, n := range numeros {
		soma += n
	}
	return soma
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	quantidadeDeNumeros := len(numerosParaSomar)
	somas = make([]int, quantidadeDeNumeros)

	for i, numeros := range numerosParaSomar {
		somas[i] = Soma(numeros)
	}

	return
}
