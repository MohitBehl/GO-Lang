package main

import "fmt"

func main() {
	arr := [9]int{10, 2, 6, 5, 7, 8, 9, 3, 4}


	sl1 := arr[1:5]
	fmt.Printf("arr :-> type %T , %v",arr,arr)
	fmt.Printf("sl1 :-> type %T , %v",sl1,sl1)
}