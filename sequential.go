package main

import "fmt"

func sequentialSearch(arr []int, target int) int {
    for i, val := range arr {
        if val == target {
            return i // Return the index of the target
        }
    }
    return -1 // Return -1 if the target is not found
}

func main() {
    arr := []int{10, 20, 30, 40, 50}
    var target int
    fmt.Print("Enter the number to search: ")
    fmt.Scan(&target)

    result := sequentialSearch(arr, target)
    if result != -1 {
        fmt.Printf("Target found at index %d\n", result)
    } else {
        fmt.Println("Target not found")
    }
}
