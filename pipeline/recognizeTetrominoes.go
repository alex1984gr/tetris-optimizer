package pipeline

import "errors"

type point struct {
	row int
	col int
}

func RecognizeTetrominoes(tetromino []string) error {
	if len(tetromino) != 4 {
		return errors.New("invalid tetromino height")
	}

	var blocks []point
	for r := 0; r < 4; r++ {
		if len(tetromino[r]) != 4 {
			return errors.New("invalid tetromino width")
		}

		for c := 0; c < 4; c++ {
			switch tetromino[r][c] {
			case '#':
				blocks = append(blocks, point{r, c})
			case '.':
				// ok
			default:
				return errors.New("invalid character")
			}
		}
	}

	if len(blocks) != 4 {
		return errors.New("invalid number of blocks")
	}

	visited := make(map[point]bool)
	stack := []point{blocks[0]}

	directions := []point{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current] {
			continue
		}

		visited[current] = true

		for _, d := range directions {
			nr := current.row + d.row
			nc := current.col + d.col

			if nr >= 0 && nr < 4 && nc >= 0 && nc < 4 {
				if tetromino[nr][nc] == '#' {
					neighbor := point{nr, nc}
					if !visited[neighbor] {
						stack = append(stack, neighbor)
					}
				}
			}
		}
	}

	if len(visited) != 4 {
		return errors.New("tetromino blocks not connected")
	}

	return nil
}
