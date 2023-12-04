/*
Ejercicio 4 - Impuestos de salario #4
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba
por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por
consola deberá decir: “Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]”,
siendo [salary] el valor de tipo int pasado por parámetro).
*/

package main

import "fmt"

func ShouldPayTaxes(salary int) (err error) {
	if salary < 150000 {
		err = fmt.Errorf("the minimum taxable amount is 150000 and the salary entered is: %d", salary)
		return
	}
	return
}

func main() {
	const salary = 125000

	if err := ShouldPayTaxes(salary); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Must pay tax")
}
