package pipeline

import "errors"

type placedBlock struct {
	row int
	col int
}

type piece struct {
	blocks []placedBlock
	letter rune
}

// Δημόσια API
func AsembleTetrominoes(tetrominoes [][]string) ([][]rune, error) {
	if tetrominoes == nil {
		return nil, errors.New("nil input")
	}

	pieces := parsePieces(tetrominoes)

	size := minimalBoardSize(len(pieces))

	for {
		board := createBoard(size)

		if solve(board, pieces, 0) {
			return board, nil
		}

		size++
	}
}

/* =======================
   Parsing helpers
======================= */

func parsePieces(tetrominoes [][]string) []piece {
	var pieces []piece

	for i, t := range tetrominoes {
		var blocks []placedBlock

		for r := 0; r < 4; r++ {
			for c := 0; c < 4; c++ {
				if t[r][c] == '#' {
					blocks = append(blocks, placedBlock{r, c})
				}
			}
		}

		normalize(&blocks)
		pieces = append(pieces, piece{
			blocks: blocks,
			letter: rune('A' + i),
		})
	}

	return pieces
}

func normalize(blocks *[]placedBlock) {
	minR, minC := 4, 4
	for _, b := range *blocks {
		if b.row < minR {
			minR = b.row
		}
		if b.col < minC {
			minC = b.col
		}
	}

	for i := range *blocks {
		(*blocks)[i].row -= minR
		(*blocks)[i].col -= minC
	}
}

/* =======================
   Solver
======================= */

func solve(board [][]rune, pieces []piece, index int) bool {
	if index == len(pieces) {
		return true
	}

	size := len(board)
	p := pieces[index]

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if canPlace(board, p, r, c) {
				place(board, p, r, c)
				if solve(board, pieces, index+1) {
					return true
				}
				remove(board, p, r, c)
			}
		}
	}

	return false
}

func canPlace(board [][]rune, p piece, r int, c int) bool {
	size := len(board)

	for _, b := range p.blocks {
		nr := r + b.row
		nc := c + b.col

		if nr < 0 || nc < 0 || nr >= size || nc >= size {
			return false
		}
		if board[nr][nc] != '.' {
			return false
		}
	}
	return true
}

func place(board [][]rune, p piece, r int, c int) {
	for _, b := range p.blocks {
		board[r+b.row][c+b.col] = p.letter
	}
}

func remove(board [][]rune, p piece, r int, c int) {
	for _, b := range p.blocks {
		board[r+b.row][c+b.col] = '.'
	}
}

/* =======================
   Board helpers
======================= */

func createBoard(size int) [][]rune {
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}

func minimalBoardSize(pieces int) int {
	cells := pieces * 4
	size := 2
	for size*size < cells {
		size++
	}
	return size
}
