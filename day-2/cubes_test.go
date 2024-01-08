package main

import (
	"testing"
)

func TestParseLin(t *testing.T) {
	tests := []struct {
		input    string
		wantGame int
		wantSets [][]int
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			1,
			[][]int{{3, 0, 4}, {6, 2, 1}, {0, 2, 0}},
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			2,
			[][]int{{1, 2, 0}, {4, 3, 1}, {1, 1, 0}},
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			3,
			[][]int{{6, 8, 20}, {5, 13, 4}, {0, 5, 1}},
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			4,
			[][]int{{6, 1, 3}, {0, 3, 6}, {15, 3, 14}},
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			5,
			[][]int{{1, 3, 6}, {2, 2, 1}},
		},
	}

	for _, test := range tests {
		gotGame, gotSets := parseLine(test.input)

		if gotGame != test.wantGame {
			t.Errorf("parseLine(%q) returned game %d, want %d", test.input, gotGame, test.wantGame)
		}

		for i, gotSet := range gotSets {
			for j, got := range gotSet {
				if got != test.wantSets[i][j] {
					t.Errorf("parseLine(%q) returned set %d, got %d, want %d", test.input, i, got, test.wantSets[i][j])
				}
			}
		}
	}
}

func TestLinePossible(t *testing.T) {
	tests := []struct {
		line       [][]int
		constraint [3]int
		want       bool
	}{
		{
			[][]int{{3, 0, 4}, {6, 2, 1}, {0, 2, 0}},
			[3]int{14, 13, 12},
			true,
		},
		{
			[][]int{{1, 2, 0}, {4, 3, 1}, {1, 1, 0}},
			[3]int{14, 13, 12},
			true,
		},
		{
			[][]int{{6, 8, 20}, {5, 13, 4}, {0, 5, 1}},
			[3]int{14, 13, 12},
			false,
		},
		{
			[][]int{{6, 1, 3}, {0, 3, 6}, {15, 3, 14}},
			[3]int{14, 13, 12},
			false,
		},
		{
			[][]int{{1, 3, 6}, {2, 2, 1}},
			[3]int{14, 13, 12},
			true,
		},
	}

	for _, test := range tests {
		got := linePossible(test.line, test.constraint)

		if got != test.want {
			t.Errorf("linePossible(%v, %v) returned %t, want %t", test.line, test.constraint, got, test.want)
		}
	}
}

func TestFewestCubes(t *testing.T) {
	tests := []struct {
		line [][]int
		want []int
	}{
		{
			[][]int{{3, 0, 4}, {6, 2, 1}, {0, 2, 0}},
			[]int{6, 2, 4},
		},
		{
			[][]int{{1, 2, 0}, {4, 3, 1}, {1, 1, 0}},
			[]int{4, 3, 1},
		},
		{
			[][]int{{6, 8, 20}, {5, 13, 4}, {0, 5, 1}},
			[]int{6, 13, 20},
		},
		{
			[][]int{{6, 1, 3}, {0, 3, 6}, {15, 3, 14}},
			[]int{15, 3, 14},
		},
		{
			[][]int{{1, 3, 6}, {2, 2, 1}},
			[]int{2, 3, 6},
		},
	}

	for _, test := range tests {
		got := fewestCubes(test.line)

		for i, got := range got {
			if got != test.want[i] {
				t.Errorf("fewestCubes(%v) returned %v, want %v", test.line, got, test.want)
			}
		}
	}
}
