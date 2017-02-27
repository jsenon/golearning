package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

/* A Channel Timeout example
select {
	case msg1 := < c1:
		fmt.Println("Message 1", msg1)
	case msg2 := < c1:
		fmt.Println("Message 2", msg1)
	case <- time.After(time.Second) * 8:
		fmt.Println("Timeout")
	default:
		fmt.Println("nothing ready")
}*/
