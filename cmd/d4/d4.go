package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

// Messied the code a bit but verrrrrrry helpful to chekc against their examples
// Allowed me to find a bug where i did not account for the final letter being an index[0]
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
	vertHits := 0
	horzHits := 0
	diagHits := 0
	// okay now lets start processing the matrix
	for i, row := range letterMatrix {
		for j, letter := range row {
			if letter == 'X' {
				vertHits += searchVert(i, j, letterMatrix, checkMatrix)
				horzHits += searchHoriz(i, j, letterMatrix, checkMatrix)
				diagHits += searchDiag(i, j, letterMatrix, checkMatrix)
			}
		}
	}

	printMatricWithCheck(letterMatrix, checkMatrix)

	log.Printf("vertHits %d, horzHits %d, diagHits %d, sum %d", vertHits, horzHits, diagHits, vertHits+horzHits+diagHits)
}

func searchVert(i int, j int, lm [][]rune, cm [][]bool) int {
	hits := 0

	//^
	if i-3 >= 0 {
		if string([]rune{lm[i][j], lm[i-1][j], lm[i-2][j], lm[i-3][j]}) == "XMAS" {
			cm[i][j] = true
			cm[i-1][j] = true
			cm[i-2][j] = true
			cm[i-3][j] = true
			hits++
		}
	}

	// v
	if len(lm) > i+3 {
		if string([]rune{lm[i][j], lm[i+1][j], lm[i+2][j], lm[i+3][j]}) == "XMAS" {
			cm[i][j] = true
			cm[i+1][j] = true
			cm[i+2][j] = true
			cm[i+3][j] = true
			hits++
		}
	}

	return hits
}

func searchHoriz(i int, j int, lm [][]rune, cm [][]bool) int {
	hits := 0
	//->
	if len(lm[i]) > j+3 {
		println()
		if string([]rune{lm[i][j], lm[i][j+1], lm[i][j+2], lm[i][j+3]}) == "XMAS" {
			cm[i][j] = true
			cm[i][j+1] = true
			cm[i][j+2] = true
			cm[i][j+3] = true
			hits++
		}
	}
	//<-
	if j-3 >= 0 {
		if string([]rune{lm[i][j], lm[i][j-1], lm[i][j-2], lm[i][j-3]}) == "XMAS" {
			cm[i][j] = true
			cm[i][j-1] = true
			cm[i][j-2] = true
			cm[i][j-3] = true
			hits++
		}
	}
	return hits
}

// Helper diagram to get my indicies right :)
//^i-x > 0                 len(lm[i]) > j+x ->
//              S(i-3, j+3)
//< j-x > 0   A(i-2, j+2)
//          M(i-1, j+1)
//        X(i,j)
//      M(i+1, j-1)
//    A(i+2, j-2)
//  S(i+3, j-3)
// v len[lm] > i+X

func searchDiag(i int, j int, lm [][]rune, cm [][]bool) int {
	hits := 0
	// -> ^ -i +j
	if i-3 >= 0 && len(lm[i-3]) > j+3 {
		if string([]rune{lm[i][j], lm[i-1][j+1], lm[i-2][j+2], lm[i-3][j+3]}) == "XMAS" {
			cm[i][j] = true
			cm[i-1][j+1] = true
			cm[i-2][j+2] = true
			cm[i-3][j+3] = true
			hits++
		}
	}

	// -> v +i +j
	if len(lm) > i+3 && len(lm[i+3]) > j+3 {
		if string([]rune{lm[i][j], lm[i+1][j+1], lm[i+2][j+2], lm[i+3][j+3]}) == "XMAS" {
			cm[i][j] = true
			cm[i+1][j+1] = true
			cm[i+2][j+2] = true
			cm[i+3][j+3] = true
			hits++
		}
	}

	// <- ^ -i -j
	if i-3 >= 0 && j-3 >= 0 {
		if string([]rune{lm[i][j], lm[i-1][j-1], lm[i-2][j-2], lm[i-3][j-3]}) == "XMAS" {
			cm[i][j] = true
			cm[i-1][j-1] = true
			cm[i-2][j-2] = true
			cm[i-3][j-3] = true
			hits++
		}
	}

	//<- v +i -j
	if len(lm) > i+3 && j-3 >= 0 {
		if string([]rune{lm[i][j], lm[i+1][j-1], lm[i+2][j-2], lm[i+3][j-3]}) == "XMAS" {
			cm[i][j] = true
			cm[i+1][j-1] = true
			cm[i+2][j-2] = true
			cm[i+3][j-3] = true
			hits++
		}
	}

	return hits
}
