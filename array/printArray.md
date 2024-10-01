## Printando elementos de uma lista

O Loop sempre se inicia aplicando o número 0 á variável sendo utilizada (com execeção á quando você quer iterar do fim para o começo).

```Go
func printArray(arr []int) {
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%d", arr[i])
    }
}
```