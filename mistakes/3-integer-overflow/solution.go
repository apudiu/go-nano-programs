package main

import (
	"fmt"
	"math"
)

func main() {
	//var n int8 = math.MaxInt8
	//fmt.Printf("%#v \n", n)            // this a signed int
	//increment(n)                       // fails here when overflowing
	//fmt.Printf("Increment: %#v \n", n) // it became unsigned & this is not detected at compile time

	fmt.Println("-------------------------------------")
	fmt.Printf("Add: %#v \n", add(7, 120))

	fmt.Println("-------------------------------------")
	fmt.Printf("Multiply: %#v \n", mul(2, 64))
}

func increment(n int8) int8 {
	if n == math.MaxInt8 {
		panic("INC: Int8 overflow!")
	}
	return n + 1
}

func add(a, b int8) int8 {
	if a > math.MaxInt8-b {
		panic("ADD: Int8 overflow!")
	}
	return a + b
}

func mul(a, b int8) int8 {
	//  If one of the operands is equal to 0, it directly returns 0.
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 { // Checks if one of the operands is equal to 1
		return result
	}

	// Checks if one of the operands is equal to math.MinInt
	if a == math.MinInt8 || b == math.MinInt8 {
		panic("MUL1: Int8 overflow!")
	}

	fmt.Printf("--> res: %#v (%d x %d), res / b = %#v \n", result, a, b, result/b)

	// Checks if the multiplication leads to an integer overflow

	// when an integer overflow happens the result will be unexpected value (of multiplication)
	// so result/b will not equal to a
	if result/b != a {
		panic("MUL2: Int8 overflow!")
	}

	return result
}
