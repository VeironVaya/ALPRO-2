package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[len(arr)/2]
    left := []int{}
    right := []int{}

    for _, v := range arr {
        if v < pivot {
            left = append(left, v)
        } else if v > pivot {
            right = append(right, v)
        }
    }

    return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
    arr := []int{10, 7, 8, 9, 1, 5}
    sortedArr := quickSort(arr)
    fmt.Println("Sorted array:", sortedArr)
}
