package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Println("Insira primeiro numero: ")
	fmt.Scan(&n1)
	fmt.Println("Insira segundo numero: ")
	fmt.Scan(&n2)
	fmt.Println("Insira terceiro numero: ")
	fmt.Scan(&n3)
	if n1 == n2 && n1 == n3 && n2 == n3 {
		fmt.Println("São iguais")
	} else if n1 <= n2 && n1 <= n3 {
		fmt.Printf("%d é menor", n1)
	} else if n2 <= n1 && n2 <= n3 {
		fmt.Printf("%d é menor", n2)
	} else {
		fmt.Printf("%d é menor", n3)
	}
}
