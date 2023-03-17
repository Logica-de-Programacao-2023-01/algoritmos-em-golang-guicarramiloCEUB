package main

import "fmt"

func main() {
	var valor float64
	fmt.Println("Insira valor do produto: ")
	fmt.Scan(&valor)
	desconto := valor * 0.90
	fmt.Println("Valor com desconto de 10%: ", desconto)
}
