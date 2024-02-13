package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	println("Part1")
	total := ComputePart1(os.Stdin)
	println(total)
	println("Part2")
	total = ComputePart2(os.Stdin)
	println(total)
}

func ComputePart1(r io.ReadSeeker) int {
	r.Seek(0, 0)
	scanner := bufio.NewScanner(r)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers, scratchedNumbers := ReadCard(line)
		points := int(GetPoints(winningNumbers, scratchedNumbers))
		if points > 0 {
			// The first match makes the card worth one point and each match after the first doubles the point value of that card.
			total += int(math.Pow(2, float64(max(points-1, 0))))
		}
	}
	return total
}

func ComputePart2(r io.ReadSeeker) int {
	copies := make(map[int]int)
	r.Seek(0, 0)
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers, scratchedNumbers := ReadCard(line)
		points := int(GetPoints(winningNumbers, scratchedNumbers))
		//fmt.Printf("points %v\n", points)
		copies[i] = copies[i] + 1 // add the current original card to the list of copies
		//fmt.Printf("copies %v\n", copies[i])
		for j := i + 1; j < i+1+points; j++ {
			// win copies of the scratchcards below the winning card equal to the number of matches (= points)
			// i.e. if we have two points, we win x copies of the next two cards
			// where x = the number of copies we have of the current card
			// i.e. if we have 4 copies of the current card, and we have two points, we win 4 copies of the next two cards
			copies[j] = copies[j] + copies[i]
		}
		i++
	}
	sum := 0
	for _, v := range copies {
		//fmt.Printf("Card %v has %v copies\n", idx, v)
		sum += v
	}
	return sum
}

// ReadCard reads string representing a card and returns the winning numbers and scratched numbers.
// ex: Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// and returns an array of winning numbers (before the "|") and an array of scratched numbers (after the "|").
func ReadCard(data string) ([]int, []int) {
	s := strings.Split(data, ":")
	s = strings.Split(s[1], "|")
	winningNumbersStr := strings.Fields(s[0])
	scratchedNumbersStr := strings.Fields(s[1])

	return stringsSliceToIntSlice(winningNumbersStr), stringsSliceToIntSlice(scratchedNumbersStr)
}

// stringsSliceToIntSlice converts a slice of strings to a slice of ints.
func stringsSliceToIntSlice(s []string) []int {
	ints := make([]int, len(s))
	for i, s := range s {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

// GetPoints returns the number of points based on how many of the scratched numbers are in the winning numbers.
func GetPoints(winningNumbers []int, scratchedNumbers []int) int {
	points := 0
	for _, number := range scratchedNumbers {
		if slices.Contains(winningNumbers, number) {
			points++
		}
	}
	return points
}
