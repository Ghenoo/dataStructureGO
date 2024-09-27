package main

func calculate(n int) int {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	return soma
}