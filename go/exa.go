package main

func main() {

	/*
		for i := 1; i <= 5; i++ {
			fmt.Println("El valor de i es", i)
		}

		x := 1
		for x <= 8 {

			fmt.Println("El valor de x es", x)
			x += 1
		}

		y := 3
		for {
			fmt.Println("El valor de y es", y)
			y++
			if y == 7 {
				break
			}
		}
	*/
	/*
		array2 := [5]int{10, 12, 12, 31, 5}
		for i := 0; i < len(array2); i++ {
			fmt.Println("el arreglo tiene tiene un", array2[i])
		}
	*/
	/*
		array2 := 0
		slice1 := []int{}

		for i := 0; i < 5; i++ {
			fmt.Println("ingrese un numero")
			fmt.Scan(&array2)
			slice1 = append(slice1, array2, array2*2)

		}

		fmt.Println("este es tu numero", slice1[0:])
	*/
	/*
		numeros := []int{12, 22, 54, 14, 5, 12, 7, 50, 75}
		n2 := numeros[:]

		copiados := copy(n2[:], numeros[:])
		sort.Ints(n2)
		fmt.Println("copiados vale: ", copiados, n2)
	*/

}
