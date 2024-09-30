package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	lectura := bufio.NewScanner(os.Stdin)

	vocales := 0

	rand.Seed(time.Now().UnixNano())

	fmt.Println("**********CREADOR DE AVATARES**********")
	fmt.Println("Ingresa una lista de nombres que te gusten:")
	lectura.Scan()
	nombres := strings.Split(lectura.Text(), " ")
	clearConsole()
	fmt.Println("**********CREADOR DE AVATARES**********")

	fmt.Println("Ingresa una lista de apellidos que te gusten:")
	lectura.Scan()
	apellidos := strings.Split(lectura.Text(), " ")
	clearConsole()
	fmt.Println("**********CREADOR DE AVATARES**********")

	avataresPosibles := len(nombres) * len(apellidos)
	nombreRandom := rand.Intn(len(nombres))
	apellidoRandom := rand.Intn(len(apellidos))

	nombreCompleto := []string{nombres[nombreRandom], apellidos[apellidoRandom]}
	nombreCompletoT := strings.Join(nombreCompleto, " ")

	fmt.Printf("Tu avatar podría tener este nombre: %q \n", strings.ToUpper(nombreCompletoT))

	fmt.Println("¿Te gustaría saber un dato curioso?")
	fmt.Println("1. ¿Cuantas vocales tiene mi avatar? 2. ¿Cuantos posibles avatares hay?")
	lectura.Scan()

	clearConsole()
	fmt.Println("**********CREADOR DE AVATARES**********")

	if lectura.Text() == "1" {
		vocales += strings.Count(nombreCompletoT, "a")
		vocales += strings.Count(nombreCompletoT, "e")
		vocales += strings.Count(nombreCompletoT, "i")
		vocales += strings.Count(nombreCompletoT, "o")
		vocales += strings.Count(nombreCompletoT, "u")
		fmt.Println("El nombre de tu avatar tiene: ", vocales, " vocales")
	} else if lectura.Text() == "2" {
		fmt.Println("Existen ", avataresPosibles, " Avatares posibles para tus nombres y apellidos favoritos")
	} else {
		fmt.Println("Bueno :'(")
	}

	fmt.Println("Presiona Enter para salir...")
	fmt.Scanln()
}
