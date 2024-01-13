package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type ScratchCard struct {
	cardNumber     int
	winningNumbers []int
	yourNumbers    []int
}

type ScratchSet struct {
	count        int
	winningCount int
	card         ScratchCard
}

func main() {
	lines := inputFileToLines(`input.txt`)
	cards := parseLines(lines)
	totalScore := scoreCards(cards)
	fmt.Println(totalScore)

	sets := parseLinesToSets(lines)
	processedSets := processSets(sets)
	totalSets := countSets(processedSets)
	fmt.Println(totalSets)
}

func inputFileToLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseLines(rawLines []string) []ScratchCard {
	cards := make([]ScratchCard, 0)
	for _, rawLine := range rawLines {
		cards = append(cards, parseLine(rawLine))
	}
	return cards
}

func parseLine(rawLine string) ScratchCard {
	re := regexp.MustCompile(`Card\s+(\d+)\:((?:\s+\d+)+) \|((?:\s+\d+)+)`)
	numbersRE := regexp.MustCompile(`\d+`)
	matches := re.FindStringSubmatch(rawLine)

	cardNumber, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(`Failed to parse cardNumber`)
	}

	winners := numbersRE.FindAllString(matches[2], -1)
	winningNumbers := make([]int, 0)
	for _, winner := range winners {
		winningNumber, err := strconv.Atoi(winner)
		if err != nil {
			log.Fatal(`Error converting winning number`)
		}
		winningNumbers = append(winningNumbers, winningNumber)
	}

	yours := numbersRE.FindAllString(matches[3], -1)
	yourNumbers := make([]int, 0)
	for _, your := range yours {
		yourNumber, err := strconv.Atoi(your)
		if err != nil {
			log.Fatal(`Error converting winning number`)
		}
		yourNumbers = append(yourNumbers, yourNumber)
	}

	card := ScratchCard{
		cardNumber:     cardNumber,
		winningNumbers: winningNumbers,
		yourNumbers:    yourNumbers,
	}
	return card
}

func scoreCards(cards []ScratchCard) int {
	totalScore := 0
	for _, card := range cards {
		winningCount := countWinningNumbers(card)
		totalScore += scoreWinningCount(winningCount)
	}
	return totalScore
}

func countWinningNumbers(card ScratchCard) int {
	winningCount := 0
	for _, yourNumber := range card.yourNumbers {
		for _, winningNumber := range card.winningNumbers {
			if yourNumber == winningNumber {
				winningCount++
				break
			}
		}
	}
	return winningCount
}

func scoreWinningCount(winningCount int) int {
	if winningCount == 0 {
		return 0
	}
	score := 1
	for x := 1; x < winningCount; x++ {
		score = score * 2
	}
	return score
}

func parseLinesToSets(rawLines []string) map[int]ScratchSet {
	scratchSets := make(map[int]ScratchSet)
	for _, rawLine := range rawLines {
		card := parseLine(rawLine)
		winningCount := countWinningNumbers(card)
		scratchSets[card.cardNumber] = ScratchSet{
			count:        1,
			winningCount: winningCount,
			card:         card,
		}
	}
	return scratchSets
}

func processSets(scratchSets map[int]ScratchSet) map[int]ScratchSet {
	for cardNumber := 1; cardNumber <= len(scratchSets); cardNumber++ {
		for inc := 1; inc <= scratchSets[cardNumber].winningCount; inc++ {
			temp := scratchSets[cardNumber+inc]
			temp.count += scratchSets[cardNumber].count
			scratchSets[cardNumber+inc] = temp
		}
	}
	return scratchSets
}

func countSets(scratchSets map[int]ScratchSet) int {
	count := 0
	for _, set := range scratchSets {
		count += set.count
	}
	return count
}
