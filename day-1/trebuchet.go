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

func concatFirstLast(text string) (int, error) {
	pattern := regexp.MustCompile(`\d`)
	numbers := pattern.FindAllString(text, -1)

	first, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, err
	}

	last, err := strconv.Atoi(numbers[len(numbers)-1])
	if err != nil {
		return 0, err
	}

	return first*10 + last, nil
}

func convertNumberStringsToDigits(text string) string {
	if text == "" {
		return ""
	}

	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for number, digit := range numbers {
		if strings.HasPrefix(text, number) {
			text = digit + text[1:]
			break
		}
	}

	return text[:1] + convertNumberStringsToDigits(text[1:])
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		oldText := scanner.Text()
		text := convertNumberStringsToDigits(oldText)
		number, err := concatFirstLast(text)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
