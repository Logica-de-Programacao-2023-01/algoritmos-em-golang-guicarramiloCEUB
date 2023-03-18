package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("Insira peso: ")
	fmt.Scan(&peso)
	libras := peso * 2.205
	fmt.Println("Valor em libras: ", libras)
}
