package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type PartCoordinate struct {
	lineNumber int
	startPos   int
	endPos     int
}

type SymbolPositions map[int][]int

func main() {
	lines := inputFileToLines(`input.txt`)
	parts := parsePartNumbers(lines)
	symbols := parseSymbolLocations(lines)
	validParts := filterValidPartNumbers(parts, symbols)

	sum := 0
	for _, partNumber := range validParts {
		sum += partNumber
	}
	fmt.Println(sum)
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

func parsePartNumbers(lines []string) map[int][]PartCoordinate {
	partNumbers := make(map[int][]PartCoordinate)
	numbersRE := regexp.MustCompile(`\d+`)
	for lineNumber, line := range lines {
		numberMatches := numbersRE.FindAllStringIndex(line, -1)
		for _, match := range numberMatches {
			value, err := strconv.Atoi(line[match[0]:match[1]])
			if err != nil {
				log.Fatal(`Error parsing number in line`)
			}
			partNumbers[value] = append(partNumbers[value], PartCoordinate{lineNumber, match[0], match[1] - 1})
		}
	}
	return partNumbers
}

func parseSymbolLocations(lines []string) SymbolPositions {
	symbols := make(SymbolPositions)
	symbolsRE := regexp.MustCompile(`[^.0-9]`)

	for lineNumber, line := range lines {
		symbolMatches := symbolsRE.FindAllStringIndex(line, -1)
		for _, match := range symbolMatches {
			symbols[lineNumber] = append(symbols[lineNumber], match[0])
		}
	}
	return symbols
}

func filterValidPartNumbers(
	parts map[int][]PartCoordinate,
	symbols SymbolPositions,
) []int {
	validPartNumbers := make([]int, 0)
	for partNumber, coordinates := range parts {
		for _, coordinate := range coordinates {
			if partNumberIsValid(coordinate, symbols) {
				validPartNumbers = append(validPartNumbers, partNumber)
			}
		}
	}
	return validPartNumbers
}

func partNumberIsValid(location PartCoordinate, symbols SymbolPositions) bool {
	if isSymbolInRange(
		symbols,
		location.lineNumber-1,
		location.startPos-1,
		location.endPos+1,
	) {
		return true
	}
	if isSymbolInRange(
		symbols,
		location.lineNumber,
		location.startPos-1,
		location.endPos+1,
	) {
		return true
	}
	if isSymbolInRange(
		symbols,
		location.lineNumber+1,
		location.startPos-1,
		location.endPos+1,
	) {
		return true
	}
	return false
}

func isSymbolInRange(
	symbols SymbolPositions,
	lineNumber int,
	startPos int,
	endPos int,
) bool {
	positions, ok := symbols[lineNumber]
	if !ok {
		return false
	}
	for _, pos := range positions {
		if pos >= startPos && pos <= endPos {
			return true
		}
	}
	return false
}
