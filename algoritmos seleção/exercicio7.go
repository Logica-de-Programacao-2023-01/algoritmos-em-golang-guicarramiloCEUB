package main

import "fmt"

func main() {
	var salario float64
	fmt.Println("insira salario: ")
	fmt.Scan(&salario)
	if salario < 1000 {
		salario *= 1.1
	} else {
		salario *= 1.05
	}
	fmt.Println(salario)
}
