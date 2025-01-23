# Formatting in Golang

## Basic Formatting

```golang
package main 

import "fmt"

func main(){

//basic formating
fmt.Printf("Hello, %s!\n", "world") // Hello, world!
fmt.Printf("I have %d apples.\n", 5) // I have 5 apples.
fmt.Printf("Pi is approximately %.2f.\n", 3.14159) // Pi is approximately 3.14.
}
```

## General Verbs
|verb|Desc|Example|
|---|---|---|
|%v|	Default format for the value|	`fmt.Printf("%v\n", 42)`|
|%+v|	Adds field names for structs|`fmt.Printf("%+v\n", struct{X int}{X: 10})`|
|%#v|	Go-syntax representation of the value|`fmt.Printf("%#v\n", 42)`|
|%T|Type of the value|`	fmt.Printf("%T\n", 42)`|
|%%|A literal percent sign|	`fmt.Printf("%%\n")`|

## Integer Formatting
|verb|Desc|Example|
|---|---|---|
|%d|	Decimal|	`	fmt.Printf("%d\n", 123)`|
|%b|	Binary|`	fmt.Printf("%b\n", 5)`|
|%o|	Octal|`	fmt.Printf("%o\n", 8)`|
|%x|Hexa (lower-case)|`		fmt.Printf("%x\n", 255)`|
|%X|Hexa (upper-case)|	`fmt.Printf("%X\n", 255)`|
|%c|Character|	`fmt.Printf("%c\n", 65)`|
|%U|Unicode|	`	fmt.Printf("%U\n", 'A')`|

## Data-Type Formatting
|verb|Desc|Example|
|---|---|---|
|%f|		Decimal point with no exponent|	`		fmt.Printf("%f\n", 3.14)`|
|%.nf|	Limits decimal places (e.g., n = 2)|	`		fmt.Printf("%.2f\n", 3.14159)`|
|%s|		String or byte slice|`		fmt.Printf("%s\n", "Go")`|
|%t|	Boolean|`		fmt.Printf("%t\n", true)`|
|%p|Base-16 pointer address|`		fmt.Printf("%p\n", &x)`|

## Width and Precision Control

```golang
fmt.Printf("|%6d|%6d|\n", 12, 345)     // |    12|   345|
fmt.Printf("|%-6d|%-6d|\n", 12, 345)   // |12    |345   |
fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // |  1.20|  3.45|
```

## Formatting Custom Types (Structs)

```golang
type Point struct {
    X, Y int
}

p := Point{10, 20}
fmt.Printf("%v\n", p)  // {10 20}
fmt.Printf("%+v\n", p) // {X:10 Y:20}
fmt.Printf("%#v\n", p) // main.Point{X:10, Y:20}

```

## Dynamic Width and Precision

```golang
fmt.Printf("%*.*f\n", 8, 2, 123.456) // "  123.46"

```



