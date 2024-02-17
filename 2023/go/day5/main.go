package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	println("Part1")
	result := ComputePart1(os.Stdin)
	println("Result:", result)
	println("Part2")
	result = ComputePart2(os.Stdin)
	println("Result:", result)
}

type mapping struct {
	destStart int
	destEnd   int
	srcStart  int
	srcEnd    int
}

type seedRange struct {
	start int
	end   int
}

func ComputePart1(r io.Reader) int {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		fmt.Println("Cannot read input data", err)
		os.Exit(1)
	}

	input := buf.String()
	seeds, conversionMaps := parseInput(input)

	// apply the conversion maps to each seed
	var results []int
	for _, seed := range seeds {
		for _, conversionMap := range conversionMaps {
			for _, mapping := range conversionMap {
				if seed >= mapping.srcStart && seed <= mapping.srcEnd {
					seed = mapping.destStart + seed - mapping.srcStart
					break
				}
			}
		}
		results = append(results, seed)
	}
	return slices.Min(results)
}

func ComputePart2(r io.ReadSeeker) int {
	r.Seek(0, 0)
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		fmt.Println("Cannot read input data", err)
		os.Exit(1)
	}

	input := buf.String()
	seeds, conversionMaps := parseInput(input)

	// seeds come now in pair
	// create a slice of seedRange
	var seedRanges []seedRange

	currentRange := seedRange{}
	for idx, seed := range seeds {
		if idx%2 == 0 {
			currentRange.start = seed
		} else {
			currentRange.end = currentRange.start + seed - 1
			seedRanges = append(seedRanges, currentRange)
			currentRange = seedRange{}
		}
	}
	// sort the seed ranges
	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})

	newRanges := seedRanges
	for idx, mappings := range conversionMaps {
		fmt.Println("Conversion map ", idx)
		var tempRanges []seedRange
		for _, r := range newRanges {
			mappedRanges := getMappedRanges(r, &mappings)
			tempRanges = append(tempRanges, mappedRanges...)
		}
		newRanges = tempRanges
	}

	//fmt.Println(newRanges)
	// order the results
	sort.Slice(newRanges, func(i, j int) bool {
		return newRanges[i].start < newRanges[j].start
	})

	// No idea why, but the first range I got was 0 to 1493865, which was not the correct answer.
	// The correct answer was the start of the second range, 1493866 (range 1493866 to 10954308)
	fmt.Println("Final ranges", newRanges[:int(math.Min(float64(len(newRanges)), 10))])
	return newRanges[0].start
}

// getMappedRanges returns a slice of seedRange applied to the mappings.
// the input seedRange might get split to fit a mapping
func getMappedRanges(r seedRange, mappings *[]mapping) []seedRange {
	var newRanges []seedRange

	// Iterate over the mappings, which are sorted
Loop:
	for _, mapping := range *mappings {
		offset := mapping.destStart - mapping.srcStart
		switch {
		case mapping.srcEnd < r.start:
			// mapping ends before the range
			continue Loop

		case mapping.srcStart > r.end:
			// mapping is after range, we can break (we iterate over sorted mappings, next mappings will be after as well)
			// 			  |-------| 	mapping
			// ^------^					range
			break Loop

		case mapping.srcStart <= r.start && mapping.srcEnd >= r.end:
			// range fully included in mapping, apply the offset
			// |------------|	mapping
			//   ^--------^		range
			newRanges = append(newRanges, seedRange{start: r.start + offset, end: r.end + offset})
			break Loop

		case mapping.srcStart <= r.start:
			// range is half (left half) in a mapping
			// |----------|			mapping
			//      ^----------^	range
			// cut the range in two, apply the mapping on the included half
			newRanges = append(newRanges, seedRange{start: r.start + offset, end: mapping.destEnd})
			if r.end > mapping.srcEnd {
				// and recursively apply the method on the right half
				rightRange := seedRange{start: mapping.srcEnd + 1, end: r.end}
				newRanges = append(newRanges, getMappedRanges(rightRange, mappings)...)
			}
			break Loop

		case mapping.srcStart >= r.start:
			// range is half (right half) in a mapping
			// 		|-----------|	mapping
			// ^---------^			range
			// cut the range at least in two, keep the left-most part as-is (no mapping applied, as we iterate over sorted mappings)
			newRanges = append(newRanges, seedRange{start: r.start, end: mapping.srcStart - 1})

			// middle part, apply the mapping
			newRanges = append(newRanges, seedRange{start: mapping.destStart, end: mapping.srcEnd + offset})

			if mapping.srcEnd < r.end {
				// range is too big, cut the right part and recursively apply the method on it
				rightRange := seedRange{start: mapping.srcEnd + 1, end: r.end}
				newRanges = append(newRanges, getMappedRanges(rightRange, mappings)...)
			}
			break Loop
		}
	}
	if len(newRanges) == 0 {
		newRanges = append(newRanges, r)
	}
	return newRanges
}

// parseInput parses a string and returns a slice of ints representing the seeds,
// and a slice of slices of mappings, representing each conversion map.
func parseInput(input string) ([]int, [][]mapping) {
	// split the seeds (first line) and the rest of the data
	s := strings.SplitN(input, "\n", 2)
	seedLine, remaining := s[0], s[1]
	seedsStr := strings.TrimSpace(strings.Split(seedLine, ":")[1])
	seeds := stringSliceToIntSlice(strings.Split(seedsStr, " "))

	// consume empty lines
	remaining = strings.TrimSpace(remaining)

	// create the conversion maps
	var conversionMaps [][]mapping
	for _, data := range strings.Split(remaining, "\n\n") {
		// remove the map name
		data := strings.SplitN(data, "\n", 2)[1]
		var mappings []mapping
		for _, line := range strings.Split(data, "\n") {
			n := strings.Split(string(line), " ")
			destStart := atoiOrPanic(n[0])
			srcStart := atoiOrPanic(n[1])
			len := atoiOrPanic(n[2])
			mappings = append(mappings, mapping{
				destStart: destStart,
				destEnd:   destStart + len - 1,
				srcStart:  srcStart,
				srcEnd:    srcStart + len - 1,
			})
		}
		// sort the mappings
		sort.Slice(mappings, func(i, j int) bool {
			return mappings[i].srcStart < mappings[j].srcStart
		})
		conversionMaps = append(conversionMaps, mappings)
	}
	return seeds, conversionMaps
}

func stringSliceToIntSlice(s []string) []int {
	var r []int
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Cannot convert seed to int", v)
			os.Exit(1)
		}
		r = append(r, n)
	}
	return r
}

func atoiOrPanic(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Cannot convert string to int", s)
		os.Exit(1)
	}
	return n
}
