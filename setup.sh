#!/bin/bash

# Determine the shell configuration file
if [ -n "$BASH_VERSION" ]; then
    SHELL_RC="$HOME/.bashrc"
elif [ -n "$ZSH_VERSION" ]; then
    SHELL_RC="$HOME/.zshrc"
else
    echo "Unsupported shell"
    exit 1
fi

# Add the goto shell function to the shell configuration file
if ! grep -q "function goto" "$SHELL_RC"; then
    echo "Adding the goto shell function to $SHELL_RC"
    cat << 'EOF' >> "$SHELL_RC"
# Goto
goto() {
    dir=$($HOME/goto/goto-bin "$@")

    if [ -d "$dir" ]; then
        echo "Changing directory to: $dir"
        cd "$dir"
    else
        echo "No directory selected or invalid directory: $dir"
    fi
}
EOF
    echo "Done! Restart your shell or run 'source $SHELL_RC' to start using the goto command"
else
    echo "The goto shell function is already defined in $SHELL_RC"
fi