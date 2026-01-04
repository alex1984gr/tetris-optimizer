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
	_, err := pipeline.AsembleTetrominoes(nil)
	if err == nil {
		t.Fatal("expected error for nil input")
	}
}
