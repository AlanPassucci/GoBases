/*
Ejercicio 5 -  Impuestos de salario #5
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
El mismo tendrá que indicar “Error: the worker cannot have worked less than 80 hours per month”.
*/

package main

import (
	"errors"
	"fmt"
)

func calculateSalary(hoursOfWork int, valuePerHour float64) (salary float64, err error) {
	if hoursOfWork < 80 {
		err = errors.New("the worker cannot have worked less than 80 hours per month")
		return
	}

	salary = float64(hoursOfWork) * valuePerHour

	if salary > 149000 {
		salary -= salary * .1
	}

	return
}

func main() {
	salary, err := calculateSalary(240, 1000)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Salary: $", salary)
}
