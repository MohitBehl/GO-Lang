package main

import "fmt"

func prepSqMatrix(n int) [][]int {
	arr := make([][]int, n)

	for i := 0; i < 3; i++ {
		arr[i] = make([]int, n)
	}
	return arr
}

func prepMatrix(n int, m int) [][]int {
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	return arr
}

func display(mat [][]int) {
	fmt.Println("-----Matrix Display-----")
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Print(mat[i][j],"\t")
		}
		fmt.Println();
	}
	fmt.Println("-------------------------")
}