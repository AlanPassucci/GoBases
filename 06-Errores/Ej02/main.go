/*
Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: salary is less than 10000"
y lanzalo en caso de que “salary” sea menor o igual a 10000. La validación debe ser hecha con la función “Is()” dentro del “main”.
*/

package main

import (
	"errors"
	"fmt"
)

const ErrorLowSalary = "Error: salary is less than 10000"

type TaxError struct {
	msg string
}

func (te TaxError) Error() string {
	return te.msg
}

func main() {
	salary := 8000

	if salary <= 10000 {
		taxError := TaxError{
			msg: ErrorLowSalary,
		}

		if errors.Is(taxError, TaxError{ErrorLowSalary}) {
			panic(taxError)
		}
	}

	fmt.Println("Salary is greater than 10000")
}
