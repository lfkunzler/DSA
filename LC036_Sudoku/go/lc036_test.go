package lc036

import "testing"

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		expected bool
	}{
		{
			name: "Valid sudoku board",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidSudoku(tt.board)
			if got != tt.expected {
				t.Errorf("isValidSudoku() for %s board returned %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}
