package main

import "fmt"

func main() {
	var dias float64
	var diaria float64
	fmt.Println("Insira dias trabalhados: ")
	fmt.Scan(&dias)
	fmt.Println("Insira valor da diaria: ")
	fmt.Scan(&diaria)
	var salario = diaria * dias
	fmt.Println("Salario: ", salario)
}
