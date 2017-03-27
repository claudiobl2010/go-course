package main

import "fmt"

// exemplo de função privada
func Add(x, y int) (int, int) {
	return x + y, 0
}

func main() {
    result, other := Add(10, 20)
	fmt.Println("O resultado de 10 + 20 é", result)
	fmt.Println("O parametro estra e", other)
}
