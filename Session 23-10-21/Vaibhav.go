package main

import (
	"fmt"
	//	"time"
)

func printing(name string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, " -- ", i)
	}
	channel <- "done"
}
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	channel3 := make(chan string)
	go printing("tortoise", channel1)
	// time.Sleep(time.Microsecond/10)
	go printing("abc", channel3)
	go printing("rabbit", channel2)
	<-channel1
	<-channel3
	<-channel2
	// fmt.Println("asd")
	//time.Sleep(100 * time.Millisecond)
}
