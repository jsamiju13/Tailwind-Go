package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("**********SUPER CONVERSOR DE UNIDADES**********")

	lectura := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese Kilometros, Kilogramos y Horas:")
	lectura.Scan()
	datos := strings.Split(lectura.Text(), " ")

	kmStr := datos[0]
	kgStr := datos[1]
	hourStr := datos[2]

	km, _ := strconv.Atoi(kmStr)
	miles := float64(km) * 0.621371

	kg, _ := strconv.Atoi(kgStr)
	pounds := float64(kg) * 2.20462

	hour, _ := strconv.Atoi(hourStr)
	seconds := float64(hour) * 3600

	fmt.Println("**********SUPER CONVERSOR DE UNIDADES**********")
	fmt.Printf("%d kilómetros son %.2f millas.\n", km, miles)
	fmt.Printf("%d kilogramos son %.3f libras.\n", kg, pounds)
	fmt.Printf("%d horas son %.2f segundos.\n", hour, seconds)

	fmt.Println("Presiona Enter para salir...")
	fmt.Scanln()
}
