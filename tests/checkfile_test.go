package tests

import (
	"testing"
	"tetris-optimizer/pipeline"
)

func TestCheckFile_ValidSingleTetromino(t *testing.T) {
	lines := []string{
		"#...",
		"#...",
		"#...",
		"#...",
	}

	err := pipeline.CheckFile(lines)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCheckFile_ValidMultipleTetrominoes(t *testing.T) {
	lines := []string{
		"#...",
		"#...",
		"#...",
		"#...",
		"",
		"....",
		"....",
		"..##",
		"..##",
	}

	err := pipeline.CheckFile(lines)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCheckFile_InvalidCharacter(t *testing.T) {
	lines := []string{
		"#..@",
		"#...",
		"#...",
		"#...",
	}

	err := pipeline.CheckFile(lines)
	if err == nil {
		t.Fatalf("Expected error for invalid character, got nil")
	}
}

func TestCheckFile_MissingEmptyLineBetweenTetrominoes(t *testing.T) {
	lines := []string{
		"#...",
		"#...",
		"#...",
		"#...",
		"....",
		"....",
		"..##",
		"..##",
	}

	err := pipeline.CheckFile(lines)
	if err == nil {
		t.Fatalf("Expected error for missing empty line between tetrominoes, got nil")
	}
}

func TestCheckFile_NotEnoughLinesForTetromino(t *testing.T) {
	lines := []string{
		"#...",
		"#...",
		"#...",
	}

	err := pipeline.CheckFile(lines)
	if err == nil {
		t.Fatalf("Expected error for incomplete tetromino, got nil")
	}
}

func TestCheckFile_ExtraEmptyLineAtEndAllowed(t *testing.T) {
	lines := []string{
		"#...",
		"#...",
		"#...",
		"#...",
		"",
	}

	err := pipeline.CheckFile(lines)
	if err != nil {
		t.Fatalf("Expected no error , got %v", err)
	}
}

func TestCheckFile_EmptyInput(t *testing.T) {
	var lines []string

	err := pipeline.CheckFile(lines)
	if err != nil {
		t.Fatalf("Expected no error for empty input, got %v", err)
	}
}
