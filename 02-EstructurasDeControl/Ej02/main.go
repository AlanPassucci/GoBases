/*
Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.
Es necesario realizar una aplicación que reciba estas variables y que imprima un mensaje de acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import "fmt"

func main() {
	var (
		age               int     = 25
		isEmployed        bool    = true
		monthsOfAntiquity int     = 13
		salary            float32 = 85000
		accessible        bool
	)

	if age < 23 {
		fmt.Println("No tiene la edad necesaria para acceder a un préstamo.")
	} else {
		fmt.Println("Tiene la edad necesaria.")
		if !isEmployed {
			fmt.Println("Usted debe ser empleado para acceder a un préstamo.")
		} else {
			fmt.Println("Cumple con el requisito de ser empleado.")
			if monthsOfAntiquity < 12 {
				fmt.Println("La antiguedad debe ser mayor a 1 año.")
			} else {
				fmt.Println("Su antiguedad es válida para acceder a un préstamo.")
				accessible = true
			}
		}
	}

	if accessible {
		fmt.Println("Cumple con todos los requisitos para acceder a un préstamo.")
		if salary < 100001 {
			fmt.Println("A su vez, deberá pagar intereses.")
		}
	} else {
		fmt.Println("No se le puede otorgar un préstamo debido a que no cumple con todos los requisitos.")
	}
}
