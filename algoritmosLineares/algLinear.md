## Algoritmo Linear

Um algoritmo linear em Go percorre a entrada uma única vez, com complexidade O(n), realizando operações em cada elemento.

```Go
func calculate(n int) int {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	return soma
 }
 ```
