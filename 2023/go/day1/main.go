package main

import (
	"fmt"
	"os"
	"strings"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	fmt.Println("Day1")
	result, err := ComputePart1("part1_input.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Result for part 1 is %d\n", result)
	result, err = ComputePart2("part2_input.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Result for part 2 is %d\n", result)
}

func ComputePart1(input_file string) (int, error) {
	var content, err = os.ReadFile(input_file)
	if err != nil {
		return 0, fmt.Errorf("Could not read file %s", input_file)
	}
	var input = string(content)
	var lines = strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		firstNum := -1
		lastNum := 0
		for _, c := range line {
			if IsNum(c) {
				if firstNum == -1 {
					firstNum = NumVal(c)
				}
				lastNum = NumVal(c)
			}
		}
		if firstNum != -1 {
			total += (firstNum * 10) + lastNum
		}
	}
	return total, nil
}

func ComputePart2(input_file string) (int, error) {
	var content, err = os.ReadFile(input_file)
	if err != nil {
		return 0, fmt.Errorf("Could not read file %s", input_file)
	}

	var input = string(content)
	var lines = strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		firstNum := -1
		lastNum := 0
		for i, c := range line {
			if IsNum(c) {
				if firstNum == -1 {
					firstNum = NumVal(c)
				}
				lastNum = NumVal(c)
			}
			v := startsWithNumber(line[i:])
			if v != -1 {
				if firstNum == -1 {
					firstNum = v
				}
				lastNum = v
			}
		}

		if firstNum != -1 {
			total += (firstNum * 10) + lastNum
		}
	}
	return total, nil
}

func IsNum(c rune) bool {
	return c >= '0' && c <= '9'
}

func NumVal(c rune) int {
	return int(c - '0')
}

// startsWithNumber returns the number if the string starts with a number (in its written form), otherwise -1
func startsWithNumber(s string) int {
	for k, v := range numbers {
		if strings.HasPrefix(s, k) {
			return v
		}
	}
	return -1
}
