package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

type GameData struct {
	Game    int
	Tirages map[int]map[string]int
}

func parseGameData(input string) GameData {
	var gameData GameData
	gameData.Tirages = make(map[int]map[string]int)

	reGameNumber := regexp.MustCompile(`Game (\d+):`)
	gameNumberMatch := reGameNumber.FindStringSubmatch(input)
	if len(gameNumberMatch) > 1 {
		gameData.Game, _ = strconv.Atoi(gameNumberMatch[1])
	}

	input = reGameNumber.ReplaceAllString(input, "")
	tirages := strings.Split(input, ";")

	for i, tirage := range tirages {
		tirage = strings.TrimSpace(tirage)
		if tirage == "" {
			continue
		}

		gameData.Tirages[i] = make(map[string]int)
		details := strings.Split(tirage, ",")
		for _, detail := range details {
			detail = strings.TrimSpace(detail)
			reDetails := regexp.MustCompile(`(\d+) (\w+)`)
			detailsMatch := reDetails.FindStringSubmatch(detail)
			if len(detailsMatch) > 2 {
				count, _ := strconv.Atoi(detailsMatch[1])
				color := detailsMatch[2]
				gameData.Tirages[i][color] = count
			}
		}
	}

	return gameData
}

func getPower(game GameData) int {
	var totalBlue int
	var totalRed int
	var totalGreen int

	var biggestBlue int
	var biggestRed int
	var biggestGreen int

	var actualBlue int
	var actualRed int
	var actualGreen int

	for _, tirage := range game.Tirages {
		totalBlue += tirage["blue"]
		totalRed += tirage["red"]
		totalGreen += tirage["green"]

		if tirage["blue"] > biggestBlue {
			biggestBlue = tirage["blue"]
		}
		if tirage["red"] > biggestRed {
			biggestRed = tirage["red"]
		}
		if tirage["green"] > biggestGreen {
			biggestGreen = tirage["green"]
		}
	}

	if totalBlue < biggestBlue {
		actualBlue = totalBlue
	} else {
		actualBlue = biggestBlue
	}

	if totalRed < biggestRed {
		actualRed = totalRed
	} else {
		actualRed = biggestRed
	}

	if totalGreen < biggestGreen {
		actualGreen = totalGreen
	} else {
		actualGreen = biggestGreen
	}

	return actualBlue * actualRed * actualGreen
}

func main() {
	lines, err := readInputContent("input.txt")
	if err != nil {
		fmt.Println("Erreur à la lecture du fichier:", err)
		return
	}

	var total int

	for _, line := range lines {
		gameData := parseGameData(line)
		total += getPower(gameData)
	}

	fmt.Println(total)
}
