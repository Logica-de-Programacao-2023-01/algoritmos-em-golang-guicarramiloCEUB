package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Println("Insira primeiro numero: ")
	fmt.Scan(&n1)
	fmt.Println("Insira segundo numero: ")
	fmt.Scan(&n2)
	if n1 > n2 {
		fmt.Printf("%d é maior", n1)
	} else if n2 > n1 {
		fmt.Printf("%d é maior", n2)
	} else {
		fmt.Printf("%d e %d são iguais", n1, n2)
	}
}
