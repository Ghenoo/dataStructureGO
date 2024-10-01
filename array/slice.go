package main

import "fmt"

func main() {
	var Slice = []int{1, 2, 3, 4, 5,5125125,924,233}

	printArray(Slice)

}

func printArray(list []int) {
	for _, item := range list {
		fmt.Printf("%d ", item)
	}
}