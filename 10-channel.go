package main
import (
	"fmt"
	"time"
)


func pinger (c chan<- string) {
	message := "ping "
	for i:=0 ;; i++ {
		c <- fmt.Sprint(message, i)
	}
}

func ponger (c chan<- string) {
	messag := "pong "
	for i:=0 ;; i++ {
		c <- fmt.Sprint(messag, i)
	}
}

func printer (c <-chan string) {
	for {
		fmt.Println(<- c)
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	input := ""
	fmt.Scanln(&input)

}