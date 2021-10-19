package main

import (
	"fmt"
	"sort"
)
func main() {
	n := 0
	arr := []int{}

	// array input
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		var vl = 0
		fmt.Scan(&vl)
		arr = append(arr, vl)
	}

	// element input
	ele:=0
	fmt.Scan(&ele)

	sort.Ints(arr)
	obj := SearchStruct{"binary"}

	idx,err := obj.searchInArray(arr,ele)
	// idx,err := linearSearch(arr,ele)
	fmt.Println(idx,"---",err)

	// sort.Ints(arr)

	// idx,err = binarySearch(arr,ele)
	// fmt.Println(idx,"---",err)
}