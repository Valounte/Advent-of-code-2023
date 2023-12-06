package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	Time     int
	Distance int
}

func readInputContent(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseRaceData(grid []string) Race {
	var race Race
	var timeStr, distanceStr string

	for _, line := range grid {
		parts := strings.Fields(line)
		for i := 1; i < len(parts); i++ {
			if strings.HasPrefix(line, "Time:") {
				timeStr += parts[i]
			} else if strings.HasPrefix(line, "Distance:") {
				distanceStr += parts[i]
			}
		}
	}

	var err error
	race.Time, err = strconv.Atoi(timeStr)
	if err != nil {
		return Race{}
	}
	race.Distance, err = strconv.Atoi(distanceStr)
	if err != nil {
		return Race{}
	}

	return race
}

func getWinningTime(race Race) int {
	maxTime := race.Time
	if maxTime <= 0 {
		return 0
	}

	recordCount := 0

	for buttonTime := 0; buttonTime <= maxTime; buttonTime++ {
		distance := buttonTime * (maxTime - buttonTime)

		if distance > race.Distance {
			recordCount++
		}
	}

	return recordCount
}

func main() {
	grid := readInputContent("input.txt")
	race := parseRaceData(grid)
	startTime := time.Now()
	recordValues := getWinningTime(race)
	println(recordValues)
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	executionTimeInMinutes := executionTime.Seconds()

	fmt.Printf("Execution time: %.2f seconds\n", executionTimeInMinutes)
}
