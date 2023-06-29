package main

import "fmt"

func main() {
	var n1, n2, n3 int
	var numeros []int
	var order []int
	fmt.Println("insira numero: ")
	fmt.Scan(&n1)
	numeros = append(numeros, n1)
	fmt.Println("insira numero: ")
	fmt.Scan(&n2)
	numeros = append(numeros, n1)
	fmt.Println("insira numero: ")
	fmt.Scan(&n3)
	numeros = append(numeros, n1)
	menor := 0
	for i := 0; i < len(numeros); i++ {
		if numeros[i] < menor {
			menor = numeros[i]
		}
	}
}
