//go:build ignore

package main

import (
	"fmt"
	"tetris-optimizer/pipeline"
)

func main() {
	tetrominoes := [][]string{
		{
			"##..",
			"##..",
			"....",
			"....",
		},
		{
			"##..",
			"##..",
			"....",
			"....",
		},
	}

	board, err := pipeline.AsembleTetrominoes(tetrominoes)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("len(board)=", len(board))
	for _, row := range board {
		fmt.Println(string(row))
	}
}
