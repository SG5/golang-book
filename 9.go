package main
import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Human struct {
	name string
}

type Robot struct {
	Human
	model string
}

type Person interface {
	say()
}

func main() {
	c := new(Circle)
	*c = Circle{1, 10, 5}

	fmt.Println(*c)
	c.x = 3.0

	fmt.Println(circleArea(*c), c.area())

	r := new(Robot)
	r.name = "Arnold"
	h := new(Human)
	h.name = "Bobby"
	h2 := Human{"Ted"}

	chat(r, h, &h2)
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (h *Human) say() {
	fmt.Printf("Hello, my name is %s\r\n", h.name)
}

func chat (persons ...Person) {
	for _, p := range persons {
		p.say()
	}
}