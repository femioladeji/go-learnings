package main

import (
	"fmt"
	"regexp"
)

func caesarCipher(s string, k int32) string {
	re := regexp.MustCompile(`[a-zA-Z]`)
	encrypted := re.ReplaceAllStringFunc(s, func(character string) string {
		asciiNum := int32(character[0])
		if asciiNum <= 90 {
			return string(((asciiNum - 65 + k) % 26) + 65)
		}
		return string(((asciiNum - 97 + k) % 26) + 97)
	})
	return encrypted
}

func main() {
	text := "There's-a-starman-waiting-in-the-sky"
	k := int32(3)
	fmt.Println(caesarCipher(text, k))
}
