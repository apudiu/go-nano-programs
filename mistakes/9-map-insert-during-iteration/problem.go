package main

import "fmt"

func main() {
	m := map[int]bool{
		1: true,
		2: false,
		3: true,
	}

	for i, v := range m {
		if v { // insert in loop
			m[i+10] = true
		}
	}

	fmt.Printf("[ UNPREDICTABLE ]: %#v \n", m)
}
