package main

import "fmt"

func main() {
	people := map[string]float64{
		"wilson": 25.5,
		"andre":  26.5,
	}

	people["junior"] = 10.5

	for name, age := range people {
		fmt.Printf("Chave %s do mapa valor: %.2f\n", name, age)
	}

	delete(people, "wilson")
	fmt.Printf("Posição não alocada: %.2f\n", people["wilson"])
}
