package main

import (
	"fmt"
	"math"
	"strings"
)

func encryption(s string) string {
	strippedText := strings.ReplaceAll(s, " ", "")
	columns := int(math.Ceil(math.Sqrt(float64(len(strippedText)))))
	rows := columns
	if (rows * columns) < len(strippedText) {
		rows += 1
	}
	envryptedText := ""
	for i := 0; i < columns; i++ {
		for j := i; j < len(strippedText); j += columns {
			envryptedText += string(strippedText[j])
		}
		envryptedText += " "
	}
	return strings.Trim(envryptedText, " ")
}

func main() {
	text := "feed the dog"
	fmt.Println(encryption(text))
}
