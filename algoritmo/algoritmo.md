## Algoritmo

- Conjunto de instruções

Seguimos este exemplo para facilitar o entendimento:

Encontramos um livro de receita e nele é passado uma forma de fazer chá.

    Esquentar água
    Colocar o chá na caneca
    Adicionar água quente
    Adicionar chá quente nas outras canecas
    Você precisa de açucar? 
    - Se sim, adicionar açucar nas cancas
    - Se não, não faça nada

```GO
package main

// Printe o valor médio de 3 números recebidos.
// Somar os 3 números e dividir por 3.
// Guargar em uma variável e printar.

func printMedia(a, b, c int) {
	soma := a + b + c
	media := soma / 3
	print(media)
}
```