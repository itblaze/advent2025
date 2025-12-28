package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// declare vars over here
	// fileName := "test.txt"
	fileName := "advent.txt"
	filePath := "/Users/samanyanga/Dev/projects/advent2025/day4/data"
	dataArr := make([]string, 0, 10)

	println("Beginning of advent of code day 4...")
	fmt.Printf("Reading file: %v from location %v\n", fileName, filePath)

	file, _ := os.Open(fmt.Sprintf("%v/%v", filePath, fileName))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataArr = append(dataArr, scanner.Text())
	}

	fmt.Printf("Solution Part 1: %d\n", solution1(dataArr))
	fmt.Printf("Solution Part 2: %d\n", solution2(dataArr))

}

func solution1(dataArr []string) int {
	// convert []string --> [][]rune
	iLen := len(dataArr)
	jLen := len(dataArr[0])
	accessible := 0
	tp := 0

	for i := range len(dataArr) {
		for j := range len(dataArr[i]) {
			// is co-ordinate a roll of paper
			if dataArr[i][j] != '@' {
				continue
			}
			var char rune
			// i-1, j
			if isValid(i-1, j, iLen, jLen) {
				char = rune(dataArr[i-1][j])
				if isAt(char) {
					tp++
				}
			}
			// i+1, j
			if isValid(i+1, j, iLen, jLen) {
				char = rune(dataArr[i+1][j])
				if isAt(char) {
					tp++
				}
			}
			// i, j-1
			if isValid(i, j-1, iLen, jLen) {
				char = rune(dataArr[i][j-1])
				if isAt(char) {
					tp++
				}
			}
			// i, j+1
			if isValid(i, j+1, iLen, jLen) {
				char = rune(dataArr[i][j+1])
				if isAt(char) {
					tp++
				}
			}

			// i-1, j-1
			if isValid(i-1, j-1, iLen, jLen) {
				char = rune(dataArr[i-1][j-1])
				if isAt(char) {
					tp++
				}
			}
			// i-1, j+1
			if isValid(i-1, j+1, iLen, jLen) {
				char = rune(dataArr[i-1][j+1])
				if isAt(char) {
					tp++
				}
			}
			// i+1, j-1
			if isValid(i+1, j-1, iLen, jLen) {
				char = rune(dataArr[i+1][j-1])
				if isAt(char) {
					tp++
				}
			}
			// i+1, j+1
			if isValid(i+1, j+1, iLen, jLen) {
				char = rune(dataArr[i+1][j+1])
				if isAt(char) {
					tp++
				}
			}

			if tp < 4 {
				accessible++
			}
			tp = 0

		}
	}
	return accessible
}

func isValid(i, j, iMax, jMax int) bool {
	if i < 0 || i >= iMax || j < 0 || j >= jMax {
		return false
	}
	return true
}

func isAt(testChar rune) bool {
	return testChar == '@'
}

func toRuneDoubleSlice(dataArr []string) *[][]rune {
	var rneArr [][]rune
	rneArr = make([][]rune, 0, 10)

	for _, line := range dataArr {
		runeArray := []rune(line)
		rneArr = append(rneArr, runeArray)
	}
	return &rneArr
}

func solution2(dataArr []string) int {
	rneArr := toRuneDoubleSlice(dataArr)
	totalCount := 0

	for {
		count := accessTissuesSln2(rneArr)
		if count == 0 {
			break
		}
		totalCount += count
	}
	return totalCount

}

func accessTissuesSln2(rneArr *[][]rune) int {
	iLen := len(*rneArr)
	jLen := len((*rneArr)[0])
	accessible := 0
	tp := 0

	for i := range iLen {
		for j := range jLen {
			// is co-ordinate a roll of paper
			if (*rneArr)[i][j] != '@' {
				continue
			}
			var char rune
			// i-1, j
			if isValid(i-1, j, iLen, jLen) {
				char = (*rneArr)[i-1][j]
				if isAt(char) {
					tp++
				}
			}
			// i+1, j
			if isValid(i+1, j, iLen, jLen) {
				char = (*rneArr)[i+1][j]
				if isAt(char) {
					tp++
				}
			}
			// i, j-1
			if isValid(i, j-1, iLen, jLen) {
				char = (*rneArr)[i][j-1]
				if isAt(char) {
					tp++
				}
			}
			// i, j+1
			if isValid(i, j+1, iLen, jLen) {
				char = (*rneArr)[i][j+1]
				if isAt(char) {
					tp++
				}
			}

			// i-1, j-1
			if isValid(i-1, j-1, iLen, jLen) {
				char = (*rneArr)[i-1][j-1]
				if isAt(char) {
					tp++
				}
			}
			// i-1, j+1
			if isValid(i-1, j+1, iLen, jLen) {
				char = (*rneArr)[i-1][j+1]
				if isAt(char) {
					tp++
				}
			}
			// i+1, j-1
			if isValid(i+1, j-1, iLen, jLen) {
				char = (*rneArr)[i+1][j-1]
				if isAt(char) {
					tp++
				}
			}
			// i+1, j+1
			if isValid(i+1, j+1, iLen, jLen) {
				char = (*rneArr)[i+1][j+1]
				if isAt(char) {
					tp++
				}
			}

			if tp < 4 {
				accessible++

				// mark the spot with a .
				(*rneArr)[i][j] = '.'
			}
			tp = 0

		}
	}
	return accessible
}
