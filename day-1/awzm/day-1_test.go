package main

import (
	"reflect"
	"testing"
)

// unit test for sum using test tables
func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "one element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: 3,
		},
		{
			name:     "three elements",
			input:    []int{1, 2, 3},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sum(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

// unit test for largest using test tables
func TestLargest(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "one element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: 2,
		},
		{
			name:     "three elements",
			input:    []int{1, 2, 3},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := largest(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

// unit test for parseEmptyLines using test tables
func TestParseEmptyLines(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "one line",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "two lines",
			input:    "hello\n\nhello",
			expected: []string{"hello", "hello"},
		},
		{
			name:     "three lines",
			input:    "hello\n\nhello\n\nhello",
			expected: []string{"hello", "hello", "hello"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := parseEmptyLines(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

// unit test for parseNumbers using test tables
func TestParseNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []int{},
		},
		{
			name:     "one line",
			input:    "1",
			expected: []int{1},
		},
		{
			name:     "two lines",
			input:    "1\n2",
			expected: []int{1, 2},
		},
		{
			name:     "three lines",
			input:    "1\n2\n3",
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := parseNumbers(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

// unit test parseInput using test tables
func TestParseInput(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected [][]int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: [][]int{},
		},
		{
			name:     "one line",
			input:    "1",
			expected: [][]int{{1}},
		},
		{
			name:  "two lines",
			input: "1\n\n2",

			expected: [][]int{{1}, {2}},
		},
		{
			name:     "three lines",
			input:    "1\n\n2\n\n3",
			expected: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := parseInput(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
