package main

// import block: bring in packages used by the program
import (
	"fmt" // formatting and printing
	"os"  // access to command-line args and Stdout

	"tetris-optimizer/pipeline" // local package with program logic
)

// main is the program entry point
func main() {
	// verify exactly one command-line argument (input file path)
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	// read file lines from provided path
	lines, err := pipeline.ReadText(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	// validate file format (tetromino blocks and separators)
	if err := pipeline.CheckFile(lines); err != nil {
		fmt.Println("ERROR")
		return
	}

	// parse the entire file into normalized Tetromino structs
	tetrominoes, err := pipeline.RecognizeAllTetrominoes(lines, true)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	// assemble tetrominoes onto a minimal board
	board, err := pipeline.AsembleTetrominoesFromTetrominoes(tetrominoes)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	// print the resulting board to standard output
	if err := pipeline.PrintOutput(board, os.Stdout); err != nil {
		fmt.Println("ERROR")
	}
}
