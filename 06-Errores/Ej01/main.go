/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: the salary entered does not reach the taxable minimum" y lanzalo en caso de que “salary” sea menor a 150.000.
De lo contrario, tendrás que imprimir por consola el mensaje “Must pay tax”.
*/

package main

import "fmt"

type TaxError struct{}

func (me *TaxError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

func main() {
	salary := 150000

	if salary < 150000 {
		taxError := TaxError{}
		panic(taxError.Error())
	}

	fmt.Println("Must pay tax")
}
