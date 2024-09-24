## Time Complexity

- De forma simples, é a forma de ter uma noção de quanto tempo seu algoritmo levaria para terminar sua execução.

- O input recebido nos ajuda a ter uma ideia de como nosso algoritmo vai performar (exemplo: se for um número maior, tende a demorar mais.)

## Exemplo

```Go
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

```