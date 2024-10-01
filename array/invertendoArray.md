## Invertendo um array

Em diversos cenários esse desafio de inverter um array aparece, principalmente em Hackerrank ou Leetcode.

```Go
    func reverseArray(arr []int) []int {
        start := 0
        end := len(arr) -1

        for start < end {
            arr[start], arr[end] = arr[end], arr [start]
            start++
            end--
        }

        return arr
    }

```