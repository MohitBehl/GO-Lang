package main

func binarySearch(arr []int, ele int) (int, string) {
	if ele < 0 {
		return -1, "wrong input format"
	}
	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		// fmt.Printf("lo : %v , hi : %v , mid : %v , val : %v \n",lo,hi,mid,arr[mid])

		if arr[mid] == ele {
			return mid, "no error"
		} else if arr[mid] < ele {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1, "element not found"
}
