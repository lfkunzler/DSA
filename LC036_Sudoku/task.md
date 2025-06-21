# Description
```
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
```

# Approach
```
We need 3 maps to validate lines, columns and 3x3 squares.
- board[i][j] when j++ we go through the columns of the board
- board[i][j] when i++ we go through the lines of the board
- for 3x3 squares, we use the following equation: (row / 3) * 3 + (col / 3)
```

## Time complexity
O(n + n^2) -> O(n^2)

## Space complexity
O(3n^2) -> O(n^2)
