package tests

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"tetris-optimizer/pipeline"
)

func createTempFile(t *testing.T, content string) string {
	t.Helper()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	return filePath
}

func TestReadText_ValidFile(t *testing.T) {
	content := "....\n....\n....\n....\n"
	path := createTempFile(t, content)

	lines, err := pipeline.ReadText(path)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{
		"....",
		"....",
		"....",
		"....",
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Fatalf("Expected %v, got %v", expected, lines)
	}
}

func TestReadText_FileDoesNotExist(t *testing.T) {
	_, err := pipeline.ReadText("non_existent_file.txt")

	if err == nil {
		t.Fatalf("Expected error for non-existent file, got nil")
	}
}

func TestReadText_EmptyFile(t *testing.T) {
	path := createTempFile(t, "")

	_, err := pipeline.ReadText(path)

	if err == nil {
		t.Fatal("expected error for empty file, got nil")
	}
}

func TestReadText_OnlyNewLines(t *testing.T) {
	path := createTempFile(t, "\n\n\n")

	lines, err := pipeline.ReadText(path)

	if err != nil {
		t.Fatalf("Expected no error for file with blank lines, got %v", err)
	}

	expected := []string{"", "", ""}
	if !reflect.DeepEqual(lines, expected) {
		t.Fatalf("Expected %v, got %v", expected, lines)
	}
}

func TestReadText_TrailingNewLineAllowed(t *testing.T) {
	content := "....\n....\n"
	path := createTempFile(t, content)

	lines, err := pipeline.ReadText(path)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{
		"....",
		"....",
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Fatalf("Expected %v, got %v", expected, lines)
	}
}

func TestReadText_WhitespacePreserved(t *testing.T) {
	content := "....\n..#.\n"
	path := createTempFile(t, content)

	lines, err := pipeline.ReadText(path)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{
		"....",
		"..#.",
	}
	if !reflect.DeepEqual(lines, expected) {
		t.Fatalf("Expected %v, got %v", expected, lines)
	}
}

func TestReadText_UnicodeCharactersAllowed(t *testing.T) {
	content := "αβγδ\n####\n"
	path := createTempFile(t, content)

	_, err := pipeline.ReadText(path)

	if err != nil {
		t.Fatalf("Expected no error with unicode, got %v", err)
	}
}
