# goto
`goto` is a shell function that adds quick navigation to directories using fuzzy search powered by `fzf`. Jump to directories by matching patterns anywhere within directory names.

## Features
- Fuzzy Matching: Match directories containing the search pattern anywhere in their names.
- Interactive Selection: Use `fzf` for real-time, interactive directory selection.

## Prerequisites  
- `fzf`: A general-purpose command-line fuzzy finder.

## Installation
macOS:
```shell
brew install fzf
```

Linux (Debian/Ubuntu):
```shell
sudo apt-get install fzf
```

## Usage
>[!WARNING]
>`goto` is slow and doesn't have a search depth. Beware when using, it may take a long time...

### Basic Command
```shell
goto <pattern>
```

### Example
`cd` into an old project directory.
```shell
cd old-project
```

Find that random directory that you need.
```shell
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
