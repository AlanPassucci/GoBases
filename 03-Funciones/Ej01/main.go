/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y
si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/

package main

import (
	"fmt"
	"gobases/03-Funciones/Ej01/calculate_tax"
)

func main() {
	//calculateAgain := "s"
	var salary float64 = 750000
	var tax float64 = calculate_tax.CalculateTax(salary)
	var netSalary float64 = salary - tax

	fmt.Printf("\nSalario bruto: $%.2f", salary)
	fmt.Printf("\nImpuestos: $%.2f", tax)
	fmt.Printf("\nSalario neto: $%.2f\n", netSalary)
	/*
		for calculateAgain == "s" {
			fmt.Print("Ingrese su sueldo: ")
			fmt.Scanln(&salary)
			for salary < 1 {
				fmt.Print("Error. Reingrese su sueldo: ")
				fmt.Scanln(&salary)
			}

			fmt.Printf("\nSalario bruto: $%.2f", salary)
			tax = calculateTax(salary)
			fmt.Printf("\nImpuestos: $%.2f", tax)
			netSalary = salary - tax
			fmt.Printf("\nSalario neto: $%.2f", netSalary)

			fmt.Print("\n\nCalcular otro? ('s' para si) ")
			fmt.Scanln(&calculateAgain)
		}*/
}
