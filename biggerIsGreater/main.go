package main

import (
	"fmt"
	"strings"
)

func biggerIsGreater(w string) string {
	endPad := []string{string(w[len(w)-1])}
	i := len(w) - 2
	found := false
	for i >= 0 && found != true {
		if w[i] < w[i+1] {
			found = true
		} else {
			endPad = append(endPad, string(w[i]))
			i--
		}
	}
	if found == false {
		return "no answer"
	}
	j := 0
	for j < len(endPad) {
		if endPad[j] > string(w[i]) {
			break
		}
		j++
	}
	answer := w[0:i] + endPad[j]
	endPad[j] = string(w[i])
	answer += strings.Join(endPad, "")
	return answer
}

func main() {
	answer := biggerIsGreater("abcd")
	fmt.Println(answer)
}
