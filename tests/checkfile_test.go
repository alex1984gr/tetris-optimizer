package tests

import (
	"testing"
	"tetris-optimizer/pipeline"
)

func TestCheckFile_ValidSingleTetromino(t *testing.T) {
	// Single-column tetromino (vertical line) should be valid
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
	// Two tetrominoes separated by an empty line should be valid
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
	// Presence of an invalid character should trigger an error
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
	// Two tetromino blocks without an empty separator should fail
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
	// Not having 4 lines for a tetromino should cause an error
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
	// Trailing empty line after a tetromino is allowed
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
	// Empty input should be considered valid (no tetrominoes)
	var lines []string

	err := pipeline.CheckFile(lines)
	if err != nil {
		t.Fatalf("Expected no error for empty input, got %v", err)
	}
}
