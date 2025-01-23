package main

import "fmt"

func main(){

}
//Right-Angled Triangle Pattern
func printRightAngledTriangle(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

Inverted Right-Angled Triangle Pattern
func printInvertedTriangle(n int) {
    for i := n; i >= 1; i-- {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}


//Equilateral Triangle Pattern
func printEquilateralTriangle(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

//Diamond Pattern
func printDiamond(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    for i := n - 1; i >= 1; i-- {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

//Pyramid of Numbers
func printNumberPyramid(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= i; k++ {
            fmt.Printf("%d ", k)
        }
        fmt.Println()
    }
}

//Reverse Pyramid of Numbers
func printReverseNumberPyramid(n int) {
    for i := n; i >= 1; i-- {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= i; k++ {
            fmt.Printf("%d ", k)
        }
        fmt.Println()
    }
}

//Square Pattern
func printSquare(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

//Hollow Square Pattern
func printHollowSquare(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if i == 1 || i == n || j == 1 || j == n {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

//Checkerboard Pattern
func printCheckerboard(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if (i+j)%2 == 0 {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

//Right-Angled Triangle of Numbers
func printNumberTriangle(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d ", j)
        }
        fmt.Println()
    }
}


