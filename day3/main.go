package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	// declare vars over here
	// fileName := "test.txt"
	fileName := "advent.txt"
	filePath := "/Users/samanyanga/Dev/projects/advent2025/day3/data"
	dataArr := make([]string, 0, 10)

	println("Beginning of advent of code day 3...")
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
	var totalVoltage int
	for _, val := range dataArr {
		// fmt.Printf("The first letter of input %v is %c\n", idx, val[0])
		// fmt.Printf("The value of the number is %v\n", val[1:])
		bankLen := len(val)
		leftPtr := bankLen - 2
		rightPtr := bankLen - 1

		for i := bankLen - 2; i >= 0; i-- {
			valAtI := val[i] - '0'
			valtAtLeftPtr := val[leftPtr] - '0'

			if valAtI >= (valtAtLeftPtr) {
				leftPtr = i

			}
			// fmt.Printf("Previous value: %v new left pointer: %v & left point => %v\n", valtAtLeftPtr, valAtI, leftPtr)
		}

		for i := bankLen - 1; i >= 0; i-- {
			if i <= leftPtr {
				break
			}
			valAtI := val[i] - '0'
			valAtRightPtr := val[rightPtr] - '0'

			if valAtI >= valAtRightPtr {
				rightPtr = i
			}

		}
		computedValue := (val[leftPtr]-'0')*10 + val[rightPtr] - '0'
		// fmt.Printf("Max voltage ---> %d\n", computedValue)
		totalVoltage += int(computedValue)
	}
	return totalVoltage
}

func solution2(dataArr []string) int {
	var totalVoltage int
	for _, val := range dataArr {
		// fmt.Printf("The first letter of input %v is %c\n", idx, val[0])
		// fmt.Printf("The value of the number is %v\n", val[1:])

		maxVoltage := extractVoltage(val, 12)
		fmt.Printf("Max voltage from %v is ---> %d\n", val, maxVoltage)
		totalVoltage += int(maxVoltage)
	}
	return totalVoltage
}

func extractVoltage(batteries string, remaining int) int {

	batteryLen := len(batteries)
	if remaining == batteryLen {
		voltage, _ := strconv.Atoi(batteries)
		fmt.Printf("Remaining voltage: %v\n", voltage)
		return voltage
	}
	if remaining == 0 {
		return 0
	}

	subArraySize := (batteryLen + 1) - remaining // +1 because we are inclusive of the nth number

	ptr := 0

	for i := 0; i < subArraySize; i++ {
		fmt.Printf("SubArraySize: %v Batteries %v Remaining %v\n", subArraySize, batteries, remaining)
		valAtI := batteries[i] - '0'
		valAtPtr := batteries[ptr] - '0'

		if valAtI > valAtPtr {
			ptr = i
		}

	}

	ptrVal := batteries[ptr] - '0'
	fmt.Printf("Chosen value is -> %v \n", (int(math.Pow(10, float64(remaining))) * int(ptrVal)))
	ptrValEx := int(math.Pow(10, float64(remaining-1))) * int(ptrVal)
	return ptrValEx + extractVoltage(batteries[(1+ptr):], remaining-1)
}
