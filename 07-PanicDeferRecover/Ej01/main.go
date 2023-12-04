/*
Ejercicio 1 - Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones.
Para ello, cuentan con todo el detalle necesario en un archivo .txt.
Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo,
no han pasado el archivo a leer por nuestro programa.
Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt” (recordá lo visto sobre el pkg “os”).
Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al
intentar leer un archivo que no existe, mostrando el mensaje “The indicated file was not found or is damaged”.
Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/

package main

import (
	"fmt"
	"os"
)

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {
	data, err := readFile("./customers.txt")

	defer fmt.Println("Ejecución finalizada!")

	if err != nil {
		panic("the indicated file was not found or is damaged")
	}

	fmt.Println("Contenido del archivo:\n", data)
}
