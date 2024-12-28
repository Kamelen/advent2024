package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func printMatrix(matrix [][]rune) {
	// Print the 2D array
	fmt.Println("2D array of letters:")
	for i, row := range matrix {
		fmt.Printf("Row %d: %c\n", i, row)
	}
}

func main() {
	// Read in the d1input.txt file
	file, err := os.Open("d4input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var letterMatrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineRunes := []rune{}

		// Process each character in the line
		for _, r := range line {
			if !unicode.IsSpace(r) {
				lineRunes = append(lineRunes, r)
			}
		}
		letterMatrix = append(letterMatrix, lineRunes)
	}

	printMatrix(letterMatrix)
	// okay now lets start processing the matrix
	for i, row := range letterMatrix {
		for j, letter := range row {
			if letter == 'X' {
				searchVert(i, j, letterMatrix)
				searchHoriz(i, j, letterMatrix)
				searchDiag(i, j, letterMatrix)
			}
		}
	}
}

func searchVert(i int, j int, lm [][]rune) int {
	hits := 0

	//^
	if i-3 > 0 {
		if string([]rune{lm[i][j], lm[i-1][j], lm[i-2][j], lm[i-3][j]}) == "XMAS" {
			hits++
		}
	}

	// v
	if len(lm) > i+3 {
		if string([]rune{lm[i][j], lm[i+1][j], lm[i+2][j], lm[i+3][j]}) == "XMAS" {
			hits++
		}
	}
	if hits > 0 {
		fmt.Printf("v[%d][%d]: %d\n", i, j, hits)
	}
	return hits
}

func searchHoriz(i int, j int, lm [][]rune) int {
	hits := 0
	//->
	if len(lm[i]) > j+3 {
		println()
		if string([]rune{lm[i][j], lm[i][j+1], lm[i][j+2], lm[i][j+3]}) == "XMAS" {
			hits++
		}
	}
	//<-
	if j-3 > 0 {
		if string([]rune{lm[i][j], lm[i][j-1], lm[i][j-2], lm[i][j-3]}) == "XMAS" {
			hits++
		}
	}
	if hits > 0 {
		fmt.Printf("h[%d][%d]: %d\n", i, j, hits)
	}
	return hits
}

func searchDiag(i int, j int, lm [][]rune) int {
	return 0
}
