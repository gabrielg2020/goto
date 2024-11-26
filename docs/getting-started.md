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

[Back](index.md)