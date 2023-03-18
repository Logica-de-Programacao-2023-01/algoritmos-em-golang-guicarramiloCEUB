package main

import "fmt"

func main() {
	var idade int
	fmt.Println("insira idade: ")
	fmt.Scan(&idade)
	dias := idade * 365
	fmt.Println("Voce viveu ", dias, " dias!")
}
