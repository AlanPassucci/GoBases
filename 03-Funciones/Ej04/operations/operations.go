package operations

import "errors"

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

type OperationFunc func(califications ...int) (float64, error)

func Operation(operation string) (OperationFunc, error) {
	switch operation {
	case Minimum:
		return CalculateMinFunc, nil
	case Average:
		return CalculateAverageFunc, nil
	case Maximum:
		return CalculateMaxFunc, nil
	default:
		return nil, errors.New("error: undefined calculation")
	}
}

func CalculateMinFunc(califications ...int) (float64, error) {
	if len(califications) == 0 {
		return 0, errors.New("error: not enough values")
	}

	minCalification := float64(califications[0])
	for _, calification := range califications {
		if minCalification > float64(calification) {
			minCalification = float64(calification)
		}
	}
	return minCalification, nil
}

func CalculateAverageFunc(califications ...int) (float64, error) {
	if len(califications) == 0 {
		return 0, errors.New("error: not enough values")
	}

	sum := 0
	for _, calification := range califications {
		sum += calification
	}
	result := float64(sum) / float64(len(califications))
	return result, nil
}

func CalculateMaxFunc(califications ...int) (float64, error) {
	if len(califications) == 0 {
		return 0, errors.New("error: not enough values")
	}

	maxCalification := float64(califications[0])
	for _, calification := range califications {
		if maxCalification < float64(calification) {
			maxCalification = float64(calification)
		}
	}
	return maxCalification, nil
}
