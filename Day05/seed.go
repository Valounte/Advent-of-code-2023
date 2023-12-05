package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	Destination int
	Source      int
	Range       int
}

type Map struct {
	ID    int
	Lines []Line
}

func readInputContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return builder.String(), nil
}

func mapGrid(grid string) ([]int, []Map) {
	var seeds []int
	var maps []Map
	lines := strings.Split(grid, "\n")

	currentMapID := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "seeds:") {
			seedsString := strings.Split(line, ":")[1]
			seeds = extractNumbers(seedsString)
		} else if strings.Contains(line, "map:") {
			currentMapID++
			maps = append(maps, Map{ID: currentMapID})
		} else if currentMapID > 0 && line != "" {
			numbers := extractNumbers(line)
			if len(numbers) == 3 {
				maps[currentMapID-1].Lines = append(maps[currentMapID-1].Lines, Line{
					Destination: numbers[0],
					Source:      numbers[1],
					Range:       numbers[2],
				})
			}
		}
	}
	return seeds, maps
}

func extractNumbers(s string) []int {
	var numbers []int
	for _, field := range strings.Fields(s) {
		if number, err := strconv.Atoi(field); err == nil {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func calculateValue(seed int, maps []Map) int {
	for _, m := range maps {
		for _, line := range m.Lines {
			if seed >= line.Source && seed <= line.Source+line.Range {
				seed = convertSourceToDestination(seed, line.Source, line.Destination, line.Range)
				break
			}
		}
	}
	return seed
}

func convertSourceToDestination(seed int, sourceStart int, destinationStart int, rangeLength int) int {
	if seed >= sourceStart && seed < sourceStart+rangeLength {
		difference := seed - sourceStart
		return destinationStart + difference
	}
	return -1
}

func addToSeedsArray(firstSeed int, seedRange int) []int {
	var seeds []int
	for i := firstSeed; i < firstSeed+seedRange; i++ {
		seeds = append(seeds, i)
	}
	return seeds
}

func main() {
	grid, _ := readInputContent("input.txt")
	seeds, maps := mapGrid(grid)

	var newSeeds []int
	for i := 0; i < len(seeds); i += 2 {
		if i+1 < len(seeds) {
			newSeeds = append(newSeeds, addToSeedsArray(seeds[i], seeds[i+1])...)
		}
	}
	seeds = append(seeds, newSeeds...)

	var nearerValue int

	for _, newSeed := range newSeeds {
		value := calculateValue(newSeed, maps)
		if (nearerValue == 0 || value < nearerValue) && value > 0 {
			nearerValue = value
		}
	}

	println(nearerValue)
}
