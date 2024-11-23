package main

import (
    "fmt"
    "math"
)

func main() {
    var n int8 = math.MaxInt8
    fmt.Printf("%#v \n", n) // this a signed int
    increment(n)            // fails here when overflowing
    fmt.Printf("%#v \n", n) // it became unsigned & this is not detected at compile time
}

func increment(n int8) int8 {
    if n == math.MaxInt8 {
        panic("Int8 overflow!")
    } // sdf
    return n + 1
}
