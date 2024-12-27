package main

import (
	"advent2024/internal/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

type Report struct {
	levels []int
}

func (r *Report) HasSafeLevels() bool {
	asc := false
	desc := false

	for i := 1; i < len(r.levels); i++ {
		currLevel := r.levels[i]
		diff := currLevel - r.levels[i-1]

		if diff > 0 {
			asc = true
		} else {
			desc = true
		}

		// If unstable levels are detected then we exist early since it can't recover
		absDiff := utils.Abs(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	// Check here in the end if both asc and desc are detected, this is not safe so we return false
	if asc && desc {
		return false
	}

	// No issues detected, levels are safe
	return true
}

func reportFromString(stringLevels []string) Report {
	intLevels, _ := utils.StringToIntList(stringLevels)
	return Report{
		levels: intLevels,
	}
}

// So, a report only counts as safe if both of the following are true:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.

func main() {
	// Read in the d1input.txt file
	file, err := os.Open("d2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reports []Report

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levelsString := strings.Fields(line)
		report := reportFromString(levelsString)
		reports = append(reports, report)
	}

	numberOfSafeLevels := 0
	for _, report := range reports {
		if report.HasSafeLevels() {
			numberOfSafeLevels++
		}
	}

	log.Println(numberOfSafeLevels)
}
