package tests

import (
	"testing"

	"tetris-optimizer/pipeline"
)

func boardToStrings(board [][]rune) []string {
	var out []string
	for _, row := range board {
		out = append(out, string(row))
	}
	return out
}

func TestAsembleTetrominoes_SingleSquare(t *testing.T) {
	// Single 2x2 square tetromino should assemble into a 2x2 board filled with 'A'
	tetrominoes := [][]string{
		{
			"##..",
			"##..",
			"....",
			"....",
		},
	}

	board, err := pipeline.AsembleTetrominoes(tetrominoes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := boardToStrings(board)

	expected := []string{
		"AA",
		"AA",
	}

	if len(result) != len(expected) {
		t.Fatalf("expected board size %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	}
}

func TestAsembleTetrominoes_TwoSquares(t *testing.T) {
	// Two separate 2x2 squares should assemble side-by-side using letters A and B
	tetrominoes := [][]string{
		{
			"##..",
			"##..",
			"....",
			"....",
		},
		{
			"##..",
			"##..",
			"....",
			"....",
		},
	}

	board, err := pipeline.AsembleTetrominoes(tetrominoes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := boardToStrings(board)

	expected := []string{
		"AABB",
		"AABB",
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	}
}

func TestAsembleTetrominoes_NeedsLargerBoard(t *testing.T) {
	// Two 4-long line tetrominoes require a board with height at least 4
	tetrominoes := [][]string{
		{
			"####",
			"....",
			"....",
			"....",
		},
		{
			"####",
			"....",
			"....",
			"....",
		},
	}

	board, err := pipeline.AsembleTetrominoes(tetrominoes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(board) < 4 {
		t.Fatalf("expected board size >= 4, got %d", len(board))
	}
}

func TestAsembleTetrominoes_InvalidInput(t *testing.T) {
	// Passing nil should return an error
	_, err := pipeline.AsembleTetrominoes(nil)
	if err == nil {
		t.Fatal("expected error for nil input")
	}
}
