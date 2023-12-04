package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	winningCards []int
	myCards      []int
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

func (g *Game) processCard(line string) error {
	re := regexp.MustCompile(`\d+`)
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return fmt.Errorf("invalid line format: %s", line)
	}
	winningNumbers := re.FindAllString(parts[0], -1)
	for _, numStr := range winningNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}
		g.winningCards = append(g.winningCards, num)
	}

	myNumbers := re.FindAllString(parts[1], -1)
	for _, numStr := range myNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}
		g.myCards = append(g.myCards, num)
	}

	return nil
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}

func (g *Game) checkWinning() float64 {
	score := 0.5
	for i := 1; i < len(g.winningCards); i++ {
		num := g.winningCards[i]
		if !contains(g.myCards, num) {
			continue
		}
		score *= 2
	}
	if score == 0.5 {
		return 0
	}

	return score
}

func main() {
	grid, _ := readInputContent("input.txt")
	// grid := []string{
	// 	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	// 	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	// 	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	// 	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	// 	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	// 	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	// }

	var total int

	for _, line := range grid {
		g := Game{}
		err := g.processCard(line)
		if err != nil {
			fmt.Println(err)
		}
		total += int(g.checkWinning())
	}
	fmt.Println(total)
}
