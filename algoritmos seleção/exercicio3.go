package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite o numero: ")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Printf("%d é par", num)
	} else {
		fmt.Printf("%d é impar", num)
	}
}
