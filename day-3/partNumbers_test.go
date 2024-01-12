package main

import (
	"reflect"
	"testing"
)

var input = []string{
	`467..114..`,
	`...*......`,
	`..35..633.`,
	`......#...`,
	`617*......`,
	`.....+.58.`,
	`..592.....`,
	`......755.`,
	`...$.*....`,
	`.664.598..`,
}

var input2 = []string{
	`12.......*..`,
	`+.........34`,
	`.......-12..`,
	`..78........`,
	`..*....60...`,
	`78..........`,
	`.......23...`,
	`....90*12...`,
	`............`,
	`2.2......12.`,
	`.*.........*`,
	`1.1.......56`,
}

func TestSumPartNumbers(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{input, 4361},
		{input2, 413},
	}

	for _, test := range tests {
		lines := parseLines(test.input)
		got := sumPartNumbers(lines)
		if got != test.want {
			t.Errorf("sumPartNumbers(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		input    string
		wantLine ParsedLine
	}{
		{
			`467..114..`,
			ParsedLine{
				map[int][]int{
					467: {0, 2},
					114: {5, 7},
				},
				[]int{},
			},
		},
		{
			`...*......`,
			ParsedLine{
				map[int][]int{},
				[]int{3},
			},
		},
		{
			`617*......`,
			ParsedLine{
				map[int][]int{
					617: {0, 2},
				},
				[]int{3},
			},
		},
		{
			`2.2......12.`,
			ParsedLine{
				map[int][]int{
					617: {0, 2},
				},
				[]int{3},
			},
		},
	}

	for _, test := range tests {
		got := parseLine(test.input)
		if !reflect.DeepEqual(got, test.wantLine) {
			t.Errorf("parseLine(%q) = %v, want %v", test.input, got, test.wantLine)
		}
	}
}

func TestIsSymbolInRange(t *testing.T) {
	tests := []struct {
		inputLine ParsedLine
		start     int
		end       int
		want      bool
	}{
		{
			ParsedLine{
				map[int][]int{
					467: {0, 2},
					114: {5, 7},
				},
				[]int{},
			},
			0,
			2,
			false,
		},
		{
			ParsedLine{
				map[int][]int{},
				[]int{3},
			},
			2,
			3,
			true,
		},
		{
			ParsedLine{
				map[int][]int{},
				[]int{3},
			},
			3,
			4,
			true,
		},
	}

	for _, test := range tests {
		got := isSymbolInRange(test.inputLine, test.start, test.end)
		if got != test.want {
			t.Errorf("isSymbolInRange(%v, %d, %d) = %t, want %t", test.inputLine, test.start, test.end, got, test.want)
		}
	}
}

func TestSumPartNumbersInLine(t *testing.T) {
	parsedLines := parseLines(input)
	tests := []struct {
		inputLine int
		want      int
	}{
		{0, 467},
		{1, 0},
		{2, 668},
		{3, 0},
		{4, 617},
		{5, 0},
		{6, 592},
		{7, 755},
		{8, 0},
		{9, 664 + 598},
	}

	for _, test := range tests {
		got := sumPartNumbersInLine(test.inputLine, parsedLines)
		if got != test.want {
			t.Errorf("sumPartNumbersInLine(%v) = %d, want %d", test.inputLine, got, test.want)
		}
	}
}
