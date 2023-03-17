package main

import "fmt"

func main() {
	var x, y, z float64
	fmt.Println("Insira numero: ")
	fmt.Scan(&x)
	fmt.Println("Insira numero: ")
	fmt.Scan(&y)
	fmt.Println("Insira numero: ")
	fmt.Scan(&z)
	media := ((x * 2) + (y * 3) + (z * 5)) / 10
	fmt.Println("Media ponderada : ", media)
}
