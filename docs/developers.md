# Developers Guide

This guide explains the architecture, key decisions, and trade-offs made during the development of `goto`.

---

## Project Structure

| Directory/File          | Description                                           |
| ----------------------- | ----------------------------------------------------- |
| **`cmd/main.go`**       | Entry point for the application.                      |
| **`cmd/pkg/search.go`** | Logic for directory searching using `find` and `fzf`. |
| **`cmd/pkg/config.go`** | Logic for dealing with user configurations.           |
| **`setup.sh`**          | Shell integration setup script.                       |
| **`test/`**             | Unit tests.                                           |

---

## Key Decisions
### 1. Why Go?
- **Reason**: Go provides a lightweight, fast, and cross-platform environment for CLI tools.
- **Trade-offs**: Limited third-party libraries compared to other languages, but we prioritize simplicity.

### 2. Why `find` and `fzf`?
- `find` is robust and handles directory traversal efficiently.
- `fzf` provides a fast, well-supported fuzzy search interface.

### 3. Configuration Format
- **YAML** was chosen for its readability and flexibility.
- **Alternatives Considered**: JSON (too verbose), INI (limited flexibility).

### 4. Shell Integration Approach
- Wrapper functions in `.bashrc` and `.zshrc` enable seamless navigation.
- **Challenge**: Go binaries cannot directly change the shell's state, so outputting the path to a wrapper script was the solution.

---

## Future Considerations
- Add bookmarks for frequently used directories.
- Directory aliases for common directions.
- Regex and advanced filtering.
- Enhanced fuzzy search customisation.
- Explore direct Fish shell support.
- Use of parallelisation and Goroutines.
- Config Validator.
- `goto` files:
  - Maybe open with preferred text editor?
  - Maybe jump to file directory?
  - Maybe both?

---

[Back](index.md)