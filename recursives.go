package main

import "fmt"

//fibonacci
func fibonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

//Factorial
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

//Sum of Natural Numbers
func sumOfNaturalNumbers(n int) int {
    if n == 0 {
        return 0
    }
    return n + sumOfNaturalNumbers(n-1)
}

//Power Function
func power(x, n int) int {
    if n == 0 {
        return 1
    }
    return x * power(x, n-1)
}

//Greatest Common Divisor (GCD)
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

//Binary Search
func binarySearch(arr []int, left, right, target int) int {
    if left > right {
        return -1
    }
    mid := left + (right-left)/2
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        return binarySearch(arr, mid+1, right, target)
    }
    return binarySearch(arr, left, mid-1, target)
}

//Reverse a String
func reverseString(s string) string {
    if len(s) == 0 {
        return ""
    }
    return reverseString(s[1:]) + string(s[0])
}

//Tower of Hanoi
func towerOfHanoi(n int, from, to, aux string) {
    if n == 1 {
        fmt.Printf("Move disk 1 from %s to %s\n", from, to)
        return
    }
    towerOfHanoi(n-1, from, aux, to)
    fmt.Printf("Move disk %d from %s to %s\n", n, from, to)
    towerOfHanoi(n-1, aux, to, from)
}

