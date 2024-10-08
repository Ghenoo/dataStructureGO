## Análise de Algoritmos

Não existe apenas uma forma de resolver um certo problema, existem várias formas de fazerem isso, nem sempre o jeito mais fácil vai ser o mais perfomático.

Exemplo: 
 Dada uma solicitação de somar os primeiros números naturais até N, temos:

 - N = 4 / OUTPUT = (1+2+3+4) = 10 
 - N = 5 / OUTPUT = (1+2+3+4+5) = 15


## Exemplo em código: 

```GO
package main

//Jeito fácil, mas menos performático

func naturalNumbers(n int) (soma int) {
	for i := 1; i <= n; i++ {
		soma += i
	}
	return
}
```

```GO
package main

// Jeito "Difícil", mais performático.

func numerosNaturais(n int) int {
	return n * (n + 1) / 2
}
```