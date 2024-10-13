package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Busqueda secuencial con for each")
	fmt.Println("Ingrese el nombre de una fruta a buscar")
	lectura := bufio.NewScanner(os.Stdin)
	lectura.Scan()

	fruits := []string{"manzana", "pera", "cambur", "granada", "uva", "mango", "guanabana", "kiwi", "mora", "tomate", "fresa", "parchita", "melon", "patilla", "naranja", "mandarina"}
	on := 0
	for _, i := range fruits {
		if lectura.Text() == i {
			on = 1
		}
	}

	if on == 1 {
		fmt.Println("Su fruta ha sido encontrada!")
	} else {
		fmt.Println("Su fruta no fué encontrada..")
	}

	fmt.Println("Presiona Enter para salir...")
	fmt.Scanln()
}
