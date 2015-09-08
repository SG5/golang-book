package math

func Average(numbers []float64) float64 {
	total := float64(0)
	for _, x := range numbers {
		total += x
	}
	return total / float64(len(numbers))
}
