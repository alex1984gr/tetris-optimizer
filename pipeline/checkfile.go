package pipeline

import "errors"

// CheckFile validates that lines define zero or more tetromino blocks
// each of four lines of 4 characters ('.' or '#'), separated by empty lines.
func CheckFile(lines []string) error {
	// index of the current line being validated
	lineIndex := 0

	// iterate until we've validated all provided lines
	for lineIndex < len(lines) {

		// skip any leading empty lines
		if lines[lineIndex] == "" {
			lineIndex++
			continue
		}

		// ensure there are at least 4 lines remaining for a tetromino
		if lineIndex+3 >= len(lines) {
			return errors.New("incomplete tetromino")
		}

		// validate the 4 lines of this tetromino
		for i := 0; i < 4; i++ {
			line := lines[lineIndex+i]

			// each line must be exactly 4 characters long
			if len(line) != 4 {
				return errors.New("invalid line length")
			}

			// each character must be either '.' or '#'
			for _, ch := range line {
				if ch != '.' && ch != '#' {
					return errors.New("invalid character")
				}
			}
		}

		// move past the 4 lines we just validated
		lineIndex += 4

		// after a tetromino, if more lines exist, the next must be an empty separator
		if lineIndex < len(lines) {
			if lines[lineIndex] != "" {
				return errors.New("missing empty line between tetrominoes")
			}
			lineIndex++
		}
	}

	return nil
}
