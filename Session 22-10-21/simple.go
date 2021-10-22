package main

import (
	"fmt"
	"time"
)

func saysHello(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v says hello %vth time \n",name,i)
		
	}
}
func main() {
	go saysHello("alex")
	go saysHello("jerry")

	time.Sleep(5*time.Second)
}