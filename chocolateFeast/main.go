package main

import (
	"fmt"
)

func chocolateFeast(n int32, c int32, m int32) int32 {
	remainingWraps := n / c
	chocolate := remainingWraps
	for remainingWraps >= m {
		newChocolate := remainingWraps / m
		remainingWraps = newChocolate + remainingWraps%m
		chocolate += newChocolate
	}
	return chocolate
}

func main() {
	answer := chocolateFeast(43203, 60, 5)
	// answer := chocolateFeast(15, 3, 2) // 9
	fmt.Println(answer)
}

// 2,2
// 16728
