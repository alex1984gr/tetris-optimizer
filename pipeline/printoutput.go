package pipeline

import (
	"errors"
	"io"
)

// PrintOutput prints the final tetromino board to the given writer.
func PrintOutput(board [][]rune, w io.Writer) error {
	if w == nil {
		return errors.New("writer is nil")
	}

	if len(board) == 0 {
		return nil
	}

	for _, row := range board {
		for _, cell := range row {
			_, err := w.Write([]byte(string(cell)))
			if err != nil {
				return err
			}
		}

		_, err := w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
