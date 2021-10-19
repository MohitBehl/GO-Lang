package main

import "sort"

type SearchStruct struct {
	algotype string
}

func (obj SearchStruct) searchInArray(arr []int, ele int) (int, string) {
	if obj.algotype == "linear" {
		return linearSearch(arr, ele)
	} else {
		sort.Ints(arr)
		return binarySearch(arr, ele)
	}
}