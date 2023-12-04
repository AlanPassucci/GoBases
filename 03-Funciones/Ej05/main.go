/*
Ejercicio 5 - Calcular cantidad de alimento
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne
una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

package main

import (
	"fmt"
	"gobases/03-Funciones/Ej05/animal"
)

func main() {
	animalTarantula, msg := animal.Animal(animal.Tarantula)
	if msg != nil {
		panic(msg)
	}
	animalHamster, msg := animal.Animal(animal.Hamster)
	if msg != nil {
		panic(msg)
	}
	animalDog, msg := animal.Animal(animal.Dog)
	if msg != nil {
		panic(msg)
	}
	animalCat, msg := animal.Animal(animal.Cat)
	if msg != nil {
		panic(msg)
	}

	var amount float64
	amount += animalTarantula(10)
	amount += animalHamster(10)
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Printf("Comida necesaria: %.2f kg\n", amount)
}
