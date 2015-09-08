package main
import (
	"time"
	"fmt"
)

func main () {
	fmt.Println("start")
	Sleep(3)
	fmt.Println("end after 3 seconds")
}

func Sleep (seconds time.Duration) {
	select {
	case <- time.After(time.Second * seconds):
	}
}