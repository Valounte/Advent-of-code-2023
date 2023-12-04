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

func indexOf(word string, numWords []string) int {
	for i, w := range numWords {
		if w == word {
			return i
		}
	}
	return -1
}

func formatCalibration(line string) []string {
	numDict := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for word, replacement := range numDict {
		line = strings.ReplaceAll(line, word, replacement)
	}

	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(line, -1)

	if len(matches) > 1 {
		return []string{matches[0], matches[len(matches)-1]}
	} else if len(matches) == 1 {
		return []string{matches[0], matches[0]}
	}
	return []string{}
}

func calculateCalibration(calibration []string) int {
	var firstIntValue int
	var lastIntValue int

	for i, char := range calibration {
		var num int
		var err error

		if num, err = strconv.Atoi(char); err != nil {
			switch char {
			case "zero":
				num = 0
			case "one":
				num = 1
			case "two":
				num = 2
			case "three":
				num = 3
			case "four":
				num = 4
			case "five":
				num = 5
			case "six":
				num = 6
			case "seven":
				num = 7
			case "eight":
				num = 8
			case "nine":
				num = 9
			}
		}

		if i == 0 {
			firstIntValue = num
		} else {
			lastIntValue = num
		}
	}

	concatenated := strconv.Itoa(firstIntValue) + strconv.Itoa(lastIntValue)
	result, _ := strconv.Atoi(concatenated)
	return result
}

func main() {
	lines, err := readInputContent("input.txt")
	if err != nil {
		fmt.Println("Erreur à la lecture du fichier:", err)
		return
	}

	var formattedCalibration []string
	var calibrationValue int

	for _, line := range lines {
		formattedCalibration = formatCalibration(line)
		calibrationValue += calculateCalibration(formattedCalibration)
	}
	println("Calibration value:", calibrationValue)
}
