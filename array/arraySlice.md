## Array vs Slice

Para termos um array, conseguimos definir pelo valor := []int {}

<div align="center"><img src="array.png"></div>

Assim sendo a finalização desse modo:

<div align="center"><img src="resultado.png"></div>


## Slice

Diferentemente do array, definimos o slice sem tamanho. Criando uma lista literalmente vazia. Assim sendo mais versátil pois permite colocar 1 milhão de itens permanentes, mas devemos ter cuidado para não dar index *index out of bounds* (índice que está fora do limite) fazendo quebrar a aplicação.

<div align="center"><img src="array.png"></div>

- Para adicionar itens no slice, devemos usar **append**.

<div align="center"><img src="sliceexemplo.png"></div>