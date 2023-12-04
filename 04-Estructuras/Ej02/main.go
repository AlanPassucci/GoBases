/*
Ejercicio 2 - Employee
Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a
gestionar correctamente dichos empleados. Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composición con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar
el método PrintEmployee().
Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
*/

package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Legajo: %v\nPosicion: %v\nDNI: %v\nNombre: %v\nFecha nacimiento: %v\n\n",
		e.ID, e.Position, e.Person.ID, e.Person.Name, e.Person.DateOfBirth)
}

func main() {
	person := Person{
		ID:          12345678,
		Name:        "Juan Perez",
		DateOfBirth: "27/10/1995",
	}

	employee := Employee{
		ID:       1000,
		Position: "Software Developer",
		Person:   person,
	}

	employee.PrintEmployee()
}
