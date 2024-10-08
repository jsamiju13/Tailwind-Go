package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var nombre string

	//edad := 0

	lectura := bufio.NewScanner(os.Stdin)

	fmt.Println("ingrese sus nombres")

	lectura.Scan()
	nombre = lectura.Text()

	picado := strings.Split(nombre, " ")

	fmt.Println(" la palabra picada es:", picado)

	comparador := strings.Compare(picado[0], picado[1])

	if comparador == 1 {
		fmt.Printf("el texto %1s es mayor al %2s", picado[0], picado[1])
	} else if comparador == -1 {
		fmt.Printf("el texto %2s es mayor al %1s", picado[0], picado[1])
	} else {
		fmt.Println("ambos textos son iguales")
	}

	/*
		fmt.Println("ingrese su cedula")

		lectura.Scan()
		cedula = lectura.Text()

		fmt.Println("ingrese su direccion")

		lectura.Scan()
		direccion = lectura.Text()

		//fmt.Print("su nombre es:", strings.ToUpper(nombre), " su cedula es :", cedula, "y su direccion es: ", direccion)
		fmt.Printf("tu primer nombre es: %1s  su cedula es: %2s y su direccion es: %3s y %q", picado[0], cedula, direccion, picado[1])
	*/
}
