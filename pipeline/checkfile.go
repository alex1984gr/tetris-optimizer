package pipeline

import "errors"

func CheckFile(lines []string) error {
	if len(lines) == 0 {
		return errors.New("empty input")
	}

	lineIndex := 0
	hasContent := false

	for lineIndex < len(lines) {

		if lines[lineIndex] == "" {
			lineIndex++
			continue
		}

		if lineIndex+3 >= len(lines) {
			return errors.New("incomplete tetromino")
		}

		for i := 0; i < 4; i++ {
			line := lines[lineIndex+i]

			if len(line) != 4 {
				return errors.New("invalid line length")
			}

			for _, ch := range line {
				if ch != '.' && ch != '#' {
					return errors.New("invalid character")
				}
			}
		}

		hasContent = true
		lineIndex += 4

		if lineIndex < len(lines) {
			if lines[lineIndex] != "" {
				return errors.New("missing empty line between tetrominoes")
			}
			lineIndex++
		}
	}

	if !hasContent {
		return errors.New("no tetromino data")
	}

	return nil
}
