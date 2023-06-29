package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Println("primeiro numero: ")
	fmt.Scan(&n1)
	fmt.Println("segundo numero: ")
	fmt.Scan(&n2)
	if n1 >= 0 && n2 >= 0 {
		fmt.Println(n1 * n2)
	} else if n1 < 0 || n2 < 0 {
		fmt.Println(n1 + n2)
	}
}
