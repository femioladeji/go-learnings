package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func readCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvr := csv.NewReader(file)
	records, readErr := csvr.ReadAll()
	if err != nil {
		return nil, readErr
	}
	return records, nil
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "Specify CSV file, defaults to problems.csv")
	limit := flag.Int("limit", 30, "How many seconds do you want to use in answering the quiz")
	flag.Parse()
	records, err := readCSVFile(*csvFile)
	if err != nil {
		log.Fatalln("An error occurred", ":", err)
		return
	}
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0
	for i, question := range records {
		fmt.Printf("Problem #%v: %s = ", i+1, question[0])
		responseChannel := make(chan string)
		go func() {
			var response string
			fmt.Scanf("%s\n", &response)
			responseChannel <- response
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou got %v right out of %v", correct, len(records))
			return
		case response := <-responseChannel:
			if response == question[1] {
				correct++
			}
		}
	}
	fmt.Printf("You got %v right out of %v", correct, len(records))
}
