package main

import (
	"fmt"
)

func main() {

	precioEntradas := 32.75

	entradas := 0.0
	dia := ""
	metodo := ""
	película := ""

	costoEntradas := 0.0

	fmt.Println("**********Bienvenido*********\n\n¿Que película desea ver?")
	fmt.Print("-")
	fmt.Scanln(&película)

	fmt.Println("¿cuantas entradas desea?")
	fmt.Print("-")
	fmt.Scanln(&entradas)

	fmt.Println("¿Para que día desea la función?")
	fmt.Print("-")
	fmt.Scanln(&dia)

	fmt.Println("Qué método desea para pagar?")
	fmt.Print("-")
	fmt.Scanln(&metodo)
	fmt.Println("")

	costoEntradas = entradas * precioEntradas

	fmt.Println("**********FACTURA**********")
	fmt.Println("Película:", película)
	fmt.Println("Entradas:", entradas)
	fmt.Println("Día elegido:", dia)
	fmt.Println("Método de pago:", metodo)
	fmt.Printf("Costo total: %.2f\n", costoEntradas)
	fmt.Println("**********FACTURA**********")
	fmt.Println("")
	fmt.Println("Presiona Enter para salir...")
	fmt.Scanln()
}
