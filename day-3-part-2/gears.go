package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type GearPosition [2]int

type PartInfo struct {
	number   int
	startPos int
	endPos   int
}

func main() {
	lines := inputFileToLines(`input.txt`)
	gears := findGears(lines)
	partPositions := getPartPositions(lines)
	gearRatios := calculateGearRatios(gears, partPositions)
	sum := sumGearRatios(gearRatios)
	log.Println(sum)
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

func findGears(lines []string) []GearPosition {
	gears := make([]GearPosition, 0)
	for lineNumber, line := range lines {
		for position, char := range line {
			if char == '*' {
				gears = append(gears, GearPosition{lineNumber, position})
			}
		}
	}
	return gears
}

func getPartPositions(lines []string) map[int][]PartInfo {
	partNumbers := make(map[int][]PartInfo)

	numbersRE := regexp.MustCompile(`\d+`)
	for lineNumber, line := range lines {
		numberMatches := numbersRE.FindAllStringIndex(line, -1)
		for _, match := range numberMatches {
			value, err := strconv.Atoi(line[match[0]:match[1]])
			if err != nil {
				log.Fatal(`Error parsing number in line`)
			}
			partNumbers[lineNumber] = append(partNumbers[lineNumber], PartInfo{value, match[0], match[1] - 1})
		}
	}
	return partNumbers
}

func calculateGearRatios(gears []GearPosition, partPositions map[int][]PartInfo) []int {
	gearRatios := make([]int, 0)
	for _, gear := range gears {
		gearRatio := calculateGearRatio(gear, partPositions)
		if gearRatio != 0 {
			gearRatios = append(gearRatios, gearRatio)
		}
	}
	return gearRatios
}

func calculateGearRatio(gear GearPosition, partPositions map[int][]PartInfo) int {
	adjacentParts := findAdjacentParts(gear, partPositions)
	if len(adjacentParts) != 2 {
		return 0
	}
	gearRatio := adjacentParts[0] * adjacentParts[1]
	return gearRatio
}

func findAdjacentParts(gear GearPosition, partPositions map[int][]PartInfo) []int {
	parts := make([]int, 0)
	parts = append(parts, findPartsTouchingPoint(gear[1], gear[0]-1, partPositions)...)
	parts = append(parts, findPartsTouchingPoint(gear[1], gear[0], partPositions)...)
	parts = append(parts, findPartsTouchingPoint(gear[1], gear[0]+1, partPositions)...)
	return parts
}

func findPartsTouchingPoint(point int, lineNumber int, partPositions map[int][]PartInfo) []int {
	parts := make([]int, 0)
	partsOnLine := partPositions[lineNumber]
	for _, part := range partsOnLine {
		if part.startPos <= point+1 && part.endPos >= point-1 {
			parts = append(parts, part.number)
		}
	}
	return parts
}

func sumGearRatios(gearRatios []int) int {
	sum := 0
	for _, gearRatio := range gearRatios {
		sum += gearRatio
	}
	return sum
}
