package main

import "fmt"

func countingSort(arr []int) []int {
    max := 0
    for _, v := range arr {
        if v > max {
            max = v
        }
    }

    count := make([]int, max+1)
    for _, v := range arr {
        count[v]++
    }

    sorted := []int{}
    for i, c := range count {
        for c > 0 {
            sorted = append(sorted, i)
            c--
        }
    }

    return sorted
}

func main() {
    arr := []int{4, 2, 2, 8, 3, 3, 1}
    sortedArr := countingSort(arr)
    fmt.Println("Sorted array:", sortedArr)
}
