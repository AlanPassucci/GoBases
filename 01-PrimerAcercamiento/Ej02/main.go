/*
Ejercicio 2 - Clima
Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.
Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
*/

package main

import "fmt"

func main() {
	var temperature float32 = 21.7
	var humidity int = 75
	var pressure float32 = 1008.5

	fmt.Printf("Temperatura: %v °C\nHumedad: %v %%\nPresión: %v hPa\n", temperature, humidity, pressure)
}
