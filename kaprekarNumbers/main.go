package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isNumberKaprekar(number int) bool {
	d := len(strconv.Itoa(number))
	squared := math.Pow(float64(number), 2)
	divisor := int(math.Pow(10, float64(d)))
	r := int(squared) % divisor
	l := int(squared) / divisor
	return (r + l) == number
}

func kaprekarNumbers(p int32, q int32) {
	answer := ""
	for i := p; i <= q; i++ {
		if isNumberKaprekar(int(i)) == true {
			answer += strconv.Itoa(int(i)) + " "
		}
	}
	if answer == "" {
		fmt.Println("INVALID RANGE")
	}
	fmt.Println(strings.TrimRight(answer, ""))
}

func main() {
	kaprekarNumbers(1, 100)
}
