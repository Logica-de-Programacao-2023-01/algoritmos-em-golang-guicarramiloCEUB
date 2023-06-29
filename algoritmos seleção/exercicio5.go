package main

import "fmt"

func main() {
	var num int
	fmt.Println("digite numero: ")
	fmt.Scan(&num)
	if num%5 == 0 && num%3 == 0 {
		fmt.Println("multiplo de 3 e 5")
	}
}
