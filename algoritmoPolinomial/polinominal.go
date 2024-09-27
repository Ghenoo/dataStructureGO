package main

import "fmt"

func printValues(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			print(fmt.Sprintf("i = %d, j = %d", i, j))
		}
		print("finalizando")
	}

	print("fora do loop")
}