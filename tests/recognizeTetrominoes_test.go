package tests

import (
	"testing"
	"tetris-optimizer/pipeline"
)

func TestRecognizeTetrominoes_ValidSquare(t *testing.T) {
	// A valid 2x2 square tetromino should pass validation
	tetromino := []string{
		"##..",
		"##..",
		"....",
		"....",
	}

	err := pipeline.RecognizeTetrominoes(tetromino)
	if err != nil {
		t.Fatalf("expected valid tetromino, got error: %v", err)
	}
}

func TestRecognizeTetrominoes_ValidLine(t *testing.T) {
	// A valid straight 4-long line tetromino should pass validation
	tetromino := []string{
		"####",
		"....",
		"....",
		"....",
	}

	err := pipeline.RecognizeTetrominoes(tetromino)
	if err != nil {
		t.Fatalf("expected valid line tetromino, got error: %v", err)
	}
}

func TestRecognizeTetrominoes_InvalidCount(t *testing.T) {
	// Tetromino with wrong number of '#' should fail
	tetromino := []string{
		"###.",
		"....",
		"....",
		"....",
	}

	err := pipeline.RecognizeTetrominoes(tetromino)
	if err == nil {
		t.Fatal("expected error for invalid # count, got nil")
	}
}

func TestRecognizeTetrominoes_DisconnectedBlocks(t *testing.T) {
	// Disconnected blocks (not orthogonally connected) should fail
	tetromino := []string{
		"#.#.",
		"....",
		"#...",
		"....",
	}

	err := pipeline.RecognizeTetrominoes(tetromino)
	if err == nil {
		t.Fatal("expected error for disconnected blocks")
	}
}

func TestRecognizeTetrominoes_DiagonalIsInvalid(t *testing.T) {
	// Diagonal-only blocks are not valid tetromino shapes
	tetromino := []string{
		"#...",
		".#..",
		"..#.",
		"...#",
	}

	err := pipeline.RecognizeTetrominoes(tetromino)
	if err == nil {
		t.Fatal("expected error for diagonal-only connections")
	}
}
