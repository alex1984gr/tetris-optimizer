# How It Works

1. The user provides a text file path as a command-line argument.
2. `main.go` reads the file using `ReadText`.
3. `CheckFile` validates that:
   - Each tetromino has exactly 4 blocks.
   - Only valid characters are used (`#` and `.`).
   - Proper empty line separation exists between tetrominoes.
4. `RecognizeTetrominoes` maps each tetromino's `#` positions into a `Tetromino` struct.
5. `AsembleTetrominoes` places the tetrominoes side by side with letters `A`, `B`, `C`, etc.
6. `PrintOutput` prints the board line by line.
7. `debug_run.go` can be run to test the pipeline with **hardcoded tetrominoes** for debugging purposes.

- If any validation fails, or file cannot be read: the program prints `ERROR`.
- Output is deterministic and testable.
- `debug_run.go` bypasses files to allow **fast debugging**.
