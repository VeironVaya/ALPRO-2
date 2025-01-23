package main

import "fmt"

func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2 // Avoid overflow

        if arr[mid] == target {
            return mid // Target found
        } else if arr[mid] < target {
            left = mid + 1 // Search the right half
        } else {
            right = mid - 1 // Search the left half
        }
    }
    return -1 // Target not found
}

func main() {
    arr := []int{10, 20, 30, 40, 50} // Ensure this array is sorted
    var target int
    fmt.Print("Enter the number to search: ")
    fmt.Scan(&target)

    result := binarySearch(arr, target)
    if result != -1 {
        fmt.Printf("Target found at index %d\n", result)
    } else {
        fmt.Println("Target not found")
    }
}
