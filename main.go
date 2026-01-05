package main

import (
	"fmt"
	"os"

	"tetris-optimizer/pipeline"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	lines, err := pipeline.ReadText(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	if err := pipeline.CheckFile(lines); err != nil {
		fmt.Println("ERROR")
		return
	}

	tetrominoes := pipeline.RecognizeAllTetrominoes(lines)

	board, err := pipeline.AsembleTetrominoesFromTetrominoes(tetrominoes)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	if err := pipeline.PrintOutput(board, os.Stdout); err != nil {
		fmt.Println("ERROR")
	}
}
