package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func readInputContent(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseRaceData(grid []string) ([]Race, error) {
	var races []Race

	for _, line := range grid {
		parts := strings.Fields(line)

		for i := 1; i < len(parts); i++ {
			value, err := strconv.Atoi(parts[i])
			if err != nil {
				return nil, err
			}

			if len(races) < i {
				races = append(races, Race{})
			}

			if strings.HasPrefix(line, "Time:") {
				races[i-1].Time = value
			} else if strings.HasPrefix(line, "Distance:") {
				races[i-1].Distance = value
			}
		}
	}

	return races, nil
}

func getWinningTime(race Race) int {
	maxTime := race.Time
	var recordValues []int
	var distance int

	for buttonTime := 0; buttonTime <= maxTime; buttonTime++ {
		speed := buttonTime

		for j := 0; j < maxTime-buttonTime; j++ {
			distance += speed
		}

		if distance > race.Distance {
			recordValues = append(recordValues, buttonTime)
		}

		// Calculer et afficher le pourcentage d'avancement
		percentage := float64(buttonTime) / float64(maxTime) * 100
		fmt.Printf("Progress: %.2f%%\n", percentage)

		distance = 0
	}

	return len(recordValues)
}

func main() {
	grid, _ := readInputContent("input.txt")
	// grid := []string{
	// 	"Time:      7  15   30",
	// 	"Distance:  9  40  200",
	// }

	races, _ := parseRaceData(grid)
	var recordValues int

	for _, race := range races {
		number := getWinningTime(race)
		if recordValues == 0 {
			recordValues = number
			continue
		}

		if recordValues != 0 {
			recordValues *= number
			continue
		}
	}

	println(recordValues)
}
