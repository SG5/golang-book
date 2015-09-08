package main

import (
	"fmt"
	m "./math"
)

func main () {
	numbers := []float64{1, 2, 3, 4}

	fmt.Println(m.Average(numbers))
}