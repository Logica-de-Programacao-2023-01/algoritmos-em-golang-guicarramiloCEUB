package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("Primeiro numero: ")
	fmt.Scan(&x)
	fmt.Println("Primeiro segundo: ")
	fmt.Scan(&y)
	fmt.Println("Primeiro terceiro: ")
	fmt.Scan(&z)
	soma := x + y + z
	fmt.Println("Soma dos tres numeros: ", soma)
}
