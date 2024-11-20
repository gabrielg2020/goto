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
`cd` into an old project directory.
```bash
cd old-project
```

Find that random directory that you need.
```bash
goto random-directory
```
- Select `random-directory` in fuzzy finder.

## Features to be added
- Speed improvements, most likely with `fd`.
- Gracefully deal with empty inputs.
- Exclude directories.
- Limit search depth.
- Add base directory argument.
- Add customisations with a `~/.goto.conf` file.
  - Add/Remove excluded directories.
  - Change search depth.
  - Change default directory.
