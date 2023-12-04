/*
Ejercicio 4 - Calcular estadísticas
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as
estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros
y devuelva el cálculo que se indicó en la función anterior.
*/

package main

import (
	"fmt"
	"gobases/03-Funciones/Ej04/operations"
)

func main() {
	minFunc, err := operations.Operation(operations.Minimum)
	if err != nil {
		panic(err)
	}
	averageFunc, err := operations.Operation(operations.Average)
	if err != nil {
		panic(err)
	}
	maxFunc, err := operations.Operation(operations.Maximum)
	if err != nil {
		panic(err)
	}

	minValue, err := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	if err != nil {
		panic(err)
	}
	averageValue, err := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if err != nil {
		panic(err)
	}
	maxValue, err := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if err != nil {
		panic(err)
	}

	fmt.Println("Minimo:", minValue)
	fmt.Println("Promedio:", averageValue)
	fmt.Println("Maximo:", maxValue)
}
