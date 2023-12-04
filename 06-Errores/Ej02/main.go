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

type ErrNotTaxableSalary struct{}

func (err ErrNotTaxableSalary) Error() string {
	return "salary is less than 10000"
}

func ShouldPayTaxes(salary int) (err error) {
	if salary <= 10000 {
		err = ErrNotTaxableSalary{}
		return
	}
	return
}

func main() {
	const salary = 8000

	if err := ShouldPayTaxes(salary); err != nil {
		if errors.Is(err, ErrNotTaxableSalary{}) {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("an unexpected error ocurred")
		return
	}

	fmt.Println("Must pay tax")
}
