package main

import "fmt"

func main() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice1 = []int{1,2,3}
	slice2 = make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)


	var mapExample map[string]uint8

	mapExample = make(map[string]uint8)

	mapExample["Test"] = 5;
	fmt.Println(mapExample)

	name, result := mapExample["q"]

	fmt.Println(name, result)

	target := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	min := target[0];
	for _, value := range target {
		if value < min {
			min = value
		}
	}

	fmt.Println(min)
}