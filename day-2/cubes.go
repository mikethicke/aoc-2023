package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func parseLine(line string) (int, [][]int) {
	var parsedLine [][]int

	lineParts := strings.Split(line, `:`)
	if len(lineParts) != 2 {
		log.Fatal(`Error parsing lineParts`)
	}

	gameNumber, err := strconv.Atoi(strings.TrimPrefix(lineParts[0], `Game `))
	if err != nil {
		log.Fatal(`Error parsing Game number`)
	}

	blueRE := regexp.MustCompile(`(\d+) blue`)
	greenRE := regexp.MustCompile(`(\d+) green`)
	redRE := regexp.MustCompile(`(\d+) red`)

	sets := strings.Split(lineParts[1], `;`)
	for _, set := range sets {
		parsedSet := []int{0, 0, 0}

		matches := blueRE.FindStringSubmatch(set)
		if len(matches) == 2 {
			count, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(`Failed to parse blue cube count`)
			}
			parsedSet[0] = count
		}

		matches = greenRE.FindStringSubmatch(set)
		if len(matches) == 2 {
			count, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(`Failed to parse blue cube count`)
			}
			parsedSet[1] = count
		}

		matches = redRE.FindStringSubmatch(set)
		if len(matches) == 2 {
			count, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(`Failed to parse blue cube count`)
			}
			parsedSet[2] = count
		}

		parsedLine = append(parsedLine, parsedSet)
	}

	return gameNumber, parsedLine
}

func linePossible(line [][]int, constraint [3]int) bool {
	for _, set := range line {
		if set[0] > constraint[0] {
			return false
		}
		if set[1] > constraint[1] {
			return false
		}
		if set[2] > constraint[2] {
			return false
		}
	}
	return true
}

func fewestCubes(line [][]int) []int {
	fewest := make([]int, 3)
	for _, set := range line {
		for i, count := range set {
			if fewest[i] < count {
				fewest[i] = count
			}
		}
	}
	return fewest
}

func main() {
	lines := inputFileToLines(`input.txt`)
	sum := 0
	for _, line := range lines {
		gameNumber, parsedLine := parseLine(line)
		if linePossible(parsedLine, [3]int{14, 13, 12}) {
			sum += gameNumber
		}
	}

	fmt.Println(`Sum of possible lines: ` + strconv.Itoa(sum))

	sum = 0
	for _, line := range lines {
		_, parsedLine := parseLine(line)
		fewest := fewestCubes(parsedLine)
		power := fewest[0] * fewest[1] * fewest[2]
		sum += power
	}

	fmt.Println(`Sum of fewest cubes: ` + strconv.Itoa(sum))
}
