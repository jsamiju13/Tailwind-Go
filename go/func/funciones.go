package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clearConsole() {
	var clearCmd *exec.Cmd
	clearCmd = exec.Command("cmd", "/c", "cls")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func trapecio(a, b, c float32) {

	area := ((b + a) * c) / 2

	fmt.Println("el area de su trapecio es: ", area)
}

func trapecio2(a, b, c float32) (r1 float32, r2 bool) {

	r1 = ((b + a) * c) / 2
	return r1, true

}

func main() {

	var a float32
	var b float32
	var c float32
	fmt.Println("ingrese la base menor, la base mayor y la altura")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	clearConsole()

	res, err := trapecio2(a, b, c)

	if err == false {
		fmt.Println("error")
	} else {
		fmt.Println("el area de su trapecio es: ", res)
	}

}
