package main

import "fmt"

func saysHi(name string) (string,string){
	if len(name) == 0{
		return "","error"
	}

	return name+" says hello","no error"
}
func saysBye() {
	fmt.Println("says bye")
}
func main() {
	msg,err := saysHi("alex")
	arr := []int{10,20,30,5,45,6}
	if err == "error"{
		fmt.Println("wrong input")
	}else{
		fmt.Println(msg)
	}
	saysBye()
}