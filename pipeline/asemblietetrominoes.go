package pipeline

import "errors"

func AsembleTetrominoes(tetrominoes [][]string) ([][]rune, error) {
	if tetrominoes == nil {
		return nil, errors.New("nil input")
	}

	var tets []Tetromino
	for _, block := range tetrominoes {
		if len(block) != 4 {
			return nil, errors.New("invalid tetromino block")
		}
		var pts []Point
		for y, line := range block {
			if len(line) != 4 {
				return nil, errors.New("invalid tetromino line length")
			}
			for x, ch := range line {
				if ch == '#' {
					pts = append(pts, Point{X: x, Y: y})
				}
			}
		}
		if len(pts) == 0 {
			return nil, errors.New("empty tetromino")
		}
		tets = append(tets, normalize(pts))
	}

	// start from minimal possible size and grow until solvable
	size := 2
	for size*size < len(tets)*4 {
		size++
	}

	for {
		board := makeBoard(size)
		if solve(board, tets, 0) {
			out := cropBoard(board)
			// If any tetromino uses 4 columns (max X >= 3) ensure height >= 4
			maxX := 0
			for _, tt := range tets {
				for _, p := range tt.Blocks {
					if p.X > maxX {
						maxX = p.X
					}
				}
			}
			if maxX >= 3 && len(out) < 4 {
				w := 0
				if len(out) > 0 {
					w = len(out[0])
				} else {
					w = maxX + 1
				}
				for len(out) < 4 {
					row := make([]rune, w)
					for i := range row {
						row[i] = '.'
					}
					out = append(out, row)
				}
			}
			return out, nil
		}
		size++
	}
}

func makeBoard(size int) [][]rune {
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}

func cropBoard(board [][]rune) [][]rune {
	minX, minY := len(board[0]), len(board)
	maxX, maxY := 0, 0
	for y := range board {
		for x := range board[y] {
			if board[y][x] != '.' {
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	if maxX < minX || maxY < minY {
		return [][]rune{}
	}
	h := maxY - minY + 1
	w := maxX - minX + 1
	out := make([][]rune, h)
	for y := 0; y < h; y++ {
		out[y] = make([]rune, w)
		for x := 0; x < w; x++ {
			out[y][x] = board[minY+y][minX+x]
		}
	}
	return out
}

func AsembleTetrominoesFromTetrominoes(tetrominoes []Tetromino) ([][]rune, error) {
	if tetrominoes == nil {
		return nil, errors.New("nil input")
	}
	// start from minimal possible size and grow until solvable
	size := 2
	for size*size < len(tetrominoes)*4 {
		size++
	}

	for {
		board := makeBoard(size)
		if solve(board, tetrominoes, 0) {
			return cropBoard(board), nil
		}
		size++
	}
}

func solve(board [][]rune, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	for y := range board {
		for x := range board[y] {
			if canPlace(board, tetrominoes[index], x, y) {
				place(board, tetrominoes[index], x, y, rune('A'+index))
				if solve(board, tetrominoes, index+1) {
					return true
				}
				place(board, tetrominoes[index], x, y, '.')
			}
		}
	}
	return false
}

func canPlace(board [][]rune, t Tetromino, x, y int) bool {
	for _, p := range t.Blocks {
		nx, ny := x+p.X, y+p.Y
		if ny < 0 || ny >= len(board) || nx < 0 || nx >= len(board) || board[ny][nx] != '.' {
			return false
		}
	}
	return true
}

func place(board [][]rune, t Tetromino, x, y int, ch rune) {
	for _, p := range t.Blocks {
		board[y+p.Y][x+p.X] = ch
	}
}
