package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
	positive := 0
	negative := 0
	zero := 0
	arrayLength := len(arr)
	for i := 0; i < arrayLength; i++ {
		if arr[i] < 0 {
			negative += 1
		} else if arr[i] > 0 {
			positive += 1
		} else {
			zero += 1
		}
	}
	fmt.Printf("%.6f\n", float64(positive)/float64(arrayLength))
	fmt.Printf("%.6f\n", float64(negative)/float64(arrayLength))
	fmt.Printf("%.6f\n", float64(zero)/float64(arrayLength))
}

func main() {
	numberArray := []int32{3, 5, -2, 0}
	plusMinus(numberArray)
}
