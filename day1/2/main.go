package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// declare vars over here
	// fileName := "test.txt"
	fileName := "advent1.txt"
	filePath := "/home/posytron/learn/advent/1/data"
	count := 0
	countOfZeroes := 0
	countOfZeroDuringRotations := 0
	countOfCross := 0
	dial := 50
	dataArr := make([]string, 0, 10)

	println("Beginning of advent of code day 1...")
	fmt.Printf("Reading file: %v from location %v\n", fileName, filePath)

	file, _ := os.Open(fmt.Sprintf("%v/%v", filePath, fileName))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataArr = append(dataArr, scanner.Text())
		count++
	}

	for _, val := range dataArr {
		// fmt.Printf("The first letter of input %v is %c\n", idx, val[0])
		// fmt.Printf("The value of the number is %v\n", val[1:])

		direction := val[0]
		valAsInt, err := strconv.Atoi(val[1:])
		if err != nil {
			panic("Critical error with input")
		}
		switch direction {
		case 'L':
			valAsInt = -valAsInt
		}
		// fmt.Printf("Value of dial currently: %d\n", dial)
		// fmt.Printf("The value of number is %v\n", valAsInt)

		// adjust for numbers bigger than 100
		div := valAsInt / 100

		if div > 0 {
			fmt.Printf("Div value : %d, countOfZeroDuringRations: %d \n", div, countOfZeroDuringRotations)
		}
		countOfZeroDuringRotations = countOfZeroDuringRotations + abs(div)
		valAsInt = valAsInt % 100

		prev := dial
		dial = dial + valAsInt

		if dial == 0 || dial == 100 {
			countOfZeroes++
		}
		// re-adjust dial
		switch {
		case dial < 0:
			dial = dial + 100
			if prev != 0 {
				countOfCross++
				fmt.Println("Crossed Zero")
			}

		case dial > 100:
			dial = dial - 100
			if prev != 0 {
				countOfCross++
				fmt.Println("Crossed Zero")
			}

		case dial == 100:
			dial = 0
		default:
			// fmt.Println("Found nothing")
		}

	}
	fmt.Printf("Number of zeros is settled %v \n", countOfZeroes)
	fmt.Printf("Number of zeros during multiple rotations is %d \n", countOfZeroDuringRotations)
	fmt.Printf("Number of zeros during cross is %v \n", countOfCross)
	fmt.Printf("Total is %v \n", countOfCross+countOfZeroes+countOfZeroDuringRotations)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
