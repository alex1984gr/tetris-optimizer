# Tetris Optimizer

Tetris Optimizer is a program written in Go that takes a text file containing tetrominoes and places them into the smallest possible square.

Each tetromino is identified by its blocks and represented in the final board using uppercase letters (A, B, C, etc). Empty spaces are represented with dots (.).

The project is designed with a modular architecture and is fully developed using Test Driven Development (TDD).

---------------------------------------------------------------------

PROJECT PURPOSE

- Understanding backtracking algorithms
- Implementing a clean data processing pipeline
- Writing clear and maintainable Go code
- Strict input validation
- Designing systems with clear separation of responsibilities

---------------------------------------------------------------------

PROJECT STRUCTURE

- readtext.go  
  Reads the input file and returns the file lines

- checkfile.go  
  Validates the format of the input file

- recognizeTetrominoes.go  
  Identifies valid tetrominoes and converts them into an internal representation

- asemblietetrominoes.go  
  Places all tetrominoes into the smallest possible square using backtracking

- printoutput.go  
  Prints the final result

- tests/  
  Contains unit tests for each module

---------------------------------------------------------------------

EXECUTION FLOW (PIPELINE)

Input File  
↓  
Read Text  
↓  
Check File  
↓  
Recognize Tetrominoes  
↓  
Assemble Tetrominoes  
↓  
Print Output

---------------------------------------------------------------------

INPUT FILE FORMAT

- Each tetromino consists of exactly 4 lines
- Each line has exactly 4 characters
- Allowed characters: # and .
- There must be an empty line between tetrominoes
- One empty line is allowed at the end of the file

Example:

....
..##
..##
....

####
....
....
....

---------------------------------------------------------------------

OUTPUT EXAMPLE

AABB
AABB
CCCC
....

---------------------------------------------------------------------

RUNNING TESTS

From the project root directory:

go test ./tests -v

For debuging :

go run debug_run.go

---------------------------------------------------------------------

RUNNING THE PROGRAM

go run main.go <path_to_input_file>

If the file format is invalid, the program prints:

ERROR

---------------------------------------------------------------------

WHAT THIS PROJECT COVERS

- Test Driven Development (TDD)
- Unit testing
- Data validation
- Backtracking algorithms
- Clean separation of concerns
- Safe and explicit error handling

---------------------------------------------------------------------
