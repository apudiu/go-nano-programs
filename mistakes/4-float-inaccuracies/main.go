package main

import "fmt"

func main() {
    n := 500

    fmt.Printf("F1: %#v \n", f1(n))
    fmt.Printf("F2: %#v \n", f2(n))
}

func f1(n int) float64 {
    result := 10_000.
    for i := 0; i < n; i++ {
        result += 1.0001
    }
    return result
}

func f2(n int) float64 {
    result := 0.
    for i := 0; i < n; i++ {
        result += 1.0001
    }
    return result + 10_000.
}
