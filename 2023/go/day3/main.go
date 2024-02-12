package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	println("part1")
	total := ComputePart1(os.Stdin)
	fmt.Println(total)
	println("part2")
	total = ComputePart2(os.Stdin)
	fmt.Println(total)
}

func ComputePart1(r io.ReadSeeker) int {
	matrix := getMatrix(r)
	if len(matrix) == 0 {
		return 0
	}
	numLines := len(matrix)
	numCols := len(matrix[0])

	// Iterate over the matrix to find numbers,
	// and look around to check if there is a symbol.
	// If there is, add the number to the total.
	total := 0
	for i := 0; i < numLines; i++ {
		numberStart := -1 // index of the start of a numver
		numberBuf := ""   // buffer for the number
		for j := 0; j < numCols; j++ {
			c := matrix[i][j]
			if IsNum(c) {
				if numberStart == -1 {
					numberStart = j
				}
				numberBuf += c
			}
			if !IsNum(c) || j == numCols-1 {
				if numberStart != -1 {
					//fmt.Printf("Found number %v\n", numberBuf)
					//fmt.Printf("Looking around %v-%v\n", numberStart, max(j-1, 0))
					if HasSymbolAround(&matrix, i, numberStart, min(j-1, len(matrix[i])-1)) {
						num, err := strconv.Atoi(numberBuf)
						if err != nil {
							fmt.Printf("Cannot parse %v", numberBuf)
							os.Exit(1)
						}
						total += num
					}
				}
				numberStart = -1
				numberBuf = ""
			}
		}
	}
	return total
}

func ComputePart2(r io.ReadSeeker) int {
	// See https://github.com/hyper-neutrino/advent-of-code/blob/main/2023/day03p2.py
	// Find "*", then check around it to find numbers.
	matrix := getMatrix(r)
	if len(matrix) == 0 {
		return 0
	}
	numLines := len(matrix)
	numCols := len(matrix[0])

	// Iterate over the matrix to find the stars which are adjacent to exactly two part numbers.
	// Multiply these numbers together, and add them to the total.
	total := 0
	for i := 0; i < numLines; i++ {
		for j := 0; j < numCols; j++ {
			c := matrix[i][j]
			if c == "*" {
				total += getMultipliedPartNumbers(&matrix, i, j)
			}
		}
	}
	return total
}

// getMatrix returns a two-dimensional array (matrix) based on the reader.
func getMatrix(r io.ReadSeeker) [][]string {
	// reset the reader
	r.Seek(0, 0)
	scanner := bufio.NewScanner(r)

	// first determine the size of the 2 dimensional array (matrix)
	numLines := 0
	numCols := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > numCols {
			numCols = len(line)
		}
		numLines++
	}
	fmt.Printf("numLines %v, numCols %v\n", numLines, numCols)

	// reset the reader, and fill the matrix
	r.Seek(0, 0)
	scanner = bufio.NewScanner(r)
	matrix := make([][]string, numLines)
	i := 0
	for scanner.Scan() {
		matrix[i] = make([]string, numCols)
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			matrix[i][j] = string(line[j])
		}
		i++
	}
	return matrix
}

func IsNum(c string) bool {
	return c >= "0" && c <= "9"
}

// HasSymbolAround looks into the matrix to see if there is a symbol around the start-end section (any character except for a dot)
// returns true if there is a symbol around the section
// For instance, given the matrix :
//
// . . . . @ .
// . . 1 2 . .
// . . . . . .
//
// calling this method with i=1, start=2, end=3 (the position of "12") will return true
// because there's an @ symbol diagonally from "12".
func HasSymbolAround(matrix *[][]string, i, start, end int) bool {
	// look before the number
	if start > 0 && (*matrix)[i][start-1] != "." && !IsNum((*matrix)[i][start-1]) {
		return true
	}
	// look after the number
	if end < len((*matrix)[i])-1 && (*matrix)[i][end+1] != "." && !IsNum((*matrix)[i][end+1]) {
		return true
	}
	// look above the number
	if i > 0 {
		for j := max(start-1, 0); j <= min(end+1, len((*matrix)[i-1])-1); j++ {
			if (*matrix)[i-1][j] != "." && !IsNum((*matrix)[i-1][j]) {
				return true
			}
		}
	}
	// look below the number
	if i < len(*matrix)-1 {
		for j := max(start-1, 0); j <= min(end+1, len((*matrix)[i+1])-1); j++ {
			if (*matrix)[i+1][j] != "." && !IsNum((*matrix)[i+1][j]) {
				return true
			}
		}
	}
	return false
}

// getMultipliedPartNumbers returns the multiplied part numbers around the star,
// only if the star has exactly two numbers around it.
// Otherwise returns 0
// For instance, given the matrix :
//
// . . . . . .
// . . . 2 . .
// . . * . . .
// . . 7 . . .
//
// calling this method with i=2, j=2 (the position of "*") will return 14
// because 7 and 2 are both adjacent to the star..
func getMultipliedPartNumbers(matrix *[][]string, i, j int) int {
	var numbers []int
	// look before the star
	if j > 0 && IsNum((*matrix)[i][j-1]) {
		numbers = append(numbers, getNumAt(matrix, i, j-1, true))
	}
	// look after the star
	if j < len((*matrix)[i])-1 && IsNum((*matrix)[i][j+1]) {
		numbers = append(numbers, getNumAt(matrix, i, j+1, false))
	}
	// look above the star
	if i > 0 {
		// three positions to check
		// ex:
		// x y z .
		// . * . .
		// . . . .
		if IsNum((*matrix)[i-1][j]) {
			// top (y) spot is a number, rewind as long as there's a number or we're at column 0
			z := j - 1
			for {
				if z == 0 || !IsNum((*matrix)[i-1][z]) {
					break
				}
				z--
			}
			numbers = append(numbers, getNumAt(matrix, i-1, z+1, false))
		} else {
			if IsNum((*matrix)[i-1][j-1]) {
				numbers = append(numbers, getNumAt(matrix, i-1, j-1, true))
			}
			if j < len((*matrix)[i])-1 && IsNum((*matrix)[i-1][j+1]) {
				numbers = append(numbers, getNumAt(matrix, i-1, j+1, false))
			}
		}
	}
	// look below the star
	if i < len(*matrix)-1 {
		// three positions to check
		// ex:
		// . . . .
		// . * . .
		// x y z .
		if IsNum((*matrix)[i+1][j]) {
			// below (y) spot is a number, rewind as long as there's a number or we're at column 0
			z := j - 1
			for {
				if z == 0 || !IsNum((*matrix)[i+1][z]) {
					break
				}
				z--
			}
			numbers = append(numbers, getNumAt(matrix, i+1, z+1, false))
		} else {
			if IsNum((*matrix)[i+1][j-1]) {
				numbers = append(numbers, getNumAt(matrix, i+1, j-1, true))
			}
			if j < len((*matrix)[i])-1 && IsNum((*matrix)[i+1][j+1]) {
				numbers = append(numbers, getNumAt(matrix, i+1, j+1, false))
			}
		}
	}
	fmt.Printf("Numbers %v\n", numbers)
	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}
	return 0
}

func getNumAt(matrix *[][]string, i, j int, reverse bool) int {
	buf := ""
	for {
		if j < 0 || j >= len((*matrix)[i]) || !IsNum((*matrix)[i][j]) {
			break
		}
		if reverse {
			buf = (*matrix)[i][j] + buf
			j--
		} else {
			buf += (*matrix)[i][j]
			j++
		}
	}
	number, err := strconv.Atoi(buf)
	if err != nil {
		fmt.Printf("Could not parse number %v\n", buf)
		os.Exit(1)
	}
	return number
}
