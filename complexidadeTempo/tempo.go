package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	findSum(9999999999999)
	timeEnd := time.Now()

	fmt.Println(timeEnd.Sub(timeStart).Milliseconds())
}

func findSum(n int) int {
	return n * (n + 1) / 2
}

func findSumIterate(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}