#!/bin/bash

function goto() {
    if [ -z "$1" ]; then
        echo "Usage: goto <directory_name_pattern>"
        return 1
    fi
    local pattern="$1"
    local dir
    dir=$(find . -type d -iname '*' -print 2>/dev/null \
        | fzf --query="$patten" --height 40% --reverse --prompt="Goto> ")
    if [ -n "$dir" ]; then
        cd "$dir" || echo "Error: Cannot change directory to $dir"
    else
        echo "No directory selected."
    fi
}
