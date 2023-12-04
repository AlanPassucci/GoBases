package calculate_salary

func CalculateSalary(minutes uint, category string) (salary float64) {
	hours := float64(minutes) / 60
	switch category {
	case "A":
		salary = hours * 3000
		salary += salary * .5
	case "B":
		salary = hours * 1500
		salary += salary * .2
	case "C":
		salary = hours * 1000
	default:
		salary = 0
	}
	return
}
