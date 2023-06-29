package main

import "fmt"

func main() {
	var idade int
	fmt.Println("insira idade: ")
	fmt.Scanln(&idade)
	if idade <= 9 {
		fmt.Printf("mirim")
	} else if idade > 9 && idade < 14 {
		fmt.Println("infantil")
	} else if idade > 13 && idade < 18 {
		fmt.Println("juvenil")
	} else {
		fmt.Println("adulto")
	}
}
