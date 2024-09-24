package main

//Jeito fácil, mas menos performático

func naturalNumbers(n int) (soma int) {
	for i := 1; i <= n; i++ {
		soma += i
	}
	return
}