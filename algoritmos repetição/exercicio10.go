package main

import "fmt"

func main() {
	maior := 0
	var num int
	for i := 0; i > 0; i++ {
		fmt.Println("insira numero: ")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > maior {
			maior = num
		}
	}
	fmt.Println(maior)
}
