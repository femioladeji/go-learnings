package main

import (
	"flag"
	"fmt"
	"regexp"
)

func camelCase(text string) int32 {
	re := regexp.MustCompile(`[A-Z]`)
	return int32(len(re.FindAllString(text, -1))) + 1
}

func main() {
	text := flag.String("word", "oneTwoThree", "The camelCase word you want to know the number of words in it")
	flag.Parse()
	fmt.Println(camelCase(*text))
}
