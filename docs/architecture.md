# Tetris Optimizer Architecture

## Overview

The project is structured in a **pipeline** style:

ReadText -> CheckFile -> RecognizeTetrominoes -> AsembleTetrominoes -> PrintOutput


- Each step is isolated, testable, and reusable.
- Errors are propagated upstream and result in a single `ERROR` output in main.

---

## Modules

### pipeline/

1. **readtext.go**: Reads the text file line by line.
2. **checkfile.go**: Validates tetromino format and line consistency.
3. **recognizeTetrominoes.go**: Converts raw text into `Tetromino` structs.
4. **asemblietetrominoes.go**: Places tetrominoes on the board.
5. **printoutput.go**: Prints the board to a provided `io.Writer`.

### tests/

- Contains `_test.go` files for each module.
- Follows **TDD** principle: tests written first, then implementation.

### debug_run.go

- A **helper script** to quickly run the pipeline without reading files.
- Uses **hardcoded tetrominoes**.
- Useful for **step-by-step debugging**.

---

## Data Flow

1. **Input**: Text file with tetrominoes.
2. **Validation**: `CheckFile` ensures proper format.
3. **Recognition**: `RecognizeTetrominoes` converts `#` positions into `Tetromino` structs.
4. **Assembly**: `AsembleTetrominoes` generates the board matrix with letters.
5. **Output**: `PrintOutput` prints the matrix to console or other `io.Writer`.
6. **Debug**: `debug_run.go` can be used to bypass file input and see intermediate steps.
