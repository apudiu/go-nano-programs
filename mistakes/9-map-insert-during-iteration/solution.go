package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[int]bool{
		1: true,
		2: false,
		3: true,
	}

	m2 := maps.Clone(m)

	for i, v := range m {
		if v { // insert in loop
			m2[i+10] = true
		}
	}

	fmt.Printf("[ PREDICTABLE ]: %#v \n", m2)
}
