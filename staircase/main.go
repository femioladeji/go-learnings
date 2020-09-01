package main

import (
	"fmt"
)

func staircase(n int32) {
	for i := n - 1; i >= 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print(" ")
		}
		for x := n - i; x > 0; x-- {
			fmt.Print("#")
		}
		if i > 0 {
			fmt.Println()
		}
	}
}

func main() {
	staircase(4)
}
