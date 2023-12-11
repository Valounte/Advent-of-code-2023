package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type galaxy struct {
	number int
	x      int
	y      int
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

func findEmptyRows(grid []string) map[int]bool {
	emptyRows := make(map[int]bool)
	for i, row := range grid {
		if strings.Trim(row, ".") == "" {
			emptyRows[i] = true
		}
	}
	return emptyRows
}

func findEmptyCols(grid []string) map[int]bool {
	emptyCols := make(map[int]bool)
	for i := 0; i < len(grid[0]); i++ {
		colEmpty := true
		for j := 0; j < len(grid); j++ {
			if grid[j][i] != '.' {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyCols[i] = true
		}
	}
	return emptyCols
}

func adjustCoordinates(x, y int, emptyRows, emptyCols map[int]bool) (int, int) {
	adjustedX, adjustedY := x, y
	for k := 0; k < y; k++ {
		if emptyRows[k] {
			adjustedY += 999999
		}
	}
	for k := 0; k < x; k++ {
		if emptyCols[k] {
			adjustedX += 999999
		}
	}
	return adjustedX, adjustedY
}

func expandUniverse(grid []string) []galaxy {
	emptyRows := findEmptyRows(grid)
	emptyCols := findEmptyCols(grid)
	galaxies := []galaxy{}
	hashCount := 0

	for y, row := range grid {
		for x, ch := range row {
			if ch == '#' {
				adjustedX, adjustedY := adjustCoordinates(x, y, emptyRows, emptyCols)
				galaxies = append(galaxies, galaxy{number: hashCount, x: adjustedX, y: adjustedY})
				hashCount++
			}
		}
	}

	return galaxies
}

func calcDistance(g1, g2 galaxy) int {
	return int(math.Abs(float64(g1.x-g2.x)) + math.Abs(float64(g1.y-g2.y)))
}

func main() {
	grid := readInputContent("input.txt")

	galaxies := expandUniverse(grid)
	var total int
	for i, g := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			distance := calcDistance(g, g2)
			total += distance
		}
	}

	fmt.Println("Total:", total)
}
