/*
Ejercicio 3 - Registrando clientes
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
Los datos requeridos son:
File
Name
ID
Phone number
Home

Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array.
En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: client already exists”, y continuar con la ejecución del programa normalmente.
Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a
registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que
ser del tipo error para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato,
ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes:
“End of execution” y “Several errors were detected at runtime”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente
para el caso de error retornado).
*/

package main

import (
	"errors"
	"fmt"
)

type Client struct {
	File        string
	Name        string
	ID          int
	PhoneNumber int
	Home        string
}

func (c Client) AddClient(clientsList []Client) ([]Client, error) {
	for _, client := range clientsList {
		if client.ID == c.ID {
			return nil, errors.New("client already exists")
		}
	}
	return append(clientsList, c), nil
}

func (c Client) IsValidClient() (bool, error) {
	if c.File == "" || c.Name == "" || c.ID == 0 || c.PhoneNumber == 0 || c.Home == "" {
		return false, errors.New("all client data must have non-zero values")
	}
	return true, nil
}

func main() {
	clients := []Client{
		{File: "./1000.file.txt", Name: "Juan Perez", ID: 1000, PhoneNumber: 1123456789, Home: "Calle 1"},
		{File: "./1001.file.txt", Name: "Maria Lopez", ID: 1001, PhoneNumber: 1122334455, Home: "Calle 2"},
		{File: "./1002.file.txt", Name: "Carlos Sanchez", ID: 1002, PhoneNumber: 1198765432, Home: "Calle 3"},
	}
	newClient := Client{
		File: "./1003.file.txt", Name: "Natalie Gonzalez", ID: 1001, PhoneNumber: 1199887766, Home: "Calle 4",
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Several errors were detected at runtime\n", err)
		}
		fmt.Println("End of execution")
	}()

	isValidClient, err := newClient.IsValidClient()
	if err != nil {
		panic(err.Error())
	}

	if isValidClient {
		newClientsList, err := newClient.AddClient(clients)
		if err != nil {
			panic(err.Error())
		}

		clients = newClientsList
		fmt.Println(clients)
	}
}
