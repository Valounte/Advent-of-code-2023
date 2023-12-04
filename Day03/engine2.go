package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func sumPartNumbers(grid []string) int {
	sum := 0
	for y, line := range grid {
		for x := 0; x < len(line); x++ {
			char := line[x]
			if char == '*' {
				var values []int
				checkedPositions := make(map[int]bool)
				checkedPositionsDown := make(map[int]bool)
				//check left
				if x > 0 && unicode.IsDigit(rune(line[x-1])) {
					end := x - 1
					for end >= 0 && unicode.IsDigit(rune(line[end])) {
						end--
					}
					numberStr := line[end+1 : x]
					number, _ := strconv.Atoi(numberStr)
					values = append(values, number)
				}

				//check right
				if x < len(line)-1 && unicode.IsDigit(rune(line[x+1])) {
					end := x + 1
					for end < len(line) && unicode.IsDigit(rune(line[end])) {
						end++
					}
					numberStr := line[x+1 : end]
					number, _ := strconv.Atoi(numberStr)
					values = append(values, number)
				}

				//check up
				if y > 0 {
					for deltaX := -1; deltaX <= 1; deltaX++ {
						newX := x + deltaX
						if newX >= 0 && newX < len(grid[y-1]) && unicode.IsDigit(rune(grid[y-1][newX])) {
							if !checkedPositions[newX] {
								numberStr, endX := extractNumber(grid, y-1, newX)
								if number, err := strconv.Atoi(numberStr); err == nil {
									values = append(values, number)
								}
								for pos := newX; pos < endX; pos++ {
									checkedPositions[pos] = true
								}
							}
						}
					}
				}

				//check down
				if y < len(grid)-1 {
					for deltaX := -1; deltaX <= 1; deltaX++ {
						newX := x + deltaX
						if newX >= 0 && newX < len(grid[y+1]) && unicode.IsDigit(rune(grid[y+1][newX])) {
							if !checkedPositionsDown[newX] {
								numberStr, endX := extractNumber(grid, y+1, newX)
								if number, err := strconv.Atoi(numberStr); err == nil {
									values = append(values, number)
								}
								for pos := newX; pos < endX; pos++ {
									checkedPositionsDown[pos] = true
								}
							}
						}
					}
				}

				if len(values) == 2 {
					fmt.Println("Found", values[0], "*", values[1])
					sum += values[0] * values[1]
				}
			}
		}
	}

	return sum
}

func extractNumber(grid []string, y, x int) (string, int) {
	start, end := x, x
	for start >= 0 && unicode.IsDigit(rune(grid[y][start])) {
		start--
	}
	start++
	for end < len(grid[y]) && unicode.IsDigit(rune(grid[y][end])) {
		end++
	}
	return grid[y][start:end], end
}

func main() {
	grid, _ := readInputContent("input.txt")
	total := sumPartNumbers(grid)
	fmt.Println("Total sum of part numbers:", total)
}
