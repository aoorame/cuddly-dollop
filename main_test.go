package main

import (
	"testing"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		inputSize   int
		expectedLen int
	}{
		{"Negative size", -3, 0},
		{"Zero size", 0, 0},
		{"Positive size", 5, 5},
		{"Large size", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.inputSize)
			if len(result) != tt.expectedLen {
				t.Errorf("Для размера %d ожидали длину %d, но получили %d",
					tt.inputSize, tt.expectedLen, len(result))
			}
		})
	}
	t.Run("Check elements range", func(t *testing.T) {
		size := 500
		result := generateRandomElements(size)

		for i, val := range result {
			if val < 1 || val > 1000 {
				t.Errorf("Элемент под индексом %d равен %d, что выходит за рамки [1, 1000]", i, val)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Nil slice",
			input:    nil,
			expected: 0,
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "All positive numbers",
			input:    []int{1, 5, 3, 9, 2},
			expected: 9,
		},
		{
			name:     "All negative numbers",
			input:    []int{-10, -5, -23, -2},
			expected: -2,
		},
		{
			name:     "Mixed numbers",
			input:    []int{-10, 0, 5, -2, 14, 3},
			expected: 14,
		},
		{
			name:     "Duplicate maximums",
			input:    []int{7, 2, 7, 5, 7},
			expected: 7,
		},
		{
			name:     "All identical elements",
			input:    []int{3, 3, 3, 3},
			expected: 3,
		},
		{
			name:     "Maximum is the second element",
			input:    []int{10, 50, 20, 5, 30},
			expected: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maximum(tt.input)
			if actual != tt.expected {
				t.Errorf("maximum(%v) = %d; ожидали %d", tt.input, actual, tt.expected)
			}
		})
	}
}
