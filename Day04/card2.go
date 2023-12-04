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
	id           int
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
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid line format: %s", line)
	}

	cardID, err := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "Card")[1]))
	if err != nil {
		return err
	}
	g.id = cardID

	cardParts := strings.Split(parts[1], "|")
	if len(cardParts) != 2 {
		return fmt.Errorf("invalid card format: %s", line)
	}

	winningNumbers := re.FindAllString(cardParts[0], -1)
	for _, numStr := range winningNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}
		g.winningCards = append(g.winningCards, num)
	}

	myNumbers := re.FindAllString(cardParts[1], -1)
	for _, numStr := range myNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}
		g.myCards = append(g.myCards, num)
	}

	return nil
}

func (g *Game) checkWinning() int {
	var winningNumber int
	for i := 0; i < len(g.winningCards); i++ {
		num := g.winningCards[i]
		if !contains(g.myCards, num) {
			continue
		}
		winningNumber += 1
	}

	return winningNumber
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}

func getCardById(games []Game, id int) *Game {
	for _, g := range games {
		if g.id == id {
			return &g
		}
	}
	return nil
}

func main() {
	grid, _ := readInputContent("input.txt")
	var games []Game
	for _, line := range grid {
		g := Game{}
		err := g.processCard(line)
		if err != nil {
			fmt.Println(err)
		}
		games = append(games, g)
	}
	for i := 0; i < len(games); i++ {
		g := games[i]
		winningNumber := g.checkWinning()
		if winningNumber != 0 {
			for j := g.id + 1; j <= g.id+winningNumber; j++ {
				g2 := getCardById(games, j)
				if g2 == nil {
					continue
				}
				games = append(games, *g2)
			}
		}
	}

	fmt.Println(len(games))
}
