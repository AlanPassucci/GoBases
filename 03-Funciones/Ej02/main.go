/*
Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/

package main

import (
	"fmt"
	"gobases/03-Funciones/Ej02/calculate_avg"
)

func main() {
	ratings := []uint{8, 7, 10}
	fmt.Printf("Promedio de notas: %.2f\n", calculate_avg.CalculateAvg(ratings...))
}
