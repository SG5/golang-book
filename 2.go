package main


import (
	"fmt"
)

func main() {
	const RATE float32 = .3048;
	feet := 6;
	fmt.Println(float32(feet) * RATE);
}