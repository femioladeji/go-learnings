package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func convertNumberToWords(number int32) string {
	digits := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	if number <= 20 {
		return digits[number]
	}
	unitPlace := number % 20
	return "twenty " + digits[unitPlace]
}

func timeInWords(h int32, m int32) string {
	minuteText := " minutes"
	if m == 59 || m == 1 {
		minuteText = " minute"
	}
	if m <= 30 {
		hour := convertNumberToWords(h)
		if m == 0 {
			return hour + " o' clock"
		}
		if m == 15 {
			return "quarter past " + hour
		}
		if m == 30 {
			return "half past " + hour
		}
		return convertNumberToWords(m) + minuteText + " past " + hour
	}
	hour := convertNumberToWords(h + 1)
	if h == 12 {
		hour = convertNumberToWords(1)
	}
	if m == 45 {
		return "quarter to " + hour
	}
	return convertNumberToWords(60-m) + minuteText + " to " + hour
}

func main() {
	time := "5:01"
	timeSplit := strings.Split(time, ":")
	hour, err := strconv.Atoi(timeSplit[0])
	minute, err := strconv.Atoi(timeSplit[1])
	if err != nil {
		log.Fatal(err)
	}
	answer := timeInWords(int32(hour), int32(minute))
	fmt.Println(answer)
}
