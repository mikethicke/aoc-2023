// This approach failed, but I'm leaving it here anyways.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type ParsedLine struct {
	numbers map[int][]int
	symbols []int
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

func parseLines(rawLines []string) []ParsedLine {
	parsedLines := make([]ParsedLine, 0)
	for _, rawLine := range rawLines {
		parsedLines = append(parsedLines, parseLine(rawLine))
	}
	return parsedLines
}

func parseLine(input string) ParsedLine {
	numbersRE := regexp.MustCompile(`\d+`)
	numbers := make(map[int][]int)
	numberMatches := numbersRE.FindAllStringIndex(input, -1)
	for _, match := range numberMatches {
		value, err := strconv.Atoi(input[match[0]:match[1]])
		if err != nil {
			log.Fatal(`Error parsing number in line`)
		}
		numbers[value] = []int{match[0], match[1] - 1}
	}

	symbolsRE := regexp.MustCompile(`[^.0-9]`)
	symbolMatches := symbolsRE.FindAllStringIndex(input, -1)
	symbols := make([]int, 0)
	for _, match := range symbolMatches {
		symbols = append(symbols, match[0])
	}

	line := ParsedLine{
		numbers,
		symbols,
	}

	return line
}

func sumPartNumbers(lines []ParsedLine) int {
	sum := 0
	for lineNumber := 0; lineNumber < len(lines); lineNumber++ {
		sum += sumPartNumbersInLine(lineNumber, lines)
	}
	return sum
}

func sumPartNumbersInLine(lineNumber int, lines []ParsedLine) int {
	sum := 0
	for partNumber, location := range lines[lineNumber].numbers {
		if lineNumber > 0 {
			if isSymbolInRange(lines[lineNumber-1], location[0]-1, location[1]+1) {
				sum += partNumber
				continue
			}
		}
		if isSymbolInRange(lines[lineNumber], location[0]-1, location[1]+1) {
			sum += partNumber
			continue
		}
		if lineNumber < len(lines)-1 {
			if isSymbolInRange(lines[lineNumber+1], location[0]-1, location[1]+1) {
				sum += partNumber
			}
		}
	}
	return sum
}

func isSymbolInRange(line ParsedLine, startPosition int, endPosition int) bool {
	for _, symbolPosition := range line.symbols {
		if symbolPosition >= startPosition && symbolPosition <= endPosition {
			return true
		}
	}
	return false
}

func main() {
	lines := inputFileToLines(`input.txt`)
	parsedLines := parseLines(lines)
	sum := sumPartNumbers(parsedLines)
	fmt.Println(sum)
}
