package pipeline

// Point holds X,Y coordinates for a single block within a tetromino.
type Point struct {
	X int // column index
	Y int // row index
}

// Tetromino is a normalized shape represented by its block points.
type Tetromino struct {
	Blocks []Point // list of block coordinates relative to (0,0)
}

// RecognizeAllTetrominoes parses full-file lines separated by empty lines
// and returns all normalized Tetromino structs found.
func RecognizeAllTetrominoes(lines []string) []Tetromino {
	var tetrominoes []Tetromino
	var current []Point // collected points for the current tetromino

	row := 0 // row index within the current tetromino block
	for _, line := range lines {
		// empty line separates tetrominoes
		if line == "" {
			if len(current) > 0 {
				tetrominoes = append(tetrominoes, normalize(current))
				current = nil
				row = 0
			}
			continue
		}

		// record '#' positions in the current row
		for col, ch := range line {
			if ch == '#' {
				current = append(current, Point{X: col, Y: row})
			}
		}
		row++
	}

	// append final tetromino if file didn't end with empty line
	if len(current) > 0 {
		tetrominoes = append(tetrominoes, normalize(current))
	}

	return tetrominoes
}

// RecognizeTetrominoes validates a single tetromino (4 lines) and returns
// an error describing problems, or nil on success. This matches tests.
func RecognizeTetrominoes(lines []string) error {
	// must be exactly 4 lines
	if len(lines) != 4 {
		return &ErrInvalidTetromino{Msg: "must be 4 lines"}
	}
	var pts []Point
	for y, line := range lines {
		// each line must be 4 chars wide
		if len(line) != 4 {
			return &ErrInvalidTetromino{Msg: "line length must be 4"}
		}
		for x, ch := range line {
			if ch == '#' {
				pts = append(pts, Point{X: x, Y: y})
			}
		}
	}
	// exactly 4 blocks required for a tetromino
	if len(pts) != 4 {
		return &ErrInvalidTetromino{Msg: "must contain exactly 4 blocks"}
	}
	// check connectivity: perform a simple DFS/BFS from first block
	seen := map[Point]bool{}
	stack := []Point{pts[0]}
	seen[pts[0]] = true
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// orthogonal neighbors
		neighbors := []Point{{p.X + 1, p.Y}, {p.X - 1, p.Y}, {p.X, p.Y + 1}, {p.X, p.Y - 1}}
		for _, n := range neighbors {
			for _, q := range pts {
				if q == n && !seen[q] {
					seen[q] = true
					stack = append(stack, q)
				}
			}
		}
	}
	// all 4 blocks must be reachable
	if len(seen) != 4 {
		return &ErrInvalidTetromino{Msg: "blocks are not connected"}
	}
	return nil
}

// ErrInvalidTetromino is returned when a tetromino fails validation.
type ErrInvalidTetromino struct{ Msg string }

func (e *ErrInvalidTetromino) Error() string { return e.Msg }

// normalize shifts all points so minimum X and Y become 0, returning a Tetromino.
func normalize(blocks []Point) Tetromino {
	minX, minY := blocks[0].X, blocks[0].Y

	for _, p := range blocks {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	for i := range blocks {
		blocks[i].X -= minX
		blocks[i].Y -= minY
	}

	return Tetromino{Blocks: blocks}
}
