/*
Ejercicio 1 - Testear el impuesto del salario
La  empresa de chocolates que anteriormente necesitaba calcular el impuesto
de sus empleados al momento de depositar el sueldo de los mismos ahora nos
solicitó validar que los cálculos de estos impuestos están correctos.
Para esto nos encargaron el trabajo de realizar los test correspondientes para:
Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/

package calculate_tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	t.Run("Validar calculo de impuesto para sueldo menor a 50000",
		func(t *testing.T) {
			var salary float64 = 45000
			var expectedResult float64 = 0

			obtainedResult := CalculateTax(salary)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 0")
		},
	)

	t.Run("Validar calculo de impuesto para sueldo mayor a 50000",
		func(t *testing.T) {
			var salary float64 = 100000
			var expectedResult float64 = 17000

			obtainedResult := CalculateTax(salary)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 17000")
		},
	)

	t.Run("Validar calculo de impuesto para sueldo mayor a 150000",
		func(t *testing.T) {
			var salary float64 = 200000
			var expectedResult float64 = 54000

			obtainedResult := CalculateTax(salary)

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 54000")
		},
	)
}
