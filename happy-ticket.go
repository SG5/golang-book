package main

import (
    "fmt"
)

type Ticket struct {
    number uint32
}

func main() {
    ticket := new(Ticket)
    ticket.number = 121112
    fmt.Println(ticket.isHappy())
}

func (t *Ticket) isHappy () bool {
    firstHalf := uint32(t.number / 1000)
    secondHalf := uint32(t.number % 1000)

    return (firstHalf % 100 + firstHalf % 10 + firstHalf)

    fmt.Println(firstHalf, secondHalf)
    return true
}