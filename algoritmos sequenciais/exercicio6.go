package main

import "fmt"

func main() {
	var salario float64
	fmt.Println("Insira salario: ")
	fmt.Scan(&salario)
	ajustado := salario * 1.15
	fmt.Println("Salario com ajuste de 15%: ", ajustado)
}
