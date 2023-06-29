package main

import "fmt"

func main() {
	for i := 0; i < 31; i++ {
		if i%3 == 0 {
			fmt.Printf("%d \n", i)
		}
	}
}
