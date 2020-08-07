package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	matrixFileName := flag.String("f", "", "matrix file name")
	word := flag.String("w", "", "word to look for")
	flag.Parse()

	fmt.Printf("reading matrix from %q\n", *matrixFileName)

	matrix := readMatrix(*matrixFileName)

	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Printf("Looking for %q in matrix\n", *word)

	if exists(matrix, []rune(*word)) {
		fmt.Printf("Found %q in matrix\n", *word)
		return
	}
	fmt.Printf("Did not find %q in matrix\n", *word)
}

// exists determines if the argument word can be
// found in board according to the rules.
func exists(board [][]rune, word []rune) bool {
	used := make(map[int]bool)
	for i, line := range board {
		for j, r := range line {
			if r == word[0] {
				pair := [2]int{i, j}
				used[coded(pair, board)] = true
				if backtrack(1, board, word[1:], i, j, used) {
					return true
				}
				used[coded(pair, board)] = false
			}
		}
	}

	return false
}

// backtrack recursively looks for word in board
// <x,y> position of last letter found
// word is what's left of the word to find,
// so backtrack tries to find word[0] in legal
// localtions on board.
// Map used is true for locations that have alread been looked at
// and found to contain a letter of the whole word.
func backtrack(depth int, board [][]rune, word []rune, x, y int, used map[int]bool) bool {
	if len(word) == 0 {
		return true
	}
	lookPairs := findPairs(board, x, y, used)
	for _, pair := range lookPairs {
		i, j := pair[0], pair[1]
		if board[i][j] == word[0] {
			used[coded(pair, board)] = true
			if backtrack(depth+1, board, word[1:], i, j, used) {
				return true
			}
			used[coded(pair, board)] = false
		}
	}

	return false
}

// readMatrix reads in a 2-D slice of runes from a newline-terminated,
// comma-separated text file
func readMatrix(fileName string) [][]rune {
	buffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var matrix [][]rune

	lines := bytes.Split(buffer, []byte{'\n'})
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var runes []rune
		fields := bytes.Split(line, []byte{','})
		for _, field := range fields {
			runes = append(runes, rune(field[0]))
		}
		matrix = append(matrix, runes)
	}
	return matrix
}

// findPairs locates usable <x,y> slots in the matrix,
// not off the matrix left, right, top or bottom,
// and not previously used. There can be a max of 4:
//       y-1  y  y+1
//  x-1       1
//  x     2   ?   3
//  x+1       4
//
//  "?" is the location of the last letter found, <x,y>
// takes advantage of Go's copy semantics
// findPairs exists so that func backtrack doesn't have to
// do a ton of checks on potential locations. All checks
// centralized here.
func findPairs(board [][]rune, x, y int, used map[int]bool) [][2]int {
	var pairs [][2]int
	var pair [2]int

	// x-1, y
	pair[0], pair[1] = x-1, y
	if x-1 >= 0 && !used[coded(pair, board)] {
		pairs = append(pairs, pair)
	}

	// x, y-1
	pair[0], pair[1] = x, y-1
	if y-1 >= 0 && !used[coded(pair, board)] {
		pairs = append(pairs, pair)
	}

	// x, y+1
	pair[0], pair[1] = x, y+1
	if y+1 < len(board[x]) && !used[coded(pair, board)] {
		pairs = append(pairs, pair)
	}

	// x+1, y
	pair[0], pair[1] = x+1, y
	if x+1 < len(board) && !used[coded(pair, board)] {
		pairs = append(pairs, pair)
	}

	return pairs
}

// coded returns a unique int for each <x,y> position
// on the board. Essentially a base N integer, where
// N is len(board[i])
func coded(pair [2]int, board [][]rune) int {
	f := len(board[pair[0]])
	return pair[0]*len(board)*f + pair[1]
}
