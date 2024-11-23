package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("len: %d, cap: %d v: %#v \n", len(s1), cap(s1), s1)

	//update(s1[:2])// can modify values out of len
	update(s1[:2:2]) // can't modify values out of len, another option can be copy the src slice & pass it

	fmt.Printf("len: %d, cap: %d v: %#v \n", len(s1), cap(s1), s1)
}

func update(l []int) {
	// modify values out of len
	_ = append(l, 5)
}
