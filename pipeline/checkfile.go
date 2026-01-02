package pipeline

import "errors"

func CheckFile(lines []string) error {
	lineIndex := 0

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

		lineIndex += 4

		if lineIndex < len(lines) {
			if lines[lineIndex] != "" {
				return errors.New("missing empty line between tetrominoes")
			}
			lineIndex++
		}
	}

	return nil
}
