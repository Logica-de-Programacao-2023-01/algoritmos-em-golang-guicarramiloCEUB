package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Println("Insira peso: ")
	fmt.Scan(&peso)
	fmt.Println("Insira altura: ")
	fmt.Scan(&altura)
	imc := peso / (altura * altura)
	fmt.Println("IMC: ", imc)
}
