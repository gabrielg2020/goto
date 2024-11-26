---
layout: default
title: Home
---

# Welcome to Goto

`goto` is a lightweight terminal tool that simplifies directory navigation.

---

## Contents
- [Getting Started](getting-started.md)
- [Usage](usage.md)
- [Configuration](configuration.md)
- [Developers](developers.md)
- [Change Log](changelog.md)

---

# Getting Started

## Installation
### Using a Package Manager
Coming soon!

### Manual Installation
#### 1. Clone the repository:
```bash
git clone https://github.com/gabrielg2020/goto.git
cd goto
```

#### 2. Build the binary:
```bash
go build -o goto ./cmd/
```

#### 3. Move the binary to a directory in your `PATH`:
```bash
mv goto /usr/local/bin/
```

#### 4. Install dependencies
Please make sure that these are installed.
- [fzf](https://github.com/junegunn/fzf)

---

### Shell Integration
#### Automated Integration
Run the provided setup script to automate this setup:
```bash
./setup.sh
```

### Manual Integration
Add the following function to your `.bashrc` or `.zshrc`:
```bash
goto() {
    dir=$(/usr/local/bin/goto "$@")

    if [ -d "$dir" ]; then
        echo "Changing directory to: $dir"
        cd "$dir"
    else
        echo "No directory selected or invalid directory: $dir"
    fi
}
```

---

### Verify Installation
>[!NOTE]
> Coming soon!

```bash
goto --help
```

---

# Usage

## Basic Command
```bash
goto <pattern>
```

### Example
```bash
goto src
```

----

## Options

| Option      | Description                           |
| ----------- | ------------------------------------- |
| `-d`        | Specify maximum search depth.         |
| `-h`        | Display help message.                 |
| `<pattern>` | Fuzzy search pattern for directories. |

---

# Configuration

`goto` allows you to customize its behavior using a configuration file. This file provides options to define default search settings and exclude specific directories from the results.

---

## Config File Location

The configuration file is located at: `~/.goto_config.yaml`

This file is optional. If it doesn't exist, `goto` uses sensible defaults.

---

## Config File Format

The configuration file is written in YAML format. Below is an example configuration:

```yaml
maxDepth: <integer>
excludeDirs:
  - <string>
  - <string>
```

## Available Options

|    Option     | Description                                 | Default                  |
| ------------- | ------------------------------------------- | ------------------------ |
|  `maxDepth`   | Maximum search depth.                       | `5`                      |
| `excludeDirs` | Directories to be excluded from the search. | `".git", "node_modules"` |

---

## Example Configuration 

Here's a complete example of a typical `.goto_config.yaml` file:
```yaml
maxDepth: 5
excludeDirs:
  - build
  - dist
  - .cache
```

---

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

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this project adheres to [Semantic Versioning](https://semver.org/).

---

## [1.0.0] - TBD
### Added
Coming soon!

---

[Back](index.md)