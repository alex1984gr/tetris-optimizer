package tests

import( 
	"os"
	"reflect"
	"testing"
	"../src"
	"
)

func createTempFile(t *testing.T, content string) string {
    t.Helper()

	tmpDir := t.TempDir()
	filePath := Filepath.Join(tmpDir, "test.txt")

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	return filePath
}

func TestReadText_Validfile(t *testing.T) {
    content := "....\n....\n....\n....\n"
	path := createTempFile(t, content)
	
	lines, err := ReadText(path)

	if err != nil{
		t.Errorf("Expected no error, got %v", err)
	}

	expected := []string {
	"....",
	"....",
	"....",
	"....",
	}
	
	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected %v, got %v", expected, lines)
	}
}

func TestReadText_FileDoesNotExist(t *testing.T) {
	path := ReadText("non_existent_file.txt")

	if err == nil {
		t.Errorf("Expected error for non-existent file, got nil")
	}
}
	
func TestReadText_EmptyFile(t *testing.T) {
	path := createTempFile(t, "")

	_, err := ReadText(path)

	if err == nil {
		t.Fatal("expected error for empty file, got nil")
	}	
}
