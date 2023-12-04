/*
Ejercicio 3 - A qué mes corresponde
Realizar una aplicación que reciba  una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.
*/

package main

import "fmt"

func main() {
	var month int

	fmt.Printf("Ingrese número de mes (1-12): ")
	fmt.Scanln(&month)
	for month < 1 || month > 12 {
		fmt.Printf("Error. Reingrese número de mes (1-12): ")
		fmt.Scanln(&month)
	}

	switch month {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	}
}
