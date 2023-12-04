package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getCalibrationValue(line string) int {
	var firstIntValue *int
	var lastIntValue *int

	for _, runeValue := range line {
		char := string(runeValue)
		if num, err := strconv.Atoi(char); err == nil {
			if firstIntValue == nil {
				firstIntValue = new(int)
				*firstIntValue = num
			} else {
				lastIntValue = new(int)
				*lastIntValue = num
			}
		}
	}

	if lastIntValue == nil {
		lastIntValue = firstIntValue
	}

	concatenated := strconv.Itoa(*firstIntValue) + strconv.Itoa(*lastIntValue)
	result, _ := strconv.Atoi(concatenated)
	return result
}

func main() {
	lines, err := readInputContent("input.txt")
	if err != nil {
		fmt.Println("Erreur à la lecture du fichier:", err)
		return
	}

	var calibrationValue int

	for _, line := range lines {
		calibrationValue += getCalibrationValue(line)
	}

	fmt.Println("Calibration value:", calibrationValue)
}
