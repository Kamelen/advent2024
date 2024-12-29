package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

// Messied the code a bit but verrrrrrry helpful to chekc against their examples
func printMatricWithCheck(matrix [][]rune, checks [][]bool) {
	for i, row := range matrix {
		for j, letter := range row {
			if !checks[i][j] {
				print(".")
			} else {
				print(string(letter))
			}
		}
		print("\n")
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
	var checkMatrix [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineRunes := []rune{}
		lineChecks := []bool{}

		// Process each character in the line
		for _, r := range line {
			if !unicode.IsSpace(r) {
				lineRunes = append(lineRunes, r)
				lineChecks = append(lineChecks, false)
			}
		}
		letterMatrix = append(letterMatrix, lineRunes)
		checkMatrix = append(checkMatrix, lineChecks)
	}

	// using separate hit sums just for debugging
	diagHits := 0
	// okay now lets start processing the matrix
	for i, row := range letterMatrix {
		for j, letter := range row {
			if letter == 'A' {
				diagHits += searchDiag(i, j, letterMatrix, checkMatrix)
			}
		}
	}

	printMatricWithCheck(letterMatrix, checkMatrix)

	log.Printf("diagHits %d", diagHits)
}

// Helper diagram to get my indicies right :)
//^i-x > 0                 len(lm[i]) > j+x ->
//
//< j-x > 0
//                M   M        S(i-1,j-1)   S(i-1)(j+1)                   AND any combination of both
//                  A                   A(i,j)
//                S   S       M(i+1, j-1)   M(i+1, j+1)
//
//
// v len[lm] > i+X

func searchDiag(i int, j int, lm [][]rune, cm [][]bool) int {
	hits := 0

	// diag 1 hit? can be SAM or MAS
	//\
	// \
	//  \
	log.Printf("diag1 %d %d", i, j)
	if i-1 >= 0 && j-1 >= 0 && len(lm) > i+1 && len(lm[i+1]) > j+1 {
		diag := string([]rune{lm[i-1][j-1], lm[i][j], lm[i+1][j+1]})
		if diag == "MAS" || diag == "SAM" {
			cm[i-1][j-1] = true
			cm[i][j] = true
			cm[i+1][j+1] = true
			hits++
		}
	}

	// diag 2 hit? can be SAM or MAS
	//     /
	//   /
	// /
	if i-1 >= 0 && len(lm[i-1]) > j+1 && j-1 >= 0 && len(lm) > i+1 {
		diag := string([]rune{lm[i+1][j-1], lm[i][j], lm[i-1][j+1]})
		if diag == "MAS" || diag == "SAM" {
			cm[i+1][j-1] = true
			cm[i][j] = true
			cm[i-1][j+1] = true
			hits++
		}
	}

	// We only count a hit if both diags are found, needs to be a X of MAS
	if hits == 2 {
		return 1
	} else {
		return 0
	}
}
