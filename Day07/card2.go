package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards    []int
	bid      int
	handType int
}

var typeClassement = map[string]int{
	"high card":       1,
	"one pair":        2,
	"two pairs":       3,
	"three of a kind": 4,
	"full house":      5,
	"four of a kind":  6,
	"five of a kind":  7,
}

var classement = []int{
	'A': 13, 'K': 12, 'Q': 11, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5,
	'4': 4, '3': 3, '2': 2, 'J': 1,
}

func parseGrid(grid []string) []hand {
	var hands []hand
	for _, line := range grid {
		parts := strings.Split(line, " ")
		cardsPart := parts[0]
		bidPart := parts[1]

		var h hand
		for _, r := range cardsPart {
			h.cards = append(h.cards, classement[r])
		}

		bid, err := strconv.Atoi(bidPart)
		if err != nil {
			continue
		}
		h.bid = bid

		hands = append(hands, h)
	}
	return hands
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

func sortHands(hands []*hand) []*hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		}

		for k := 0; k < len(hands[i].cards); k++ {
			if hands[i].cards[k] != hands[j].cards[k] {
				return hands[i].cards[k] < hands[j].cards[k]
			}
		}

		return false
	})
	return hands
}

func determineBestHandWithJacks(frequency map[int]int, jacksCount int) int {
	bestHand := 0

	for card := range frequency {
		tempFreq := make(map[int]int)
		for k, v := range frequency {
			tempFreq[k] = v
		}
		tempFreq[card] += jacksCount

		handType := getHandType(tempFreq)
		if handType > bestHand {
			bestHand = handType
		}
	}

	if bestHand == 0 {
		frequency[0] = jacksCount
		bestHand = getHandType(frequency)
	}

	return bestHand
}

func getHandType(frequency map[int]int) int {
	var freqs []int
	for _, freq := range frequency {
		freqs = append(freqs, freq)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))

	switch {
	case freqs[0] == 5:
		return typeClassement["five of a kind"]
	case freqs[0] == 4:
		return typeClassement["four of a kind"]
	case freqs[0] == 3 && freqs[1] == 2:
		return typeClassement["full house"]
	case freqs[0] == 3:
		return typeClassement["three of a kind"]
	case freqs[0] == 2 && freqs[1] == 2:
		return typeClassement["two pairs"]
	case freqs[0] == 2:
		return typeClassement["one pair"]
	default:
		return typeClassement["high card"]
	}
}

func determineHandType(h *hand) int {
	frequency := make(map[int]int)
	jacksCount := 0

	for _, card := range h.cards {
		if card == 1 {
			jacksCount++
		} else {
			frequency[card]++
		}
	}

	if jacksCount > 0 {
		return determineBestHandWithJacks(frequency, jacksCount)
	}

	return getHandType(frequency)
}

func main() {
	grid := readInputContent("input.txt")
	hands := parseGrid(grid)

	handsPtrs := make([]*hand, len(hands))
	for i := range hands {
		handsPtrs[i] = &hands[i]
		handsPtrs[i].handType = determineHandType(handsPtrs[i])
	}

	sortedHands := sortHands(handsPtrs)
	for _, h := range sortedHands {
		fmt.Println("Mise:", h.bid, "Type:", h.handType, "Cartes:", h.cards)
	}
	var total int
	for i, h := range sortedHands {
		count := (i + 1) * h.bid
		total += count
	}
	fmt.Println(total)
}
