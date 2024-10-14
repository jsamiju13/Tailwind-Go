package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Promedio de una lista")

	lectura := bufio.NewScanner(os.Stdin)
	words := []string{}
	numbers := 0
	fmt.Println("Ingrese elementos numericos para sacar el promedio")
	lectura.Scan()
	words = strings.Split(lectura.Text(), " ")

	for _, i := range words {
		num, _ := strconv.Atoi(i)
		numbers += num
	}

	numbers /= len(words)

	fmt.Println("El promedio de sus numeros es: ", numbers)
	fmt.Println("Presiona Enter para salir...")
	fmt.Scanln()
}
