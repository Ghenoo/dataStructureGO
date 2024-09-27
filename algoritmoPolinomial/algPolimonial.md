## Algoritmo Polinominal

Mais formalmente, um algoritmo é considerado polinomial se o tempo de execução T(n) for limitado por um polinômio de n, onde "n" representa o tamanho da entrada. Por exemplo, se T(n) for da forma I(n) = O(n^^2) ou T(n)= O(n^3), então o algoritmo é considerado polinomial.

## Exemplo

```Go
func printValues(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			print(fmt.Sprintf("i = %d, j = %d", i, j))
		}
		print("finalizando")
	}

	print("fora do loop")
}
```