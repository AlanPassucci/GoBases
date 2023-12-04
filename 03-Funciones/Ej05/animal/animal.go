package animal

import "errors"

const (
	Tarantula = "tarantula"
	Hamster   = "hamster"
	Dog       = "dog"
	Cat       = "cat"
)

type AnimalFunc func(qty int) float64

func Animal(animal string) (AnimalFunc, error) {
	switch animal {
	case Tarantula:
		return CalculateTarantulaFood, nil
	case Hamster:
		return CalculateHamsterFood, nil
	case Dog:
		return CalculateDogFood, nil
	case Cat:
		return CalculateCatFood, nil
	default:
		return nil, errors.New("error: animal not found")
	}
}

func CalculateTarantulaFood(qty int) float64 {
	return float64(qty) * .15
}

func CalculateHamsterFood(qty int) float64 {
	return float64(qty) * .25
}

func CalculateDogFood(qty int) float64 {
	return float64(qty) * 10
}

func CalculateCatFood(qty int) float64 {
	return float64(qty) * 5
}
