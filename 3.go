package main
import "fmt"

func main() {
	var printed bool

	for i:=1; i <= 100; i++ {
		printed = false

		if i % 3 == 0 {
			fmt.Print("Fizz");
			printed = true
		}
		if i % 5 == 0 {
			fmt.Print("Buzz");
			printed = true
		}

		if !printed {
			fmt.Print(i)
		}

		fmt.Println()
	}
}