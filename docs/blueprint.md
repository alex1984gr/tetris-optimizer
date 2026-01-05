# Blueprint of Tetris Optimizer

## Modules

main.go
└─ pipeline/
├─ readtext.go
├─ checkfile.go
├─ recognizeTetrominoes.go
├─ asemblietetrominoes.go
└─ printoutput.go
└─ tests/
├─ readtext_test.go
├─ checkfile_test.go
├─ recognizeTetrominoes_test.go
├─ asemblietetrominoes_test.go
└─ printoutput_test.go
└─ debug_run.go


## Data Structure

- `Tetromino` struct:
  - `Blocks []Point`
- `Point` struct:
  - `X int`
  - `Y int`
- Board:
  - `[][]string` or `[][]rune` depending on stage

## Flow

File -> ReadText -> CheckFile -> RecognizeTetrominoes -> AsembleTetrominoes -> PrintOutput -> Console
Debug: debug_run.go bypasses File input for quick testing


- Each module is isolated.
- Easy to unit test.
- Easy to swap implementations if needed.
