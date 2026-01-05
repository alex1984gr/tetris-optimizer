package tests

import (
	"bytes"
	"testing"

	"tetris-optimizer/pipeline"
)

func TestPrintOutput_ValidBoard(t *testing.T) {
	// Valid 4x4 board with two 2x2 blocks labeled A and B
	board := [][]rune{
		{'A', 'A', '.', '.'},
		{'A', 'A', '.', '.'},
		{'.', '.', 'B', 'B'},
		{'.', '.', 'B', 'B'},
	}

	var buf bytes.Buffer

	err := pipeline.PrintOutput(board, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "" +
		"AA..\n" +
		"AA..\n" +
		"..BB\n" +
		"..BB\n"

	if buf.String() != expected {
		t.Errorf("unexpected output:\nexpected:\n%s\ngot:\n%s", expected, buf.String())
	}
}

func TestPrintOutput_EmptyBoard(t *testing.T) {
	// Nil board should produce no output
	var buf bytes.Buffer

	err := pipeline.PrintOutput(nil, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if buf.String() != "" {
		t.Errorf("expected empty output, got %q", buf.String())
	}
}

func TestPrintOutput_NilWriter(t *testing.T) {
	// Passing nil writer should return an error
	board := [][]rune{
		{'A'},
	}

	err := pipeline.PrintOutput(board, nil)
	if err == nil {
		t.Fatal("expected error for nil writer")
	}
}
