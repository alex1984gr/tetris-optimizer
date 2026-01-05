# Coding Standards

1. **Language**: Go (1.20+ recommended)
2. **File Naming**: Lowercase with underscores for readability (e.g., `readtext.go`).
3. **Function Names**: CamelCase starting with uppercase if exported, lowercase if internal.
4. **Error Handling**:
   - Always return `error` if something can fail.
   - In `main.go`, errors print `ERROR`.
5. **Testing**:
   - Follow TDD: write `_test.go` first.
   - Each function in `pipeline/` has a corresponding test.
6. **Imports**:
   - Only standard Go packages.
   - Group by standard library first.
7. **Code Style**:
   - `go fmt` for formatting.
   - Avoid global variables.
   - Keep functions < 50 lines if possible.
8. **Debugging**:
   - Use `debug_run.go` for fast iterative testing.
   - Avoid editing production code just for debugging.
