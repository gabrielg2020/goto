# goto
`goto` is a shell function that adds quick navigation to directories using fuzzy search powered by `fzf`. Jump to directories by matching patterns anywhere within directory names.

## Features
- Fuzzy Matching: Match directories containing the search pattern anywhere in their names.
- Interactive Selection: Use `fzf` for real-time, interactive directory selection.

## Prerequisites  
- `fzf`: A general-purpose command-line fuzzy finder.
- `git` (optional).

## Installation
### 1. Install `fzf`
macOS:
```bash
brew install fzf
```

Linux (Debian/Ubuntu):
```bash
sudo apt-get install fzf
```
### 2. Clone the Repository into Your Home Directory
```bash
git clone https://github.com/gabrielg2020/goto.git
```

### 3. Source the goto Function in Your Shell Configuration
Add following line to your `~/.bashrc` or `~/zshrc` file:
```bash
# Source the goto function
if [ -f ~/goto/src/goto.sh ]; then
    source ~/goto/src/goto.sh
fi
```

#### 3.1 Reload Shell Configuration
```bash
source ~/.bashrc    # For Bash
# or
source ~/.zshrc     # For Zsh
```

## Usage
>[!WARNING]
>`goto` is slow and doesn't have a search depth. Beware when using, it may take a long time...

### Basic Command
```bash
goto <pattern>
```

### Example
Lets say you have a project that has a directory structure like this
```
project/
├─ sizzle/
│  ├─ foo/
├─ grizzle/
│  ├─ fizz/
├─ swizzle/
├─ fizzle/
├─ ...
```

#### `cd` into project directory
```bash
cd project
```

#### Find all directories that start with `f`
```bash
goto f
```
- Lists `sizzle/foo`, `grizzle/fizz`, `fizzle`.

#### Find all directories that end with `izzle`
```bash
goto izzle
```
- Lists `sizzle`, `grizzle`, `swizzle`, `fizzle`.

#### Auto-`cd` into `swizzle/foo`
```bash
goto swizzle
```
- Now in the `swizzle/foo` directory.

## Features to be added
- Speed improvements, most likely with `fd`.
- Gracefully deal with empty inputs.
- Exclude directories.
- Limit search depth.
- Add base directory argument.
