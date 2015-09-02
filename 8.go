package main

import "fmt"

func main () {
	x:= 10

	fmt.Println(x)
	zero(x)
	fmt.Println(x)
	zeroPtr(&x)
	fmt.Println(x)

	newX := new(int)
	zeroPtr(newX)
	fmt.Println(*newX)

	y := 1.5
	square(&y)
	fmt.Println(y)

	xx := 1
	yy := 2
	swap(&xx, &yy)
	fmt.Println(xx, yy)
}

func zero (number int) {
	number = 0
}

func zeroPtr(number *int) {
	*number = 0
}

func square (x *float64) {
	*x = *x * *x
}

func swap (x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}