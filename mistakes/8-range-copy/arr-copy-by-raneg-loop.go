package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}

	for i, v := range a {

		// NOTICE: this isn't printed as the updated array (a) is copied before loop start
		a[2] = 10

		if i == 2 {
			fmt.Printf("-> %#v \n", v)
		}
	}

	fmt.Printf("%#v \n", "----------------------------------")

	for i, v := range &a {

		// NOTICE: this is printed as the updated array (a) because range is copying the address not the value
		a[2] = 10

		if i == 2 {
			fmt.Printf("-> %#v \n", v)
		}
	}

}
