/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle
de los datos de cada uno de ellos/as, de la siguiente manera:

Name: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y
que tenga un método detalle.
*/

package main

import "fmt"

type Student struct {
	Name            string
	LastName        string
	DNI             int
	DateOfAdmission string
}

func (s Student) Details() string {
	return fmt.Sprintf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", s.Name, s.LastName, s.DNI, s.DateOfAdmission)
}

type Detailable interface {
	Details() string
}

func main() {
	students := []Student{
		{
			Name: "Juan", LastName: "Perez", DNI: 12345678, DateOfAdmission: "01/03/2023",
		},
		{
			Name: "Maria", LastName: "Lopez", DNI: 87654321, DateOfAdmission: "01/03/2023",
		},
	}

	for _, student := range students {
		fmt.Println(student.Details())
	}
}
