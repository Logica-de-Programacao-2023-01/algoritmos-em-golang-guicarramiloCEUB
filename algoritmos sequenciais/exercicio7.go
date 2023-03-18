package main

import "fmt"

func main() {
	var num int
	fmt.Println("insira numero: ")
	fmt.Scan(&num)
	fmt.Println("Sucessor: ", (num + 1), "\nAntecessor: ", (num - 1))
}
