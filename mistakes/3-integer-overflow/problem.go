package main

import (
	"fmt"
	"math"
)

func main() {
	var n int8 = math.MaxInt8
	fmt.Printf("%#v \n", n) // this a signed int
	n++
	fmt.Printf("%#v \n", n) // it became unsigned & this is not detected at compile time
}
