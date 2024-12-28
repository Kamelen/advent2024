package main

import (
	"advent2024/internal/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

func extractMulNumbers(line string) (int, int) {
	// Trying to find out of the given substring contains "x,y)"
	shouldContainMul := strings.Split(line, ")")[0]
	numbersAsString := strings.Split(shouldContainMul, ",")
	numbers, _ := utils.StringToIntList(numbersAsString)

	// Just checking for safety to exclude strange strings lime mul(655somethingblä)
	if len(numbers) == 2 {
		return numbers[0], numbers[1]
	} else {
		return 0, 0
	}
}

func calcMulsInLine(line string) int {
	sum := 0
	enabled := true

	l := len(line)
	for i := 0; i < len(line); i++ {
		// Check if we find a do or dont
		if l > i+7 && line[i:i+7] == "don't()" {
			println("dont")
			enabled = false
		} else if l > i+4 && line[i:i+4] == "do()" {
			println("do")
			enabled = true
		}

		// make sure to check that the line is still big enough before checking for mul(
		if l > i+3 && line[i:i+4] == "mul(" {
			mul1, mul2 := extractMulNumbers(line[i+4:])
			if enabled {
				sum += mul1 * mul2
			}
		}
	}

	return sum
}

func main() {
	// Read in the d1input.txt file
	file, err := os.Open("d3input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	oneLine := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		oneLine += line
	}

	sum += calcMulsInLine(oneLine)
	println(sum)
}
