#!/bin/bash

function goto() {
    # Check argument
    if [ -z "$1" ]; then
        echo "Usage: goto <directory_name_pattern>"
        return 1
    fi

    # Setup local variables
    local pattern="$1"
    local dir

    # Find directory and pipe to fzf
    dir=$(find . -type d -iname '*' -print 2>/dev/null \
        | fzf --query="$pattern" --height 40% --reverse --prompt="Goto> ")
    if [ -n "$dir" ]; then
        cd "$dir" || echo "Error: Cannot change directory to $dir"
    else
        echo "No directory selected."
    fi
}
