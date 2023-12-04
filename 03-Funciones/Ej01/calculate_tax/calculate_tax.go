package calculate_tax

func CalculateTax(salary float64) (tax float64) {
	tax = 0
	if salary > 50000 {
		tax = salary * .17
		if salary > 150000 {
			tax = salary * .27
		}
	}
	return
}
