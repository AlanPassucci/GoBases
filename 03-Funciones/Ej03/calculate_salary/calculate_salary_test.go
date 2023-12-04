/*
Ejercicio 3 - Test del salario
La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios, por ello nos piden realizar una serie de tests sobre nuestro programa. Necesitaremos realizar las siguientes pruebas en nuestro código:
Calcular el salario de la categoría “A”.
Calcular el salario de la categoría “B”.
Calcular el salario de la categoría “C”.
*/

package calculate_salary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSalary(t *testing.T) {
	t.Run("Validar calculo de salario para categoria A",
		func(t *testing.T) {
			var minutes uint = 9600
			category := "A"
			var expectedResult float64 = 720000

			obtainedResult := CalculateSalary(minutes, category)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 720000")
		},
	)

	t.Run("Validar calculo de salario para categoria B",
		func(t *testing.T) {
			var minutes uint = 9600
			category := "B"
			var expectedResult float64 = 288000

			obtainedResult := CalculateSalary(minutes, category)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 288000")
		},
	)

	t.Run("Validar calculo de salario para categoria C",
		func(t *testing.T) {
			var minutes uint = 9600
			category := "C"
			var expectedResult float64 = 160000

			obtainedResult := CalculateSalary(minutes, category)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 160000")
		},
	)
}
