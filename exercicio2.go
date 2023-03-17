package main

import "fmt"

func main() {
	var num int
	fmt.Println("Insira numero:")
	fmt.Scan(&num)
	dobro := num * 2
	triplo := num * 3
	quadruplo := num * 4
	fmt.Println("Dobro: ", dobro)
	fmt.Println("Triplo: ", triplo)
	fmt.Println("Quadruplo: ", quadruplo)
}
