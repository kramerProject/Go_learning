package main

import (
	"fmt"
)

func main() {
	states := make([]string, 26, 26)
	states = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia",
	"Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
	"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
	"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
	"Santa Catarina", "São Paulo", "Sergipe", "Tocantins",
	}

	for state := 0; state < len(states); state++ {
		fmt.Printf("state: %v\n", states[state])
	}
	fmt.Printf("cap: %v and len: %v", cap(states), len(states))
}