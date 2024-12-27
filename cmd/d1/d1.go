package main

import (
	"advent2024/internal/utils"
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read in the d1input.txt file
	file, err := os.Open("d1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Fields(line)
		utils.Assert(len(pair) == 2, "must only be 2 in a pair: "+strings.Join(pair, "  "))
		leftNumber, _ := strconv.Atoi(pair[0])
		rightNumber, _ := strconv.Atoi(pair[1])
		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort each list by lowest to highest leftNumber
	sort.Ints(left)
	sort.Ints(right)

	// calculate distances for each pair
	distances := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		distance := utils.Abs(left[i] - right[i])

		distances = append(distances, distance)
	}
	// sum distances
	sum := utils.SumInts(distances)
	log.Println(sum)
}
