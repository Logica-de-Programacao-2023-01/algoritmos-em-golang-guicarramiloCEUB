package main

import "fmt"

func main() {
	var sexo string
	var peso, altura float64
	fmt.Println("Informe sexo (m/f): ")
	fmt.Scan(&sexo)
	fmt.Println("Informe peso: ")
	fmt.Scan(&peso)
	fmt.Println("Informe altura: ")
	fmt.Scan(&altura)
	imc := peso / (altura * altura)
	fmt.Println(imc)
	if sexo == "m" {
		if imc >= 20.7 && imc <= 26.4 {
			fmt.Println("normal")
		} else if imc < 20.7 {
			fmt.Println("abaixo")
		} else {
			fmt.Println("acima")
		}
	} else if sexo == "f" {
		if imc >= 19.1 && imc <= 25.8 {
			fmt.Println("normal")
		} else if imc < 19.1 {
			fmt.Println("abaixo")
		} else {
			fmt.Println("acima")
		}
	}
}
