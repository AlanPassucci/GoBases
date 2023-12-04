package calculate_avg

func CalculateAvg(ratings ...uint) (avg float64) {
	var sum uint = 0
	for _, rating := range ratings {
		sum += rating
	}
	avg = float64(sum) / float64(len(ratings))
	return
}
