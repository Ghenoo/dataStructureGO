package main

import "fmt"

func mainn() {
	var arr = []int{1, 2, 3, 4, 5,55,977,78888}

	printArr(arr)
}

func printArr(arr []int) {
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
}
