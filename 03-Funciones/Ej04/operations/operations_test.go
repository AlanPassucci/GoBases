/*
Ejercicio 4 - Testear el cálculo de estadísticas
Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google, el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:
Realizar test para calcular el mínimo de calificaciones.
Realizar test para calcular el máximo de calificaciones.
Realizar test para calcular el promedio de calificaciones.
*/

package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperations(t *testing.T) {
	t.Run("Validar el minimo de notas",
		func(t *testing.T) {
			operationFunc, err := Operation(Minimum)
			if err != nil {
				panic(err)
			}
			var expectedResult float64 = 2

			obtainedResult, err := operationFunc(2, 3, 3, 4, 10, 2, 4, 5)
			if err != nil {
				panic(err)
			}

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 2")
		},
	)

	t.Run("Validar el promedio de notas",
		func(t *testing.T) {
			operationFunc, err := Operation(Average)
			if err != nil {
				panic(err)
			}
			var expectedResult float64 = 3

			obtainedResult, err := operationFunc(2, 3, 3, 4, 1, 2, 4, 5)
			if err != nil {
				panic(err)
			}

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 3")
		},
	)

	t.Run("Validar el maximo de notas",
		func(t *testing.T) {
			operationFunc, err := Operation(Maximum)
			if err != nil {
				panic(err)
			}
			var expectedResult float64 = 5

			obtainedResult, err := operationFunc(2, 3, 3, 4, 1, 2, 4, 5)
			if err != nil {
				panic(err)
			}

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 5")
		},
	)
}
