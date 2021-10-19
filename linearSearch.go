package main

func linearSearch(arr []int,ele int) (int,string){
	if ele < 0 {
		return -1,"wrong input format"
	}
	for i:=0;i<len(arr);i++{
		if arr[i] == ele{
			return i,"no error"
		}
	}
	return -1,"element not found"
}