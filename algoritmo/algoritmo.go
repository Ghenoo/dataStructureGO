package main

// Printe o valor médio de 3 números recebidos.
// Somar os 3 números e dividir por 3.
// Guargar em uma variável e printar.

func printMedia(a, b, c int) {
	soma := a + b + c
	media := soma / 3
	print(media)
}