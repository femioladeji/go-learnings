package main

import (
	"fmt"
)

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	var positions []int32
	for i, v := range scores {
		if i == 0 {
			positions = append(positions, 1)
		} else if v == scores[i-1] {
			positions = append(positions, positions[i-1])
		} else {
			positions = append(positions, positions[i-1]+1)
		}
	}
	var alicePositions []int32
	lastIndex := len(scores) - 1
	index := 0
	for index < len(alice) && lastIndex >= 0 {
		score := scores[lastIndex]
		if alice[index] < score {
			alicePositions = append(alicePositions, positions[lastIndex]+1)
			index++
		} else if score == alice[index] {
			alicePositions = append(alicePositions, positions[lastIndex])
			index++
		} else {
			lastIndex--
		}
	}
	for index < len(alice) {
		alicePositions = append(alicePositions, 1)
		index++
	}
	return alicePositions
}

func main() {
	scores := []int32{100, 90, 90, 80, 75, 60}

	alice := []int32{25, 50, 65, 77, 90, 102, 120, 120}
	fmt.Println(climbingLeaderboard(scores, alice))
}
