package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type CubeSet struct {
	Red, Blue, Green int
}

func main() {
	println("Part1")
	r := os.Stdin
	sum := ComputePart1(r)
	println("Part1 sum is", sum)
	r.Seek(0, 0)
	sum = ComputePart2(r)
	println("Part2 sum is", sum)
}

func ComputePart1(r io.Reader) int {
	// Array of CubeSet for each Game ID
	var games = parseGames(r)

	sum := 0
	for gameId, cubeSet := range games {
		if isPossible(cubeSet) {
			sum += gameId
		}
	}
	return sum
}

func ComputePart2(r io.Reader) int {
	// Array of CubeSet for each Game ID
	var games = parseGames(r)
	power := 0
	for _, cubeSets := range games {
		blue, red, green := 0, 0, 0
		for _, cubeSet := range cubeSets {

			if cubeSet.Blue > blue {
				blue = cubeSet.Blue
			}
			if cubeSet.Red > red {
				red = cubeSet.Red
			}
			if cubeSet.Green > green {
				green = cubeSet.Green
			}
		}
		power += blue * red * green
	}
	return power
}

// parseGames returns a map of game IDs to CubeSets based on the io.Reader
func parseGames(r io.Reader) map[int][]CubeSet {
	scanner := bufio.NewScanner(r)

	games := make(map[int][]CubeSet)
	for scanner.Scan() {
		line := scanner.Text()

		gameId, cubeSet := parseLine(line)
		games[gameId] = cubeSet
	}
	return games
}

// parseLine parses a line of input in the form of
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// and returns the Game ID and an array of CubeSet
func parseLine(line string) (int, []CubeSet) {
	// Split the line into Game ID and CubeSets
	split := strings.Split(line, ":")
	var gameId int
	fmt.Sscanf(split[0], "Game %d", &gameId)

	var cubes = parseSets(split[1])

	return gameId, cubes
}

// parseSets parses a string of CubeSets in the form of
// 1 green, 3 red, 6 blue
// and returns an array of CubeSet
func parseSets(line string) []CubeSet {
	var cubes = []CubeSet{}
	for _, cubeSetStr := range strings.Split(line, ";") {
		sets := strings.Split(cubeSetStr, ", ")
		cubSet := CubeSet{}

		for _, set := range sets {
			// ex: 1 green
			set = strings.Trim(set, " ")
			count := 0
			color := ""
			fmt.Sscanf(set, "%d %s", &count, &color)
			switch color {
			case "red":
				cubSet.Red = count
			case "blue":
				cubSet.Blue = count
			case "green":
				cubSet.Green = count
			default:
				println("Unknown color")
			}
		}
		cubes = append(cubes, cubSet)
	}
	return cubes
}

func isPossible(sets []CubeSet) bool {
	for _, set := range sets {
		if set.Red > 12 || set.Green > 13 || set.Blue > 14 {
			return false
		}
	}
	return true
}
