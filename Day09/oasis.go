package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func parseGrid(grid []string) [][]int {
	var result [][]int

	for _, line := range grid {
		var nums []int
		strNumbers := strings.Split(line, " ")
		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Printf("Error converting '%s' to int: %s\n", strNum, err)
				continue
			}
			nums = append(nums, num)
		}
		result = append(result, nums)
	}

	return result
}

func getDifferenceValues(line []int) []int {
	var newLine []int
	for i := 1; i < len(line); i++ {
		newLine = append(newLine, line[i]-line[i-1])
	}

	return newLine
}

func processLineRecursively(line []int, result *[][]int) {
	*result = append(*result, line)

	if isAllZeros(line) {
		return
	}

	newLine := getDifferenceValues(line)
	processLineRecursively(newLine, result)
}

func isAllZeros(line []int) bool {
	for _, value := range line {
		if value != 0 {
			return false
		}
	}
	return true
}

func calculateValue(lines [][]int) int {
	var result int
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if i == len(lines)-1 {
			continue
		}
		result = result + line[len(line)-1]
	}

	return result
}

func main() {
	grid := readInputContent("input.txt")
	// grid := []string{
	// 	"0 3 6 9 12 15",
	// 	"1 3 6 10 15 21",
	// 	"10 13 16 21 30 45",
	// }
	parsedGrid := parseGrid(grid)
	fmt.Println(parsedGrid)
	var count int
	for _, line := range parsedGrid {
		var processedLines [][]int
		processLineRecursively(line, &processedLines)
		count += calculateValue(processedLines)
	}

	fmt.Println("Count:", count)
}
