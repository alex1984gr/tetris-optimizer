package pipeline

import (
	"bufio"  // for reading file line-by-line
	"errors" // for constructing error values
	"os"     // for opening files
)

// ReadText opens the file at filePath and returns its lines.
func ReadText(filePath string) ([]string, error) {
	// open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// ensure file is closed when function returns
	defer file.Close()

	// scanner reads the file line-by-line
	scanner := bufio.NewScanner(file)
	var lines []string

	// collect each scanned line into the lines slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// if file had no lines, return an error
	if len(lines) == 0 {
		return nil, errors.New("file is empty")
	}

	return lines, nil
}
