# goto

`goto` is a lightweight terminal tool that simplifies directory navigation.

## Features
- Quickly navigate directories with fuzzy search.
- Customise behaviour with configuration options.
- Compatible with Bash and Zsh shells.

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

### Verify Installation
>[!NOTE]
> Coming soon!
```bash
goto --help
```

## Usage
### Basic Command
```bash
goto <pattern>
```

#### Example
```bash
goto src
```

### Options
| Option | Description |
| :---------: | :------------------------------------ |
| `-d` | Specify maximum search depth. |
| `<pattern>` | Fuzzy search pattern for directories. |

## Configuration
`goto` uses a configuration file located at `~/.goto_config.yaml`. This file is optional and allows you to customise the tool's behaviour.

### Example Configuration
```yaml
maxDepth: 3
excludeDirs:
  - node_modules
  - .git
  - .vscode
```

### Options
| Option | Description | Default |
| :-----------: | :------------------------------------------ | :----------------------- |
| `maxDepth` | Maximum search depth. | `5` |
| `excludeDirs` | Directories to be excluded from the search. | `".git", "node_modules"` |

## Contributing
We welcome contributions! Please open an issue or pull request with your changes.

### developers.md
The [developers.md](https://github.com/gabrielg2020/goto/blob/main/developers.md) file explains some key ideas and decisions, future consideration and extra knowledge of this project.

### Change Log
[Change Log](https://github.com/gabrielg2020/goto/blob/main/CHANGELOG.md)

### Documentation
Coming Soon! 

## License
The MIT Licence (MIT)

Copyright (c) 2023 - 2024 Gabriel Guimaraes. All Rights Reserved.