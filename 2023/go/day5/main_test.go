package main

import (
	"os"
	"strings"
	"testing"
)

func TestComputePart1(t *testing.T) {
	const expected = 35
	const input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
	actual := ComputePart1(strings.NewReader(input))
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func TestComputePart2(t *testing.T) {
	const expected = 46
	const input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
	actual := ComputePart2(strings.NewReader(input))
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func TestComputePart2Mapping1(t *testing.T) {
	const expected = 2
	const input = `seeds: 0 10

seed-to-soil map:
15 0 2`
	actual := ComputePart2(strings.NewReader(input))
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func TestComputePart2Mapping2(t *testing.T) {
	const expected = 0
	const input = `seeds: 10 20

seed-to-soil map:
0 10 2`
	actual := ComputePart2(strings.NewReader(input))
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func TestComputePart2Mapping3(t *testing.T) {
	const expected = 0
	const input = `seeds: 10 20

seed-to-soil map:
0 12 2`
	actual := ComputePart2(strings.NewReader(input))
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func TestComputePart2Mapping4(t *testing.T) {
	const expected = 0
	const input = `seeds: 10 20

seed-to-soil map:
0 19 2`
	actual := ComputePart2(strings.NewReader(input))
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func TestComputePart2Mapping5(t *testing.T) {
	const expected = 2
	const input = `seeds: 10 20

seed-to-soil map:
0 8 30`
	actual := ComputePart2(strings.NewReader(input))
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func TestRealComputePart2(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening file")
	}
	result := ComputePart2(f)
	const expected = 1493866
	if result != expected {
		t.Errorf("actual %v != expected %v", result, expected)
	}
}
