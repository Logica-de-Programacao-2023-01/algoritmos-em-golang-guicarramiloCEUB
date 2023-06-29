package main

import "fmt"

func main() {
	soma := 0
	count := 0
	var num int
	for i := 0; i > 0; i++ {
		fmt.Println("insira numero: ")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		soma += num
		count += 1
	}
	fmt.Println(soma / count)
}
