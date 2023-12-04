/*
Ejercicio 3 - Impuestos de salario #3
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”.
*/

package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotTaxableSalary = "salary is less than 10000"
)

func ShouldPayTaxes(salary int) (err error) {
	if salary <= 10000 {
		err = errors.New(ErrNotTaxableSalary)
		return
	}
	return
}

func main() {
	const salary = 8000

	if err := ShouldPayTaxes(salary); err != nil {
		switch err.Error() {
		case ErrNotTaxableSalary:
			fmt.Println(err.Error())
			return
		default:
			fmt.Println("an unexpected error ocurred")
			return
		}
	}

	fmt.Println("Must pay tax")
}
