package main

import (
	"fmt"
	"time"
)

func saysHello(name string,channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v says hello %vth time \n",name,i)
		time.Sleep(100*time.Millisecond)
	}

	channel <- "done"
}
func printing(channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100*time.Millisecond)
	}
	channel <- "done"
}
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go saysHello("jerry",channel1)
	go printing(channel2)
	
	<- channel1
	<- channel2
}