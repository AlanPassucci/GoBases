/*
El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente, por lo que ahora
nos corresponde realizar los test correspondientes:
Calcular el promedio de las notas de los alumnos.
*/

package calculate_avg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSalary(t *testing.T) {
	ratings := []uint{8, 6, 10}
	var expectedResult float64 = 8

	obtainedResult := CalculateAvg(ratings...)

	assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 8")
}
