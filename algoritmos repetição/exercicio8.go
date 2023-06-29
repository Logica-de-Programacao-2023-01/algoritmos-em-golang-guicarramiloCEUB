package main

import "fmt"

func main() {
	var num int
	for i := num; i > 0; i-- {
		if num%i == 0 {
			fmt.Printf("%d \n", i)
		}
	}
}
