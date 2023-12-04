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

func isSymbol(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}

func isSymbolAdjacent(grid []string, y, x int) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			newY, newX := y+dy, x+dx
			if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[newY]) && isSymbol(rune(grid[newY][newX])) {
				return true
			}
		}
	}
	return false
}

func sumPartNumbers(grid []string) int {
	sum := 0
	for y, line := range grid {
		for x := 0; x < len(line); x++ {
			char := line[x]
			if unicode.IsDigit(rune(char)) {
				end := x
				for end < len(line) && unicode.IsDigit(rune(line[end])) {
					end++
				}
				numberStr := line[x:end]
				number, _ := strconv.Atoi(numberStr)

				isAdjacent := false
				for i := x; i < end; i++ {
					if isSymbolAdjacent(grid, y, i) {
						isAdjacent = true
						break
					}
				}
				if isAdjacent {
					sum += number
				}
				x = end - 1
			}
		}
	}
	return sum
}

func main() {
	grid, _ := readInputContent("input.txt")
	total := sumPartNumbers(grid)
	fmt.Println("Total sum of part numbers:", total)
}
