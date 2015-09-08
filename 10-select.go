package main
import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "func 1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for{
			c2 <- "func 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func () {
		var msg string
		for {
			select {
			case msg = <- c1:
				fmt.Println(msg)
			case msg = <- c2:
				fmt.Println(msg)
			}
		}
	}()

	var in string;
	fmt.Scanln(&in)
}
