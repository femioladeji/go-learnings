package main

import (
	"fmt"
)

type Cord struct {
	row int
	col int
}

type NextCords struct {
	currentIndex  int
	possibleCords []Cord
}

func getStartingPoint(matrix []string) Cord {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'M' {
				return Cord{row: i, col: j}
			}
		}
	}
	// will never get here because there's always one `m`
	return Cord{row: 0, col: 0}
}

func isCordinateValid(matrix []string, row int, col int, lastVisited []Cord) bool {
	if matrix[row][col] != 'X' {
		for i := len(lastVisited) - 1; i >= 0; i-- {
			if row == lastVisited[i].row && col == lastVisited[i].col {
				return false
			}
		}
		return true
	}
	return false
}

func getAllPossibleNextPoints(matrix []string, point Cord, lastVisited []Cord) []Cord {
	row := point.row
	column := point.col
	var cordinates []Cord
	if row >= 0 && row < len(matrix)-1 && isCordinateValid(matrix, row+1, column, lastVisited) {
		cordinates = append(cordinates, Cord{row + 1, column})
	}
	if row > 0 && isCordinateValid(matrix, row-1, column, lastVisited) {
		cordinates = append(cordinates, Cord{row - 1, column})
	}
	if column >= 0 && column < len(matrix[0])-1 && isCordinateValid(matrix, row, column+1, lastVisited) {
		cordinates = append(cordinates, Cord{row, column + 1})
	}
	if column > 0 && isCordinateValid(matrix, row, column-1, lastVisited) {
		cordinates = append(cordinates, Cord{row, column - 1})
	}
	return cordinates
}

func countLuck(matrix []string, k int32) string {
	point := getStartingPoint(matrix)
	wandCount := 0
	var possiblePaths []NextCords
	visitedPaths := []Cord{{row: -1, col: -1}}
	for matrix[point.row][point.col] != '*' {
		nextPossibleCords := getAllPossibleNextPoints(matrix, point, visitedPaths)
		visitedPaths = append(visitedPaths, point)
		if len(nextPossibleCords) == 1 {
			point = nextPossibleCords[0]
		} else if len(nextPossibleCords) > 1 {
			wandCount++
			possiblePaths = append(possiblePaths, NextCords{currentIndex: 0, possibleCords: nextPossibleCords})
			point = nextPossibleCords[0]
		} else {
			// track back
			// get list of last possible paths
			i := len(possiblePaths) - 1
			for i >= 0 {
				visitedPaths = visitedPaths[:len(visitedPaths)-1]
				currentIndex := possiblePaths[i].currentIndex
				if currentIndex+1 < len(possiblePaths[i].possibleCords) {
					point = possiblePaths[i].possibleCords[currentIndex+1]
					possiblePaths[i].currentIndex++
					break
				} else {
					possiblePaths = possiblePaths[:i]
					wandCount--
					i--
				}
			}
		}
	}
	if int32(wandCount) == k {
		return "Impressed"
	}
	return "Oops!"
}

func main() {
	// matrix := []string{".X.X......X", ".X*.X.XXX.X", ".XX.X.XM...", "......XXXX."}
	// matrix := []string{"*.M...", ".X.X.X", "XXX..."}
	matrix := []string{"*..M..", ".X.X.X", "XXX..."}

	fmt.Println(countLuck(matrix, 2))
}
