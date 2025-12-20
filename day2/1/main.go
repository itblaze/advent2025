package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare vars over here
	// fileName := "test.txt"
	fileName := "advent2.txt"
	filePath := "/home/posytron/learn/advent/day2/1/data"

	println("Beginning of advent of code day 2...")
	fmt.Printf("Reading file: %v from location %v\n", fileName, filePath)

	file, _ := os.Open(fmt.Sprintf("%v/%v", filePath, fileName))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fileContents := scanner.Text()

	fmt.Printf("File contents read from file: %s\n", fileContents)

	ranges := strings.Split(fileContents, ",")
	invalidIds := make([]int, 0, 10)

	for _, rng := range ranges {
		r := strings.Split(rng, "-")

		lower := r[0]
		upper := r[1]

		lowerInt, _ := strconv.Atoi(lower)
		upperInt, _ := strconv.Atoi(upper)

		for i := lowerInt; i <= upperInt; i++ {
			iString := strconv.Itoa(i)
			iLen := len(iString)

			// fmt.Printf("A string => %v and B string is %v\n", a, b)

			if isInvalid(i, iLen) {
				fmt.Printf("This number is considered invalid: %d \n", i)
				invalidIds = append(invalidIds, i)
			}

		}
	}
	total := 0
	for _, val := range invalidIds {
		total = total + val
	}
	fmt.Printf("Invalid number total is: %v\n", total)

}

func isInvalid(candidate int, size int) bool {
	if size == 0 {
		return false
	}

	candidateString := strconv.Itoa(candidate)
	candidateLen := len(candidateString)

	isDivisible := candidateLen % size
	if isDivisible != 0 {
		return isInvalid(candidate, size-1) || false
	}

	chunkSize := candidateLen / size

	if chunkSize < 2 {
		return isInvalid(candidate, size-1) || false
	}

	isValid := true
	for i := 0; i < candidateLen; i = size + i {
		if i == 0 {
			continue
		}
		// curr = i  and prev i - size
		prev := i - size
		curr := i
		// compare str[curr] & str[prev]

		comparison := strings.Compare(candidateString[curr:curr+size], candidateString[prev:prev+size])
		if comparison == 0 {
			// they are equal
			continue
		} else {
			// they are not equal
			isValid = false
			break
		}
	}
	if !isValid {
		return isInvalid(candidate, size-1) || false
	}
	return true
}
