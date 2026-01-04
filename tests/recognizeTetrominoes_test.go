package tests

import (
	"testing"
)

func TestRecognizeTetrominoes_ValidSquare(t *testing.T) {
	tetromino := []string{
		"##..",
		"##..",
		"....",
		"....",
	}

	err := RecognizeTetrominoes(tetromino)
	if err != nil {
		t.Fatalf("expected valid tetromino, got error: %v", err)
	}
}

func TestRecognizeTetrominoes_ValidLine(t *testing.T) {
	tetromino := []string{
		"####",
		"....",
		"....",
		"....",
	}

	err := RecognizeTetrominoes(tetromino)
	if err != nil {
		t.Fatalf("expected valid line tetromino, got error: %v", err)
	}
}

func TestRecognizeTetrominoes_InvalidCount(t *testing.T) {
	tetromino := []string{
		"###.",
		"....",
		"....",
		"....",
	}

	err := RecognizeTetrominoes(tetromino)
	if err == nil {
		t.Fatal("expected error for invalid # count, got nil")
	}
}

func TestRecognizeTetrominoes_DisconnectedBlocks(t *testing.T) {
	tetromino := []string{
		"#.#.",
		"....",
		"#...",
		"....",
	}

	err := RecognizeTetrominoes(tetromino)
	if err == nil {
		t.Fatal("expected error for disconnected blocks")
	}
}

func TestRecognizeTetrominoes_DiagonalIsInvalid(t *testing.T) {
	tetromino := []string{
		"#...",
		".#..",
		"..#.",
		"...#",
	}

	err := RecognizeTetrominoes(tetromino)
	if err == nil {
		t.Fatal("expected error for diagonal-only connections")
	}
}
