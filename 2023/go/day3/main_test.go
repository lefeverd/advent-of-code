package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	const expected = 4361
	const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	actual := ComputePart1(strings.NewReader(input))
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestPart1EndLine(t *testing.T) {
	const expected = 4362
	const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755*
...$.*...1
.664.598..`
	actual := ComputePart1(strings.NewReader(input))
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	const expected = 467835
	const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	actual := ComputePart2(strings.NewReader(input))
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestGetNumAt(t *testing.T) {
	const expected = 123
	var input = [][]string{{".", ".", "1", "2", "3", "."}}
	actual := getNumAt(&input, 0, 2, false)
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetNumAtReverse(t *testing.T) {
	const expected = 123
	var input = [][]string{{".", ".", "1", "2", "3", "."}}
	actual := getNumAt(&input, 0, 4, true)
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
