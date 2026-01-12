package pipeline

// Point represents the X,Y coordinates of a single block ('#')
// inside a tetromino grid.
type Point struct {
	X int // column index
	Y int // row index
}

// Tetromino represents a normalized tetromino shape.
// All blocks are shifted so the top-left block starts at (0,0).
type Tetromino struct {
	Blocks []Point
}

// ErrInvalidTetromino is returned when a tetromino
// does not satisfy validation rules.
type ErrInvalidTetromino struct {
	Msg string
}

// Error implements the built-in error interface.
func (e *ErrInvalidTetromino) Error() string {
	return e.Msg
}

// RecognizeAllTetrominoes parses all file lines, separates tetrominoes
// by empty lines, validates them if strict mode is enabled,
// and returns a slice of normalized Tetromino structs.
func RecognizeAllTetrominoes(lines []string, strict bool) ([]Tetromino, error) {
	var tetrominoes []Tetromino // final result
	var currentLines []string   // lines for the current tetromino

	// Iterate over every line in the file
	for _, line := range lines {

		// An empty line separates tetrominoes
		if line == "" {
			// If we have collected lines for a tetromino
			if len(currentLines) > 0 {

				// In strict mode, validate the tetromino
				if strict {
					if err := RecognizeTetrominoes(currentLines); err != nil {
						return nil, err
					}
				}

				// Parse and normalize the tetromino
				tetrominoes = append(tetrominoes, parseAndNormalize(currentLines))
				currentLines = nil // reset for the next tetromino
			}
			continue
		}

		// Non-empty lines belong to the current tetromino
		currentLines = append(currentLines, line)
	}

	// Handle the last tetromino if file does not end with an empty line
	if len(currentLines) > 0 {

		if strict {
			if err := RecognizeTetrominoes(currentLines); err != nil {
				return nil, err
			}
		}

		tetrominoes = append(tetrominoes, parseAndNormalize(currentLines))
	}

	return tetrominoes, nil
}

// RecognizeTetrominoes validates a single tetromino.
// It checks dimensions, block count, and connectivity.
func RecognizeTetrominoes(lines []string) error {

	// A tetromino must be exactly 4 lines tall
	if len(lines) != 4 {
		return &ErrInvalidTetromino{Msg: "must be exactly 4 lines"}
	}

	var blocks []Point // stores positions of all '#' blocks

	// Iterate through each row
	for y, line := range lines {

		// Each line must be exactly 4 characters wide
		if len(line) != 4 {
			return &ErrInvalidTetromino{Msg: "each line must be 4 characters wide"}
		}

		// Iterate through characters in the line
		for x, ch := range line {
			if ch == '#' {
				blocks = append(blocks, Point{X: x, Y: y})
			}
		}
	}

	// A valid tetromino must contain exactly 4 blocks
	if len(blocks) != 4 {
		return &ErrInvalidTetromino{Msg: "tetromino must contain exactly 4 blocks"}
	}

	// Connectivity check using DFS (only orthogonal neighbors)
	visited := map[Point]bool{}
	stack := []Point{blocks[0]}
	visited[blocks[0]] = true

	// Depth-first search
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Possible neighbors (up, down, left, right)
		neighbors := []Point{
			{p.X + 1, p.Y},
			{p.X - 1, p.Y},
			{p.X, p.Y + 1},
			{p.X, p.Y - 1},
		}

		// Check if neighbors exist in the block list
		for _, n := range neighbors {
			for _, b := range blocks {
				if b == n && !visited[b] {
					visited[b] = true
					stack = append(stack, b)
				}
			}
		}
	}

	// All 4 blocks must be connected
	if len(visited) != 4 {
		return &ErrInvalidTetromino{Msg: "blocks are not connected"}
	}

	return nil
}

// parseAndNormalize converts raw tetromino lines into
// a normalized Tetromino struct.
func parseAndNormalize(lines []string) Tetromino {
	var blocks []Point

	// Collect all '#' block positions
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				blocks = append(blocks, Point{X: x, Y: y})
			}
		}
	}

	return normalize(blocks)
}

// normalize shifts all block coordinates so that
// the minimum X and Y values become 0.
func normalize(blocks []Point) Tetromino {

	// Initialize minimum coordinates
	minX := blocks[0].X
	minY := blocks[0].Y

	// Find the minimum X and Y among all blocks
	for _, b := range blocks {
		if b.X < minX {
			minX = b.X
		}
		if b.Y < minY {
			minY = b.Y
		}
	}

	// Shift all blocks relative to (0,0)
	for i := range blocks {
		blocks[i].X -= minX
		blocks[i].Y -= minY
	}

	return Tetromino{Blocks: blocks}
}
