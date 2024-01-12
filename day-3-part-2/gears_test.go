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

func TestFindGears(t *testing.T) {
	tests := []struct {
		input []string
		want  []GearPosition
	}{
		{input, []GearPosition{{1, 3}, {4, 3}, {8, 5}}},
	}

	for _, test := range tests {
		got := findGears(test.input)

		if len(got) != len(test.want) {
			t.Errorf("len(findGears()) = %d, want %d", len(got), len(test.want))
		}

	}
}

func TestFindAdjacentParts(t *testing.T) {
	tests := []struct {
		input GearPosition
		want  []int
	}{
		{
			GearPosition{1, 3},
			[]int{467, 35},
		},
		{
			GearPosition{4, 3},
			[]int{617},
		},
		{
			GearPosition{8, 5},
			[]int{755, 598},
		},
	}

	parts := getPartPositions(input)

	for _, test := range tests {
		got := findAdjacentParts(test.input, parts)

		if len(got) != len(test.want) {
			t.Errorf("len(findAdjacentParts()) = %d, want %d", len(got), len(test.want))
		}

		sort.Ints(got)
		sort.Ints(test.want)
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("findAdjacentParts() = %v, want %v", got, test.want)
				break
			}
		}
	}
}
