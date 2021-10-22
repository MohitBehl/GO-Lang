package main

import "fmt"


func fib(n int) []int{
	res := []int{}
	f := 0
	s := 1

	i := 1
	for i < n{
		res = append(res, f)

		th := f + s
		f  = s
		s = th

		i++
	}
	return res
}
func main() {
	n := 0

	fmt.Scan(&n)

	res := fib(n)

	fmt.Println(res)
}