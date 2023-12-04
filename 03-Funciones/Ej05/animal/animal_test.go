/*
Ejercicio 5 - Calcular cantidad de alimento
El refugio de animales envió una queja ya que el cálculo total de alimento a comprar no fue el correcto y compraron menos alimento del que necesitaban. Para mantener satisfecho a nuestro cliente deberemos realizar los test necesarios para verificar que todo funcione correctamente:
Verificar el cálculo de la cantidad de alimento para los perros.
Verificar el cálculo de la cantidad de alimento para los gatos.
Verificar el cálculo de la cantidad de alimento para los hamster.
Verificar el cálculo de la cantidad de alimento para las tarántulas.
*/

package animal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimal(t *testing.T) {
	t.Run("Validar el calculo de comida para tarantulas",
		func(t *testing.T) {
			animalFunc, err := Animal(Tarantula)
			if err != nil {
				panic(err)
			}
			var expectedResult float64 = 1.5

			obtainedResult := animalFunc(10)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 1.5")
		},
	)

	t.Run("Validar el calculo de comida para hamsters",
		func(t *testing.T) {
			animalFunc, err := Animal(Hamster)
			if err != nil {
				panic(err)
			}
			var expectedResult float64 = 2.5

			obtainedResult := animalFunc(10)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 2.5")
		},
	)

	t.Run("Validar el calculo de comida para perros",
		func(t *testing.T) {
			animalFunc, err := Animal(Dog)
			if err != nil {
				panic(err)
			}
			var expectedResult float64 = 100

			obtainedResult := animalFunc(10)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 100")
		},
	)

	t.Run("Validar el calculo de comida para gatos",
		func(t *testing.T) {
			animalFunc, err := Animal(Cat)
			if err != nil {
				panic(err)
			}
			var expectedResult float64 = 50

			obtainedResult := animalFunc(10)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 50")
		},
	)
}
