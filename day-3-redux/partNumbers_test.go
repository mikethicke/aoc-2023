package main

import (
	"sort"
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

func TestFilterValidPartNumbers(t *testing.T) {
	tests := []struct {
		input []string
		want  []int
	}{
		{input, []int{467, 35, 633, 617, 592, 755, 664, 598}},
		{input2, []int{12, 34, 12, 78, 78, 23, 90, 12, 2, 2, 12, 1, 1, 56}},
	}

	for testNum, test := range tests {
		parts := parsePartNumbers(test.input)
		symbols := parseSymbolLocations(test.input)
		got := filterValidPartNumbers(parts, symbols)

		sort.Ints(got)
		sort.Ints(test.want)

		if len(got) != len(test.want) {
			t.Errorf("Testnum %v: len(filterValidPartNumbers()) = %v, want %v", testNum, len(got), len(test.want))
			continue
		}

		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("Testnum %v: filterValidPartNumbers() = %v, want %v", testNum, got, test.want)
				break
			}
		}
	}
}

func TestSumValidPartNumbers(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{input, 4361},
		{input2, 413},
	}

	for testNum, test := range tests {
		parts := parsePartNumbers(test.input)
		symbols := parseSymbolLocations(test.input)
		got := filterValidPartNumbers(parts, symbols)

		sum := 0
		for _, partNumber := range got {
			sum += partNumber
		}
		if sum != test.want {
			t.Errorf("Testnum %v: sum(filterValidPartNumbers()) = %v, want %v", testNum, sum, test.want)
		}
	}
}
