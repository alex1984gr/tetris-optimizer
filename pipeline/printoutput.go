package pipeline

import (
	"errors" // for error creation
	"io"     // io.Writer type
)

// PrintOutput writes the board rows to the provided io.Writer, one line per row.
func PrintOutput(board [][]rune, w io.Writer) error {
	// writer must be provided
	if w == nil {
		return errors.New("nil writer")
	}
	// nil board means nothing to print
	if board == nil {
		return nil
	}
	// write each row as a string followed by newline
	for _, row := range board {
		if len(row) == 0 {
			_, err := w.Write([]byte("\n"))
			if err != nil {
				return err
			}
			continue
		}
		// convert rune slice to string and write
		_, err := w.Write([]byte(string(row)))
		if err != nil {
			return err
		}
		// write newline after the row
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
