package main

import (
	"testing"
)

func TestConcatFirstLast(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"sixrrmlkptmc18zhvninek", 18},
	}

	for _, test := range tests {
		got, err := concatFirstLast(test.input)
		if err != nil {
			t.Errorf("concatFirstLast(%q) returned error %q", test.input, err)
		}
		if got != test.want {
			t.Errorf("concatFirstLast(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestConvertNumberStringsToDigits(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"two1nine", "219"},
		{"eightwothree", "8wo3"},
		{"abcone2threexyz", "abc123xyz"},
		{"xtwone3four", "x2ne34"},
		{"4nineeightseven2", "49872"},
		{"zoneight234", "z1ight234"},
		{"7pqrstsixteen", "7pqrst6teen"},
	}

	for _, test := range tests {
		got := convertNumberStringsToDigits(test.input)
		if got != test.want {
			t.Errorf("convertNumberStringsToDigits(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestSumConvertedNumberStrings(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	sum := 0
	for _, test := range tests {
		text := convertNumberStringsToDigits(test.input)
		got, err := concatFirstLast(text)
		sum += got
		if err != nil {
			t.Errorf("concatFirstLast(%q) returned error %q", test.input, err)
		}
		if got != test.want {
			t.Errorf("convertNumberStringsToDigits(%q) = %q, want %q", test.input, got, test.want)
		}
	}

	if sum != 281 {
		t.Errorf("sum = %d, want %d", sum, 281)
	}
}
